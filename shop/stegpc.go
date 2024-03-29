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

	"github.com/sugawarayuuta/sonnet"

	helpers "jsapi-scraper/helpers"
)

var StegRegex = regexp.MustCompile(` - |\s+\(?(\d+/)?\d+\s*[GT]B|\s+\(?\d+(\.\d+)?"|\s+\(?20[12]\d\)?|\s+\(?[2345]G\)?| Dual SIM| Vegan Leather|\s+CH$`)

var StegCleanFn = func(name string) string {
	name = strings.NewReplacer("Enterprise Edition", "EE").Replace(name)

	name = regexp.MustCompile(`^Renewd | \(?(SM-)?[AGMS]\d{3}[A-Z]*(/DSN)?\)?| XT\d{4}-\d+`).ReplaceAllString(name, "")

	if loc := StegRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	name = strings.NewReplacer(" E e", " e", " E ", " E", " G ", " G", "Samsung Samsung ", "Samsung ", "Xiaomi Xiaomi ", "Xiaomi ").Replace(name)

	return helpers.Lint(name)
}

// POST https://www.steg-electronics.ch/de/Filter/SetCheckboxFilter
// application/x-www-form-urlencoded; charset=UTF-8
// [Android]	id=22917&sectionId=7145&mainId=29123&isRadioButton=False&productListCacheId=11853&productListType=0&queryString=
// [Apple iOS]	id=22918&sectionId=7145&mainId=29123&isRadioButton=False&productListCacheId=11853&productListType=0&queryString=

// POST https://www.steg-electronics.ch/de/Filter/SetSlidableFilter
// application/x-www-form-urlencoded; charset=UTF-8
// [Preis]		id=2&sectionId=0&minValue=50.00&maxValue=300.00&fromSetIndex=0&toSetIndex=475&productListCacheId=11853&productListType=0&queryString=

func XXX_stegpc(isDryRun bool) IShop {
	const _name = "Steg Electronics"
	const _url = "https://www.steg-electronics.ch/de/product/list/11853?sortKey=preisasc&smsc=200&p=%d"

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

	type _Body struct {
		NewProductList string `json:"newProductList"`
		PagesCount     int    `json:"pagesCount"`
	}

	var _result []_Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	for p := 1; p <= 10; p++ {
		fn := fmt.Sprintf("shop/stegpc.%d.json", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			url := fmt.Sprintf(_url, p)

			resp, err := http.Post(url, "application/json", bytes.NewBuffer(nil))
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, url)
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
		// fmt.Println(BytesToString(_body))

		var body _Body
		if err := sonnet.Unmarshal(_body, &body); err != nil { // Parse []byte to go struct pointer
			panic(err)
		}
		// fmt.Println(body.NewProductList)

		doc := parse(body.NewProductList)

		if productList := traverse(doc, "article", "class", "product-element"); productList != nil {
			// fmt.Println(productList)

			for item := productList; item != nil; item = item.NextSibling.NextSibling {
				// fmt.Println(item)

				_product := _Response{}

				productId, _ := attr(item.Attr, "data-product-id")
				if _debug {
					fmt.Println(productId)
				}
				_product.code = productId

				// percentage := traverse(item, "div", "class", "percentage")
				// // fmt.Println(percentage)

				// discount, _ := text(percentage)
				// discount = strings.TrimSpace(discount)
				// if _debug {
				// 	fmt.Println(discount)
				// }
				// _product.discount = discount

				imageTitleLink := traverse(item, "a", "class", "link-detail")
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

				model := StegCleanFn(_product.title)
				if _debug {
					fmt.Println(model)
				}
				_product.model = model

				currentPrice := traverse(item, "div", "class", "generalPrice")
				// fmt.Println(currentPrice)

				price, _ := text(currentPrice)
				if _debug {
					fmt.Println(price)
				}

				if _price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(price, ".-", ".00"), "'", ""), 32); err != nil {
					panic(err)
				} else {
					_product.oldPrice = float32(_price)
				}

				if insteadPrice := traverse(item, "div", "class", "insteadPrice"); insteadPrice != nil {
					// fmt.Println(insteadPrice)

					itemText := traverse(insteadPrice, "text", "", "")

					price, _ := text(itemText)
					price = strings.TrimSpace(strings.ReplaceAll(price, "statt", ""))
					if _debug {
						fmt.Println(price)
					}

					if _price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(price, ".-", ".00"), "'", ""), 32); err != nil {
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

		if body.PagesCount <= p {
			break
		}
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result))
		for _, _product := range _result {
			// fmt.Println(_product)

			_title := _product.title
			_model := _product.model

			if Skip(_model) {
				continue
			}

			if _tests {
				testCases[_title] = _model
			}

			_retailPrice := _product.oldPrice
			_price := _retailPrice
			if _product.price > 0 {
				_price = _product.price
			}

			_savings := _price - _retailPrice
			_discount := 100 - ((100 / _retailPrice) * _price)

			product := &Product{
				Code:  _name + "//" + _product.code,
				Name:  _title,
				Model: _model,

				RetailPrice: _retailPrice,
				Price:       _price,
				Savings:     _savings,
				Discount:    _discount,

				URL: _product.link,
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
