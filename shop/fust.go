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
)

// https://www.fust.ch/de/r/pc-tablet-handy/smartphone-145.html?shop_comparatorkey=9-1&shop_nrofrecs=12&brand=Fairphone%7CGoogle%7CHuawei%7CMotorola%7CNokia%7CNothing%20Phones%7COnePlus%7COppo%7CRealme%7CSamsung%7CXiaomi
// https://www.fust.ch/de/r/pc-tablet-handy/smartphone/samsung-galaxy-789.html?shop_comparatorkey=9-1&shop_nrofrecs=12
// https://www.fust.ch/de/r/pc-tablet-handy/smartphone/huawei-smartphone-809.html?shop_comparatorkey=9-1&shop_nrofrecs=12
// https://www.fust.ch/de/r/pc-tablet-handy/smartphone/xiaomi-smartphone-808.html?shop_comparatorkey=9-1&shop_nrofrecs=12
// https://www.fust.ch/de/r/pc-tablet-handy/smartphone/oppo-smartphone-1010.html?shop_comparatorkey=9-1&shop_nrofrecs=12
// https://www.fust.ch/de/r/pc-tablet-handy/smartphone/weitere-smartphones-und-handy-366.html?shop_comparatorkey=9-1&shop_nrofrecs=12
// https://www.fust.ch/de/r/pc-tablet-handy/smartphone/apple-iphone-530.html?shop_comparatorkey=7-1&shop_nrofrecs=12&brand=Apple&ff1878=4G%20%2F%20LTE%7C5G&price=%7B%22from%22%3A399%2C%22to%22%3A500%7D&showAllFacets=true

var FustRegex = regexp.MustCompile(`(\s*[-,–]\s+)|(\d+\s*GB?)\b|\s+((EE )?Enterprise Edition( CH)?|Arctic Bleen|Astral|Awesome|Black|(New )?(Blk|Slv)|Champagne|Charcoal|Cloudy|Cosmo|Blue|Frost|Galactic|Green|Grey|Ice|Marine|Midnight|Moonlight|Ocean|Pepper Grey|Pink|Shadow|Space|Starlight|Sunset|Titan|White|black|cosmic|gold|schwarz|starry|c\.teal|e\.graphite|n\.blue|CH)`)

var FustCleanFn = func(name string) string {
	if loc := FustRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	name = strings.ReplaceAll(name, "Samsung Speicherkarte + ", "")
	name = strings.ReplaceAll(name, "Speicherkarte + ", "")
	name = strings.ReplaceAll(name, "Phones ", "")

	s := strings.Split(name, " ")

	if s[0] == "Fust" {
		name = strings.ReplaceAll(name, "Fust ", "Inoi ")
	}

	if s[0] == "Reno" || s[0] == "Oppo" {
		name = regexp.MustCompile(`Reno\s*(\d)\s*(\w)?`).ReplaceAllString(name, "Reno$1 $2")
	}

	if s[0] == "Apple" {
		name = strings.NewReplacer(" 2020", " (2020)", " 2022", " (2022)", " 2nd Gen", " (2020)", " 3rd Gen", " (2022)").Replace(name)
	} else {
		// Remove year component for all other than Apple.
		name = regexp.MustCompile(`\s+\(?20[12]\d\)?`).ReplaceAllString(name, "")
	}

	return strings.TrimSpace(name)
}

