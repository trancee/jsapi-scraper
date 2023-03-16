package shop

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func XXX_alternate(isDryRun bool) IShop {
	const _name = "alternate"
	// const _url = "https://www.alternate.ch/Smartphone/Smartphone-Marken?t=18356&s=price_asc&filter_-2=true&filter_416=177&filter_1653=1"
	// const _url = "https://www.alternate.ch/Alle-Smartphones?t=18352&filter_-2=true&filter_16536=5&s=price_asc&page=%d"
	const _url = "https://www.alternate.ch/Smartphone/Smartphone-Marken?t=18356&filter_416=177&filter_-2=true&filter_16536=5&s=price_asc&page=%d"

	const _debug = false

	type _Response struct {
		code  string
		title string

		link string

		oldPrice float32
		price    float32
	}

	var _result []_Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	for p := 1; p <= 5; p++ {
		fn := fmt.Sprintf("shop/alternate.%d.html", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			resp, err := http.Get(fmt.Sprintf(_url, p))
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

		doc := parse(string(_body))

		productList := traverse(doc, "div", "class", "grid-container")
		// fmt.Println(productList)

		for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling.NextSibling {
			// fmt.Println(item)

			_product := _Response{}

			link, _ := attr(item.Attr, "href")
			if _debug {
				fmt.Println(link)
			}
			_product.link = link

			_parts := strings.Split(link, "/")
			code := _parts[len(_parts)-1]
			if _debug {
				fmt.Println(code)
			}
			_product.code = code

			productPicture := traverse(item, "img", "class", "productPicture")
			// fmt.Println(productPicture)

			title, _ := attr(productPicture.Attr, "alt")
			title = strings.Split(strings.ReplaceAll(title, ", Handy", ""), ",")[0]
			if brand := strings.Split(title, " "); brand[0] == "realme" {
				title = strings.ReplaceAll(title, "-", "")
			}
			if _debug {
				fmt.Println(title)
			}
			_product.title = title

			if Skip(title) {
				continue
			}

			currentPrice := traverse(item, "span", "class", "price")
			// fmt.Println(currentPrice)

			price, _ := text(currentPrice)
			price = strings.ReplaceAll(strings.TrimSpace(strings.TrimPrefix(price, "CHF")), ",", ".")
			if _debug {
				fmt.Println(price)
			}

			if _price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(price, ".-", ".00"), "'", ""), 32); err != nil {
				panic(err)
			} else {
				_product.oldPrice = float32(_price)
			}

			if _debug {
				fmt.Println()
			}

			_result = append(_result, _product)
		}
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
