package shop

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func XXX_melectronics(isDryRun bool) IShop {
	const _name = "melectronics"
	const _url = "https://www.melectronics.ch/jsapi/v1/de/products/search/category/3421317829?q=:price-asc:special:Aktion&pageSize=20&currentPage=0"

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
			panic(err)
		}
		defer resp.Body.Close()

		if body, err := io.ReadAll(resp.Body); err != nil {
			panic(err)
		} else {
			_body = body
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

		fmt.Printf("-- %s (%d)\n", _name, len(_result.Products))
		for _, product := range _result.Products {
			product := Product{
				Code: _name + "//" + product.Code,
				Name: product.Brand.Name + " " + product.Name,

				RetailPrice: product.SuggestedRetailPrice.Value,
				Price:       product.Price.Value,
				Savings:     product.Price.Value - product.SuggestedRetailPrice.Value,
				Discount:    product.PercentageReduction,

				URL: s.ResolveURL(product.URL).String(),
			}

			if s.IsWorth(&product) {
				products = append(products, &product)
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
