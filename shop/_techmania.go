package shop

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Not feasible due to Cloudflare Anti-Bot Protection

func XXX_techmania() IShop {
	const _name = "techmania"
	const _url = "https://www.techmania.ch/de/product/list/smartphones-11853?smsc=100"

	const _debug = true

	type _Response struct {
		code  string
		title string

		link string

		oldPrice float32
		price    float32
	}

	var _result []_Response

	resp, err := http.Get(_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// body, err := os.ReadFile("shop/alltron.html")
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(string(body))

	doc := parse(string(body))

	productList := traverse(doc, "ul", "class", "cds-ProductList")
	// fmt.Println(productList)

	for item := productList.FirstChild; /*.NextSibling*/ item != nil; item = item.NextSibling /*.NextSibling*/ {
		// fmt.Println(item)

		_product := _Response{}

		productBox := traverse(item, "div", "class", "cds-ProductBox")
		// fmt.Println(productBox)

		productKey := traverse(productBox, "span", "class", "cds-CopyableText-Content")
		// fmt.Println(productKey)

		productId, _ := text(productKey)
		if _debug {
			fmt.Println(productId)
		}
		_product.code = productId

		imageLink := traverse(productBox, "a", "class", "cds-ProductBox-ImageLink")
		// fmt.Println(imageLink)

		link, _ := attr(imageLink.Attr, "href")
		if _debug {
			fmt.Println(link)
		}
		_product.link = link

		itemTitle := traverse(productBox, "a", "class", "cds-ProductBox-Title")
		// fmt.Println(itemTitle)

		title, _ := attr(itemTitle.Attr, "title")
		if _debug {
			fmt.Println(title)
		}
		_product.title = title

		productPrice := traverse(productBox, "div", "class", "cds-ProductBox-Price")
		// fmt.Println(productPrice)

		productPriceValue := traverse(productPrice, "div", "class", "cds-Price-Value")
		// fmt.Println(productPriceValue)

		price, _ := text(productPriceValue.FirstChild.NextSibling.NextSibling.NextSibling)
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

		nil,

		_parseFn,
	)
}
