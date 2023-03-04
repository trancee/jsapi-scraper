package shop

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type IShop interface {
	Fetch() []Product
	ResolveURL(refURL string) *url.URL
}

type shop struct {
	name string
	url  string

	baseURL *url.URL

	result any
	// products []Product

	parseFn parseFn
}

type Product struct {
	Code string
	Name string

	RetailPrice float32
	Price       float32
	Savings     float32
	Discount    float32

	URL string
}

type parseFn func() []Product

func NewShop(_name string, _url string, _response any, _parseFn parseFn) IShop {
	_baseURL, err := url.Parse(_url)
	if err != nil {
		panic(err)
	}

	return shop{
		name: _name,
		url:  _url,

		baseURL: _baseURL,

		result: _response,

		parseFn: _parseFn,
	}
}

func (s shop) Fetch() []Product {
	resp, err := http.Get(s.url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))

	if body[0] != '{' {
		body = body[2:(len(body) - 2)]
		// fmt.Println(string(body))
	}

	// var result Response
	if err := json.Unmarshal(body, &s.result); err != nil { // Parse []byte to go struct pointer
		panic(err)
	}
	// fmt.Printf("%v\n", s.result)

	// for _, product := range s.parseFn() {
	// 	fmt.Printf("%#v\n", product)
	// }
	return s.parseFn()
}

func (s shop) ResolveURL(refURL string) *url.URL {
	ref, err := url.Parse(refURL)
	if err != nil {
		panic(err)
	}

	return s.baseURL.ResolveReference(ref)
}
