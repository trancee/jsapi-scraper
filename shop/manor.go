package shop

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

var ManorRegex = regexp.MustCompile(`, |\d\+\d+GB|\s+\(?[2345]G\)?|\d+(,\d)? cm| \d\.\d|(\d+\s*GB?)|\s+20[12]\d|(SM-)?[AFGMS]\d{3}[BFR]?(\/DSN?)?| XT\d{4}-\d+|\s+EE |\s+(Enterprise Edition( CH)?)| Dual`)

var ManorCleanFn = func(name string) string {
	name = strings.NewReplacer(" NOK ", " ", " Smartphone Pack ", " ", " Smartphone Bundle ", " ", " Pack Smartphone Vivo", " ", " G ", " ").Replace(name)

	name = regexp.MustCompile(`\s{2,}`).ReplaceAllString(name, " ")

	if loc := ManorRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	s := strings.Split(name, " ")

	if s[0] == "MOTOROLA" {
		name = strings.ReplaceAll(name, "Moto E ", "Moto ")
	}

	if s[0] == "OPPO" {
		name = regexp.MustCompile(`Reno\s*(\d)\s*(\w)?`).ReplaceAllString(name, "Reno$1 $2")
	}

	return strings.TrimSpace(name)
}

func XXX_manor(isDryRun bool) IShop {
	const _name = "Manor"
	_url := fmt.Sprintf("https://www.manor.ch/_next/data/g5Ai52cssZcOwaX4irIEt/de/shop/multimedia/telefonie-navigation/smartphones/c/telephone-navigation-smartphones.json?priceValue=>%.f+|+<%.f&sort=PRICE_VALUE_ASC&slug=shop&slug=multimedia&slug=telefonie-navigation&slug=smartphones&slug=c&slug=telephone-navigation-smartphones", ValueMinimum, ValueMaximum)

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

	type _Body struct {
		PageProps struct {
			InitialApolloState map[string]json.RawMessage `json:"initialApolloState"`
		} `json:"pageProps"`
	}

	type _Query struct {
		SearchProducts struct {
			TotalResults int `json:"totalResults"`
			Page         int `json:"page"`
			PageSize     int `json:"pageSize"`
			TotalPages   int `json:"totalPages"`
			Products     []struct {
				Ref string `json:"__ref"`
			} `json:"products"`
		} `json:"searchProducts"`
	}

	type _Product struct {
		Code string `json:"code"`
		Link string `json:"link"`

		BrandName string `json:"brandName"`
		Name      string `json:"name"`

		Titles struct {
			First  string `json:"first"`
			Second string `json:"second"`
			Third  string `json:"third"`
		} `json:"titles"`

		PriceValue struct {
			Amount float32 `json:"amount"`
		} `json:"priceValue"`
		OriginalPrice struct {
			Amount float32 `json:"amount"`
		} `json:"originalPrice"`
	}

	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	var _results []_Response

	for p := 1; p <= 5; p++ {
		fn := fmt.Sprintf("shop/manor.%d.json", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			page := ""
			if p > 1 {
				page = fmt.Sprintf("&page=%d", p)
			}
			resp, err := http.Get(fmt.Sprintf("%s%s", _url, page))
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, fmt.Sprintf("%s%d", _url, p))
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

		var body _Body
		if err := json.Unmarshal(_body, &body); err != nil { // Parse []byte to go struct pointer
			panic(err)
		}

		var query _Query
		if err := json.Unmarshal(body.PageProps.InitialApolloState["ROOT_QUERY"], &query); err != nil {
			panic(err)
		}
		// fmt.Println(query)

		for _, k := range query.SearchProducts.Products {
			var product _Product
			if err := json.Unmarshal(body.PageProps.InitialApolloState[k.Ref], &product); err != nil {
				panic(err)
			}
			// fmt.Println(product)

			name := product.Name
			brand := strings.Split(name, " ")[0]
			if strings.ToUpper(product.BrandName) == strings.ToUpper(brand) {
				name = strings.ReplaceAll(name, brand, "")
			}
			title := product.BrandName + " " + product.Titles.Second + " " + name
			model := ManorCleanFn(title)
			// fmt.Println(model)

			_results = append(_results, _Response{
				code:  product.Code,
				title: title,
				model: model,

				link: product.Link,

				oldPrice: product.OriginalPrice.Amount,
				price:    product.PriceValue.Amount,
			})
		}

		if query.SearchProducts.Page >= p {
			break
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

			if _tests {
				testCases[_title] = _model
			}

			var _savings float32
			var _discount float32

			_retailPrice := _product.oldPrice
			_price := _retailPrice
			if _product.price > 0 {
				_price = _product.price
			}
			if _retailPrice > 0 {
				_savings = _price - _retailPrice
				_discount = 100 - ((100 / _retailPrice) * _price)
			}

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
		_url,

		_parseFn,
	)
}
