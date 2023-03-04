package shop

import (
	"fmt"
	"regexp"
	"strings"
)

func XXX_interdiscount() IShop {
	const _name = "Interdiscount"
	const _url = "https://www.interdiscount.ch/idocc/occ/id/products/search?currentPage=0&pageSize=100&query=%3Aprice-asc%3AcategoryPath%3A%2F1%2F400%2F4100%3AcategoryPath%3A%2F1%2F400%2F4100%2F411000%3AhasPromoLabel%3Atrue&lang=de"

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
				FixPrice bool `json:"fixPrice"`
			} `json:"prices"`
		} `json:"productPriceData"`

		Orderable bool `json:"productOrderable"`
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
			var _retailPrice float32
			var _price float32
			var _savings float32
			var _discount float32

			for _, price := range product.Price.Prices {
				if price.FixPrice {
					_price = price.FinalPrice.Value
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
				_productUrl := fmt.Sprintf("https://www.interdiscount.ch/de/telefonie-tablet-smartwatch/smartphone/smartphone--c411000/%s--p%s", _productName, product.Code)

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
