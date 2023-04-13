package shop

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var MiStoreRegex = regexp.MustCompile(`\s+(\d+/(GB)?)?\d+GB|\s+20[12]\d|\s+[2345]G`)

var MiStoreCleanFn = func(name string) string {
	if loc := MiStoreRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return strings.TrimSpace(name)
}

func XXX_mistore(isDryRun bool) IShop {
	const _name = "Mi Store"
	const _url = "https://mi-store.ch/produkt-kategorie/smartphones/?orderby=price"

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	type _Response struct {
		code  string
		title string
		model string

		link string

		oldPrice float32
		price    float32
	}

	var _result []_Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	fn := "shop/mi-store.html"

	if isDryRun {
		if body, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else {
			_body = body
		}
	} else {
		var formData = []byte(`mogo-woocommerce-products-per-page=96`)

		resp, err := http.Post(_url, "application/x-www-form-urlencoded", bytes.NewBuffer(formData))
		if err != nil {
			// panic(err)
			fmt.Println(err)
			return NewShop(
				_name,
				_url,

				nil,
			)
		}
		defer resp.Body.Close()

		if body, err := io.ReadAll(resp.Body); err != nil {
			// panic(err)
			fmt.Println(err)
			return NewShop(
				_name,
				_url,

				nil,
			)
		} else {
			_body = body
		}

		os.WriteFile(path+fn, _body, 0664)
	}
	// fmt.Println(string(_body))

	doc := parse(string(_body))

	if productList := traverse(doc, "div", "class", "tt-product-view"); productList != nil {
		// fmt.Println(productList)

		for item := productList.FirstChild; /*.NextSibling*/ item != nil; item = item.NextSibling.NextSibling {
			// fmt.Println(item)

			if contains(item.Attr, "class", "outofstock") {
				continue
			}

			_product := _Response{}

			buttons := traverse(item, "div", "class", "tt-product__buttons")
			// fmt.Println(buttons)

			buttonsCart := traverse(buttons, "a", "class", "tt-product__buttons_cart")
			// fmt.Println(buttonsCart)

			productId, _ := attr(buttonsCart.Attr, "data-product_id")
			if _debug {
				fmt.Println(productId)
			}
			_product.code = productId

			content := traverse(item, "div", "class", "tt-product__content")
			// fmt.Println(content)

			contentLink := traverse(content, "a", "href", "")
			// fmt.Println(contentLink)

			link, _ := attr(contentLink.Attr, "href")
			if _debug {
				fmt.Println(link)
			}
			_product.link = link

			title, _ := text(contentLink)
			if _debug {
				fmt.Println(title)
			}
			if s := strings.Split(title, " "); len(s) > 0 {
				if strings.ToUpper(s[0]) != "XIAOMI" {
					title = "Xiaomi " + title
				}
			}
			_product.title = title

			if Skip(title) {
				continue
			}

			model := MiStoreCleanFn(_product.title)
			if _debug {
				fmt.Println(model)
			}
			_product.model = model

			if _tests {
				testCases[_product.title] = _product.model
			}

			productPrice := traverse(item, "span", "class", "tt-product__price")
			// fmt.Println(productPrice)

			productAmount := traverse(productPrice, "span", "class", "amount")
			// fmt.Println(productAmount)

			price, _ := text(productAmount.FirstChild.NextSibling)
			if _debug {
				fmt.Println(price)
			}

			if _price, err := strconv.ParseFloat(strings.ReplaceAll(price, "'", ""), 32); err != nil {
				panic(err)
			} else {
				_product.oldPrice = float32(_price)
			}

			if ins := traverse(productPrice, "ins", "", ""); ins != nil {
				productAmount := traverse(ins, "span", "class", "amount")
				// fmt.Println(productAmount)

				price, _ := text(productAmount.FirstChild.NextSibling)
				if _debug {
					fmt.Println(price)
				}

				if _price, err := strconv.ParseFloat(strings.ReplaceAll(price, "'", ""), 32); err != nil {
					panic(err)
				} else {
					_product.price = float32(_price)
				}
			}

			if _debug {
				fmt.Println()
			}

			_result = append(_result, _product)
		}
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result))
		for _, product := range _result {
			// fmt.Println(product)

			_title := product.title
			_model := product.model

			if Skip(_model) {
				continue
			}

			_retailPrice := product.oldPrice
			_price := _retailPrice
			if product.price > 0 {
				_price = product.price
			}
			_savings := _price - _retailPrice
			_discount := 100 - ((100 / _retailPrice) * _price)

			_link := s.ResolveURL(product.link).String()

			product := &Product{
				Code:  _name + "//" + product.code,
				Name:  _title,
				Model: _model,

				RetailPrice: _retailPrice,
				Price:       _price,
				Savings:     _savings,
				Discount:    _discount,

				URL: _link,
			}

			if s.IsWorth(product) {
				products = append(products, product)
			}
		}

		if _tests {
			keys := make([]string, 0, len(testCases))

			for k := range testCases {
				keys = append(keys, k)
			}
			sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

			for _, k := range keys {
				fmt.Println("\"" + strings.ReplaceAll(k, "\"", "\\\"") + "\",")
			}
			fmt.Println()
			for _, k := range keys {
				fmt.Println("\"" + strings.ReplaceAll(testCases[k], "\"", "\\\"") + "\",")
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
