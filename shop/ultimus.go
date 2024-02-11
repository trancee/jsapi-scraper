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

var UltimusRegex = regexp.MustCompile(`(?i)Refurbished|(Outdoor-|Robustes |Rugged )?Smartphone|(2|4|6|8)GB|[345]G|\s+((EE )?Enterprise Edition( CH)?)`)
var UltimusExclusionRegex = regexp.MustCompile(`(?i)Adapter|Armband|Ch?inch|Christbaum|Etui|Halterung|Halter|Kfz|Kopfhörer|Ladegerät|Ladestation|Netzkabel|Objektiv|Robustes Smartphone$|Saugnapf|Schutzfolie|Smartphone mit 100MP Kamera|Stativ|Virtual-Reality|Wasserdicht(es)?|Weihnachtsbaum|Windschutzscheibe`)

var UltimusCleanFn = func(name string) string {
	name = strings.ReplaceAll(name, "Xioami", "Xiaomi")
	name = strings.ReplaceAll(name, "Robustes Smartphone ", "")

	if loc := UltimusRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	s := strings.Split(name, " ")

	if s[0] == "Nokia" {
		name = strings.ReplaceAll(name, "C2-2E", "C2 2nd Edition")
	}

	return helpers.Lint(name)
}

func XXX_ultimus(isDryRun bool) IShop {
	const _name = "Ultimus"
	_url := fmt.Sprintf("https://ultimus.ch/product-category/elektronik/handy-watch-und-tablet/smartphones/page/%%d/?orderby=price&price_filter_min=%.f&price_filter_max=%.f", ValueMinimum, ValueMaximum)

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

	for p := 1; p <= 10; p++ {
		fn := fmt.Sprintf("shop/ultimus.%d.html", p)

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

		if productList := traverse(doc, "ul", "class", "products"); productList != nil {
			// fmt.Println(productList)

			for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling.NextSibling {
				// fmt.Println(item)

				_product := _Response{}

				code, _ := attr(item.Attr, "data-product-id")
				if _debug {
					fmt.Println(code)
				}
				_product.code = code

				itemLink := traverse(item, "a", "class", "woocommerce-loop-product__link")
				// fmt.Println(itemLink)

				link, _ := attr(itemLink.Attr, "href")
				if _debug {
					fmt.Println(link)
				}
				_product.link = link

				itemTitle := traverse(item, "h2", "class", "woocommerce-loop-product__title")
				// fmt.Println(itemTitle)

				title, _ := text(itemTitle)
				if _debug {
					fmt.Println(title)
				}
				_product.title = title

				if Skip(title) {
					continue
				}

				if UltimusExclusionRegex.MatchString(title) {
					continue
				}

				model := UltimusCleanFn(title)
				if _debug {
					fmt.Println(model)
				}
				_product.model = model

				itemPrice := traverse(item, "span", "class", "price")
				// fmt.Println(itemPrice)

				// fmt.Println(itemPrice.FirstChild)

				if itemPrice.FirstChild.Data == "del" {
					itemPrice = itemPrice.FirstChild
				}

				if itemPrice.FirstChild.FirstChild.Data == "bdi" {
					price, _ := text(itemPrice.FirstChild.FirstChild.FirstChild.NextSibling)
					if _debug {
						fmt.Println(price)
					}

					if _price, err := strconv.ParseFloat(strings.NewReplacer(",", "", ".-", ".00", "'", "").Replace(price), 32); err != nil {
						panic(err)
					} else {
						_product.oldPrice = float32(_price)
					}
				}

				// fmt.Printf("%+v\n", itemPrice.Parent.FirstChild.NextSibling.NextSibling)

				if itemPrice.Parent.FirstChild.NextSibling.NextSibling.Data == "ins" {
					itemPrice = itemPrice.Parent.FirstChild.NextSibling.NextSibling

					if itemPrice.FirstChild.FirstChild.Data == "bdi" {
						price, _ := text(itemPrice.FirstChild.FirstChild.FirstChild.NextSibling)
						if _debug {
							fmt.Println(price)
						}

						if _price, err := strconv.ParseFloat(strings.NewReplacer(",", "", ".-", ".00", "'", "").Replace(price), 32); err != nil {
							panic(err)
						} else {
							_product.price = float32(_price)
						}
					}
				}

				if _debug {
					fmt.Println()
				}

				_result = append(_result, _product)
			}

			if pagination := traverse(doc, "nav", "class", "ultimus_pagination"); pagination != nil {
				results := pagination.FirstChild.NextSibling.NextSibling
				if result, ok := text(results); ok {
					if x := regexp.MustCompile(`Seite (\d+) von (\d+)`).FindStringSubmatch(result); x != nil && x[1] == x[2] {
						break
					}
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
		strings.ReplaceAll(_url, "/page/%d", ""),

		_parseFn,
	)
}
