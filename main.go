package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/recoilme/pudge"

	"jsapi-scraper/shop"
)

const Token = "6219604147:AAERFP-_PfSELN3-gorzE9czM6WR-3Rum-Q"
const ChatID = "1912073977"

func main() {
	isDryRun := flag.Bool("dryrun", false, "dry run (avoid making external calls)")
	flag.Parse()

	// r := regexp.MustCompile("[^a-z0-9 .]")
	// fmt.Println("regexp:", r)

	// str := "OPPO Find X3 Lite (5G, 128 GB, 6.44\", 64 MP, Blau)"
	// fmt.Println(str)

	// fmt.Println(strings.NewReplacer(" ", "-", ".", "-").Replace(r.ReplaceAllString(strings.ToLower(str), "$1")))
	// return

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println()
			fmt.Printf("%s\n", strings.Join(strings.Split(string(debug.Stack()), "\n")[7:], "\n"))

			if isDryRun != nil && *isDryRun {
			} else {
				if _, err := SendMessage(fmt.Sprintf("%v", err)); err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	wg := sync.WaitGroup{}

	_products := map[string]*[]*shop.Product{}

	for _, _shop := range []shop.IShop{
		shop.XXX_alltron(isDryRun),
		shop.XXX_alternate(isDryRun),
		shop.XXX_brack(isDryRun),
		shop.XXX_conrad(isDryRun),
		shop.XXX_foletti(isDryRun),
		shop.XXX_fust(isDryRun),
		shop.XXX_interdiscount(isDryRun),
		shop.XXX_mediamarkt(isDryRun),
		shop.XXX_mediamarkt_refurbished(isDryRun),
		shop.XXX_melectronics(isDryRun),
		shop.XXX_microspot(isDryRun),
		shop.XXX_mobilezone(isDryRun),
		shop.XXX_stegpc(isDryRun),
	} {
		wg.Add(1)

		_products[_shop.Name()] = nil

		go func(_shop shop.IShop) {
			defer wg.Done()

			_products[_shop.Name()] = _shop.Fetch()
		}(_shop)
	}

	wg.Wait()

	db, err := pudge.Open("products", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ids := map[string]bool{}

	if keys, err := db.Keys(nil, 0, 0, true); err != nil {
		panic(err)
	} else {
		for _, key := range keys {
			ids[string(key)] = true
		}
	}
	// fmt.Println(ids)

	for name, products := range _products {
		fmt.Println()

		_num := 0
		if products != nil {
			_num = len(*products)
		}
		_name := fmt.Sprintf("%s (%d)", name, _num)

		fmt.Printf("%s\n%s\n", _name, strings.Repeat("=", len(_name)))

		if products != nil {
			for _, product := range *products {
				id := product.Code
				delete(ids, id)

				_name := product.Name
				if product.Quantity > 0 {
					_name += fmt.Sprintf(" (%d)", product.Quantity)
				}
				priceLine := ""
				if product.Price != product.RetailPrice {
					priceLine = fmt.Sprintf("%8.2f %8.2f %3d%%", product.Price, product.Savings, int(product.Discount))
				}

				fmt.Printf("%-69s %8.2f %22s %s\n", _name, product.RetailPrice, priceLine, product.URL)

				notify := false

				var oldProduct shop.Product
				if ok, _ := db.Has(id); ok {
					db.Get(id, &oldProduct)
				}

				if oldProduct != *product {
					db.Set(id, product)

					notify = true
				}

				if notify {
					if isDryRun != nil && !*isDryRun {
						priceLine := ""
						if product.Price != product.RetailPrice {
							priceLine = fmt.Sprintf("%-8.2f %-8.2f %-3d%%", product.Price, product.Savings, int(product.Discount))
						}

						productLine := fmt.Sprintf("%s\n%-8.2f %s\n\n%s", _name, product.RetailPrice, priceLine, product.URL)

						if _, err := SendMessage(productLine); err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}

	// fmt.Println(ids)
	for id := range ids {
		db.Delete(id)
	}
}

func getURL() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", Token)
}

func SendMessage(text string) (bool, error) {
	// Global variables
	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", getURL())
	body, _ := json.Marshal(map[string]string{
		"chat_id": ChatID,
		"text":    text,
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Body
	_, err = io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	// Return
	return true, nil
}
