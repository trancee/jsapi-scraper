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

var BrackRegex = regexp.MustCompile(`(\s*[-,]\s+)|(\d+\s*GB?)|\s+20[12]\d|\s+((EE )?Enterprise Edition( CH)?)`)

var BrackCleanFn = func(name string) string {
	name = strings.ReplaceAll(name, " Phones ", " ")

	if loc := BrackRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return strings.TrimSpace(name)
}

func XXX_brack(isDryRun bool) IShop {
	const _name = "Brack"
	// const _url = "https://www.brack.ch/it-multimedia/telefonie-kommunikation/mobiltelefone/smartphone?filter%5BArt%5D%5B%5D=offer&filter%5BArt%5D%5B%5D=intropromotion&filter%5BArt%5D%5B%5D=occassion&filter%5BArt%5D%5B%5D=new&sortProducts=priceasc&query=*"
	// const _url = "https://www.brack.ch/it-multimedia/telefonie-kommunikation/mobiltelefone/smartphone?limit=192&sortProducts=priceasc&query=*"
	_url := fmt.Sprintf("https://www.brack.ch/it-multimedia/telefonie-kommunikation/mobiltelefone/smartphone?filter[availability][]=Verfügbar&filter[price_standard][]=%f~~~%f&sortProducts=priceasc&query=*", ValueMinimum, ValueMaximum)

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	type Callout int

	const (
		_ Callout = iota

		New
		Action
		Trade
		Sustainability
	)

	type _Response struct {
		code  string
		title string
		model string

		link string

		oldPrice float32
		price    float32

		quantity int

		callout Callout
	}

	var _result []_Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	fn := "shop/brack.html"

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

	doc := parse(string(_body))

	if productList := traverse(doc, "ul", "class", "productList"); productList != nil {
		// fmt.Println(productList)

		for item := productList.FirstChild; /*.NextSibling*/ item != nil; item = item.NextSibling /*.NextSibling*/ {
			// fmt.Println(item)
			if !contains(item.Attr, "class", "product-card") {
				continue
			}
			// item := traverse(items, "li", "class", "productList__item")
			// fmt.Println(item)

			_product := _Response{}

			sku, _ := attr(item.Attr, "data-sku")
			if _debug {
				fmt.Println(sku)
			}
			_product.code = sku

			itemAvailability := traverse(item, "span", "class", "stock__amount")
			// fmt.Println(itemManufacturer)

			amount, _ := attr(itemAvailability.Attr, "data-amount")
			if _debug {
				fmt.Println(amount)
			}
			if _amount, err := strconv.Atoi(amount); err != nil {
				panic(err)
			} else {
				_product.quantity = _amount
			}

			// itemManufacturer := traverse(item, "span", "class", "productList__itemManufacturer")
			// // fmt.Println(itemManufacturer)

			// manufacturer, _ := text(itemManufacturer)
			// // fmt.Println(manufacturer)

			itemImage := traverse(item, "img", "class", "productList__itemImage")
			// fmt.Println(itemImage)

			title, _ := attr(itemImage.Attr, "title")
			title = strings.ReplaceAll(strings.ReplaceAll(title, " - ", " "), "Fairphone Fairphone", "Fairphone")
			if _debug {
				fmt.Println(title)
			}
			_product.title = title

			if Skip(title) {
				continue
			}

			model := BrackCleanFn(title)
			if _debug {
				fmt.Println(model)
			}
			_product.model = model

			imageTitleLink := traverse(item, "a", "class", "product__imageTitleLink")
			// fmt.Println(imageTitleLink)

			link, _ := attr(imageTitleLink.Attr, "href")
			if _debug {
				fmt.Println(link)
			}
			_product.link = link

			if itemCallout := traverse(item, "span", "class", "callout__text"); itemCallout != nil {
				// fmt.Println(itemCallout)

				callout, _ := attr(itemCallout.Attr, "class")
				switch strings.ReplaceAll(callout, "callout__text callout__color", "") {
				case "New":
					_product.callout = New
				case "Action":
					_product.callout = Action
				case "Trade":
					_product.callout = Trade
				case "Sustainability":
					_product.callout = Sustainability
				default:
					panic(callout)
				}
			}

			// itemTitle := traverse(imageTitleLink, "span", "class", "productList__itemTitle")
			// // fmt.Println(itemTitle)

			// title, _ := text(itemTitle)
			// fmt.Println(manufacturer + " " + title)

			itemPrice := traverse(item, "div", "class", "productList__itemPrice")
			// fmt.Println(itemPrice)

			if itemOldPrice := traverse(itemPrice, "span", "class", "productList__itemOldPrice"); itemOldPrice != nil {
				// fmt.Println(itemOldPrice)

				oldPrice, _ := attr(itemOldPrice.Attr, "content")
				if _debug {
					fmt.Println(oldPrice)
				}

				if _price, err := strconv.ParseFloat(oldPrice, 32); err != nil {
					panic(err)
				} else {
					_product.oldPrice = float32(_price)
				}
			}

			// currentPrice := traverse(itemPrice, "div", "class", "currentPrice")
			// fmt.Println(currentPrice)

			if currentPrice := traverse(itemPrice, "em", "class", "js-currentPriceValue"); currentPrice != nil {
				// fmt.Println(currentPrice)

				price, _ := attr(currentPrice.Attr, "content")
				if _debug {
					fmt.Println(price)
				}

				if _price, err := strconv.ParseFloat(price, 32); err != nil {
					panic(err)
				} else {
					_product.price = float32(_price)
				}
			}

			// if specialOffer := traverse(itemPrice, "em", "class", "specialOffer"); specialOffer != nil {
			// 	fmt.Println(specialOffer)

			// 	price, _ := attr(specialOffer.Attr, "content")
			// 	fmt.Println(price)
			// }

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

			if _tests {
				testCases[_title] = _model
			}

			switch product.callout {
			case New:
				_title += " [N]"
			case Action:
				_title += " [P]"
			case Trade:
				_title += " [R]"
			case Sustainability:
				_title += " [S]"
			}

			_retailPrice := product.price
			_price := _retailPrice
			if product.oldPrice > 0 {
				_retailPrice = product.oldPrice
			} else if product.callout == Trade {
				for _, _product := range _result {
					if _product.callout != Trade && _product.title == product.title {
						_retailPrice = _product.price
					}
				}
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
