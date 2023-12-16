package shop

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/sugawarayuuta/sonnet"

	helpers "jsapi-scraper/helpers"
)

var AckermannRegex = regexp.MustCompile(`(?i)(,\s*)?\d+\s*GB|(,\s*)?\(?[2345]G\)?| LTE`)

var AckermannCleanFn = func(name string) string {
	// name = strings.NewReplacer("", "").Replace(name)

	if loc := AckermannRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	s := strings.Split(name, " ")

	if s[0] == "Apple" {
		if s[1] != "iPhone" {
			name = strings.ReplaceAll(name, "Apple ", "Apple iPhone ")
		}
	}

	if s[0] == "Samsung" {
		if strings.HasSuffix(name, "Xcover Pro") {
			name = strings.ReplaceAll(name, "Xcover Pro", "XCover 6 Pro")
		} else if strings.HasSuffix(name, "Xcover") {
			name = strings.ReplaceAll(name, "Xcover", "XCover 5")
		}
	}

	return helpers.Lint(name)
}

func XXX_ackermann(isDryRun bool) IShop {
	const _name = "Ackermann"
	_url := fmt.Sprintf("https://www.ackermann.ch/_next/data/shopping_app/de/technik/multimedia/smartphones-telefone.json?o=price-asc&f=%s&categories=technik&categories=multimedia&categories=smartphones-telefone", base64.StdEncoding.EncodeToString(StringToBytes(fmt.Sprintf(`{"filter_Produkttyp_1":["fb_prdkt.p1_smrt.hn_38"],"filter_price":["%.f-%.f"]}`, ValueMinimum, ValueMaximum))))

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

	type _Body struct {
		PageProps struct {
			Fallback struct {
				SearchApiResult struct {
					SearchResult struct {
						Request struct {
							Count int `json:"count"` // 72
							Start int `json:"start"` // 144
						} `json:"request"`
						Result struct {
							Count    int `json:"count"` // 162
							Products []struct {
								Brand struct {
									Image string `json:"image"` // https://bilder.ackermann.ch/marken/ackermannch/samsung.gif
									Name  string `json:"name"`  // Samsung
								} `json:"brand"`
								MasterSku   string `json:"masterSku"`   // AKLBB1660796293
								Name        string `json:"name"`        // Samsung Smartphone »Samsung Galaxy A13«, light blue, 16,72 cm/6,6 Zoll, 128 GB...
								NameNoBrand string `json:"nameNoBrand"` // Smartphone »Samsung Galaxy A13«, light blue, 16,72 cm/6,6 Zoll, 128 GB Speicherplatz,...
								Variations  []struct {
									VariationName string `json:"variationName"` // light blue
									OldPrice      struct {
										Currency string  `json:"currency"` // CHF
										Value    float32 `json:"value"`
										UVP      bool    `json:"uvp"`
									} `json:"oldPrice"`
									Price struct {
										Currency   string  `json:"currency"` // CHF
										Value      float32 `json:"value"`
										Saving     int     `json:"saving"`
										SavingType string  `json:"savingType"` // CURRENCY
									} `json:"price"`
									ProductUrl string `json:"productUrl"` // /p/samsung-smartphone-samsung-galaxy-a13/AKLBB1660796293?sku=9682996614&nav-c=134922#nav-i=12&ref=mba
									ArtNo      string `json:"artNo"`      // 96829966
									Sku        string `json:"sku"`        // 9682996614-0-1660796293
								} `json:"variations"`
							} `json:"products"`
						} `json:"result"`
					} `json:"searchresult"`
				} `json:"search-api-result"`
			} `json:"fallback"`
		} `json:"pageProps"`
	}

	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	var _results []_Response

	for p := 1; p <= 5; p++ {
		fn := fmt.Sprintf("shop/ackermann.%d.json", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			page := ""
			if p > 1 {
				page = fmt.Sprintf("&p=%d", p)
			}

			url := fmt.Sprintf("%s%s", _url, page)

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

		var body _Body
		if err := sonnet.Unmarshal(_body, &body); err != nil { // Parse []byte to go struct pointer
			panic(err)
		}

		reName := regexp.MustCompile(`»(.*?)«`)

		for _, product := range body.PageProps.Fallback.SearchApiResult.SearchResult.Result.Products {
			name := product.Name
			brand := product.Brand.Name

			if matches := reName.FindStringSubmatch(name); len(matches) > 1 {
				title := matches[1]

				s := strings.Split(title, " ")
				if s[0] != brand {
					title = brand + " " + title
				}

				// fmt.Println(title)
				model := AckermannCleanFn(title)
				// fmt.Println(model)
				// fmt.Println()

				// for _, variation := range product.Variations {
				if variation := product.Variations[0]; len(product.Variations) > 0 {
					result := _Response{
						code:  variation.Sku,
						title: title,
						model: model,

						link: variation.ProductUrl,

						oldPrice: variation.OldPrice.Value,
						price:    variation.Price.Value,
					}
					// fmt.Println(result)

					_results = append(_results, result)
				}
			}
		}

		if body.PageProps.Fallback.SearchApiResult.SearchResult.Request.Start+body.PageProps.Fallback.SearchApiResult.SearchResult.Request.Count >= body.PageProps.Fallback.SearchApiResult.SearchResult.Result.Count {
			break
		}
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_results))
		for _, _product := range _results {
			// fmt.Println(_product)

			_title := _product.title
			_model := _product.model

			if Skip(_model) {
				if _debug {
					fmt.Println()
				}

				continue
			}

			if _tests {
				testCases[_title] = _model
			}

			var _savings float32
			var _discount float32

			_retailPrice := max(_product.oldPrice, _product.price)
			_price := _retailPrice
			if _product.price > 0 {
				_price = _product.price
			}
			if _debug {
				fmt.Println(_retailPrice)
				fmt.Println(_price)
			}

			if _price > 0 {
				_savings = _price - _retailPrice
				_discount = 100 - ((100 / _retailPrice) * _price)
			}
			if _debug {
				fmt.Println(_savings)
				fmt.Println(_discount)
			}

			_link := s.ResolveURL(_product.link).String()
			if _debug {
				fmt.Println(_link)
				fmt.Println()
			}

			product := &Product{
				Code:  _name + "//" + _product.code,
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
