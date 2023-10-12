package shop

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html"

	"github.com/sugawarayuuta/sonnet"

	helpers "jsapi-scraper/helpers"
)

var GalaxusRegex = regexp.MustCompile(`, | [+-] |\s+\d+\/\d+|\s*\d+G?\+\d+G?|\s*\(?(\s*[+\/]\s*)?((2|4|6|8|12|16)(GB)?\s*[+\/]\s*)?(2|4|6|8|12|16|32|64|128|256)\s*GB\)?|\d+G\/\d+G|\s+\(?20[12]\d\)?|\s+[45]g|\s+X\d{3}F|\s+\(V\d{4}\)|\d{4,} mAh|\s+\(?(\d{1,2}[., ])?\d+( Zoll|\")\)?|\s+(1\d[., ])?\d+\s*cm|\s+\(?\d\.\d+( Zoll|\")\s*\)?| DS\s*\d|\s+((EE )?Enterprise Edition( CH)?)| Master( Edition)?| DE| EU| LTE| NFC| OLED| (Dual|DUAL)[ -](Sim|SIM)|\/BLUE|GREEN |( Sky)? [Bb]lue| Cosmic Aurora| Elegant Black| Force Touch| Grey|(\/?LASER)? BLACK| Midnight Blue| Midnight Space| \(?Ocean Blue\)?| Pastel Lime| Pearl White|Space Silver| Viva Magenta| bamboo green| black| dark green| hellblau| midday dream| midnight blue`)

var GalaxusCleanFn = func(name string) string {
	if loc := GalaxusRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	name = regexp.MustCompile(`\s+[2345]G(\s+EU|\s+\d)?(\s+NE)?(\s+Phone)?|\s+I9505|\s+[A]\d{3}[B]| XT\d{4}-\d+|\/(2|4|6|8|12)\/(64|128|256)|( Blackview| Graues)? Smart(fon|phone)( Blackview| oppo| ZTE)?| Smartfon|^Vodafone |^TIM |^TE Connectivity |HON DS |OPP DS | Snapdragon| Black| 2 ”| MOBILE PHONE| SMARTPHONE( MOTOROLA)?|Motorola Smartfon | Handy| OEM| TCT| VoLTE| \+ Huawei| Outdoor| Bluetooth Speaker| Android| Limited|(Honor)? Telefon(as)?|Inapa|\(Snapdragon\)|( Porsche)? Design| czarny| pomarańczowy| zielony| Supplier did not provide product name`).ReplaceAllString(name, "")
	name = strings.NewReplacer("Xiaomi M5s", "Xiaomi Poco M5s", "Note9", "Note 9", "Nokia Nokia ", "Nokia ", "Edge30", "Edge 30", "Rephone Rephone", "Rephone", "A1 Plus", "A1+", "Master Edition", "Master", "SAM DS ", "SAMSUNG ", "GAL ", "GALAXY ", "HOT205G", "HOT 20 5G ", "SE2020", "SE 2020", "TCL 40 40SE", "TCL 40SE", "Xiaomi Xia ", "Xiaomi ", "Motorola 41", "Motorola Moto G41", " CE3", " CE 3", "A57s 4", "A57s", "2nd Gen", "2020").Replace(name)
	name = strings.TrimSpace(name)

	s := strings.Split(name, " ")

	if s[0] == "Emporia" {
		name = strings.NewReplacer("Super Easy", "SUPEREASY").Replace(name)
	}

	if s[0] == "Honor" || s[0] == "HONOR" {
		name = regexp.MustCompile(`(?i)([X]\d[a]?)\s+\d$`).ReplaceAllString(name, "$1")
	}

	if s[0] == "Infinix" {
		name = strings.Split(name, "5G")[0]
		name = strings.ReplaceAll(name, " INFINIX", "")
		name = strings.ReplaceAll(name, " Infinix", "")
	}

	if s[0] == "Motorola" {
		name = strings.ReplaceAll(name, "G31 4", "G31")
		name = strings.ReplaceAll(name, "G42 4", "G42")

		name = strings.ReplaceAll(name, "Motorola Motorola ", "Motorola ")
		name = strings.ReplaceAll(name, "Moto E Moto ", "Moto ")
		name = strings.ReplaceAll(name, "Moto E moto ", "Moto ")
		name = strings.ReplaceAll(name, "Motorola 30 ", "Motorola Edge 30 ")
	}

	if s[0] == "POCO" || s[0] == "Poco" {
		name = "Xiaomi" + " " + name
	}

	if s[0] == "realme" {
		name = strings.ReplaceAll(name, "realme SM ", "")
		name = strings.ReplaceAll(name, "-Y", "Y")
	}

	if s[0] == "Xiaomi" {
		if s[1] == "Samsung" || s[1] == "Honor" || s[1] == "Xiaomi" {
			name = strings.Replace(name, "Xiaomi ", "", 1)
		}

		name = regexp.MustCompile(`Redmi\s*(\d+)`).ReplaceAllString(name, "Redmi $1")

		if strings.HasSuffix(name, "NOTE 1") {
			name = strings.ReplaceAll(name, "NOTE 1", "NOTE 12")
		}

		if strings.HasSuffix(name, "Note 1") {
			name = strings.ReplaceAll(name, "Note 1", "Note 11S")
		}
	}

	if s[0] == "Renewd" {
		if s[1] == "iPhone" {
			name = strings.ReplaceAll(name, "Renewd", "Apple")
		}
	}

	if s[0] == "ZTE" {
		name = strings.ReplaceAll(name, "s 3", "s")
	}

	return helpers.Lint(name)
}

