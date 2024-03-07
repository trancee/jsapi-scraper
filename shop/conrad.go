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

	"github.com/sugawarayuuta/sonnet"

	helpers "jsapi-scraper/helpers"
)

var ConradRegex = regexp.MustCompile(`\s*[-,]\s+|\W\+\s+|EU |\d+\s*GB|\s*\d+G|\s+\(Version 20[12]\d\)|\s+\(Grade [A-Z]\)|\s+(((Senioren-|senior |Industrie |Outdoor )?Smartphone)|((EE )?Enterprise Edition( CH)?)|Satellite|Ex-geschütztes Handy|Fusion( Holiday Edition)?|Refurbished|\(PRODUCT\) RED™|Weiß)`)

var ConradCleanFn = func(name string) string {
	name = strings.NewReplacer(" Phones ", " ", " Mini iPhone", " Mini", "Edge20", "Edge 20", "Edge Neo 40", "Edge 40 Neo", "Samsung XCover", "Samsung Galaxy XCover", "Renewd® ", "", "refurbished", "", "5G Smartphone", "").Replace(name)

	if loc := ConradRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	s := strings.Split(name, " ")

	if s[0] == "Gigaset" {
		name = strings.ReplaceAll(name, " Gigaset", "")
	}

	return helpers.Lint(name)
}

