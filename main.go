package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/recoilme/pudge"

	"jsapi-scraper/shop"
)

const Token = "6219604147:AAERFP-_PfSELN3-gorzE9czM6WR-3Rum-Q"
const ChatID = "1912073977"

func main() {
	isInit := flag.Bool("init", false, "initialize database")
	flag.Parse()

	shop.XXX_steg()
	return

	// r := regexp.MustCompile("[^a-z0-9 .]")
	// fmt.Println("regexp:", r)

	// str := "OPPO Find X3 Lite (5G, 128 GB, 6.44\", 64 MP, Blau)"
	// fmt.Println(str)

	// fmt.Println(strings.NewReplacer(" ", "-", ".", "-").Replace(r.ReplaceAllString(strings.ToLower(str), "$1")))
	// return

	// Close all database on exit
	defer pudge.CloseAll()

	for _, _shop := range []shop.IShop{shop.XXX_conrad(), shop.XXX_melectronics(), shop.XXX_microspot(), shop.XXX_mobilezone(), shop.XXX_interdiscount()} {
		for _, product := range _shop.Fetch() {
			productLine := fmt.Sprintf("%s\n%8.2f %8.2f %8.2f %3.f%%\n%s", product.Name, product.RetailPrice, product.Price, product.Savings, product.Discount, _shop.ResolveURL(product.URL))
			fmt.Println(strings.ReplaceAll(productLine, "\n", "\t"))

			var oldProduct shop.Product
			pudge.Get("products", product.Code, &oldProduct)
			oldProductLine := fmt.Sprintf("%s\n%8.2f %8.2f %8.2f %3.f%%\n%s", oldProduct.Name, oldProduct.RetailPrice, oldProduct.Price, oldProduct.Savings, oldProduct.Discount, _shop.ResolveURL(oldProduct.URL))
			if productLine != oldProductLine {
				fmt.Println(strings.ReplaceAll(oldProductLine, "\n", "\t"))
				fmt.Println()
				pudge.Set("products", product.Code, product)

				if !*isInit {
					_, err := SendMessage(productLine)
					if err != nil {
						panic(err)
					}
				}
			}
		}
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
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	// fmt.Println(string(body))

	// Return
	return true, nil
}
