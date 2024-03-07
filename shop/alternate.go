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

	helpers "jsapi-scraper/helpers"
)

var AlternateRegex = regexp.MustCompile(`(\s*[-,]\s+)|(\d+\s*GB?[^T])|\s+\(SM-A\d+\)`)

var AlternateCleanFn = func(name string) string {
	name = strings.NewReplacer("Enterprise Edition", "EE").Replace(name)

	if loc := AlternateRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)
}

func XXX_alternate(isDryRun bool) IShop {
	const _name = "alternate"
	// const _url = "https://www.alternate.ch/Smartphone/Smartphone-Marken?t=18356&s=price_asc&filter_-2=true&filter_416=177&filter_1653=1"
	// const _url = "https://www.alternate.ch/Alle-Smartphones?t=18352&filter_-2=true&filter_16536=5&s=price_asc&page=%d"
	// const _url = "https://www.alternate.ch/Smartphone/Smartphone-Marken?t=18356&filter_416=177&filter_-2=true&filter_16536=5&s=price_asc&page=%d"
	// const _url = "https://www.alternate.ch/Smartphone/Smartphone-Marken?t=18356&filter_416=177&filter_16536=5&s=price_asc&page=%d"
	// _url := fmt.Sprintf("https://www.alternate.ch/Smartphone/Smartphone-Marken?t=18356&filter_416=177&filter_16536=5&s=price_asc&pr1=%.f&pr2=%.f&page=%%d", ValueMinimum, ValueMaximum)
	_url := fmt.Sprintf("https://www.alternate.ch/Alle-Smartphones?t=18352&s=price_asc&pr1=%.f&pr2=%.f&filter_-2=true&filter_16536=5&page=%%d", ValueMinimum, ValueMaximum)

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	type _Response struct {
		code  string
		title string
		model string

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

	for p := 1; p <= 20; p++ {
		fn := fmt.Sprintf("shop/alternate.%d.html", p)

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
		// fmt.Println(BytesToString(_body))

		doc := parse(BytesToString(_body))

		if productList := traverse(doc, "div", "class", "grid-container"); productList != nil {
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

				model := AlternateCleanFn(title)
				if _debug {
					fmt.Println(model)
				}
				_product.model = model

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

			results := traverse(doc, "div", "class", "col-md-auto")
			if result, ok := text(results); ok {
				if x := regexp.MustCompile(`(\d+)-(\d+) von (\d+)`).FindStringSubmatch(result); x != nil && x[2] == x[3] {
					break
				}
			}
		}
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result))
		for _, _product := range _result {
			// fmt.Println(_product)

			_title := _product.title
			_model := _product.model

			if Skip(_model) {
				continue
			}

			if _tests {
				testCases[_title] = _model
			}

			_retailPrice := _product.oldPrice
			_price := _retailPrice
			if _product.price > 0 {
				_price = _product.price
			}

			_savings := _price - _retailPrice
			_discount := 100 - ((_price * 100) / _retailPrice)

			product := &Product{
				Code:  _name + "//" + _product.code,
				Name:  _title,
				Model: _model,

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
