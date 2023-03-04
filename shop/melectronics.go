package shop

import "fmt"

func XXX_melectronics() IShop {
	const _name = "melectronics"
	const _url = "https://www.melectronics.ch/jsapi/v1/de/products/search/category/3421317829?q=%3Aprice-asc%3Aspecial%3AAktion&pageSize=20&currentPage=0"

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

	_parseFn := func() []Product {
		products := []Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result.Products))
		for _, product := range _result.Products {
			products = append(products, Product{
				Code: _name + "//" + product.Code,
				Name: product.Brand.Name + " " + product.Name,

				RetailPrice: product.SuggestedRetailPrice.Value,
				Price:       product.Price.Value,
				Savings:     product.Price.Value - product.SuggestedRetailPrice.Value,
				Discount:    product.PercentageReduction,

				URL: product.URL,
			})
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
