package shop

import (
	"fmt"
	"strconv"
)

func XXX_mobilezone() IShop {
	const _name = "mobilezone"
	const _url = "https://search.epoq.de/inbound-servletapi/getSearchResult?full&ff=e%3Aalloc_THEME&fv=alle_handys&ff=c%3Aanzeigename&fv=Handys&ff=e%3AisPriceVariant&fv=0&callback=X&tenantId=mobilezone-ch-2019&sessionId=f87cc9415cf968d4d633dd6d15f812ca&orderBy=e%3Asorting_price&order=asc&limit=100&offset=0&style=compact&format=json&query=*"

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
				Value string `json:"$"`
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

	_parseFn := func() []Product {
		products := []Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result.Result.Findings.Products))
		for _, product := range _result.Result.Findings.Products {
			if _sale, err := strconv.ParseBool(product.MatchItem.Sale.Value); err != nil {
				panic(err)
			} else if _sale {
				// fmt.Println(product)

				_retailPrice, err := strconv.ParseFloat(product.MatchItem.OldPrice.Value, 32)
				if err != nil {
					panic(err)
				}
				_price, err := strconv.ParseFloat(product.MatchItem.Price.Value, 32)
				if err != nil {
					panic(err)
				}
				_savings := float32(_price - _retailPrice)
				_discount := float32(100 - ((100 / _retailPrice) * _price))

				if _price > 0 && _discount >= 10 {
					products = append(products, Product{
						Code: _name + "//" + product.MatchItem.Code.Value,
						Name: product.MatchItem.Description.Value,

						RetailPrice: float32(_retailPrice),
						Price:       float32(_price),
						Savings:     _savings,
						Discount:    _discount,

						URL: product.MatchItem.Link.Value,
					})
				}
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
