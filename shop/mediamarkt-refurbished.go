package shop

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func XXX_mediamarkt_refurbished(isDryRun *bool) IShop {
	const _name = "Mediamarkt (Refurbished)"
	const _url = "https://refurbished.mediamarkt.ch/ch_de/unsere-refurbished-smartphones?is_in_stock=1&product_list_order=price&product_list_limit=100"

	const _debug = false

	type _Response struct {
		code  string
		title string

		link string

		oldPrice float32
		price    float32
	}

	var _result []_Response
	var _body []byte

	fn := "shop/mediamarkt-refurbished.html"

	if isDryRun != nil && *isDryRun {
		if body, err := os.ReadFile(fn); err != nil {
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

		os.WriteFile(fn, _body, 0664)
	}
	// fmt.Println(string(_body))

	doc := parse(string(_body))

	productList := traverse(doc, "ol", "class", "products")
	// fmt.Println(productList)

	for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling.NextSibling {
		// item := traverse(productList, "li", "class", "product")
		// fmt.Println(item)

		_product := _Response{}

		product := traverse(item, "div", "class", "product-item-details")
		// fmt.Println(product)

		itemLink := traverse(product, "a", "class", "product-item-link")
		// fmt.Println(itemLink)

		link, _ := attr(itemLink.Attr, "href")
		if _debug {
			fmt.Println(link)
		}
		_product.link = link

		itemTitle := traverse(itemLink, "span", "class", "is-refurb")
		// fmt.Println(itemTitle)

		title, _ := text(itemTitle)
		// fmt.Println(title)

		itemAttribute := traverse(product, "div", "class", "product-item-attribute")
		// fmt.Println(itemAttribute)

		attribute, _ := text(itemAttribute)
		// fmt.Println(attribute)
		title += " " + attribute
		if _debug {
			fmt.Println(title)
		}
		_product.title = title

		priceBox := traverse(product, "div", "class", "price-box")
		// fmt.Println(priceBox)

		productId, _ := attr(priceBox.Attr, "data-product-id")
		if _debug {
			fmt.Println(productId)
		}
		_product.code = productId

		priceWrapper := traverse(priceBox, "span", "class", "price-wrapper")
		// fmt.Println(priceWrapper)

		price, _ := attr(priceWrapper.Attr, "data-price-amount")
		if _debug {
			fmt.Println(price)
		}

		if _price, err := strconv.ParseFloat(price, 32); err != nil {
			panic(err)
		} else {
			_product.oldPrice = float32(_price)
		}

		if _debug {
			fmt.Println()
		}

		_result = append(_result, _product)
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result))
		for _, product := range _result {
			// fmt.Println(product)

			_retailPrice := product.oldPrice
			_price := _retailPrice
			if product.price > 0 {
				_price = product.price
			}
			_savings := _price - _retailPrice
			_discount := 100 - ((100 / _retailPrice) * _price)

			product := &Product{
				Code: _name + "//" + product.code,
				Name: product.title,

				RetailPrice: _retailPrice,
				Price:       _price,
				Savings:     _savings,
				Discount:    _discount,

				URL: product.link,
			}

			if s.IsWorth(product) {
				products = append(products, product)
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