func XXX_galaxus(isDryRun bool) IShop {
	const _name = "Galaxus"
	const _url = "https://www.galaxus.ch/api/graphql/product-type-filter-products"

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	type _Product struct {
		Code int    `json:"productId"`
		Name string `json:"name"`

		Description string `json:"nameProperties"`

		Category string `json:"productTypeName"`

		Brand string `json:"brandName"`
	}

	type _Offer struct {
		Price struct {
			Amount float32 `json:"amountIncl"`
		} `json:"price"`

		ShopOfferID int `json:"shopOfferId"`

		// Type string `json:"type"`

		// IsNew            bool `json:"isNew"`
		// IsSalesPromotion bool `json:"isSalesPromotion"`

		OldPrice struct {
			Price struct {
				Amount float32 `json:"amountIncl"`
			} `json:"price"`
		} `json:"insteadOfPrice"`
	}

	type _Result struct {
		Product _Product `json:"product"`
		Offer   _Offer   `json:"offer"`

		IsDefaultOffer bool `json:"isDefaultOffer"`
	}

	type _Response struct {
		Data struct {
			ProductType struct {
				FilterProducts struct {
					Products struct {
						Results []_Result `json:"results"`
					} `json:"products"`
				} `json:"filterProducts"`
			} `json:"productType"`
		} `json:"data"`
	}

	var _result []_Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	var _results []_Result

	for p := 1; p <= 5; p++ {
		fn := fmt.Sprintf("shop/galaxus.%d.json", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			jsonData := []byte(fmt.Sprintf(`[
				{
					"operationName": "PRODUCT_TYPE_FILTER_PRODUCTS",
					"variables": {
					  "productTypeId": 24,
					  "offset": %d,
					  "limit": 200,
					  "sortOrder": "LOWESTPRICE",
					  "siteId": null,
					  "filters": [
						{
							"identifier": "8279",
							"filterType": "TEXTUAL",
							"options": ["6395", "476832"]
						},
						{
						  "identifier": "pr",
						  "filterType": "NUMERICRANGE",
						  "options": [],
						  "greaterThanOrEquals": %.f,
						  "lessThanOrEquals": %.f
						}
					  ],
					  "searchTerm": null
					},
					"query": "query PRODUCT_TYPE_FILTER_PRODUCTS($productTypeId: Int!, $filters: [SearchFilter!], $sortOrder: ProductSort, $offset: Int, $siteId: String, $limit: Int, $searchTerm: String) {\n  productType(id: $productTypeId) {\n    filterProducts(\n      offset: $offset\n      limit: $limit\n      sort: $sortOrder\n      siteId: $siteId\n      filters: $filters\n      searchTerm: $searchTerm\n    ) {\n      products {\n        hasMore\n        results {\n          ...ProductWithOffer\n          __typename\n        }\n        __typename\n      }\n      counts {\n        total\n        filteredTotal\n        __typename\n      }\n      filters {\n        identifier\n        name\n        filterType\n        score\n        tooltip {\n          ...FilterTooltipResult\n          __typename\n        }\n        ...CheckboxSearchFilterResult\n        ...RangeSearchFilterResult\n        __typename\n      }\n      quickFilter {\n        options {\n          filterType\n          filterIdentifier\n          filterName\n          filterOptionIdentifier\n          filterOptionName\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment ProductWithOffer on ProductWithOffer {\n  mandatorSpecificData {\n    ...ProductMandatorSpecific\n    __typename\n  }\n  product {\n    ...ProductMandatorIndependent\n    __typename\n  }\n  offer {\n    ...ProductOffer\n    __typename\n  }\n  isDefaultOffer\n  __typename\n}\n\nfragment FilterTooltipResult on FilterTooltip {\n  text\n  moreInformationLink\n  __typename\n}\n\nfragment CheckboxSearchFilterResult on CheckboxSearchFilter {\n  options {\n    identifier\n    name\n    productCount\n    score\n    referenceValue {\n      value\n      unit {\n        abbreviation\n        __typename\n      }\n      __typename\n    }\n    preferredValue {\n      value\n      unit {\n        abbreviation\n        __typename\n      }\n      __typename\n    }\n    tooltip {\n      ...FilterTooltipResult\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment RangeSearchFilterResult on RangeSearchFilter {\n  referenceMin\n  preferredMin\n  referenceMax\n  preferredMax\n  referenceStepSize\n  preferredStepSize\n  rangeMergeInfo {\n    isBottomMerged\n    isTopMerged\n    __typename\n  }\n  referenceUnit {\n    abbreviation\n    __typename\n  }\n  preferredUnit {\n    abbreviation\n    __typename\n  }\n  rangeFilterDataPoint {\n    ...RangeFilterDataPointResult\n    __typename\n  }\n  __typename\n}\n\nfragment ProductMandatorSpecific on MandatorSpecificData {\n  isBestseller\n  isDeleted\n  showroomSites\n  sectorIds\n  hasVariants\n  __typename\n}\n\nfragment ProductMandatorIndependent on ProductV2 {\n  id\n  productId\n  name\n  nameProperties\n  productTypeId\n  productTypeName\n  brandId\n  brandName\n  averageRating\n  totalRatings\n  totalQuestions\n  isProductSet\n  images {\n    url\n    height\n    width\n    __typename\n  }\n  energyEfficiency {\n    energyEfficiencyColorType\n    energyEfficiencyLabelText\n    energyEfficiencyLabelSigns\n    energyEfficiencyImage {\n      url\n      height\n      width\n      __typename\n    }\n    __typename\n  }\n  seo {\n    seoProductTypeName\n    seoNameProperties\n    productGroups {\n      productGroup1\n      productGroup2\n      productGroup3\n      productGroup4\n      __typename\n    }\n    gtin\n    __typename\n  }\n  smallDimensions\n  basePrice {\n    priceFactor\n    value\n    __typename\n  }\n  productDataSheet {\n    name\n    languages\n    url\n    size\n    __typename\n  }\n  __typename\n}\n\nfragment ProductOffer on OfferV2 {\n  id\n  productId\n  offerId\n  shopOfferId\n  price {\n    amountIncl\n    amountExcl\n    currency\n    __typename\n  }\n  deliveryOptions {\n    mail {\n      classification\n      futureReleaseDate\n      __typename\n    }\n    pickup {\n      siteId\n      classification\n      futureReleaseDate\n      __typename\n    }\n    detailsProvider {\n      productId\n      offerId\n      quantity\n      type\n      __typename\n    }\n    __typename\n  }\n  label\n  labelType\n  type\n  volumeDiscountPrices {\n    minAmount\n    price {\n      amountIncl\n      amountExcl\n      currency\n      __typename\n    }\n    isDefault\n    __typename\n  }\n  salesInformation {\n    numberOfItems\n    numberOfItemsSold\n    isEndingSoon\n    validFrom\n    __typename\n  }\n  incentiveText\n  isIncentiveCashback\n  isNew\n  isSalesPromotion\n  hideInProductDiscovery\n  canAddToBasket\n  hidePrice\n  insteadOfPrice {\n    type\n    price {\n      amountIncl\n      amountExcl\n      currency\n      __typename\n    }\n    __typename\n  }\n  minOrderQuantity\n  __typename\n}\n\nfragment RangeFilterDataPointResult on RangeFilterDataPoint {\n  count\n  referenceValue {\n    value\n    unit {\n      abbreviation\n      __typename\n    }\n    __typename\n  }\n  preferredValue {\n    value\n    unit {\n      abbreviation\n      __typename\n    }\n    __typename\n  }\n  __typename\n}"
				  }
			]`, len(_results), ValueMinimum, ValueMaximum))

			req, err := http.NewRequest("POST", _url, bytes.NewBuffer(jsonData))
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
				return NewShop(
					_name,
					_url,

					nil,
				)
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
			req.Header.Set("Origin", "https://www.galaxus.ch")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
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

		if err := sonnet.Unmarshal(_body, &_result); err != nil {
			panic(err)
		}
		// fmt.Println(_result[0].Data.ProductType.FilterProducts.Products.Results)

		_results = append(_results, _result[0].Data.ProductType.FilterProducts.Products.Results...)

		if len(_result[0].Data.ProductType.FilterProducts.Products.Results) < 200 {
			break
		}
	}

	r := regexp.MustCompile("[^a-z0-9  -]")

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_results))
		for _, result := range _results {
			product := result.Product
			offer := result.Offer

			product.Brand = html.UnescapeString(product.Brand)
			product.Name = html.UnescapeString(product.Name)
			product.Description = html.UnescapeString(product.Description)

			_title := product.Brand + " " + product.Name
			_model := GalaxusCleanFn(_title)

			if Skip(_model) {
				continue
			}
			if _debug {
				// fmt.Println(_title)
				fmt.Println(_model)
			}

			if _tests {
				testCases[_title] = _model
			}

			var _retailPrice float32
			var _price float32
			var _savings float32
			var _discount float32

			_retailPrice = max(offer.Price.Amount, offer.OldPrice.Price.Amount)
			_price = min(offer.Price.Amount, offer.OldPrice.Price.Amount)
			if _price == 0 {
				_price = _retailPrice
			}
			if _debug {
				fmt.Println(_retailPrice)
				fmt.Println(_price)
			}

			if _savings == 0 {
				_savings = _price - _retailPrice
			}
			_discount = 100 - ((100 / _retailPrice) * _price)
			if _debug {
				fmt.Println(_savings)
				fmt.Println(_discount)
			}

			_productName := strings.NewReplacer(" ", "-", " ", "-").Replace(r.ReplaceAllString(strings.ToLower(product.Brand+"-"+product.Name+"-"+product.Description+"-"+product.Category), "$1"))
			_link := fmt.Sprintf("https://www.galaxus.ch/de/s1/product/%s-%d", _productName, product.Code)
			if !result.IsDefaultOffer {
				_link += fmt.Sprintf("?shid=%d", offer.ShopOfferID)
			}
			if _debug {
				fmt.Println(_link)
				fmt.Println()
			}

			{
				code := strconv.Itoa(product.Code)
				product := &Product{
					Code:  _name + "//" + code,
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
