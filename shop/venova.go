package shop

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/sugawarayuuta/sonnet"

	helpers "jsapi-scraper/helpers"
)

var VenovaRegex = regexp.MustCompile(`(?i)[,-]? ?(6|8|16|32|64|128|256) ?([MG]B|BG)| \d"| [45]G|\d+(,\d+)? cm|(EE )?Enterprise( Edition)?( CH)?`)

var VenovaCleanFn = func(name string) string {
	if loc := VenovaRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	name = regexp.MustCompile(`(?i)Dual[- ]SIM`).ReplaceAllString(name, "")
	name = strings.NewReplacer(" - ", " ").Replace(name)
	name = strings.TrimSpace(name)

	s := strings.Split(name, " ")

	if s[0] == "Samsung" {
		name = regexp.MustCompile(`\s+(SM-)?[AFMS]\d{3}[A-Za-z]*`).ReplaceAllString(name, "")
	}

	if s[0] == "Xiaomi" {
		name = regexp.MustCompile(`\s+\d+-\d+-\d+`).ReplaceAllString(name, "")
	}

	return helpers.Lint(name)
}

func XXX_venova(isDryRun bool) IShop {
	const PageCount = 12

	const _name = "Venova"
	_url := fmt.Sprintf("https://www.venova.ch/de/widgets/listing/listingCount/sCategory/5760?p=%%d&o=3&n=%d&min=%.f&max=%.f&loadProducts=1", PageCount, ValueMinimum, ValueMaximum)

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	type _Product struct {
		code  string
		title string
		model string

		link string

		oldPrice float32
		price    float32

		// quantity int
	}

	type _Response struct {
		TotalCount int    `json:"totalCount"`
		Pagination string `json:"pagination"`

		Listing string `json:"listing"`
	}

	var _result _Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	var _products []_Product
	_count := 0

	for p := 1; p <= 10; p++ {
		fn := fmt.Sprintf("shop/venova.%d.json", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			url := fmt.Sprintf(_url, p)

			resp, err := http.Get(url)
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, url)
				return NewShop(
					_name,
					_url,

					nil,
				)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				// panic(resp.StatusCode)
				fmt.Printf("[%s] %d: %s (%s)\n", _name, resp.StatusCode, resp.Status, resp.Request.URL)
				return NewShop(
					_name,
					_url,

					nil,
				)
			}

			if body, err := io.ReadAll(resp.Body); err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, resp.Request.URL)
				return NewShop(
					_name,
					_url,

					nil,
				)
			} else {
				_body = body
			}

			os.WriteFile(path+fn, _body, 0664)
		}
		// fmt.Println(string(_body))

		if err := sonnet.Unmarshal(_body, &_result); err != nil {
			panic(err)
		}

		doc := parse(string(_result.Listing))

		for item := traverse(doc, "div", "class", "product--box"); item != nil; item = item.NextSibling.NextSibling {
			// fmt.Println(item)

			// if itemAvailability := traverse(item, "link", "itemprop", "availability"); itemAvailability != nil {
			// 	// fmt.Println(itemAvailability)
			// 	availability, _ := attr(itemAvailability.Attr, "href")

			// 	if !strings.HasSuffix(availability, "InStock") {
			// 		continue
			// 	}
			// }

			_product := _Product{}

			imageLink := traverse(item, "a", "class", "product--image")
			// fmt.Println(imageLink)

			link, _ := attr(imageLink.Attr, "href")
			if _debug {
				fmt.Println(link)
			}
			_product.link = link

			title, _ := attr(imageLink.Attr, "data-product-name")
			if _debug {
				fmt.Println(title)
			}
			_product.title = title

			if Skip(title) {
				continue
			}

			model := VenovaCleanFn(_product.title)
			// fmt.Printf("%s\n%s\n\n", title, model)
			if _debug {
				fmt.Println(model)
			}
			_product.model = model

			if _tests {
				testCases[_product.title] = _product.model
			}

			sku, _ := attr(imageLink.Attr, "data-product-ordernumber")
			if _debug {
				fmt.Println(sku)
			}
			_product.code = sku

			price, _ := attr(imageLink.Attr, "data-product-price")
			if _debug {
				fmt.Println(price)
			}

			if _price, err := strconv.ParseFloat(price, 32); err != nil {
				panic(err)
			} else {
				_product.price = float32(_price)
			}

			if priceDiscount := traverse(item, "span", "class", "price--discount"); priceDiscount != nil {
				oldPrice, _ := text(priceDiscount)
				if _debug {
					fmt.Println(oldPrice)
				}

				oldPrice = strings.ReplaceAll(strings.Split(oldPrice, "\u00a0")[0], ",", ".")

				if _price, err := strconv.ParseFloat(oldPrice, 32); err != nil {
					panic(err)
				} else {
					_product.oldPrice = float32(_price)
				}
			}

			if priceDefault := traverse(item, "span", "class", "price--default"); priceDefault != nil {
				price, _ := text(priceDefault)
				if _debug {
					fmt.Println(price)
				}

				price = strings.ReplaceAll(strings.Split(price, "\u00a0")[0], ",", ".")

				if _price, err := strconv.ParseFloat(price, 32); err != nil {
					panic(err)
				} else {
					_product.price = float32(_price)
				}
			}

			if _debug {
				fmt.Println()
			}

			_products = append(_products, _product)
		}

		_count += PageCount

		if _count >= _result.TotalCount {
			break
		}
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_products))
		for _, product := range _products {
			// fmt.Println(product)

			_title := product.title
			_model := product.model

			if Skip(_model) {
				continue
			}

			_retailPrice := product.price
			_price := _retailPrice
			if product.oldPrice > 0 {
				_retailPrice = product.oldPrice
			}
			_savings := _price - _retailPrice
			_discount := 100 - ((_price * 100) / _retailPrice)

			_link := s.ResolveURL(product.link).String()

			product := &Product{
				Code:  _name + "//" + product.code,
				Name:  _title,
				Model: _model,

				RetailPrice: _retailPrice,
				Price:       _price,
				Savings:     _savings,
				Discount:    _discount,

				// Quantity: product.quantity,

				URL: _link,
			}

			if s.IsWorth(product) {
				products = append(products, product)
			}

			// fmt.Printf("%#v\n", product)
		}

		if _tests {
			keys := make([]string, 0, len(testCases))

			for k := range testCases {
				keys = append(keys, k)
			}
			sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

			for _, k := range keys {
				fmt.Println("\"" + strings.ReplaceAll(k, "\"", "\\\"") + "\",")
			}
			fmt.Println()
			for _, k := range keys {
				fmt.Println("\"" + strings.ReplaceAll(testCases[k], "\"", "\\\"") + "\",")
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
