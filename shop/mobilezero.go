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

var MobilezeroRegex = regexp.MustCompile(`(?i),| - |\([^\d]|(\d+/)?(2|4|6|8|16|32|64|128|256)GB|\(?[345]G\)?|Dual(-SIM)?| DS| EU$`)
var MobilezeroExclusionRegex = regexp.MustCompile(`(?i)Adapter|AirTag|Armband|Band|CABLE|Charger|Ch?inch|Christbaum|^Core|\bCover\b|Earphones|Etui|Halterung|Hülle|Kopfhörer|Ladegerät|Ladestation|Magnet|Netzkabel|Objektiv|Reiselader|S Pen|Saugnapf|Schutzfolie|SmartTag|Stand|Ständer|Stativ|Stylus|Virtual-Reality|Wasserdicht(es)?|Weihnachtsbaum`)

var MobilezeroCleanFn = func(name string) string {
	name = strings.NewReplacer("Appel", "Apple", "Blackshark", "Black Shark", "Motorla", "Motorola", "Enterprise Edition", "EE", "Enterprise Editon", "EE").Replace(name)
	name = regexp.MustCompile(`(?i)\b(Black|Blau|Gold|Graphite Grey|Grau|Pale Grey|Pink)\b`).ReplaceAllString(name, "")
	name = strings.TrimSpace(name)

	if loc := MobilezeroRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	s := strings.Split(name, " ")

	if s[0] == "Blackview" {
		name = strings.ReplaceAll(name, " 4900S", " BV4900s")
	}

	if s[0] == "Huawei" {
		name = regexp.MustCompile(`(?i)\((\d{4})\)`).ReplaceAllString(name, "$1")
	}

	if s[0] == "Motorola" {
		name = strings.ReplaceAll(name, "Razr22", "Razr 2022")
	}

	return helpers.Lint(name)
}

func XXX_mobilezero(isDryRun bool) IShop {
	const _name = "mobilezero"
	const _count = 500
	_url := fmt.Sprintf("https://www.mobilezero.ch/ajax/ProductLoader.aspx?mode=category&orderby=priceup&filterid=1549&searchterm=&skip=%%d&count=%d&languageId=2055&languageCode=de", _count)

	const _debug = true
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

	for p := 1; p <= 10; p++ {
		fn := fmt.Sprintf("shop/mobilezero.%d.html", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				break
			} else {
				_body = body
			}
		} else {
			url := fmt.Sprintf(_url, (p-1)*_count)

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

			if len(_body) == 0 {
				break
			}

			os.WriteFile(path+fn, _body, 0664)
		}
		// fmt.Println(BytesToString(_body))

		doc := parse(BytesToString(_body))

		if productList := traverse(doc, "article", "class", "shop-product-list"); productList != nil {
			// fmt.Println(productList)

			for item := productList; item != nil; item = item.NextSibling.NextSibling.NextSibling.NextSibling {
				// fmt.Println(item)

				itemAvailability := traverse(item, "span", "class", "icon-shop")
				// fmt.Println(itemAvailability)

				if contains(itemAvailability.Attr, "class", "icon-red") {
					// Skip if item out of stock
					continue
				}
				if contains(itemAvailability.Attr, "class", "icon-orange") && !contains(itemAvailability.Attr, "title", "Ab Fremdlager verfügbar") {
					// Skip if item not available
					continue
				}

				_product := _Response{}

				code, _ := attr(item.Attr, "data-productid")
				if _debug {
					fmt.Println(code)
				}
				_product.code = code

				itemLink := traverse(item, "a", "", "")
				// fmt.Println(itemLink)

				link, _ := attr(itemLink.Attr, "href")
				if _debug {
					fmt.Println(link)
				}
				_product.link = link

				itemTitle := traverse(item, "h3", "", "")
				// fmt.Println(itemTitle)

				title, _ := text(itemTitle)
				if _debug {
					fmt.Println(title)
				}
				_product.title = title

				if Skip(title) {
					continue
				}

				if MobilezeroExclusionRegex.MatchString(title) {
					continue
				}

				model := MobilezeroCleanFn(title)
				if _debug {
					fmt.Println(model)
				}
				_product.model = model

				itemPrice := itemTitle.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild
				// fmt.Println(itemPrice)

				if itemPrice.FirstChild.Data == "span" {
					itemPrice = itemPrice.FirstChild
				}

				price, _ := text(itemPrice)
				if _debug {
					fmt.Println(price)
				}

				price = strings.TrimSpace(strings.NewReplacer("CHF", "", ".-", ".00", "'", "", "Aktion", "").Replace(price))

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
		strings.ReplaceAll(_url, "/page/%d", ""),

		_parseFn,
	)
}
