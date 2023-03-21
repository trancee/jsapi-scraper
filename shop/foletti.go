package shop

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var FolettiRegex = regexp.MustCompile(`(\s*[-,]\s+)|(\d+\s*GB?)|\s+20[12]\d|\s+(Dual\W(SIM|Sim)|LTE|smartphone|Ice|Blue|Charcoal|Dark Green|Night|bamboo green|blau|denim black|elegant black|grau|lake blue|schwarz)`)

var FolettiCleanFn = func(name string) string {
	// name = strings.ReplaceAll(strings.ReplaceAll(name, " Phones ", " "), " Mini iPhone", " Mini")
	name = regexp.MustCompile(` (SM-)?[AS]\d{3}[A-Z]*(\/DSN)?| XT\d{4}-\d`).ReplaceAllString(name, "")

	if loc := FolettiRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	name = strings.ReplaceAll(name, " E ", " E")

	return strings.TrimSpace(name)
}

func XXX_foletti(isDryRun bool) IShop {
	const _name = "Foletti"
	const _url = "https://superstore.foletti.com/de/categories/it--multimedia/telekommunikation/mobiltelefone/smartphone?limit=100&sort=price|asc&listStyle=list"

	const _debug = false

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

	fn := "shop/foletti.html"

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

		if strings.Contains(title, "Wallet") {
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
		if Skip(brand) {
			continue
		}

		if !strings.EqualFold(strings.ToUpper(strings.Split(brand, " ")[0]), strings.ToUpper(strings.Split(title, " ")[0])) {
			_product.title = brand + " " + _product.title
		}

		model := FolettiCleanFn(_product.title)
		if _debug {
			fmt.Println(model)
		}
		_product.model = model

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

		if itemBadge := traverse(item, "span", "class", "badge"); itemBadge != nil {
			fmt.Println(itemBadge)

			badge, _ := text(itemBadge)
			panic(badge)
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

			_retailPrice := product.price
			_price := _retailPrice
			if product.oldPrice > 0 {
				_retailPrice = product.oldPrice
			}
			_savings := _price - _retailPrice
			_discount := 100 - ((100 / _retailPrice) * _price)

			_title := product.title
			_link := s.ResolveURL(product.link).String()

			product := &Product{
				Code:  _name + "//" + product.code,
				Name:  _title,
				Model: product.model,

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
