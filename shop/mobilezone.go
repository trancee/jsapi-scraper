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

var MobileZoneRegex = regexp.MustCompile(`\s+\(?\d+\s*GB?|\s+\(?\d+(\.\d+)?"| Dual Sim`) // |\s+\(?[2345]G\)?

var MobileZoneCleanFn = func(name string) string {
	name = strings.NewReplacer(" Phone 1 A063", " Phone (1)", " Xcover5", " XCover 5", " 5G", "").Replace(name)

	if loc := MobileZoneRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)
}

func XXX_mobilezone(isDryRun bool) IShop {
	const _name = "mobilezone"
	const _url = "https://search.epoq.de/inbound-servletapi/getSearchResult?full&ff=e:alloc_THEME&fv=alle_handys&ff=c:anzeigename&fv=Handys&ff=e:isPriceVariant&fv=0&callback=X&tenantId=mobilezone-ch-2019&sessionId=f87cc9415cf968d4d633dd6d15f812ca&orderBy=e:sorting_price&order=asc&limit=100&offset=0&style=compact&format=json&query="

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	type _Product struct {
		MatchItem struct {
			Code struct {
				Value string `json:"$"`
			} `json:"g:id"`

			Description struct {
				Value string `json:"$"`
			} `json:"g:description"`

			Link struct {
				Value string `json:"$"`
			} `json:"link"`

			Action struct {
				Value string `json:"$"`
			} `json:"c:action"`
			Sale struct {
				Value *string `json:"$"`
			} `json:"e:sale"`

			Price struct {
				Value string `json:"$"`
			} `json:"g:price"`
			OldPrice struct {
				Value string `json:"$"`
			} `json:"g:old_price"`
		} `json:"match-item"`
	}

	type _Response struct {
		Result struct {
			Findings struct {
				Products []_Product `json:"finding"`
			} `json:"findings"`
		} `json:"result"`
	}

	var _result _Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	fn := "shop/mobilezone.json"

	if isDryRun {
		if body, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else {
			_body = body
		}
	} else {
		resp, err := http.Get(_url)
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
			_body = body[2:(len(body) - 2)] // remove shitty stuff
		}

		os.WriteFile(path+fn, _body, 0664)
	}
	// fmt.Println(BytesToString(_body))

	if err := sonnet.Unmarshal(_body, &_result); err != nil {
		panic(err)
	}
	// fmt.Println(_result.Products)

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result.Result.Findings.Products))
		for _, product := range _result.Result.Findings.Products {
			// fmt.Printf("%+v\n", product)

			_title := product.MatchItem.Description.Value
			_model := MobileZoneCleanFn(_title)

			if Skip(_model) {
				continue
			}
			if _debug {
				// fmt.Println(_title)
				fmt.Println(_model)
			}

			if _tests {
				testCases[_title] = _model
			}

			if product.MatchItem.Sale.Value == nil {
				continue
			}

			var _retailPrice float32
			var _price float32
			var _savings float32
			var _discount float32

			if product.MatchItem.OldPrice.Value != "" {
				if oldPrice, err := strconv.ParseFloat(product.MatchItem.OldPrice.Value, 32); err != nil {
					panic(err)
				} else {
					_retailPrice = float32(oldPrice)
				}
			}
			if price, err := strconv.ParseFloat(product.MatchItem.Price.Value, 32); err != nil {
				panic(err)
			} else {
				_price = float32(price)

				if _retailPrice == 0 {
					_retailPrice = _price
				}
			}
			if _debug {
				fmt.Println(_retailPrice)
				fmt.Println(_price)
			}

			if _savings == 0 {
				_savings = _price - _retailPrice
			}
			_discount = 100 - ((100 / _retailPrice) * _price)
			if _debug {
				fmt.Println(_savings)
				fmt.Println(_discount)
			}

			_link := product.MatchItem.Link.Value
			if _debug {
				fmt.Println(_link)
				fmt.Println()
			}

			product := &Product{
				Code:  _name + "//" + product.MatchItem.Code.Value,
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
