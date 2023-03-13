package shop

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func XXX_mobiledevice(isDryRun bool) IShop {
	const _name = "mobiledevice"
	const _url = "https://www.mobiledevice.ch/modules/blocklayered/blocklayered-ajax.php?layered_quantity_1=1&id_category_layered=28&orderby=price&orderway=asc&n=100"

	const _debug = false

	type _Response struct {
		code  string
		title string

		link string

		oldPrice float32
		price    float32

		quantity int
	}

	var _result []_Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	fn := "shop/mobiledevice.html"

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

	type _Body struct {
		Products string `json:"productList"`
	}

	var body _Body
	{
		if err := json.Unmarshal([]byte(_body), &body); err != nil {
			panic(err)
		}
		_body = []byte(body.Products)
	}
	// fmt.Println(string(_body))

	doc := parse(string(_body))

	productList := traverse(doc, "ul", "class", "product_list")
	// fmt.Println(productList)

	for item := productList.FirstChild; item != nil; item = item.NextSibling {
		// item := traverse(items, "li", "class", "ajax_block_product")
		// fmt.Println(item)

		_product := _Response{}

		imageTitleLink := traverse(item, "a", "class", "product_img_link")
		// fmt.Println(imageTitleLink)

		link, _ := attr(imageTitleLink.Attr, "href")
		if _debug {
			fmt.Println(link)
		}
		_product.link = link

		title, _ := attr(imageTitleLink.Attr, "title")
		if _debug {
			fmt.Println(title)
		}
		_product.title = title

		if Skip(title) {
			continue
		}

		code := strings.Split(link[45:], "-")[0]
		if _debug {
			fmt.Println(code)
		}
		_product.code = code

		if itemPrice := traverse(item, "span", "class", "price"); itemPrice != nil {
			// fmt.Println(itemPrice)

			price, _ := text(itemPrice)
			price = strings.ReplaceAll(strings.ReplaceAll(price, " CHF", ""), ",", ".")
			if _debug {
				fmt.Println(price)
			}

			if _price, err := strconv.ParseFloat(price, 32); err != nil {
				panic(err)
			} else {
				_product.price = float32(_price)
			}
		}

		if itemOldPrice := traverse(item, "span", "class", "old-price"); itemOldPrice != nil {
			// fmt.Println(itemOldPrice)

			price, _ := text(itemOldPrice)
			price = strings.ReplaceAll(strings.ReplaceAll(price, " CHF", ""), ",", ".")
			if _debug {
				fmt.Println(price)
			}

			if _price, err := strconv.ParseFloat(price, 32); err != nil {
				panic(err)
			} else {
				_product.oldPrice = float32(_price)
			}
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

			_retailPrice := product.price
			_price := _retailPrice
			if product.oldPrice > 0 {
				_retailPrice = product.oldPrice
			}
			_savings := _price - _retailPrice
			_discount := 100 - ((100 / _retailPrice) * _price)

			_title := product.title
			_link := s.ResolveURL(product.link).String()

			product := &Product{
				Code: _name + "//" + product.code,
				Name: _title,

				RetailPrice: _retailPrice,
				Price:       _price,
				Savings:     _savings,
				Discount:    _discount,

				Quantity: product.quantity,

				URL: _link,
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
