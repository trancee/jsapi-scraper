package shop

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func XXX_microspot() IShop {
	const _name = "Microspot"
	const _url = "https://www.microspot.ch/mspocc/occ/msp/products/search?currentPage=0&pageSize=100&query=%3Aprice-asc%3AcategoryPath%3A%2F1%2F400%2F4100%3AcategoryPath%3A%2F1%2F400%2F4100%2F411000%3AhasPromoLabel%3Atrue&lang=de"

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

	r := regexp.MustCompile("[^a-z0-9 .-]")

	_parseFn := func() []Product {
		products := []Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result.Products))
		for _, product := range _result.Products {
			// if product.MaxOrderValue == 0 {
			// 	// No more stock
			// 	continue
			// }

			var _retailPrice float32
			var _price float32
			var _savings float32
			var _discount float32

			for _, price := range product.Price.Prices {
				if price.FixPrice {
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
						_savings = price.Savings.Value
					}
				}
			}

			if _savings == 0 {
				_savings = _price - _retailPrice
			}
			_discount = 100 - ((100 / _retailPrice) * _price)

			if _price > 0 && _discount >= 10 {
				_productName := strings.NewReplacer(" ", "-", ".", "-").Replace(r.ReplaceAllString(strings.ToLower(product.Name), "$1"))
				_productUrl := fmt.Sprintf("https://www.microspot.ch/de/telefonie-tablet-smartwatch/smartphone/smartphone--c411000/%s--p%s", _productName, product.Code)

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
