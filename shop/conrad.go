package shop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func XXX_conrad() IShop {
	const _name = "Conrad"
	const _url = "https://api.conrad.ch/search/1/v3/facetSearch/ch/de/b2c?apikey=2cHbdksbmXc6PQDkPzRVFOcdladLvH7w"

	articles := []map[string]any{}

	type _Product struct {
		Code string `json:"productId"`
		Name string `json:"title"`

		RetailPrice float32
		Price       float32
		Savings     float32
		Discount    float32

		IsBuyable bool `json:"isBuyable"`
	}

	type _Response struct {
		Meta struct {
			Total int `json:"total"`
		} `json:"meta"`

		Products *[]*_Product `json:"hits"`
	}

	var _result _Response

	var jsonData = []byte(`{
		"query": "",
		"enabledFeatures": ["and_filters", "b2b_results_count", "filters_without_values", "query_relaxation", "show_hero_products"],
		"disabledFeatures": [],
		"globalFilter": [{
			"field": "categoryId",
			"type": "TERM_OR",
			"values": "1801015"
		}],
		"facetFilter": [],
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
	}`)

	resp, err := http.Post("https://api.conrad.ch/search/1/v3/facetSearch/ch/de/b2c?apikey=2cHbdksbmXc6PQDkPzRVFOcdladLvH7w", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))

	// var result Response
	if err := json.Unmarshal(body, &_result); err != nil { // Parse []byte to go struct pointer
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
			Code string `json:"articleID"`

			Offers struct {
				Offer struct {
					Code string `json:"articleID"`

					Price struct {
						Price           float32 `json:"price"`
						CrossedOutPrice float32 `json:"crossedOutPrice"`
						SavedAmount     float32 `json:"savedAmount"`
						SavedPercentage float32 `json:"savedPercentage"`
					} `json:"price"`
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
			panic(err)
		}
		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body) // response body is []byte
		if err != nil {
			panic(err)
		}
		// fmt.Println(string(body))

		var __result _Response
		if err := json.Unmarshal(body, &__result); err != nil { // Parse []byte to go struct pointer
			panic(err)
		}
		// fmt.Println(_result.PriceAndAvailabilityFacadeResponse.Products)

		for _, _product := range __result.PriceAndAvailabilityFacadeResponse.Products.Product {
			code := strings.TrimLeft(_product.Code, "0")
			// fmt.Println(code)
			for _, product := range *_result.Products {
				if product.Code == code {
					product.RetailPrice = _product.Offers.Offer.Price.Price
					if _product.Offers.Offer.Price.CrossedOutPrice > 0 {
						product.Price = _product.Offers.Offer.Price.CrossedOutPrice
					} else {
						product.Price = _product.Offers.Offer.Price.Price
					}
					product.Savings = _product.Offers.Offer.Price.SavedAmount
					product.Discount = _product.Offers.Offer.Price.SavedPercentage
				}
			}
		}
	}

	r := regexp.MustCompile("[^a-z0-9 .-]")

	_parseFn := func() []Product {
		products := []Product{}

		// https://www.conrad.ch/de/p/samsung-galaxy-a04s-smartphone-32-gb-16-5-cm-6-5-zoll-schwarz-android-12-dual-sim-2749363.html

		fmt.Printf("-- %s (%d)\n", _name, len(*_result.Products))
		for _, product := range *_result.Products {
			// fmt.Println(product)
			_retailPrice := product.RetailPrice
			_price := product.Price
			_savings := product.Savings
			_discount := product.Discount

			if _price > 0 && _discount >= 10 {
				_productName := strings.NewReplacer(" ", "-", ".", "-").Replace(r.ReplaceAllString(strings.ToLower(product.Name), "$1"))
				_productUrl := fmt.Sprintf("https://www.conrad.ch/de/p/%s-%s.html", _productName, product.Code)

				products = append(products, Product{
					Code: _name + "//" + product.Code,
					Name: product.Name,

					RetailPrice: _retailPrice,
					Price:       _price,
					Savings:     _savings,
					Discount:    _discount,

					URL: _productUrl,
				})
			}
		}
		return products
	}

	return NewShop(
		_name,
		_url,

		&_result,

		_parseFn,
	)
}
