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

	"golang.org/x/net/html"
)

var MobileDeviceRegex = regexp.MustCompile(`\s+\(?\d+\s*GB?|\s+\(?\d+(\.\d+)?"|\s+\(?20[12]\d\)?|\s+\(?[2345]G\)?| Dual Sim`)

var MobileDeviceCleanFn = func(name string) string {
	name = regexp.MustCompile(` (SM-)?[AGMS]\d{3}[A-Z]*(\/DSN)?| XT\d{4}-\d+`).ReplaceAllString(name, "")
	name = strings.NewReplacer("Nothing Phone 1", "Nothing Phone (1)", "X Cover", "XCover").Replace(name)

	if loc := MobileDeviceRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return strings.TrimSpace(name)
}

func XXX_mobiledevice(isDryRun bool) IShop {
	const _name = "mobiledevice"
	const _url = "https://www.mobiledevice.ch/modules/blocklayered/blocklayered-ajax.php?layered_quantity_1=1&id_category_layered=28&orderby=price&orderway=asc&n=100"

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

	fn := "shop/mobiledevice.html"

	if isDryRun {
		if body, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else {
			_body = body
		}
	} else {
		/*
			client := &http.Client{
				Timeout: 20 * time.Second,
				Transport: &http.Transport{
					TLSHandshakeTimeout: 10 * time.Second,
			                TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
					CipherSuites: []uint16{
					          tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
					          tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					          tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
					          tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
					          tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					          tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			                 },
				},
			}
		*/
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

	type _Body struct {
		Products string `json:"productList"`
	}

	var body _Body
	{
		if err := json.Unmarshal([]byte(_body), &body); err != nil {
			panic(err)
		}
		_body = []byte(body.Products)
	}
	// fmt.Println(string(_body))

	doc := parse(string(_body))

	if productList := traverse(doc, "ul", "class", "product_list"); productList != nil {
		// fmt.Println(productList)

		for item := productList.FirstChild; item != nil; item = item.NextSibling {
			// item := traverse(items, "li", "class", "ajax_block_product")
			// fmt.Println(item)

			_product := _Response{}

			imageTitleLink := traverse(item, "a", "class", "product_img_link")
			// fmt.Println(imageTitleLink)

			link, _ := attr(imageTitleLink.Attr, "href")
			if _debug {
				fmt.Println(link)
			}
			_product.link = link

			title, _ := attr(imageTitleLink.Attr, "title")
			title = strings.ReplaceAll(title, "X Cover 5 G525", "XCover 5")
			if _debug {
				fmt.Println(title)
			}
			_product.title = title

			if Skip(title) {
				continue
			}

			model := MobileDeviceCleanFn(html.UnescapeString(_product.title))
			if _debug {
				fmt.Println(model)
			}
			_product.model = model

			code := strings.Split(link[45:], "-")[0]
			if _debug {
				fmt.Println(code)
			}
			_product.code = code

			if itemPrice := traverse(item, "span", "class", "price"); itemPrice != nil {
				// fmt.Println(itemPrice)

				price, _ := text(itemPrice)
				price = strings.ReplaceAll(strings.ReplaceAll(price, " CHF", ""), ",", ".")
				if _debug {
					fmt.Println(price)
				}

				if _price, err := strconv.ParseFloat(price, 32); err != nil {
					panic(err)
				} else {
					_product.price = float32(_price)
				}
			}

			if itemOldPrice := traverse(item, "span", "class", "old-price"); itemOldPrice != nil {
				// fmt.Println(itemOldPrice)

				price, _ := text(itemOldPrice)
				price = strings.ReplaceAll(strings.ReplaceAll(price, " CHF", ""), ",", ".")
				if _debug {
					fmt.Println(price)
				}

				if _price, err := strconv.ParseFloat(price, 32); err != nil {
					panic(err)
				} else {
					_product.oldPrice = float32(_price)
				}
			}

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

			_title := html.UnescapeString(product.title)
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