func XXX_fust(isDryRun bool) IShop {
	const _name = "Fust"
	const _url = "https://www.fust.ch/de/r/pc-tablet-handy/smartphone-145.html?shop_comparatorkey=9-1&shop_nrofrecs=12&brand=Fairphone%7CGoogle%7CHuawei%7CMotorola%7CNokia%7CNothing%20Phones%7COnePlus%7COppo%7CRealme%7CSamsung%7CXiaomi"

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

	for _category, _url := range map[string]string{
		"apple":   fmt.Sprintf(`https://www.fust.ch/de/r/pc-tablet-handy/smartphone/apple-iphone-530.html?shop_comparatorkey=9-1&shop_nrofrecs=60&brand=Apple&ff1878=4G%%%%20%%%%2F%%%%20LTE%%%%7C5G&price={"from":%.f,"to":%.f}&shop_recpage=%%d`, ValueMinimum, ValueMaximum),
		"samsung": fmt.Sprintf(`https://www.fust.ch/de/r/pc-tablet-handy/smartphone/samsung-galaxy-789.html?shop_comparatorkey=9-1&shop_nrofrecs=60&price={"from":%.f,"to":%.f}&shop_recpage=%%d`, ValueMinimum, ValueMaximum),
		"huawei":  fmt.Sprintf(`https://www.fust.ch/de/r/pc-tablet-handy/smartphone/huawei-smartphone-809.html?shop_comparatorkey=9-1&shop_nrofrecs=60&price={"from":%.f,"to":%.f}&shop_recpage=%%d`, ValueMinimum, ValueMaximum),
		"xiaomi":  fmt.Sprintf(`https://www.fust.ch/de/r/pc-tablet-handy/smartphone/xiaomi-smartphone-808.html?shop_comparatorkey=9-1&shop_nrofrecs=60&price={"from":%.f,"to":%.f}&shop_recpage=%%d`, ValueMinimum, ValueMaximum),
		"oppo":    fmt.Sprintf(`https://www.fust.ch/de/r/pc-tablet-handy/smartphone/oppo-smartphone-1010.html?shop_comparatorkey=9-1&shop_nrofrecs=60&price={"from":%.f,"to":%.f}&shop_recpage=%%d`, ValueMinimum, ValueMaximum),
		"other":   fmt.Sprintf(`https://www.fust.ch/de/r/pc-tablet-handy/smartphone/weitere-smartphones-und-handy-366.html?shop_comparatorkey=9-1&shop_nrofrecs=60&price={"from":%.f,"to":%.f}&shop_recpage=%%d`, ValueMinimum, ValueMaximum),
	} {
		for p := 1; p <= 3; p++ {
			fn := fmt.Sprintf("shop/fust.%s.%d.html", _category, p)

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

			if productList := traverse(doc, "ul", "id", "productslist"); productList != nil {
				// fmt.Println(productList)

				for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling.NextSibling {
					// item := traverse(productList, "li", "class", "product__list")
					// fmt.Println(item)

					if contains(item.Attr, "class", "rubric-advertising") {
						continue
					}

					if available := traverse(item, "div", "class", "stati2"); available == nil {
						if available := traverse(item, "div", "class", "stati3"); available == nil {
							continue
						}
					}

					_product := _Response{}

					productKey, _ := attr(item.Attr, "data-prj-productkey")
					// fmt.Println(productKey)
					productId := strings.Split(productKey, "_")[0]
					if _debug {
						fmt.Println(productId)
					}

					_product.code = productId

					product := traverse(item, "div", "class", "product")
					// fmt.Println(product)

					figure := traverse(product, "figure", "class", "product__overview-img")
					// fmt.Println(figure)

					imageTitleLink := traverse(figure, "a", "href", "")
					// fmt.Println(imageTitleLink)

					link, _ := attr(imageTitleLink.Attr, "href")
					if _debug {
						fmt.Println(link)
					}
					_product.link = link

					itemProducer := traverse(item, "span", "class", "producer")
					// fmt.Println(itemTitle)

					brand, _ := text(itemProducer)
					if _debug {
						fmt.Println(brand)
					}
					_product.title = brand

					itemTitle := traverse(imageTitleLink, "img", "class", "product__overview-img")
					// fmt.Println(itemTitle)

					if itemTitle == nil {
						continue
					}

					title, _ := attr(itemTitle.Attr, "title")
					title = strings.TrimSpace(title)
					if _debug {
						fmt.Println(title)
					}
					_product.title = brand + " " + title

					if Skip(_product.title) {
						continue
					}

					model := FustCleanFn(title)
					if _debug {
						fmt.Println(model)
					}
					if brand == "Samsung" && strings.Split(model, " ")[0] == "Motorola" {
						model = strings.ReplaceAll(model, "Motorola ", "")
						brand = "Motorola"
					} else if brand == "Fust" {
						brand = "Inoi"
					}
					_product.model = brand + " " + model

					itemPrice := traverse(item, "div", "class", "price-block")
					// fmt.Println(itemPrice)

					if itemOldPrice := traverse(itemPrice, "span", "class", "oldprice"); itemOldPrice != nil {
						// fmt.Println(itemOldPrice)

						oldPrice, _ := text(itemOldPrice)
						if _debug {
							fmt.Println(oldPrice)
						}

						if oldPrice := strings.ReplaceAll(strings.ReplaceAll(oldPrice, "–", "00"), "’", ""); oldPrice != "" {
							if _price, err := strconv.ParseFloat(oldPrice, 32); err != nil {
								panic(err)
							} else {
								_product.oldPrice = float32(_price)
							}
						}
					}

					currentPrice := traverse(itemPrice, "div", "class", "endprice")
					// fmt.Println(currentPrice)

					price, _ := text(currentPrice)
					if _debug {
						fmt.Println(price)
					}

					if price := strings.ReplaceAll(strings.ReplaceAll(price, "–", "00"), "’", ""); price != "" {
						if _price, err := strconv.ParseFloat(price, 32); err != nil {
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

			results := traverse(doc, "span", "class", "currentPaging")
			pageAmount := traverse(results, "span", "class", "js-pageAmount")
			if amount, ok := text(pageAmount); ok {
				if result, ok := text(results.FirstChild.NextSibling.NextSibling); ok {
					if x := regexp.MustCompile(`(\d+) von (\d+)`).FindStringSubmatch(amount + " " + result); x != nil && x[1] == x[2] {
						break
					}
				}
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

				// Quantity: product.Quantity,

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
