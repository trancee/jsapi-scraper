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

	helpers "jsapi-scraper/helpers"
)

var CashConvertersRegex = regexp.MustCompile(`(?i),|(\d+\/)?(2|4|6|8|16|32|64|128|256)\s*(GB|BG|GG|Go|G|B)|\(?[345]G\)?|NFC|LTE|Dual[- ]SIM|\+? Boîte|\((Black|Bleu Azur|Gris|Noir)\)`)

var CashConvertersCleanFn = func(name string) string {
	name = strings.NewReplacer(" - ", " ").Replace(name)
	name = regexp.MustCompile(`(?i)Portable|Reconditionné|Rouge|Téléphone(\s*:\s*)?`).ReplaceAllString(name, "")
	name = strings.TrimSpace(name)

	if loc := CashConvertersRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	s := strings.Split(name, " ")

	if s[0] == "iPhone" {
		name = strings.ReplaceAll(name, "3rd", "2022")
		name = strings.ReplaceAll(name, "2022 2022", "2022")
	}

	if s[0] == "Oppo" || s[0] == "OPPO" {
		name = strings.ReplaceAll(name, "Fond", "Find")
	}

	if s[0] == "Samsung" || s[0] == "SAMSUNG" {
		if s[1] == "Note20" {
			name = strings.ReplaceAll(name, "Note20", "Galaxy Note 20")
		}
		if s[1] == "Samsung" {
			name = strings.ReplaceAll(name, "Samsung Samsung", "Samsung")
		}

		name = strings.NewReplacer("S10PLUS", "S10 Plus", "S20FE", "S20 FE", "AO2S", "A02s", "S10 +", "S10+").Replace(name)
	}

	if s[0] == "Xiaomi" {
		if s[1] == "Redmi" {
			if s[2] == "Mi" {
				strings.ReplaceAll(name, "Redmi ", "")
			}
		}
	}

	return helpers.Lint(name)
}

func XXX_cashconverters(isDryRun bool) IShop {
	const _name = "Cash Converters"
	_url := "https://cash-converters.ch/collections/telephones?page=%d&sort_by=price-ascending&view=32ajax"

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

	nomore := false

	for p := 1; p <= 20; p++ {
		fn := fmt.Sprintf("shop/cash-converters.%d.html", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			url := fmt.Sprintf(_url, p)

			resp, err := http.Get(url)
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
		// fmt.Println(string(_body))

		doc := parse(string(_body))

		if productList := traverse(doc, "div", "class", "tt-product-listing"); productList != nil {
			// fmt.Println(productList)

			for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling {
				// fmt.Println(item)

				itemData := traverse(item, "a", "", "")

				code, _ := attr(itemData.Attr, "data-firstavavariantid")
				if _debug {
					fmt.Println(code)
				}

				quantity, _ := attr(itemData.Attr, "data-quantity")
				quantity = strings.ReplaceAll(quantity[:len(quantity)-1], code, "")[1:]
				if quantity, err := strconv.Atoi(quantity); err != nil {
					panic(err)
				} else if quantity == 0 {
					continue
				}

				_product := _Response{code: code}

				itemDescription := traverse(item, "div", "class", "tt-description")

				itemLink := traverse(itemDescription, "a", "", "")

				link, _ := attr(itemLink.Attr, "href")
				if _debug {
					fmt.Println(link)
				}
				_product.link = link

				title, _ := text(itemLink)
				if _debug {
					fmt.Println(title)
				}
				_product.title = title

				if Skip(title) {
					continue
				}

				model := CashConvertersCleanFn(title)
				if _debug {
					fmt.Println(model)
				}
				_product.model = model

				itemPrice := traverse(itemDescription, "div", "class", "tt-price")
				// fmt.Println(itemPrice)

				price, _ := text(itemPrice.FirstChild)
				price = strings.ReplaceAll(strings.TrimSpace(strings.TrimPrefix(price, "CHF")), ",", "")
				if _debug {
					fmt.Println(price)
				}

				if _price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(price, ".-", ".00"), "'", ""), 32); err != nil {
					panic(err)
				} else {
					_product.price = float32(_price)
				}

				if oldPrice, _ := text(itemPrice.FirstChild.NextSibling); oldPrice != "" {
					oldPrice = strings.ReplaceAll(strings.TrimSpace(strings.TrimPrefix(oldPrice, "CHF")), ",", "")
					if _debug {
						fmt.Println(oldPrice)
					}

					if _oldPrice, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(oldPrice, ".-", ".00"), "'", ""), 32); err != nil {
						panic(err)
					} else {
						_product.oldPrice = float32(_oldPrice)
					}
				}

				if _debug {
					fmt.Println()
				}

				nomore = _product.price > ValueMaximum

				_result = append(_result, _product)
			}

			if nomore {
				break
			}
			if nomore := traverse(doc, "div", "class", "tt_item_all_js"); nomore != nil {
				break
			}
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

			_retailPrice := _product.price
			_price := _retailPrice
			if _product.oldPrice > 0 {
				_retailPrice = _product.oldPrice
			}

			_savings := _price - _retailPrice
			_discount := 100 - ((_price * 100) / _retailPrice)

			_link := s.ResolveURL(_product.link).String()

			product := &Product{
				Code:  _name + "//" + _product.code,
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
