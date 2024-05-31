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

var MediamarktV2Regex = regexp.MustCompile(` - |(64|128)\s*GB|\s+[2345]G|\s+CH$`)

var MediamarktV2CleanFn = func(name string) string {
	name = strings.NewReplacer("ONE PLUS", "ONEPLUS", "Enterprise Edition", "EE").Replace(name)

	if loc := MediamarktV2Regex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)
}

var MediamarktV2CodeRegex = regexp.MustCompile(`-(\d+)\.html$`)

func XXX_mediamarkt_v2(isDryRun bool) IShop {
	const _name = "Mediamarkt"
	// _url := "https://www.mediamarkt.ch/api/v1/graphql"
	_url := fmt.Sprintf("https://www.mediamarkt.ch/de/category/smartphone-680815.html?filter=currentprice:%.f-%.f&sort=currentprice+asc&page=", ValueMinimum, ValueMaximum)

	// variables := `{"isDemonstrationModelAvailabilityActive":false,"isMultilingual":true,"hasMarketplace":false,"isArtificialScarcityActive":true,"isRefurbishedGoodsActive":false,"locale":"de-CH","salesLine":"Media","filters":["currentprice:69.0-500.0"],"sort":"currentprice+asc","page":1,"pimCode":"CAT_CH_MM_680815","criteoInputArgs":{"adEnvironment":"desktop","adCustomerId":"7d9c67e9-11ab-44a1-9925-25469c737332","adRetailerVisitorId":"7d9c67e9-11ab-44a1-9925-25469c737332","adOutletId":"1132","adGdpr":"1","adPositionVariantB":false}}`
	// extensions := `{"persistedQuery":{"version":1,"sha256Hash":"bca574024cc9c170a4ce8a7f3ad882cf4502c2b89a70fcb789ccd80fa518e290"},"pwa":{"salesLine":"Media","country":"CH","language":"de","globalLoyaltyProgram":true,"isOneAccountProgramActive":true,"isMdpActive":true,"isUsingXccCustomerComponent":true}}`

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

		discount float32
	}

	var _result []_Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	for p := 1; p <= 10; p++ {
		fn := fmt.Sprintf("shop/mediamarkt.%d.html", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			// url := _url

			// req, err := http.NewRequest("GET", url, nil)
			// if err != nil {
			// 	// panic(err)
			// 	fmt.Printf("[%s] %s (%s)\n", _name, err, url)
			// 	return NewShop(
			// 		_name,
			// 		_url,

			// 		nil,
			// 	)
			// }

			// req.Header.Set("Content-Type", "application/json")
			// req.Header.Set("Pragma", "no-cache")

			// req.Header.Set("Accept", "*/*")
			// req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
			// req.Header.Set("Accept-Language", "en-US,en;q=0.9,de-CH;q=0.8,de-DE;q=0.7,de;q=0.6")

			// req.Header.Set("Referer", "https://www.mediamarkt.ch/de/category/smartphone-680815.html?filter=currentprice:69.0-500.0&sort=currentprice+asc")
			// req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

			// req.Header.Set("Apollographql-Client-Name", "pwa-client")
			// req.Header.Set("Apollographql-Client-Version", "8.67.3")

			// req.Header.Set("X-Cacheable", "true")
			// req.Header.Set("X-Flow-Id", "1ca75873-988f-4b6d-a31f-254026439988")
			// req.Header.Set("X-Mms-Country", "CH")
			// req.Header.Set("X-Mms-Language", "de")
			// req.Header.Set("X-Mms-Salesline", "Media")
			// req.Header.Set("X-Operation", "CategoryV4")

			// q := req.URL.Query()
			// q.Add("operationName", "CategoryV4")
			// q.Add("variables", variables)
			// q.Add("extensions", extensions)
			// req.URL.RawQuery = q.Encode()
			// // fmt.Printf("%+v\n", req)

			// client := &http.Client{}
			// resp, err := client.Do(req)
			// if err != nil {
			// 	// panic(err)
			// 	fmt.Printf("[%s] %d: %s (%s)\n", _name, resp.StatusCode, resp.Status, resp.Request.URL)
			// 	return NewShop(
			// 		_name,
			// 		_url,

			// 		nil,
			// 	)
			// }
			// defer resp.Body.Close()
			// // fmt.Printf("%+v\n", resp)

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
		// fmt.Println(BytesToString(_body))

		doc := parse(BytesToString(_body))

		if productList := traverse(doc, "div", "data-test", "mms-search-srp-productlist"); productList != nil {
			// fmt.Println(productList)

			for item := traverse(productList, "div", "data-test", "mms-product-card"); item != nil; item = traverse(item.Parent.NextSibling, "div", "data-test", "mms-product-card") {
				// fmt.Println(item)

				_product := _Response{}

				productLink := traverse(item, "a", "data-test", "mms-product-list-item-link")
				// fmt.Println(productLink)

				link, _ := attr(productLink.Attr, "href")
				if _debug {
					fmt.Println(link)
				}
				_product.link = link

				code := MediamarktV2CodeRegex.FindStringSubmatch(link)[1]
				if _debug {
					fmt.Println(code)
				}
				_product.code = code

				productTitle := traverse(productLink, "p", "data-test", "product-title")
				// fmt.Println(productTitle)

				title, _ := text(productTitle)
				if _debug {
					fmt.Println(title)
				}
				_product.title = title

				if Skip(title) {
					if _debug {
						fmt.Println()
					}

					if item.Parent.NextSibling == nil {
						break
					}

					continue
				}

				model := MediamarktV2CleanFn(_product.title)
				if _debug {
					fmt.Println(model)
				}
				_product.model = model

				productPrice := traverse(item, "div", "data-test", "product-price")
				// fmt.Println(productPrice)

				if presentation := traverse(productPrice, "div", "role", "presentation"); presentation != nil {
					fmt.Println(presentation)

					discount, _ := text(presentation)
					if _debug {
						fmt.Println(discount)
					}

					if _discount, err := strconv.ParseFloat(discount, 32); err != nil {
						panic(err)
					} else {
						_product.discount = float32(_discount)
					}
				}

				if oldPrice := traverse(productPrice, "span", "class", "hdwkym"); oldPrice != nil {
					// fmt.Println(oldPrice)

					price, _ := text(oldPrice)
					if _debug {
						fmt.Println(price)
					}

					price = strings.TrimSpace(strings.NewReplacer("CHF", "", ".–", ".00", "'", "").Replace(price))

					if _price, err := strconv.ParseFloat(price, 32); err != nil {
						panic(err)
					} else {
						_product.oldPrice = float32(_price)
					}
				}

				if price := traverse(productPrice, "span", "class", "fnhqEi"); price != nil {
					// fmt.Println(price)

					price, _ := text(price)
					if _debug {
						fmt.Println(price)
					}

					price = strings.TrimSpace(strings.NewReplacer("CHF", "", ".–", ".00", "'", "").Replace(price))

					if _price, err := strconv.ParseFloat(price, 32); err != nil {
						panic(err)
					} else {
						_product.price = float32(_price)
					}
				}

				if _debug {
					fmt.Println()
				}

				_result = append(_result, _product)

				if item.Parent.NextSibling == nil {
					break
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

			if _tests {
				testCases[_title] = _model
			}

			_retailPrice := product.oldPrice
			_price := _retailPrice
			if product.price > 0 {
				_price = product.price
			}
			_savings := _price - _retailPrice
			// _discount := 100 - ((_price * 100) / _retailPrice)
			_discount := product.discount

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