func XXX_conrad(isDryRun bool) IShop {
	const _name = "Conrad"
	const _url = "https://api.conrad.ch/search/1/v3/facetSearch/ch/de/b2c?apikey=2cHbdksbmXc6PQDkPzRVFOcdladLvH7w"

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	articles := []map[string]any{}

	type _Product struct {
		Code string `json:"productId"`
		Name string `json:"title"`

		RetailPrice float32
		Price       float32
		Savings     float32
		Discount    float32

		Quantity int

		IsBuyable      bool `json:"isBuyable"`
		IsSpecialOffer bool

		TechnicalDetails []struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"technicalDetails"`
	}

	type _Response struct {
		Meta struct {
			Total int `json:"total"`
		} `json:"meta"`

		Products *[]*_Product `json:"hits"`
	}

	var _result _Response
	var _body []byte

	var jsonData = StringToBytes(
		fmt.Sprintf(
			`{
				"query": "",
				"enabledFeatures": ["and_filters", "b2b_results_count", "filters_without_values", "query_relaxation", "show_hero_products"],
				"disabledFeatures": [],
				"globalFilter": [{
					"field": "categoryId",
					"type": "TERM_OR",
					"values": "1801015"
				}],
				"facetFilter": [{
					"field": "price",
					"type": "RANGE",
					"values": [
						%.f,
						%.f
					]
				}],
				"sort": [{
					"field": "price",
					"order": "asc"
				}],
				"from": 0,
				"size": 200,
				"facets": [],
				"partialThreshold": 10,
				"partialQueries": 3,
				"partialQuerySize": 6
			}`,
			ValueMinimum,
			ValueMaximum,
		),
	)

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	fn := "shop/conrad.json"

	if isDryRun {
		if body, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else {
			_body = body
		}
	} else {
		resp, err := http.Post(_url, "application/json", bytes.NewBuffer(jsonData))
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

	for _, product := range *_result.Products {
		articles = append(articles, map[string]any{
			"articleID":         product.Code,
			"insertCode":        "UO",
			"calculatePrice":    true,
			"checkAvailability": true,
			"findExclusions":    true,
		})
	}

	{
		type _Product struct {
			Code string `json:"articleId"`

			Offers struct {
				Offer struct {
					Code string `json:"articleId"`

					Price struct {
						Price           float32 `json:"price"`
						CrossedOutPrice float32 `json:"crossedOutPrice"`
						SavedAmount     float32 `json:"savedAmount"`
						SavedPercentage float32 `json:"savedPercentage"`

						IsSpecialOffer string `json:"isSpecialOffer"`
					} `json:"price"`

					Availability struct {
						Quantity float32 `json:"stockQuantity"`
					} `json:"availability"`
				} `json:"offer"`
			} `json:"offers"`
		}

		type _Response struct {
			PriceAndAvailabilityFacadeResponse struct {
				Products struct {
					Product []_Product `json:"product"`
				} `json:"products"`
				// PriceAndAvailability []struct{} `json:"priceAndAvailability"`
			} `json:"priceAndAvailabilityFacadeResponse"`
		}

		fn := "shop/conrad-articles.json"

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			reqData, err := json.Marshal(map[string]any{
				"ns:inputArticleItemList": map[string]any{
					"#namespaces": map[string]any{
						"ns": "http://www.conrad.de/ccp/basit/service/article/priceandavailabilityservice/api",
					},
					"articles": articles,
				},
			})
			if err != nil {
				panic(err)
			}

			req, err := http.NewRequest("POST", "https://api.conrad.ch/price-availability/4/CQ_CH_B2C/facade?apikey=2cHbdksbmXc6PQDkPzRVFOcdladLvH7w&forceStorePrice=false&overrideCalculationSchema=GROSS", bytes.NewBuffer(reqData))
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
				return NewShop(
					_name,
					_url,

					nil,
				)
			}
			req.Header.Set("Accept", "application/json, text/plain, */*")
			req.Header.Set("Content-Type", "application/json")

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
		// fmt.Println(BytesToString(_body))

		var __result _Response
		if err := sonnet.Unmarshal(_body, &__result); err != nil {
			panic(err)
		}
		// fmt.Println(__result.PriceAndAvailabilityFacadeResponse.Products.Product)

		for _, _product := range __result.PriceAndAvailabilityFacadeResponse.Products.Product {
			code := strings.TrimLeft(_product.Code, "0")
			// fmt.Println(code)
			for _, product := range *_result.Products {
				if Skip(product.Name) {
					continue
				}

				if brand := strings.Split(product.Name, " "); strings.EqualFold(brand[0], brand[1]) {
					product.Name = strings.ReplaceAll(product.Name, " "+brand[1], "")
				} else if brand[0] == "ZTE" {
					product.Name = strings.ReplaceAll(product.Name, "ZTE Blade V40 Vita 4 Smartphone", "ZTE Blade V40 Vita Smartphone")
				}

				if product.Code == code {
					isSpecialOffer := false
					if _product.Offers.Offer.Price.IsSpecialOffer != "false" {
						isSpecialOffer = true
					}
					product.RetailPrice = _product.Offers.Offer.Price.Price
					if _product.Offers.Offer.Price.CrossedOutPrice > 0 {
						product.Price = _product.Offers.Offer.Price.CrossedOutPrice
					} else {
						product.Price = _product.Offers.Offer.Price.Price
					}
					product.Savings = _product.Offers.Offer.Price.SavedAmount
					product.Discount = _product.Offers.Offer.Price.SavedPercentage

					product.Quantity = int(_product.Offers.Offer.Availability.Quantity)

					product.IsSpecialOffer = isSpecialOffer
					// fmt.Println(product)
				}
			}
		}
	}

	r := regexp.MustCompile("[^a-z0-9 .-]")

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		// https://www.conrad.ch/de/p/samsung-galaxy-a04s-smartphone-32-gb-16-5-cm-6-5-zoll-schwarz-android-12-dual-sim-2749363.html

		fmt.Printf("-- %s (%d)\n", _name, len(*_result.Products))
		for _, product := range *_result.Products {
			// fmt.Println(product)

			_title := product.Name
			_model := ConradCleanFn(_title)

			if Skip(_model) {
				continue
			}
			if _debug {
				// fmt.Println(_title)
				fmt.Println(_model)
			}

			for _, detail := range product.TechnicalDetails {
				switch detail.Name {
				case "ATT_CALC_DISPLAY-DIAGONAL_CM",
					"ATT_LOV_SIM_CARD_TECHNOLOGY",
					"ATT_OPERATINGSYSTEM":
					for _, value := range detail.Values {
						_title = strings.ReplaceAll(_title, " "+value, "")
					}
				case "ATT_DISPLAY_DIAGONAL":
					for _, value := range detail.Values {
						_title = strings.ReplaceAll(_title, " ("+value+")", "")
					}

				}
			}

			if _tests {
				testCases[_title] = _model
			}

			_retailPrice := max(product.RetailPrice, product.Price)
			_price := min(product.RetailPrice, product.Price)
			if _debug {
				fmt.Println(_retailPrice)
				fmt.Println(_price)
			}

			_savings := product.Savings
			if _savings > 0 {
				_savings = -_savings
			}
			_discount := product.Discount
			if _debug {
				fmt.Println(_savings)
				fmt.Println(_discount)
			}

			_productName := strings.NewReplacer(" ", "-", ".", "-").Replace(r.ReplaceAllString(strings.ToLower(_title), "$1"))
			_link := fmt.Sprintf("https://www.conrad.ch/de/p/%s-%s.html", _productName, product.Code)
			if _debug {
				fmt.Println(_link)
				fmt.Println()
			}

			product := &Product{
				Code:  _name + "//" + product.Code,
				Name:  _title,
				Model: _model,

				RetailPrice: _retailPrice,
				Price:       _price,
				Savings:     _savings,
				Discount:    _discount,

				Quantity: product.Quantity,

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
