package shop

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/sugawarayuuta/sonnet"

	"golang.org/x/net/html"

	helpers "jsapi-scraper/helpers"
)

var MobileDeviceV2Regex = regexp.MustCompile(`\s+\(?\d+\s*GB?|\s+\(?\d+(\.\d+)?"|\s+\(?[2345]G\)?| Dual Sim| LTE`)

var MobileDeviceV2CleanFn = func(name string) string {
	name = regexp.MustCompile(` (SM-)?[AGMS]\d{3}[A-Z]*(/DSN)?| XT\d{4}-\d+| Master Edition`).ReplaceAllString(name, "")
	name = strings.NewReplacer("Nothing Phone 1", "Nothing Phone (1)", "X Cover", "XCover").Replace(name)

	if loc := MobileDeviceV2Regex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)
}

func XXX_mobiledevice_v2(isDryRun bool) IShop {
	const _name = "mobiledevice"
	_url := fmt.Sprintf("https://www.mobiledevice.ch/de/28-mobiltelefone?order=product.price.asc&q=Preis-CHF-%.f-%.f&page=%%d&from-xhr", ValueMinimum, ValueMaximum)

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	type _Product struct {
		ID string `json:"id_product"`

		Name string `json:"name"`

		PriceAmount        float32 `json:"price_amount"`
		RegularPriceAmount float32 `json:"regular_price_amount"`

		Link string `json:"link"`

		HasDiscount bool `json:"has_discount"`
	}

	type _Body struct {
		Products []_Product `json:"products"`

		Pagination struct {
			CurrentPage int `json:"current_page"`
			PagesCount  int `json:"pages_count"`
			TotalItems  int `json:"total_items"`
		} `json:"pagination"`
	}

	type _Body2 struct {
		Products map[int]_Product `json:"products"`

		Pagination struct {
			CurrentPage int `json:"current_page"`
			PagesCount  int `json:"pages_count"`
			TotalItems  int `json:"total_items"`
		} `json:"pagination"`
	}

	var _result []_Product
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	for p := 1; p <= 10; p++ {
		fn := fmt.Sprintf("shop/mobiledevice.%d.json", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			url := fmt.Sprintf(_url, p)
			// fmt.Println(url)

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
				return NewShop(
					_name,
					_url,

					nil,
				)
			}
			req.Header.Set("Accept", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, _url)
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
		var body2 _Body2
		{
			if err := sonnet.Unmarshal(_body, &body); err != nil {
				if err := sonnet.Unmarshal(_body, &body2); err != nil {
					panic(err)
				}

				for _, result := range body2.Products {
					_result = append(_result, result)
				}

				body.Pagination = body2.Pagination
			} else {
				_result = append(_result, body.Products...)
			}
		}
		// fmt.Println(BytesToString(body))

		if body.Pagination.CurrentPage >= body.Pagination.PagesCount {
			break
		}
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result))
		for _, product := range _result {
			// fmt.Printf("%+v\n", product)

			_title := html.UnescapeString(product.Name)
			if _debug {
				fmt.Println(_title)
			}
			_model := MobileDeviceCleanFn(_title)
			if _debug {
				fmt.Println(_model)
			}

			if Skip(_model) {
				if _debug {
					fmt.Println()
				}

				continue
			}

			if _tests {
				testCases[_title] = _model
			}

			_retailPrice := product.RegularPriceAmount
			_price := product.PriceAmount
			if _debug {
				fmt.Println(_retailPrice)
				fmt.Println(_price)
			}

			_savings := _price - _retailPrice
			_discount := 100 - ((_price * 100) / _retailPrice)
			if _debug {
				fmt.Println(_savings)
				fmt.Println(_discount)
			}

			_link := s.ResolveURL(product.Link).String()
			if _debug {
				fmt.Println(_link)
				fmt.Println()
			}

			product := &Product{
				Code:  _name + "//" + product.ID,
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
