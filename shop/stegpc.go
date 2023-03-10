package shop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func XXX_stegpc(isDryRun bool) IShop {
	const _name = "Steg Electronics"
	const _url = "https://www.steg-electronics.ch/de/product/list/11853?sortKey=preisasc&smsc=1000"

	const _debug = false

	skips := map[string]bool{
		"BEAFON":      true,
		"CROSSCALL":   true,
		"CYRUS":       true,
		"DORO":        true,
		"FELLOWES":    true,
		"GIGASET":     true,
		"JABLOCOM":    true,
		"KONTAKT":     true,
		"MAGNETOPLAN": true,
		"MAUL":        true,
		"OLYMPIA":     true,
		"PANASONIC":   true,
		"RUGGEAR":     true,
		"SIGEL":       true,
		"STYRO":       true,
		"SWISSTONE":   true,
	}

	type _Response struct {
		code  string
		title string

		link string

		oldPrice float32
		price    float32
	}

	type _Body struct {
		NewProductList string `json:"newProductList"`
	}

	var _result []_Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	fn := "shop/stegpc.json"

	if isDryRun {
		if body, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else {
			_body = body
		}
	} else {
		resp, err := http.Post(_url, "application/json", bytes.NewBuffer(nil))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if body, err := io.ReadAll(resp.Body); err != nil {
			panic(err)
		} else {
			_body = body
		}

		os.WriteFile(path+fn, _body, 0664)
	}
	// fmt.Println(string(_body))

	var body _Body
	if err := json.Unmarshal(_body, &body); err != nil { // Parse []byte to go struct pointer
		panic(err)
	}
	// fmt.Println(body.NewProductList)

	doc := parse(string(body.NewProductList))

	productList := traverse(doc, "article", "class", "product-element")
	// fmt.Println(productList)

	for item := productList; item != nil; item = item.NextSibling.NextSibling {
		// fmt.Println(item)

		_product := _Response{}

		productId, _ := attr(item.Attr, "data-product-id")
		if _debug {
			fmt.Println(productId)
		}
		_product.code = productId

		// percentage := traverse(item, "div", "class", "percentage")
		// // fmt.Println(percentage)

		// discount, _ := text(percentage)
		// discount = strings.TrimSpace(discount)
		// if _debug {
		// 	fmt.Println(discount)
		// }
		// _product.discount = discount

		imageTitleLink := traverse(item, "a", "class", "link-detail")
		// fmt.Println(imageTitleLink)

		link, _ := attr(imageTitleLink.Attr, "href")
		if _debug {
			fmt.Println(link)
		}
		_product.link = link

		title, _ := attr(imageTitleLink.Attr, "title")
		if _debug {
			fmt.Println(title)
		}
		_product.title = title

		if skip := skips[strings.ToUpper(strings.ReplaceAll(strings.Split(title, " ")[0], "-", ""))]; skip {
			continue
		}

		currentPrice := traverse(item, "div", "class", "generalPrice")
		// fmt.Println(currentPrice)

		price, _ := text(currentPrice)
		if _debug {
			fmt.Println(price)
		}

		if _price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(price, ".-", ".00"), "'", ""), 32); err != nil {
			panic(err)
		} else {
			_product.oldPrice = float32(_price)
		}

		if insteadPrice := traverse(item, "div", "class", "insteadPrice"); insteadPrice != nil {
			// fmt.Println(insteadPrice)

			itemText := traverse(insteadPrice, "text", "", "")

			price, _ := text(itemText)
			price = strings.TrimSpace(strings.ReplaceAll(price, "statt", ""))
			if _debug {
				fmt.Println(price)
			}

			if _price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(price, ".-", ".00"), "'", ""), 32); err != nil {
				panic(err)
			} else {
				_product.price = float32(_price)
			}
		}

		if _debug {
			fmt.Println()
		}

		_result = append(_result, _product)
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result))
		for _, _product := range _result {
			// fmt.Println(_product)

			_retailPrice := _product.oldPrice
			_price := _retailPrice
			if _product.price > 0 {
				_price = _product.price
			}
			_savings := _price - _retailPrice
			_discount := 100 - ((100 / _retailPrice) * _price)

			product := &Product{
				Code: _name + "//" + _product.code,
				Name: _product.title,

				RetailPrice: _retailPrice,
				Price:       _price,
				Savings:     _savings,
				Discount:    _discount,

				URL: _product.link,
			}

			if s.IsWorth(product) {
				products = append(products, product)
			}
		}

		// fmt.Printf("%#v\n", products)
		return &products
	}

	return NewShop(
		_name,
		_url,

		_parseFn,
	)
}
