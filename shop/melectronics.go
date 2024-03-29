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

var MelectronicsRegex = regexp.MustCompile(` - |\s+\(?[2345]G\)?|,?\s*,?\(?(\d+( ?GB)?\+)?\d+ ?GB\)?|\s+CH$| DS`)

var MelectronicsCleanFn = func(name string) string {
	name = strings.NewReplacer(" 3th ", " 3rd Gen. ", "A53 s", "A53s", "Enterprise Edition", "EE").Replace(name)

	if loc := MelectronicsRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)
}

func XXX_melectronics(isDryRun bool) IShop {
	const _name = "melectronics"
	// const _url = "https://www.melectronics.ch/jsapi/v1/de/products/search/category/3421317829?q=:price-asc:special:Aktion&pageSize=20&currentPage=0"
	// const _url = "https://www.melectronics.ch/jsapi/v1/de/products/search/category/3421317829?q=:price-asc:summaryAsString:Smartphone&pageSize=100"
	_url := fmt.Sprintf("https://www.melectronics.ch/jsapi/v1/de/products/search/category/3421317829?q=:price-asc:priceValue:%%5B%.f+TO+%.f%%5D:summaryAsString:Smartphone&pageSize=100", ValueMinimum, ValueMaximum)

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	type _Product struct {
		Code    string `json:"code"`
		Name    string `json:"name"`
		URL     string `json:"url"`
		Summary string `json:"summary"`

		Price struct {
			Value float32 `json:"value"`
		} `json:"price"`

		SuggestedRetailPrice struct {
			Value float32 `json:"value"`
		} `json:"suggestedRetailPrice"`

		Images []struct {
			URL string `json:"url"`
		} `json:"images"`

		Reservable bool `json:"reservable"`
		Orderable  bool `json:"orderable"`

		PercentageReduction float32 `json:"percentageReduction"`

		Brand struct {
			Name string `json:"name"`

			Image struct {
				URL string `json:"url"`
			} `json:"image"`
		} `json:"brand"`

		Preorder   bool `json:"preorder"`
		NewProduct bool `json:"newProduct"`
	}

	type _Response struct {
		CategoryName string `json:"categoryName"`

		Pagination struct {
			Page       int `json:"currentPage"`
			PerPage    int `json:"pageSize"`
			Total      int `json:"totalResults"`
			TotalPages int `json:"totalPages"`
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

	fn := "shop/melectronics.json"

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

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result.Products))
		for _, product := range _result.Products {
			_title := product.Name
			if len(product.Brand.Name) > 0 && strings.ToUpper(product.Brand.Name) != strings.ToUpper(strings.Split(_title, " ")[0]) {
				_title = product.Brand.Name + " " + _title
			}
			_model := MelectronicsCleanFn(_title)

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

			_retailPrice := product.SuggestedRetailPrice.Value
			if _retailPrice == 0 {
				_retailPrice = product.Price.Value
			}
			_price := _retailPrice
			if product.Price.Value > 0 {
				_price = product.Price.Value
			}
			if _debug {
				fmt.Println(_retailPrice)
				fmt.Println(_price)
			}

			_savings := _price - _retailPrice
			_discount := product.PercentageReduction
			if _debug {
				fmt.Println(_savings)
				fmt.Println(_discount)
			}

			_link := s.ResolveURL(product.URL).String()
			if _debug {
				fmt.Println(_link)
				fmt.Println()
			}

			product := Product{
				Code:  _name + "//" + product.Code,
				Name:  _title,
				Model: _model,

				RetailPrice: _retailPrice,
				Price:       _price,
				Savings:     _savings,
				Discount:    _discount,

				URL: _link,
			}

			if s.IsWorth(&product) {
				products = append(products, &product)
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
