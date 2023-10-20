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

var FolettiRegex = regexp.MustCompile(`\s*[-,]+\s+|\s*\(?(\d+(\s*GB)?[+/])?\d+\s*GB\)?|\s*\d+G|(2|4|6|8|12)\/(64|128|256?B?)(GB)?|\s+\(?20[12]\d\)?|\s*\d+([,.]\d+)?\s*(cm|\")|\d{4,5}\s*mAh|\s+20[12]\d|\s+(Hybrid|Dual\W(SIM|Sim)|(EE )?Enterprise( Edition)?( CH)?|LTE|NFC|smartphone|Ice|Black|(Ocean )?Blue|Charcoal|Dark Green|Dusk|Grey|HIMALAYA GREY|Light|Glowing Black|Glowing Green|Graphite Gray|Midnight Black|Mint Green|Night|Polar White|Prism Black|Prism Blue|astro black|atlantic green|bamboo green|black onyx|blau|blue|charcoal grey|cosmic black|denim black|electric graphite|elegant black|frosted grey|glowing blue|graphite grey|(gravity )?grau|ic[ey] blue|lake blue|lavender blue|matte charcoal|metallic rose|meteor black|meteorite black|meteorite grey|midnight blue|mint green|night|ocean blue|onyx gray|petrol|polar blue|sage|sandy|sky blue|stargaze white|steel blue|steel grey|sterling blue|sunburst gold|titan\/?silber|(dark )?titanium grey|schwarz|inkl\.)`)
var FolettiExclusionRegex = regexp.MustCompile(`(?i)Abdeckung|Adapter|AirTag|Armband|Band|CABLE|Charger|Ch?inch|Christbaum|Clamshell|^Core|\bCover\b|Earphones|Etui|Fernauslöser|Gimbal|Halterung|Handschuhe|HARDCASE|Headset|Hülle|Kopfhörer|Ladegerät|Ladestation|Lautsprecher|Magnet|Majestic|Näh(faden|garn)|Netzkabel|Objektiv|Reiselader|S Pen|Saugnapf|Schutzfolie|Schutzglas|SmartTag|Stand|Ständer|Stativ|Stick|Stylus|Tastatur|Virtual-Reality|Wasserdicht(es)?|Weihnachtsbaum`)

var FolettiCleanFn = func(name string) string {
	if loc := FolettiRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	// name = strings.ReplaceAll(strings.ReplaceAll(name, " Phones ", " "), " Mini iPhone", " Mini")
	name = regexp.MustCompile(` XT\d{4}-\d+|SMARTPHONE\s*|Smartfon\s*|Solutions |TIM | Mobility Motorola| Mobility| Outdoor| NE| EE|o2-Aktion `).ReplaceAllString(name, "")

	s := strings.Split(name, " ")

	if s[0] == "Nothing" {
		name = strings.ReplaceAll(name, " Phones", "")
	}

	if s[0] == "Xiaomi" {
		name = strings.ReplaceAll(name, "23021RAA2Y", "Redmi Note 12")
	}

	if s[0] == "Samsung" {
		if part := regexp.MustCompile(`\(?\s*(SM-)?[AGMS]\d{3}[A-Z]*(\/DSN?)?\)?`).FindString(name); len(part) > 0 {
			name = strings.ReplaceAll(name, part, "")

			model := regexp.MustCompile(`[AMS]\d{2}`).FindString(part)

			if !strings.Contains(name, model) {
				name = name + " " + model
			}
		}
	}

	return helpers.Lint(name)
}

