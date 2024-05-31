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

var MiStoreV2Regex = regexp.MustCompile(`\s+(\d+/(GB)?)?\d+GB|\s+20[12]\d|\s+[2345]G| \| `)
var MiStoreV2ExclusionRegex = regexp.MustCompile(`(?i)Abdeckung|Adapter|AirTag|Armband|Band|CABLE|Charger|Ch?inch|Christbaum|Clamshell|^Core|\bCover\b|Earphones|Etui|Fernauslöser|Gimbal|Halterung|Handschuhe|HARDCASE|Headset|Hülle|Kopfhörer|Ladegerät|Ladestation|Lautsprecher|Magnet|Majestic|Netzkabel|Objektiv|Reiselader|S Pen|Saugnapf|Schutzfolie|Schutzglas|SmartTag|Stand|Ständer|Stativ|Stick|Stylus|Tastatur|Virtual-Reality|Wasserdicht(es)?|Weihnachtsbaum`)

var MiStoreV2CleanFn = func(name string) string {
	// name = strings.NewReplacer("*** BUNDLE ***", "", "A53 s", "A53s").Replace(name)
	name = regexp.MustCompile(`(?i)^\s*\*{3}\s*BUNDLE\s*\*{3}`).ReplaceAllString(name, "")

	if loc := MiStoreV2Regex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)
}

func XXX_mistore_v2(isDryRun bool) IShop {
	const _name = "Mi Store"
	_url := fmt.Sprintf("https://mi-store.ch/c/phones?manufacturer=cf97e8a2c61349e3ba8f4ed2d944b0cf&min-price=%.f&max-price=%.f&order=price-asc&p=%%d", ValueMinimum, ValueMaximum)

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
		fn := fmt.Sprintf("shop/mi-store.%d.html", p)

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

		if productList := traverse(doc, "div", "class", "cms-listing-row"); productList != nil {
			// fmt.Printf("%+v\n", productList)

			for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling.NextSibling {
				// fmt.Printf("%+v\n", item)

				if contains(item.Attr, "class", "outofstock") {
					continue
				}

				_product := _Response{}

				productInfo := traverse(item, "div", "class", "product-info")
				// fmt.Printf("%+v\n", productInfo)

				productLink := traverse(productInfo, "a", "class", "product-name")

				link, _ := attr(productLink.Attr, "href")
				if _debug {
					fmt.Println(link)
				}
				_product.link = link

				productID := traverse(productInfo, "input", "name", "product-id")

				productId, _ := attr(productID.Attr, "value")
				if _debug {
					fmt.Println(productId)
				}
				_product.code = productId

				productName := traverse(productInfo, "input", "name", "product-name")

				title, _ := attr(productName.Attr, "value")
				if _debug {
					fmt.Println(title)
				}
				_product.title = title

				if Skip(title) {
					if _debug {
						fmt.Println()
					}

					continue
				}

				if MiStoreV2ExclusionRegex.MatchString(title) {
					if _debug {
						fmt.Println()
					}

					continue
				}

				model := MiStoreV2CleanFn(_product.title)
				if _debug {
					fmt.Println(model)
				}
				_product.model = model

				if _tests {
					testCases[_product.title] = _product.model
				}

				productPriceInfo := traverse(item, "div", "class", "product-price-info")
				// fmt.Printf("%+v\n", productPriceInfo)

				productPrice := traverse(productPriceInfo, "span", "class", "product-price")
				// fmt.Printf("%+v\n", productPrice)

				price, _ := text(productPrice)
				if _debug {
					fmt.Println(price)
				}

				if _price, err := strconv.ParseFloat(strings.NewReplacer("CHF\u00a0", "", "*", "", "'", "").Replace(price), 32); err != nil {
					panic(err)
				} else {
					_product.oldPrice = float32(_price)
				}

				// if ins := traverse(productPrice, "ins", "", ""); ins != nil {
				// 	productAmount := traverse(ins, "span", "class", "amount")
				// 	// fmt.Println(productAmount)

				// 	price, _ := text(productAmount.FirstChild.NextSibling)
				// 	if _debug {
				// 		fmt.Println(price)
				// 	}

				// 	if _price, err := strconv.ParseFloat(strings.ReplaceAll(price, "'", ""), 32); err != nil {
				// 		panic(err)
				// 	} else {
				// 		_product.price = float32(_price)
				// 	}
				// }

				if _debug {
					fmt.Println()
				}

				_result = append(_result, _product)

				if item.NextSibling == nil {
					break
				}
			}

			if pagination := traverse(doc, "nav", "class", "pagination-nav"); pagination != nil {
				if pageNext := traverse(pagination, "li", "class", "page-next"); pageNext != nil {
					// fmt.Printf("%+v\n", contains(pageNext.Attr, "class", "disabled"))
					if contains(pageNext.Attr, "class", "disabled") {
						break
					}
				}
			}
		}
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result))
		for _, product := range _result {
			// fmt.Println(product)

			_title := product.title
			_model := product.model

			if Skip(_model) {
				continue
			}

			_retailPrice := product.oldPrice
			_price := _retailPrice
			if product.price > 0 {
				_price = product.price
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

				URL: _link,
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
