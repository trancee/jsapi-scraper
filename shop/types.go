package shop

import (
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type IShop interface {
	Name() string

	Fetch() *[]*Product
	IsWorth(product *Product) bool

	ResolveURL(refURL string) *url.URL
}

type shop struct {
	name string
	url  string

	baseURL *url.URL

	parseFn parseFn
}

type Action struct {
	MaxPrice    float32
	MaxDiscount float32
}

type Product struct {
	Code string
	Name string

	RetailPrice float32
	Price       float32
	Savings     float32
	Discount    float32

	Quantity int

	URL string
}

type parseFn func(s IShop) *[]*Product

func NewShop(_name string, _url string, _parseFn parseFn) IShop {
	_baseURL, err := url.Parse(_url)
	if err != nil {
		panic(err)
	}

	return shop{
		name: _name,
		url:  _url,

		baseURL: _baseURL,

		parseFn: _parseFn,
	}
}

func (s shop) Name() string {
	return s.name
}

func (s shop) Fetch() *[]*Product {
	return s.parseFn(s)
}

func (s shop) IsWorth(product *Product) bool {
	return (product.Price > 0 && product.Price < 150) || (product.Discount >= 75)
}

func (s shop) ResolveURL(refURL string) *url.URL {
	ref, err := url.Parse(refURL)
	if err != nil {
		panic(err)
	}
	return s.baseURL.ResolveReference(ref)
}

func text(n *html.Node) (string, bool) {
	if n.Type == html.TextNode {
		return strings.TrimSpace(html.UnescapeString(n.Data)), true
	}
	n = n.FirstChild
	if n.Type == html.TextNode {
		return strings.TrimSpace(html.UnescapeString(n.Data)), true
	}
	return "", false
}
func attr(s []html.Attribute, key string) (string, bool) {
	for _, a := range s {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}
func contains(s []html.Attribute, key string, val string) bool {
	for _, a := range s {
		// fmt.Printf("ATTR [%v] [%v]\n", a.Key, a.Val)
		if a.Key == key {
			if val == "" || a.Val == val {
				return true
			} else {
				matched, _ := regexp.Match(`\b`+val+`\b`, []byte(a.Val))
				return matched
			}
		}
	}
	return false
}

func parse(_html string) *html.Node {
	doc, err := html.Parse(strings.NewReader(_html))
	if err != nil {
		panic(err)
	}
	return doc
}

func traverse(n *html.Node, tag string, key string, val string) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// fmt.Printf("DataAtom[%v] Data[%v] Attr[%v] Tag[%v] Type[%v]\n", c.DataAtom, c.Data, c.Attr, c.Parent.Data, c.Type)
		if c.Type == html.ElementNode && c.Data == tag && ((key != "" && contains(c.Attr, key, val)) || key == "") {
			return c
		}

		if res := traverse(c, tag, key, val); res != nil {
			return res
		}
	}
	return nil
}
