package shop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	helpers "jsapi-scraper/helpers"

	"golang.org/x/net/html"
)

var ManorRegex = regexp.MustCompile(`, |\d\+\d+GB|\s+\(?[2345]G\)?|\d+(,\d)? cm| \d[.,]\d|(\d+\s*GB?)|\s+20[12]\d|(SM-)?[AFGMS]\d{3}[BFR]?(\/DSN?)?| XT\d{4}-\d+|\s+EE |\s+(Enterprise Edition( CH)?)| Dual`)

var ManorCleanFn = func(name string) string {
	name = strings.NewReplacer(" NOK ", " ", " Smartphone Pack ", " ", " Smartphone Bundle ", " ", " Pack Smartphone Vivo", " ", "NOKIA Nokia ", "Nokia ", "OPPO OPPO ", "OPPO ", "OPPO Oppo ", "OPPO ").Replace(name)

	name = regexp.MustCompile(`\s{2,}`).ReplaceAllString(name, " ")

	if loc := ManorRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)
}

func XXX_manor(isDryRun bool) IShop {
	const _name = "Manor"
	const _url = "https://ecom-api.manor.ch/graphql"

	const _tests = false

	testCases := map[string]string{}

	type _Result struct {
		Code string `json:"code"`
		Name string `json:"name"`

		Description string `json:"description"`

		BrandID string `json:"brandId"`
		Brand   string `json:"brandName"`

		Link string `json:"link"`

		PriceValue struct {
			Amount float32 `json:"amount"`
		} `json:"priceValue"`
		OriginalPrice struct {
			Amount float32 `json:"amount"`
		} `json:"originalPrice"`

		Stock struct {
			Level int `json:"level"`
		} `json:"stock"`
	}

	type _Response struct {
		Data struct {
			SearchProducts struct {
				Products []_Result `json:"products"`

				Page     int `json:"page"`
				PageSize int `json:"pageSize"`

				TotalPages   int `json:"totalPages"`
				TotalResults int `json:"totalResults"`
			} `json:"searchProducts"`
		} `json:"data"`
	}

	var _result _Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	var _results []_Result

	for p := 1; p <= 5; p++ {
		fn := fmt.Sprintf("shop/manor.%d.json", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			jsonData := []byte(fmt.Sprintf(`{
				"operationName": "SearchProducts",
				"variables": {
					"input": {
						"numericFilters": [
							{
								"fieldName": "priceValue",
								"lowerBound": %.f,
								"upperBound": %.f
							}
						],
						"orderBy": "PRICE_VALUE_ASC",
						"page": %d,
						"pageSize": 24,
						"selectedFilters": [
							{
								"facetName": "category",
								"facetValues": [
									"telephone-navigation-smartphones"
								]
							}
						],
						"mixedRuleSearchResult": {}
					}
				},
				"query": "query SearchProducts($input: InputSearch!) {\n  searchProducts(input: $input) {\n    ...productSearchResultFields\n    __typename\n  }\n}\n\nfragment productSearchResultFields on ProductSearchResult {\n  totalResults\n  page\n  pageSize\n  totalPages\n  queryId\n  analyticsIndexName\n  products {\n    ...indexedProductFields\n    __typename\n  }\n  productListerConfig {\n    hideCount\n    __typename\n  }\n  searchUserData {\n    type\n    url\n    __typename\n  }\n  mixedRuleSearchResult {\n    introductoryOffset\n    introductoryProductCodes\n    marketplaceOffset\n    wholesaleOffset\n    __typename\n  }\n  __typename\n}\n\nfragment indexedProductFields on IndexedProduct {\n  id\n  baseCode\n  titles {\n    first\n    second\n    third\n    __typename\n  }\n  priceValue {\n    ...priceFields\n    __typename\n  }\n  originalPrice {\n    ...priceFields\n    __typename\n  }\n  uvpPrice {\n    ...priceFields\n    __typename\n  }\n  stock {\n    level\n    status\n    __typename\n  }\n  variantColors {\n    name\n    url\n    hexCode\n    variantCode\n    __typename\n  }\n  imageUrls {\n    mobile\n    tablet\n    desktop\n    __typename\n  }\n  labels {\n    id\n    text\n    backgroundColor\n    textColor\n    priority\n    __typename\n  }\n  discountLabels {\n    id\n    text\n    backgroundColor\n    textColor\n    priority\n    __typename\n  }\n  averageRating\n  brandId\n  brandName\n  category\n  color\n  size\n  code\n  description\n  isFromPrice\n  isManorProduct\n  link\n  name\n  productDisplayConfig {\n    ...productDisplayConfigFields\n    __typename\n  }\n  gtin\n  productBreadcrumbPath\n  offlineAvailabilityStatus\n  __typename\n}\n\nfragment priceFields on Price {\n  currency\n  amount\n  digits\n  formattedValue\n  priceType\n  __typename\n}\n\nfragment productDisplayConfigFields on ProductDisplayConfig {\n  hideRatings\n  hideLabels\n  hideAvailability\n  hideCarousels\n  __typename\n}"
			}`, ValueMinimum, ValueMaximum, p))

			req, err := http.NewRequest("POST", _url, bytes.NewBuffer(jsonData))
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
				return NewShop(
					_name,
					_url,

					nil,
				)
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
			req.Header.Set("Origin", "https://www.manor.ch")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
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

		if err := json.Unmarshal(_body, &_result); err != nil {
			panic(err)
		}

		_results = append(_results, _result.Data.SearchProducts.Products...)

		if _result.Data.SearchProducts.TotalPages <= p {
			break
		}
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_results))
		for _, product := range _results {
			// fmt.Println(_product)

			product.Brand = html.UnescapeString(product.Brand)
			product.Name = html.UnescapeString(product.Name)
			product.Description = html.UnescapeString(product.Description)

			_title := product.Name
			if !strings.HasPrefix(strings.ToUpper(_title), strings.ToUpper(product.Brand)) {
				_title = product.Brand + " " + _title
			}
			// fmt.Println(_title)
			_model := ManorCleanFn(_title)
			// fmt.Println(_model)
			// fmt.Println()

			if Skip(_model) {
				continue
			}

			if _tests {
				testCases[_title] = _model
			}

			var _savings float32
			var _discount float32

			_retailPrice := product.OriginalPrice.Amount
			_price := _retailPrice
			if product.PriceValue.Amount > 0 {
				_price = product.PriceValue.Amount
			}
			if _retailPrice > 0 {
				_savings = _price - _retailPrice
				_discount = 100 - ((100 / _retailPrice) * _price)
			}

			_link := s.ResolveURL(product.Link).String()

			product := &Product{
				Code:  _name + "//" + product.Code,
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
