package shop

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func XXX_brack(isDryRun *bool) IShop {
	const _name = "Brack"
	// const _url = "https://www.brack.ch/it-multimedia/telefonie-kommunikation/mobiltelefone/smartphone?filter%5BArt%5D%5B%5D=offer&filter%5BArt%5D%5B%5D=intropromotion&filter%5BArt%5D%5B%5D=occassion&filter%5BArt%5D%5B%5D=new&sortProducts=priceasc&query=*"
	const _url = "https://www.brack.ch/it-multimedia/telefonie-kommunikation/mobiltelefone/smartphone?limit=192&sortProducts=priceasc&query=*"

	const _debug = false

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

		link string

		oldPrice float32
		price    float32

		quantity int

		callout Callout
	}

	var _result []_Response
	var _body []byte

	fn := "shop/brack.html"

	if isDryRun != nil && *isDryRun {
		if body, err := os.ReadFile(fn); err != nil {
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

		os.WriteFile(fn, _body, 0664)
	}
	// fmt.Println(string(_body))

	doc := parse(string(_body))

	productList := traverse(doc, "ul", "class", "productList")
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
		if _debug {
			fmt.Println(title)
		}
		_product.title = strings.ReplaceAll(title, " - ", " ")

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

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result))
		for _, product := range _result {
			// fmt.Println(product)

			_title := product.title
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
