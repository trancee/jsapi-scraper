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

var MediamarktRegex = regexp.MustCompile(` - |\s+[2345]G|\s+((EE )?Enterprise Edition( CH)?)`)

var MediamarktCleanFn = func(name string) string {
	name = strings.NewReplacer("ONE PLUS", "ONEPLUS").Replace(name)

	if loc := MediamarktRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)

	// s := strings.Split(name, " ")

	// if s[0] == "Apple" {
	// 	name = strings.NewReplacer(" 2020", " (2020)", " 2022", " (2022)", " 2nd Gen", " (2020)", " 3rd Gen", " (2022)").Replace(name)
	// } else {
	// 	// Remove year component for all other than Apple.
	// 	name = regexp.MustCompile(`\s+\(?20[12]\d\)?`).ReplaceAllString(name, "")
	// }

	// return strings.TrimSpace(name)
}

func XXX_mediamarkt(isDryRun bool) IShop {
	const _name = "Mediamarkt"
	// const _url = "https://www.mediamarkt.ch/de/category/_smartphone-680815.html?searchParams=&sort=price&view=PRODUCTGRID"
	// const _url = "https://www.mediamarkt.ch/de/category/_smartphone-680815.html?searchParams=%2FSearch.ff%3FfilterCategoriesROOT%3DHandy%2B%2526%2BNavigation%25C2%25A7MediaCHdec680760%26filterCategoriesROOT%252FHandy%2B%2526%2BNavigation%25C2%25A7MediaCHdec680760%3DSmartphone%25C2%25A7MediaCHdec680815%26filteravailability%3D1%26filterTyp%3D___Smartphone%26channel%3Dmmchde%26followSearch%3D9782%26disableTabbedCategory%3Dtrue%26navigation%3Dtrue&sort=price&view=PRODUCTGRID&page=%d"
	// const _url = "https://www.mediamarkt.ch/de/category/_smartphone-680815.html?searchParams=/Search.ff?filterTabbedCategory%3Donlineshop%26filterCategoriesROOT%3DHandy%2B%2526%2BNavigation%25C2%25A7MediaCHdec680760%26filterCategoriesROOT%252FHandy%2B%2526%2BNavigation%25C2%25A7MediaCHdec680760%3DSmartphone%25C2%25A7MediaCHdec680815%26filteravailability%3D1%26filterTyp%3D___Smartphone%26channel%3Dmmchde%26followSearch%3D9809%26disableTabbedCategory%3Dtrue%26navigation%3Dtrue&sort=price&page="
	_url := fmt.Sprintf("https://www.mediamarkt.ch/de/category/_smartphone-680815.html?searchParams=%%2FSearch.ff%%3FfilterTabbedCategory%%3Donlineshop%%26filterCategoriesROOT%%3DHandy%%2B%%2526%%2BNavigation%%25C2%%25A7MediaCHdec680760%%26filterCategoriesROOT%%252FHandy%%2B%%2526%%2BNavigation%%25C2%%25A7MediaCHdec680760%%3DSmartphone%%25C2%%25A7MediaCHdec680815%%26filteravailability%%3D1%%26filterTyp%%3D___Smartphone%%26filtercurrentprice%%3D%.f%%2B-%%2B%.f%%26channel%%3Dmmchde%%26followSearch%%3D9650%%26disableTabbedCategory%%3Dtrue%%26navigation%%3Dtrue&sort=price&page=", ValueMinimum, ValueMaximum)

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

	for p := 1; p <= 5; p++ {
		fn := fmt.Sprintf("shop/mediamarkt.%d.html", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			url := fmt.Sprintf("%s%d", _url, p)

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

		doc := parse(string(_body))

		if productList := traverse(doc, "ul", "class", "products-grid"); productList != nil {
			// fmt.Println(productList)

			for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling.NextSibling {
				// fmt.Println(item)

				_product := _Response{}

				baseInfo := traverse(item, "div", "class", "base-info")
				if baseInfo == nil {
					continue
				}
				// fmt.Println(baseInfo)

				productKey, _ := attr(baseInfo.Attr, "data-reco-pid")
				productId := productKey[2:]
				if _debug {
					fmt.Println(productId)
				}
				_product.code = productId

				imageTitleLink := traverse(baseInfo, "a", "class", "product-link")
				// fmt.Println(imageTitleLink)

				link, _ := attr(imageTitleLink.Attr, "href")
				if _debug {
					fmt.Println(link)
				}
				_product.link = link

				title, _ := text(imageTitleLink)
				// title = strings.TrimSpace(strings.Split(strings.ReplaceAll(strings.ReplaceAll(title, " - Smartphone", ""), " \"", "\""), "(")[0])
				if _debug {
					fmt.Println(title)
				}
				_product.title = title

				if Skip(title) {
					continue
				}

				model := MediamarktCleanFn(_product.title)
				if _debug {
					fmt.Println(model)
				}
				_product.model = model

				if oldPrice := traverse(baseInfo, "div", "class", "price-old"); oldPrice != nil {
					// fmt.Println(oldPrice.FirstChild.Parent.LastChild)

					price, _ := text(oldPrice.FirstChild.Parent.LastChild)
					if _debug {
						fmt.Println(price)
					}

					if _price, err := strconv.ParseFloat(strings.ReplaceAll(price, ".-", ".00"), 32); err != nil {
						panic(err)
					} else {
						_product.oldPrice = float32(_price)
					}

					{
						currentPrice := oldPrice.Parent.NextSibling.NextSibling
						// fmt.Println(currentPrice)

						price, _ := text(currentPrice)
						if _debug {
							fmt.Println(price)
						}

						if _price, err := strconv.ParseFloat(strings.ReplaceAll(price, ".-", ".00"), 32); err != nil {
							panic(err)
						} else {
							_product.price = float32(_price)
						}
					}
				} else if currentPrice := traverse(baseInfo, "div", "class", "price"); currentPrice != nil {
					// fmt.Println(currentPrice)

					price, _ := text(currentPrice)
					if _debug {
						fmt.Println(price)
					}

					if _price, err := strconv.ParseFloat(strings.ReplaceAll(price, ".-", ".00"), 32); err != nil {
						panic(err)
					} else {
						_product.oldPrice = float32(_price)
					}
				} else {
					continue
				}

				if _debug {
					fmt.Println()
				}

				_result = append(_result, _product)
			}

			if results := traverse(doc, "li", "class", "pagination-next"); results != nil {
				if _, exists := attr(results.Attr, "data-value"); !exists {
					break
				}
			} else {
				break
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

			if _tests {
				testCases[_title] = _model
			}

			_retailPrice := product.oldPrice
			_price := _retailPrice
			if product.price > 0 {
				_price = product.price
			}
			_savings := _price - _retailPrice
			_discount := 100 - ((100 / _retailPrice) * _price)

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
