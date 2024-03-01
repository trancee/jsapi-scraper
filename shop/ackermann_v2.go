package shop

import (
	"encoding/base64"
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

var AckermannV2Regex = regexp.MustCompile(`(?i)(,\s*)?\d+\s*GB|(,\s*)?\(?[2345]G\)?| LTE`)

var AckermannV2CleanFn = func(name string) string {
	// name = strings.NewReplacer("", "").Replace(name)

	if loc := AckermannV2Regex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	s := strings.Split(name, " ")

	if s[0] == "Apple" {
		if s[1] != "iPhone" {
			name = strings.ReplaceAll(name, "Apple ", "Apple iPhone ")
		}
	}

	if s[0] == "Motorola" {
		name = strings.ReplaceAll(name, " 454", " g54")
		name = strings.ReplaceAll(name, "¹³", "13")
	}

	if s[0] == "Samsung" {
		if strings.HasSuffix(name, "Xcover Pro") {
			name = strings.ReplaceAll(name, "Xcover Pro", "XCover 6 Pro")
		} else if strings.HasSuffix(name, "Xcover") {
			name = strings.ReplaceAll(name, "Xcover", "XCover 5")
		}
	}

	return helpers.Lint(name)
}

func XXX_ackermann_v2(isDryRun bool) IShop {
	const _name = "Ackermann"
	_url := fmt.Sprintf("https://www.ackermann.ch/_next/data/shopping_app/de/technik/multimedia/smartphones-telefone.json?f=%s&o=price-asc&categories=technik&categories=multimedia&categories=smartphones-telefone", base64.StdEncoding.EncodeToString(StringToBytes(fmt.Sprintf(`{"filter_Produkttyp_1":["fb_prdkt.p1_smrt.hn_38"],"filter_price":["%.f-%.f"]}`, ValueMinimum, ValueMaximum))))

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

		discountPercentage float32
	}

	type _Data struct {
		Category struct {
			Summary struct {
				Count int `json:"totalResultCount"`
			} `json:"summary"`
			Items []struct {
				Brand struct {
					Name string `json:"name"`
				} `json:"brand"`

				StyleID   string `json:"styleId"`
				StyleName string `json:"styleName"`

				PrimaryVariationGroup struct {
					Price struct {
						Currency     string  `json:"currency"`
						CurrentPrice float32 `json:"currentPrice"`
						Discount     struct {
							PreviousPrice      float32 `json:"previousPrice"`
							DiscountPercentage float32 `json:"discountPercentage"`
						} `json:"discount"`
					} `json:"price"`

					SKU           string `json:"sku"`
					ArticleNumber string `json:"articleNumber"`

					Link string `json:"href"`
				} `json:"primaryVariationGroup"`
			} `json:"items"`
		} `json:"category"`
	}

	type _State struct {
		Data    string `json:"data"`
		HasNext bool   `json:"hasNext"`
	}

	type _Body struct {
		PageProps struct {
			UrqlState map[string]_State `json:"urqlState"`
		} `json:"pageProps"`
	}

	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	var _results []_Response

	fn := fmt.Sprintf("shop/ackermann.json")

	if isDryRun {
		if body, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else {
			_body = body
		}
	} else {
		url := _url

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
	// fmt.Println(BytesToString(_body))

	var body _Body
	if err := sonnet.Unmarshal(_body, &body); err != nil { // Parse []byte to go struct pointer
		panic(err)
	}

	for k := range body.PageProps.UrqlState {
		var _data _Data
		{
			if err := sonnet.Unmarshal(StringToBytes(body.PageProps.UrqlState[k].Data), &_data); err != nil {
				panic(err)
			}
			// fmt.Printf("%+v\n", _data)
		}

		reName := regexp.MustCompile(`»(.*?)«`)

		for _, item := range _data.Category.Items {
			name := item.StyleName
			brand := item.Brand.Name

			if matches := reName.FindStringSubmatch(name); len(matches) > 1 {
				title := matches[1]

				s := strings.Split(title, " ")
				if s[0] != brand {
					title = brand + " " + title
				}

				// fmt.Println(title)
				model := AckermannV2CleanFn(title)
				// fmt.Println(model)
				// fmt.Println()

				result := _Response{
					code:  item.PrimaryVariationGroup.SKU,
					title: title,
					model: model,

					link: item.PrimaryVariationGroup.Link,

					oldPrice: item.PrimaryVariationGroup.Price.Discount.PreviousPrice / 100.0,
					price:    item.PrimaryVariationGroup.Price.CurrentPrice / 100.0,

					discountPercentage: item.PrimaryVariationGroup.Price.Discount.DiscountPercentage,
				}
				// fmt.Println(result)

				_results = append(_results, result)
			}
		}
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_results))
		for _, _product := range _results {
			// fmt.Println(_product)

			_title := _product.title
			_model := _product.model

			if Skip(_model) {
				continue
			}

			if _debug {
				fmt.Println(_model)
			}

			if _tests {
				testCases[_title] = _model
			}

			var _savings float32
			var _discount float32

			_retailPrice := max(_product.oldPrice, _product.price)
			_price := _retailPrice
			if _product.price > 0 {
				_price = _product.price
			}
			if _debug {
				fmt.Println(_retailPrice)
				fmt.Println(_price)
			}

			if _price > 0 {
				_savings = _price - _retailPrice
				_discount = 100 - ((_price * 100) / _retailPrice)

				// _discount = _product.discountPercentage
			}
			if _debug {
				fmt.Println(_savings)
				fmt.Println(_discount)
			}

			_link := s.ResolveURL(_product.link).String()
			if _debug {
				fmt.Println(_link)
				fmt.Println()
			}

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