func XXX_foletti(isDryRun bool) IShop {
	const _name = "Foletti"
	// const _url = "https://superstore.foletti.com/de/categories/it--multimedia/telekommunikation/mobiltelefone/smartphone?limit=100&sort=price|asc&listStyle=list"
	_url := fmt.Sprintf("https://superstore.foletti.com/de/categories/it--multimedia/telekommunikation/mobiltelefone/smartphone?price-min=%.f&price-max=%.f&limit=100&sort=price|asc&listStyle=list&page=%%d", ValueMinimum, ValueMaximum)

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

	for p := 1; p <= 10; p++ {
		fn := fmt.Sprintf("shop/foletti.%d.html", p)

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

		doc := parse(string(_body))

		if productList := traverse(doc, "div", "class", "product-list-items"); productList != nil {
			// fmt.Println(productList)

			for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling.NextSibling {
				// item := traverse(items, "li", "class", "productList__item")
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
				title = strings.Split(strings.Split(title, " - ")[0], " 16.")[0]
				if _debug {
					fmt.Println(title)
				}
				_product.title = title

				if strings.Contains(title, "Wallet") || strings.Contains(title, "WALLET") {
					continue
				}
				if strings.Contains(title, "Tasche") {
					continue
				}
				if Skip(title) {
					continue
				}

				itemBrand := traverse(item, "strong", "", "")
				brand, _ := text(itemBrand)
				if _debug {
					fmt.Println(brand)
				}
				if Skip(brand) {
					continue
				}

				model := FolettiCleanFn(_product.title)
				if _debug {
					fmt.Println(model)
				}
				_product.model = model

				if brand != "o2" && !(brand == "Huawei" && strings.HasPrefix(title, "Honor")) {
					if !strings.EqualFold(strings.ToUpper(strings.Split(brand, " ")[0]), strings.ToUpper(strings.Split(title, " ")[0])) {
						_product.title = strings.ReplaceAll(brand, " Mobility", "") + " " + _product.title
					}
					if !strings.EqualFold(strings.ToUpper(strings.Split(brand, " ")[0]), strings.ToUpper(strings.Split(model, " ")[0])) {
						_product.model = strings.ReplaceAll(brand, " Mobility", "") + " " + _product.model
					}
				}

				if FolettiExclusionRegex.MatchString(model) {
					if _debug {
						fmt.Println()
					}

					continue
				}

				if _tests {
					testCases[_product.title] = _product.model
				}

				itemAvailability := traverse(item, "span", "class", "text")
				// fmt.Println(itemAvailability)

				amount, _ := text(itemAvailability)
				amount = strings.Trim(strings.Split(amount, " ")[0], ">")
				if _debug {
					fmt.Println(amount)
				}
				if amount == "Liefertermin" {
					continue
				}
				if _amount, err := strconv.Atoi(amount); err != nil {
					panic(err)
				} else {
					_product.quantity = _amount
				}

				itemArticle := traverse(item, "div", "class", "mpn")
				// fmt.Println(itemArticle)

				itemValue := traverse(itemArticle, "span", "class", "value")
				// fmt.Println(itemValue)

				sku, _ := text(itemValue)
				if _debug {
					fmt.Println(sku)
				}
				_product.code = sku

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

				// if itemBadge := traverse(item, "span", "class", "badge"); itemBadge != nil {
				// 	fmt.Println(itemBadge)

				// 	// badge, _ := text(itemBadge)
				// 	// panic(badge)
				// }

				if _debug {
					fmt.Println()
				}

				_result = append(_result, _product)
			}

			results := traverse(traverse(doc, "div", "class", "listing-number-of-results"), "em", "", "").NextSibling.NextSibling
			if result, ok := text(results); ok {
				current, _ := strconv.Atoi(result)

				results = results.NextSibling.NextSibling
				if result, ok := text(results); ok {
					total, _ := strconv.Atoi(result)

					if current >= total {
						break
					}
				}
			}

			// results := traverse(doc, "div", "class", "ps-3")
			// if result, ok := text(results); ok {
			// 	if x := regexp.MustCompile(`(\d+)‐(\d+) \/ (\d+)`).FindStringSubmatch(result); x != nil && x[2] == x[3] {
			// 		break
			// 	}
			// }
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

			_retailPrice := product.price
			_price := _retailPrice
			if product.oldPrice > 0 {
				_retailPrice = product.oldPrice
			}

			_savings := _price - _retailPrice
			_discount := 100 - ((_price * 100) / _retailPrice)

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

			// fmt.Printf("%#v\n", product)
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
