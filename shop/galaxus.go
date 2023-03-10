package shop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func XXX_galaxus(isDryRun bool) IShop {
	const _name = "Galaxus"
	const _url = "https://www.galaxus.ch/api/graphql/product-type-filter-products"

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

		Type string `json:"type"`

		IsNew            bool `json:"isNew"`
		IsSalesPromotion bool `json:"isSalesPromotion"`

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

	fn := "shop/galaxus.json"

	if isDryRun {
		if body, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else {
			_body = body
		}
	} else {
		jsonData := []byte(`[
			{
				"operationName": "PRODUCT_TYPE_FILTER_PRODUCTS",
				"variables": {
					"productTypeId": 24,
					"offset": 0,
					"limit": 200,
					"sortOrder": "LOWESTPRICE",
					"siteId": null,
					"filters": [
						{
							"identifier": "off",
							"filterType": "TEXTUAL",
							"options": [
								"Sale",
								"Secondhand"
							]
						}
					]
				},
				"query": "query PRODUCT_TYPE_FILTER_PRODUCTS($productTypeId: Int!, $filters: [SearchFilter!], $sortOrder: ProductSort, $offset: Int, $siteId: String, $limit: Int, $searchTerm: String) {\n  productType(id: $productTypeId) {\n    filterProducts(\n      offset: $offset\n      limit: $limit\n      sort: $sortOrder\n      siteId: $siteId\n      filters: $filters\n      searchTerm: $searchTerm\n    ) {\n      products {\n        hasMore\n        results {\n          ...ProductWithOffer\n          __typename\n        }\n        __typename\n      }\n      counts {\n        total\n        filteredTotal\n        __typename\n      }\n      filters {\n        identifier\n        name\n        filterType\n        score\n        tooltip {\n          ...FilterTooltipResult\n          __typename\n        }\n        ...CheckboxSearchFilterResult\n        ...RangeSearchFilterResult\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment ProductWithOffer on ProductWithOffer {\n  mandatorSpecificData {\n    ...ProductMandatorSpecific\n    __typename\n  }\n  product {\n    ...ProductMandatorIndependent\n    __typename\n  }\n  offer {\n    ...ProductOffer\n    __typename\n  }\n  isDefaultOffer\n  __typename\n}\n\nfragment FilterTooltipResult on FilterTooltip {\n  text\n  moreInformationLink\n  __typename\n}\n\nfragment CheckboxSearchFilterResult on CheckboxSearchFilter {\n  options {\n    identifier\n    name\n    productCount\n    score\n    referenceValue {\n      value\n      unit {\n        abbreviation\n        __typename\n      }\n      __typename\n    }\n    preferredValue {\n      value\n      unit {\n        abbreviation\n        __typename\n      }\n      __typename\n    }\n    tooltip {\n      ...FilterTooltipResult\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment RangeSearchFilterResult on RangeSearchFilter {\n  referenceMin\n  preferredMin\n  referenceMax\n  preferredMax\n  referenceStepSize\n  preferredStepSize\n  rangeMergeInfo {\n    isBottomMerged\n    isTopMerged\n    __typename\n  }\n  referenceUnit {\n    abbreviation\n    __typename\n  }\n  preferredUnit {\n    abbreviation\n    __typename\n  }\n  rangeFilterDataPoint {\n    ...RangeFilterDataPointResult\n    __typename\n  }\n  __typename\n}\n\nfragment ProductMandatorSpecific on MandatorSpecificData {\n  isBestseller\n  isDeleted\n  showroomSites\n  sectorIds\n  hasVariants\n  __typename\n}\n\nfragment ProductMandatorIndependent on ProductV2 {\n  id\n  productId\n  name\n  nameProperties\n  productTypeId\n  productTypeName\n  brandId\n  brandName\n  averageRating\n  totalRatings\n  totalQuestions\n  isProductSet\n  images {\n    url\n    height\n    width\n    __typename\n  }\n  energyEfficiency {\n    energyEfficiencyColorType\n    energyEfficiencyLabelText\n    energyEfficiencyLabelSigns\n    energyEfficiencyImage {\n      url\n      height\n      width\n      __typename\n    }\n    __typename\n  }\n  seo {\n    seoProductTypeName\n    seoNameProperties\n    productGroups {\n      productGroup1\n      productGroup2\n      productGroup3\n      productGroup4\n      __typename\n    }\n    gtin\n    __typename\n  }\n  smallDimensions\n  basePrice {\n    priceFactor\n    value\n    __typename\n  }\n  productDataSheet {\n    name\n    languages\n    url\n    size\n    __typename\n  }\n  __typename\n}\n\nfragment ProductOffer on OfferV2 {\n  id\n  productId\n  offerId\n  shopOfferId\n  price {\n    amountIncl\n    amountExcl\n    currency\n    __typename\n  }\n  deliveryOptions {\n    mail {\n      classification\n      futureReleaseDate\n      __typename\n    }\n    pickup {\n      siteId\n      classification\n      futureReleaseDate\n      __typename\n    }\n    detailsProvider {\n      productId\n      offerId\n      quantity\n      type\n      __typename\n    }\n    __typename\n  }\n  label\n  labelType\n  type\n  volumeDiscountPrices {\n    minAmount\n    price {\n      amountIncl\n      amountExcl\n      currency\n      __typename\n    }\n    isDefault\n    __typename\n  }\n  salesInformation {\n    numberOfItems\n    numberOfItemsSold\n    isEndingSoon\n    validFrom\n    __typename\n  }\n  incentiveText\n  isIncentiveCashback\n  isNew\n  isSalesPromotion\n  hideInProductDiscovery\n  canAddToBasket\n  hidePrice\n  insteadOfPrice {\n    type\n    price {\n      amountIncl\n      amountExcl\n      currency\n      __typename\n    }\n    __typename\n  }\n  minOrderQuantity\n  __typename\n}\n\nfragment RangeFilterDataPointResult on RangeFilterDataPoint {\n  count\n  referenceValue {\n    value\n    unit {\n      abbreviation\n      __typename\n    }\n    __typename\n  }\n  preferredValue {\n    value\n    unit {\n      abbreviation\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n"
			}
		]`)

		req, err := http.NewRequest("POST", _url, bytes.NewBuffer(jsonData))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

		client := &http.Client{}
		resp, err := client.Do(req)
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

	if err := json.Unmarshal(_body, &_result); err != nil {
		panic(err)
	}
	// fmt.Println(_result[0].Data.ProductType.FilterProducts.Products.Results)

	r := regexp.MustCompile("[^a-z0-9 ??-]")

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result[0].Data.ProductType.FilterProducts.Products.Results))
		for _, result := range _result[0].Data.ProductType.FilterProducts.Products.Results {
			product := result.Product
			offer := result.Offer

			product.Brand = html.UnescapeString(product.Brand)
			product.Name = html.UnescapeString(product.Name)
			product.Description = html.UnescapeString(product.Description)

			if Skip(product.Brand) {
				// fmt.Println("** SKIP: " + product.Brand)
				continue
			}

			var _retailPrice float32
			var _price float32
			var _savings float32
			var _discount float32

			_price = offer.Price.Amount
			_retailPrice = _price

			if offer.OldPrice.Price.Amount > 0 {
				_retailPrice = offer.OldPrice.Price.Amount
			}

			if _savings == 0 {
				_savings = _price - _retailPrice
			}
			_discount = 100 - ((100 / _retailPrice) * _price)

			// 0.03??GB, Black, 2.40\", Hybrid Dual SIM, 0.30??Mpx, 2G
			// 003-gb-black-240-hybrid-dual-sim-030-mpx-2g
			// 0-03gb-black-2-40-hybrid-dual-sim-0-30mpx-2g

			// https://www.galaxus.ch/de/s1/product/blaupunkt-fm-01-slider-2g-003-gb-black-240-hybrid-dual-sim-030-mpx-2g-smartphone-10336937?shid=979245
			// https://www.galaxus.ch/de/s1/product/blaupunkt-fm-01-slider-2g-003-gb-black-240-hybrid-dual-sim-030-mpx-2g-smartphone-10336937

			_productName := strings.NewReplacer(" ", "-", "??", "-").Replace(r.ReplaceAllString(strings.ToLower(product.Brand+"-"+product.Name+"-"+product.Description+"-"+product.Category), "$1"))
			_productUrl := fmt.Sprintf("https://www.galaxus.ch/de/s1/product/%s-%d", _productName, product.Code)
			if !result.IsDefaultOffer {
				_productUrl += fmt.Sprintf("?shid=%d", offer.ShopOfferID)
			}

			{
				code := strconv.Itoa(product.Code)
				product := &Product{
					Code: _name + "//" + code,
					Name: product.Brand + " " + product.Name,

					RetailPrice: _retailPrice,
					Price:       _price,
					Savings:     _savings,
					Discount:    _discount,

					URL: _productUrl,
				}

				if s.IsWorth(product) {
					products = append(products, product)
				}
			}
		}

		return &products
	}

	return NewShop(
		_name,
		_url,

		_parseFn,
	)
}
