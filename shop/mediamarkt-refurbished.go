package shop

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	helpers "jsapi-scraper/helpers"
)

var MediamarktRefurbishedRegex = regexp.MustCompile(`\s+\d+\s*GB?|\s+[2345]G|\s+\(?(mono|dual) sim\)?`)

var MediamarktRefurbishedCleanFn = func(name string) string {
	if loc := MediamarktRefurbishedRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)

	// s := strings.Split(name, " ")

	// if s[0] == "Samsung" {
	// 	name = regexp.MustCompile(`Note\s*(\d+)`).ReplaceAllString(name, "Note $1")
	// }

	// if s[0] == "Apple" {
	// 	name = strings.NewReplacer(" 2020", " (2020)", " 2022", " (2022)", " 2nd Gen", " (2020)", " 3rd Gen", " (2022)").Replace(name)
	// } else {
	// 	// Remove year component for all other than Apple.
	// 	name = regexp.MustCompile(`\s+\(?20[12]\d\)?`).ReplaceAllString(name, "")
	// }

	// return strings.TrimSpace(name)
}

func XXX_mediamarkt_refurbished(isDryRun bool) IShop {
	const _name = "Mediamarkt (Refurbished)"
	const _url = "https://refurbished.mediamarkt.ch/ch_de/unsere-refurbished-smartphones?is_in_stock=1&product_list_order=price&product_list_limit=100"

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

	fn := "shop/mediamarkt-refurbished.html"

	if isDryRun {
		if body, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else {
			_body = body
		}
	} else {
		resp, err := http.Get(_url)
		if err != nil {
			// panic(err)
			fmt.Printf("[%s] %s (%s)\n", _name, err, _url)
			return NewShop(
				_name,
				_url,

				nil,
			)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			// panic(resp.StatusCode)
			fmt.Printf("[%s] %d: %s (%s)\n", _name, resp.StatusCode, resp.Status, resp.Request.URL)
			return NewShop(
				_name,
				_url,

				nil,
			)
		}

		if body, err := io.ReadAll(resp.Body); err != nil {
			// panic(err)
			fmt.Printf("[%s] %s (%s)\n", _name, err, resp.Request.URL)
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

	type _Product struct {
		Category string  `json:"category"`
		Name     string  `json:"name"`
		Price    float32 `json:"price"`
	}

	type _Products struct {
		Products map[string]_Product `json:"products"`
	}

	var _products _Products
	{
		r := regexp.MustCompile(`"products": {(.*?)},\n`)
		body := "{" + strings.TrimSuffix(string(r.Find(_body)), ",\n") + "}"
		// fmt.Println(body)

		if err := json.Unmarshal([]byte(body), &_products); err != nil {
			panic(err)
		}
		// fmt.Println(_products)
	}

	doc := parse(string(_body))

	if productList := traverse(doc, "ol", "class", "products"); productList != nil {
		// fmt.Println(productList)

		for item := productList.FirstChild; /*.NextSibling*/ item != nil; item = item.NextSibling.NextSibling {
			// item := traverse(productList, "li", "class", "product")
			// fmt.Println(item)

			_product := _Response{}

			if product := traverse(item, "div", "class", "product-item-details"); product != nil {
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
				// title = strings.TrimSpace(strings.Split(title, "(")[0])
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

				item := _products.Products[productId]
				_product.title = item.Category + " " + _product.title

				model := MediamarktRefurbishedCleanFn(_product.title)
				if _debug {
					fmt.Println(model)
				}
				_product.model = model

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

				if Skip(title) {
					continue
				}

				_result = append(_result, _product)
			}
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

			if _tests {
				testCases[_title] = _model
			}

			_retailPrice := product.oldPrice
			_price := _retailPrice
			if product.price > 0 {
				_price = product.price
			}
			_savings := _price - _retailPrice
			_discount := 100 - ((100 / _retailPrice) * _price)

			product := &Product{
				Code:  _name + "//" + product.code,
				Name:  _title,
				Model: _model,

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
