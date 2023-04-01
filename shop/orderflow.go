package shop

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

var OrderflowRegex = regexp.MustCompile(`\s+\(?(\d\+)?\d+\s*GB?|\s+\(?\d+(\.\d+)?"|\s+\(?20[12]\d\)?|\s+\(?[2345]G\)?| Dual SIM| Blau| GREEN| Schwarz|(EE )?Enterprise Edition( CH)?`)

var OrderflowCleanFn = func(name string) string {
	name = strings.NewReplacer(" 4G ", " ", " 3. Gen.", " 3rd Gen").Replace(name)

	if loc := OrderflowRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return strings.TrimSpace(name)
}

func XXX_orderflow(isDryRun bool) IShop {
	const _name = "orderflow"
	const _url = "https://www.orderflow.ch/de/categories/elektronik/kommunikation/mobiltelefone?limit=100&sort=price|asc&c300101=[30010104,30010117,30010105,30010103]"

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

		quantity int
	}

	var _result []_Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	fn := "shop/orderflow.html"

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

	doc := parse(string(_body))

	productList := traverse(doc, "div", "class", "product-list-items")
	// fmt.Println(productList)

	for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling.NextSibling {
		// item := traverse(items, "div", "class", "item")
		// fmt.Println(item)

		if !contains(item.Attr, "class", "item") {
			continue
		}

		_product := _Response{}

		imageTitleLink := traverse(item, "a", "class", "")
		// fmt.Println(imageTitleLink)

		link, _ := attr(imageTitleLink.Attr, "href")
		if _debug {
			fmt.Println(link)
		}
		_product.link = link

		itemImage := traverse(item, "img", "class", "img-fluid")
		// fmt.Println(itemImage)

		title, _ := attr(itemImage.Attr, "alt")
		title = strings.Split(strings.Split(strings.Split(title, " - ")[0], " 16.")[0], " (")[0]
		if _debug {
			fmt.Println(title)
		}
		_product.title = title

		if Skip(title) {
			continue
		}

		model := OrderflowCleanFn(html.UnescapeString(_product.title))
		if _debug {
			fmt.Println(model)
		}
		_product.model = model

		code := strings.Split(link[54:], "-")[0]
		if _debug {
			fmt.Println(code)
		}
		_product.code = code

		itemFirstPrice := traverse(item, "span", "class", "first_price")
		// fmt.Println(itemFirstPrice)

		if itemOldPrice := traverse(itemFirstPrice, "span", "class", "price"); itemOldPrice != nil {
			// fmt.Println(itemOldPrice)

			price, _ := text(itemOldPrice)
			if _debug {
				fmt.Println(price)
			}

			if _price, err := strconv.ParseFloat(price, 32); err != nil {
				panic(err)
			} else {
				_product.price = float32(_price)
			}
		}

		itemSecondPrice := traverse(item, "span", "class", "second_price")
		// fmt.Println(itemSecondPrice)

		if currentPrice := traverse(itemSecondPrice, "span", "class", "price"); currentPrice != nil {
			// fmt.Println(currentPrice)

			oldPrice, _ := text(currentPrice)
			if _debug {
				fmt.Println(oldPrice)
			}

			if _price, err := strconv.ParseFloat(oldPrice, 32); err != nil {
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

			_title := html.UnescapeString(product.title)
			_model := product.model

			if Skip(_model) {
				continue
			}

			if _tests {
				testCases[_title] = _model
			}

			_retailPrice := product.price
			_price := _retailPrice
			if product.oldPrice > 0 {
				_retailPrice = product.oldPrice
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

				Quantity: product.quantity,

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
