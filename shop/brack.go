package shop

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/sugawarayuuta/sonnet"

	helpers "jsapi-scraper/helpers"
)

var BrackRegex = regexp.MustCompile(`(\s*[-,]\s+)|(\d+\s*GB?)|\s+((EE )?Enterprise Edition( CH)?)`)

var BrackCleanFn = func(name string) string {
	name = strings.NewReplacer(" Phones ", " ", "Recommerce Switzerland SA ", "", "3. Gen.", "(2022)").Replace(name)

	if loc := BrackRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	return helpers.Lint(name)

	// s := strings.Split(name, " ")

	// if s[0] == "iPhone" {
	// 	name = "Apple " + name
	// }

	// if s[0] == "Apple" {
	// 	name = strings.NewReplacer(" 2020", " (2020)", " 2022", " (2022)", " 2nd Gen", " (2020)", " 3rd Gen", " (2022)").Replace(name)
	// } else {
	// 	// Remove year component for all other than Apple.
	// 	name = regexp.MustCompile(`\s+\(?20[12]\d\)?`).ReplaceAllString(name, "")
	// }

	// return strings.TrimSpace(name)
}

func XXX_brack(isDryRun bool) IShop {
	const _name = "Brack"
	// const _url = "https://www.brack.ch/it-multimedia/telefonie-kommunikation/mobiltelefone/smartphone?filter%5BArt%5D%5B%5D=offer&filter%5BArt%5D%5B%5D=intropromotion&filter%5BArt%5D%5B%5D=occassion&filter%5BArt%5D%5B%5D=new&sortProducts=priceasc&query=*"
	// const _url = "https://www.brack.ch/it-multimedia/telefonie-kommunikation/mobiltelefone/smartphone?limit=192&sortProducts=priceasc&query=*"
	// _url := fmt.Sprintf("https://www.brack.ch/it-multimedia/telefonie-kommunikation/mobiltelefone/smartphone?filter[availability][]=Verfügbar&filter[price_standard][]=%.f~~~%.f&sortProducts=priceasc&query=*", ValueMinimum, ValueMaximum)
	// _url := fmt.Sprintf("https://www.brack.ch/api/search?uri=%%2Fit-multimedia%%2Ftelefonie-kommunikation%%2Fmobiltelefone%%2Fsmartphone%%3Ffilter%%255Bprice_standard%%255D%%255B%%255D%%3D%.f~~~%.f%%26filter%%255Bavailability%%255D%%255B%%255D%%3DVerf%%25C3%%25BCgbar%%26limit%%3D192%%26sortProducts%%3Dpriceasc", ValueMinimum, ValueMaximum)
	_url := fmt.Sprintf("https://www.brack.ch/it-multimedia/smartphone?filter[availability][]=Verfügbar&filter[price_standard][]=%.f~~~%.f&sortProducts=priceasc&query=*", ValueMinimum, ValueMaximum)

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

	type _LSPI struct {
		Inventory struct {
			Current                     int  `json:"current"`
			TargetStockQuantity         int  `json:"targetStockQuantity"`
			HasManuallySetStockQuantity bool `json:"hasManuallySetStockQuantity"`
		} `json:"inventory"`

		URL string `json:"url"`

		States struct {
			Active       bool `json:"active"`
			NewProduct   bool `json:"newProduct"`
			Sellout      bool `json:"sellout"`
			SpecialSale  bool `json:"specialSale"`
			SpecialOffer bool `json:"specialOffer"`
		} `json:"states"`
	}

	type _LSI struct {
		Src string `json:"src"`
		Alt string `json:"alt"`
	}

	type _SKU struct {
		quantity int

		price    float32
		oldPrice float32
	}
	var _skus map[string]_SKU = map[string]_SKU{}

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

	if body := regexp.MustCompile(`window.competec.products.data = {.*?};`).Find(_body); body != nil {
		data := body[32 : len(body)-1]

		var skus map[string]any
		if err := sonnet.Unmarshal(data, &skus); err != nil {
			panic(err)
		}
		// fmt.Printf("%+v\n", skus)

		for sku, value := range skus {
			// Each value is an `any` type, that is type asserted as a string
			// fmt.Printf("%s: %v\n", key, value)
			_inventory := value.(map[string]any)["inventory"]
			quantity := _inventory.(map[string]any)["current"].(float64)
			// fmt.Printf("quantity: %v\n", quantity)

			_price := value.(map[string]any)["price"]
			// fmt.Printf("%v\n", _price)
			// specialOffer := _price.(map[string]any)["specialOffer"]
			// fmt.Printf("specialOffer: %v\n", specialOffer)
			price := _price.(map[string]any)["priceWithVat"].(float64) / 100
			// fmt.Printf("price: %v\n", price)

			_standardPrice := _price.(map[string]any)["standardPrice"]
			// fmt.Printf("%v\n", _standardPrice)
			oldPrice := _standardPrice.(map[string]any)["priceWithVat"].(float64) / 100
			// fmt.Printf("oldPrice: %v\n", oldPrice)

			// fmt.Printf("%s\t%.2f\t%.2f\t%v\n", sku, price, oldPrice, quantity)

			_skus[sku] = _SKU{
				quantity: int(quantity),

				price:    float32(price),
				oldPrice: float32(oldPrice),
			}
		}

		// fmt.Println(_skus)
	}

	doc := parse(string(_body))

	for sku, value := range _skus {
		if snippet := traverse(doc, "li", "data-snippet", sku); snippet != nil {
			_product := _Response{}

			if _debug {
				fmt.Println(sku)
			}
			_product.code = sku

			oldPrice := value.oldPrice
			if _debug {
				fmt.Println(oldPrice)
			}
			_product.oldPrice = oldPrice

			price := value.price
			if _debug {
				fmt.Println(price)
			}
			_product.price = price

			_brand := ""
			if product := traverse(snippet, "div", "class", "blsaP_Iq2"); product != nil {
				if brand, ok := text(product.FirstChild.FirstChild); ok {
					if _debug {
						fmt.Println(brand)
					}
					_brand = brand
				}
			}

			var _lspl _LSPI
			if lspl := traverse(snippet, "div", "data-e-ref", "lspl"); lspl != nil {
				if data, ok := text(lspl.FirstChild); ok {
					if err := sonnet.Unmarshal([]byte(data), &_lspl); err != nil {
						panic(err)
					}
					// fmt.Println(_lspl)

					amount := _lspl.Inventory.Current
					if _debug {
						fmt.Println(amount)
					}
					_product.quantity = amount

					link := _lspl.URL
					if _debug {
						fmt.Println(link)
					}
					_product.link = link
				}
			}

			var _lsi _LSI
			if lsi := traverse(snippet, "div", "data-e-ref", "lsi"); lsi != nil {
				if data, ok := text(lsi.FirstChild); ok {
					if err := sonnet.Unmarshal([]byte(data), &_lsi); err != nil {
						panic(err)
					}
					// fmt.Println(_lsi)

					title := _brand + " " + _lsi.Alt
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
