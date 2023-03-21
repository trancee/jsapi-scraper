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

var MediamarktRegex = regexp.MustCompile(` - |\s+20[12]\d|\s+[2345]G`)

var MediamarktCleanFn = func(name string) string {
	if loc := MediamarktRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return strings.TrimSpace(name)
}

func XXX_mediamarkt(isDryRun bool) IShop {
	const _name = "Mediamarkt"
	// const _url = "https://www.mediamarkt.ch/de/category/_smartphone-680815.html?searchParams=&sort=price&view=PRODUCTGRID"
	const _url = "https://www.mediamarkt.ch/de/category/_smartphone-680815.html?searchParams=%2FSearch.ff%3FfilterCategoriesROOT%3DHandy%2B%2526%2BNavigation%25C2%25A7MediaCHdec680760%26filterCategoriesROOT%252FHandy%2B%2526%2BNavigation%25C2%25A7MediaCHdec680760%3DSmartphone%25C2%25A7MediaCHdec680815%26filteravailability%3D1%26filterTyp%3D___Smartphone%26channel%3Dmmchde%26followSearch%3D9782%26disableTabbedCategory%3Dtrue%26navigation%3Dtrue&sort=price&view=PRODUCTGRID&page="

	const _debug = false

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

	fn := "shop/mediamarkt.html"

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

	productList := traverse(doc, "ul", "class", "products-grid")
	// fmt.Println(productList)

	for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling.NextSibling {
		// fmt.Println(item)

		_product := _Response{}

		baseInfo := traverse(item, "div", "class", "base-info")
		// fmt.Println(baseInfo)

		productKey, _ := attr(baseInfo.Attr, "data-reco-pid")
		// fmt.Println(productKey)
		productId := productKey[2:]
		if _debug {
			fmt.Println(productId)
		}
		_product.code = productId

		imageTitleLink := traverse(baseInfo, "a", "class", "product-link")
		// fmt.Println(imageTitleLink)

		link, _ := attr(imageTitleLink.Attr, "href")
		if _debug {
			fmt.Println(link)
		}
		_product.link = link

		title, _ := text(imageTitleLink)
		// title = strings.TrimSpace(strings.Split(strings.ReplaceAll(strings.ReplaceAll(title, " - Smartphone", ""), " \"", "\""), "(")[0])
		if _debug {
			fmt.Println(title)
		}
		_product.title = title

		if Skip(title) {
			continue
		}

		model := MediamarktCleanFn(_product.title)
		if _debug {
			fmt.Println(model)
		}
		_product.model = model

		currentPrice := traverse(baseInfo, "div", "class", "price")
		// fmt.Println(currentPrice)

		price, _ := text(currentPrice)
		if _debug {
			fmt.Println(price)
		}

		if _price, err := strconv.ParseFloat(strings.ReplaceAll(price, ".-", ".00"), 32); err != nil {
			panic(err)
		} else {
			_product.oldPrice = float32(_price)
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

			_retailPrice := product.oldPrice
			_price := _retailPrice
			if product.price > 0 {
				_price = product.price
			}
			_savings := _price - _retailPrice
			_discount := 100 - ((100 / _retailPrice) * _price)

			_link := s.ResolveURL(product.link).String()

			product := &Product{
				Code:  _name + "//" + product.code,
				Name:  product.title,
				Model: product.model,

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

		// fmt.Printf("%#v\n", products)
		return &products
	}

	return NewShop(
		_name,
		_url,

		_parseFn,
	)
}
