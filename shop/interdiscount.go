package shop

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/sugawarayuuta/sonnet"

	helpers "jsapi-scraper/helpers"
)

var InterdiscountRegex = regexp.MustCompile(`\(\d+\s*GB?|\(\d\.\d{1,2}"|\s+[2345]G| LTE| Enterprise Edition`) // |\s+\(?20[12]\d\)?

var InterdiscountCleanFn = func(name string) string {
	if loc := InterdiscountRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	s := strings.Split(name, " ")

	if s[0] == "GIGASET" {
		name = strings.NewReplacer(" 2022", "").Replace(name)
	}

	if s[0] == "NOKIA" && s[1] == "Nokia" {
		name = strings.ReplaceAll(name, "NOKIA Nokia ", "NOKIA ")
	}

	return helpers.Lint(name)
}

func XXX_interdiscount(isDryRun bool) IShop {
	const _name = "Interdiscount"
	// const _url = "https://www.interdiscount.ch/idocc/occ/id/products/search?currentPage=0&pageSize=100&query=:price-asc:categoryPath:/1/400/4100:categoryPath:/1/400/4100/411000:hasPromoLabel:true&lang=de"
	_url := fmt.Sprintf("https://www.interdiscount.ch/idocc/occ/id/products/search?currentPage=0&pageSize=100&query=:price-asc:categoryPath:/1/400/4100:categoryPath:/1/400/4100/411000:price_min:%.f:price_max:%.f:producttype411001k:Smartphone&lang=de", ValueMinimum, ValueMaximum)

	const _tests = false

	testCases := map[string]string{}

	type _Product struct {
		Code string `json:"code"`
		Name string `json:"name"`

		Manufacturer string `json:"manufacturer"`

		PromoLabels []struct {
			Text string `json:"text"`
		} `json:"promoLabels"`

		Price struct {
			Prices []struct {
				FinalPrice struct {
					Value float32 `json:"value"`
				} `json:"finalPrice"`
				InsteadPrice struct {
					Value float32 `json:"value"`
				} `json:"insteadPrice"`
				Savings struct {
					Value float32 `json:"value"`
				} `json:"savings"`
				// Discount struct {
				// 	Value float32 `json:"value"`
				// } `json:"discount"`
				FixPrice bool      `json:"fixPrice"`
				Expires  time.Time `json:"expires"`
			} `json:"prices"`
		} `json:"productPriceData"`

		Orderable     bool `json:"productOrderable"`
		MaxOrderValue int  `json:"maxOrderValue"`
	}

	type _Response struct {
		CategoryCode string `json:"categoryCode"`

		Pagination struct {
			Page       int `json:"currentPage"`
			PerPage    int `json:"pageSize"`
			Total      int `json:"totalNumberOfResults"`
			TotalPages int `json:"numberOfPages"`
		} `json:"pagination"`

		Products []_Product `json:"products"`
	}

	var _result _Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	fn := "shop/interdiscount.json"

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
			_body = body
		}

		os.WriteFile(path+fn, _body, 0664)
	}
	// fmt.Println(BytesToString(_body))

	if err := sonnet.Unmarshal(_body, &_result); err != nil {
		panic(err)
	}
	// fmt.Println(_result.Products)

	r := regexp.MustCompile("[^a-z0-9 .-]")

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result.Products))
		for _, product := range _result.Products {
			_title := product.Name
			_model := InterdiscountCleanFn(_title)

			if Skip(_model) {
				continue
			}

			if _tests {
				testCases[_title] = _model
			}

			var _retailPrice float32
			var _price float32
			var _savings float32
			var _discount float32

			for _, price := range product.Price.Prices {
				_price = price.FinalPrice.Value

				if price.FixPrice {
					_retailPrice = _price

					if !price.Expires.IsZero() {
						if time.Now().Before(price.Expires) {
							if _price > 0 {
								_retailPrice = price.FinalPrice.Value
							} else {
								_price = price.FinalPrice.Value
							}
						}
					} else {
						_price = price.FinalPrice.Value
					}
				} else {
					_retailPrice = price.FinalPrice.Value

					if price.InsteadPrice.Value > 0 {
						_price = price.FinalPrice.Value
						_retailPrice = price.InsteadPrice.Value
						_savings = -price.Savings.Value
					}
				}
			}

			if _savings == 0 {
				_savings = _price - _retailPrice
			}
			_discount = 100 - ((100 / _retailPrice) * _price)

			_productName := strings.NewReplacer(" ", "-", ".", "-").Replace(r.ReplaceAllString(strings.ToLower(_title), "$1"))
			_productUrl := fmt.Sprintf("https://www.interdiscount.ch/de/telefonie-tablet-smartwatch/smartphone/smartphone--c411000/%s--p%s", _productName, product.Code)

			product := &Product{
				Code:  _name + "//" + product.Code,
				Name:  _title,
				Model: _model,

				RetailPrice: _retailPrice,
				Price:       _price,
				Savings:     _savings,
				Discount:    _discount,

				URL: _productUrl,
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
