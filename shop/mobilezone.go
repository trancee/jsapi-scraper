package shop

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func XXX_mobilezone(isDryRun bool) IShop {
	const _name = "mobilezone"
	const _url = "https://search.epoq.de/inbound-servletapi/getSearchResult?full&ff=e:alloc_THEME&fv=alle_handys&ff=c:anzeigename&fv=Handys&ff=e:isPriceVariant&fv=0&callback=X&tenantId=mobilezone-ch-2019&sessionId=f87cc9415cf968d4d633dd6d15f812ca&orderBy=e:sorting_price&order=asc&limit=100&offset=0&style=compact&format=json&query="

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
			panic(err)
		}
		defer resp.Body.Close()

		if body, err := io.ReadAll(resp.Body); err != nil {
			panic(err)
		} else {
			_body = body[2:(len(body) - 2)] // remove shitty stuff
		}

		os.WriteFile(path+fn, _body, 0664)
	}
	// fmt.Println(string(_body))

	if err := json.Unmarshal(_body, &_result); err != nil {
		panic(err)
	}
	// fmt.Println(_result.Products)

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result.Result.Findings.Products))
		for _, product := range _result.Result.Findings.Products {
			if product.MatchItem.Sale.Value == nil {
				continue
			}

			if _sale, err := strconv.ParseBool(*product.MatchItem.Sale.Value); err != nil {
				panic(err)
			} else if _sale {
				// fmt.Println(product)

				var _retailPrice float32
				var _price float32
				var _savings float32
				var _discount float32

				if oldPrice, err := strconv.ParseFloat(product.MatchItem.OldPrice.Value, 32); err != nil {
					panic(err)
				} else {
					_retailPrice = float32(oldPrice)
				}
				if price, err := strconv.ParseFloat(product.MatchItem.Price.Value, 32); err != nil {
					panic(err)
				} else {
					_price = float32(price)
				}

				if _savings == 0 {
					_savings = _price - _retailPrice
				}
				_discount = 100 - ((100 / _retailPrice) * _price)

				product := &Product{
					Code: _name + "//" + product.MatchItem.Code.Value,
					Name: product.MatchItem.Description.Value,

					RetailPrice: _retailPrice,
					Price:       _price,
					Savings:     _savings,
					Discount:    _discount,

					URL: product.MatchItem.Link.Value,
				}

				if s.IsWorth(product) {
					products = append(products, product)
				}
			}
		}

		return &products
	}

	return NewShop(
		_name,
		_url,

		_parseFn,
	)
}
