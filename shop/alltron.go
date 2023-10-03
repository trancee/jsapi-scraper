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

	helpers "jsapi-scraper/helpers"
)

var AlltronRegex = regexp.MustCompile(`(\s*[-,]\s+)|(\d+\s*GB?)|\s+((EE )?Enterprise Edition( CH)?)`)

var AlltronCleanFn = func(name string) string {
	name = strings.NewReplacer(" Phones ", " ", "Recommerce Switzerland SA ", "", "3. Gen.", "(2022)").Replace(name)

	if loc := AlltronRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)

	// s := strings.Split(name, " ")

	// if s[0] == "iPhone" {
	// 	name = "Apple " + name
	// }

	// if s[0] == "Apple" {
	// 	name = strings.NewReplacer(" 2020", " (2020)", " 2022", " (2022)", " 2nd Gen", " (2020)", " 3rd Gen", " (2022)").Replace(name)
	// } else {
	// 	// Remove year component for all other than Apple.
	// 	name = regexp.MustCompile(`\s+\(?20[12]\d\)?`).ReplaceAllString(name, "")
	// }

	// return strings.TrimSpace(name)
}

func XXX_alltron(isDryRun bool) IShop {
	const _name = "Alltron"
	// const _url =     "https://alltron.ch/api/v1/catalog/search?path=/telco-ucc/mobiltelefonie/smartphones/smartphone&limit=192&sortProducts=priceasc&filters=availability:::Verfügbar&searchEarlyFilter=true&format=json"
	_url := fmt.Sprintf("https://alltron.ch/api/v1/catalog/search?path=/telco-ucc/mobiltelefonie/smartphones/smartphone&limit=192&sortProducts=priceasc&filters=availability:::Verfügbar|||price_standard:::%.f~~~%.f&searchEarlyFilter=true&format=json", ValueMinimum, ValueMaximum)

	const _api = "https://alltron.ch/api/v1/products/multiple-tiles/"

	const _tests = false

	testCases := map[string]string{}

	type _Product struct {
		Description struct {
			Title string `json:"title"`
		} `json:"description"`

		Settings struct {
			IsBuyable      bool `json:"isBuyable"`
			IsNew          bool `json:"isNew"`
			IsOccasion     bool `json:"isOccasion"`
			IsSellout      bool `json:"isSellout"`
			IsSpecialOffer bool `json:"isSpecialOffer"`
			IsSpecialOrder bool `json:"isSpecialOrder"`

			Quantity int `json:"quantity"`

			ShouldShowPrice bool `json:"shouldShowPrice"`
		} `json:"settings"`

		EffectivePricing struct {
			UserPrice float32 `json:"userPrice"`
			ListPrice float32 `json:"listPrice"`
			MainPrice float32 `json:"mainPrice"`
		} `json:"effectivePricing"`

		SKU string `json:"sku"`
	}

	type _Response struct {
		ProductsCount int `json:"productsCount"`

		Products []_Product `json:"products"`
	}

	var _result _Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	fn := "shop/alltron.json"

	if isDryRun {
		if body, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else {
			_body = body
		}
	} else {
		req, err := http.NewRequest("GET", _url, nil)
		if err != nil {
			// panic(err)
			fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
			return NewShop(
				_name,
				_url,

				nil,
			)
		}
		req.Header.Set("Accept-Language", "de")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")

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
			panic(err)
		} else {
			_body = body
		}

		os.WriteFile(path+fn, _body, 0664)
	}
	// fmt.Println(string(_body))

	if err := sonnet.Unmarshal(_body, &_result); err != nil {
		panic(err)
	}

	_skus := []string{}
	for _, product := range _result.Products {
		_skus = append(_skus, product.SKU)
	}
	// fmt.Println(_skus)

	fn = "shop/alltron-products.json"

	if isDryRun {
		if body, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else {
			_body = body
		}
	} else {
		req, err := http.NewRequest("GET", _api+strings.Join(_skus, ","), nil)
		if err != nil {
			// panic(err)
			fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
			return NewShop(
				_name,
				_url,

				nil,
			)
		}
		req.Header.Set("Accept-Language", "de")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			// panic(err)
			fmt.Printf("[%s] %s (%s)\n", _name, err, _api+strings.Join(_skus, ","))
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

	_products := make(map[string]*_Product)
	if err := sonnet.Unmarshal(_body, &_products); err != nil {
		panic(err)
	}
	// fmt.Println(_products)

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_products))
		for _, product := range _products {
			// fmt.Println(product)
			if product != nil {
				_title := strings.ReplaceAll(product.Description.Title, "Fairphone Fairphone", "Fairphone")
				_model := AlltronCleanFn(_title)

				if Skip(_model) {
					continue
				}

				if _tests {
					testCases[_title] = _model
				}

				if product.Settings.IsNew {
					_title += " [N]"
				} else if product.Settings.IsSpecialOffer {
					_title += " [P]"
				} else if product.Settings.IsOccasion {
					_title += " [R]"
				} else if product.Settings.IsSellout {
					_title += " [S]"
				}

				_retailPrice := product.EffectivePricing.ListPrice
				_price := _retailPrice
				if product.EffectivePricing.MainPrice > 0 {
					_retailPrice = product.EffectivePricing.MainPrice
				}

				_savings := _price - _retailPrice
				_discount := 100 - ((100 / _retailPrice) * _price)

				_link := s.ResolveURL("https://alltron.ch/de/product/" + product.SKU).String()

				product := &Product{
					Code:  _name + "//" + product.SKU,
					Name:  _title,
					Model: _model,

					RetailPrice: _retailPrice,
					Price:       _price,
					Savings:     _savings,
					Discount:    _discount,

					Quantity: product.Settings.Quantity,

					URL: _link,
				}

				if s.IsWorth(product) {
					products = append(products, product)
				}
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
