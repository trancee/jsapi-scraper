package shop

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

const ValueDiscount = 50.0
const ValueWorth = 100.0

const ValueMinimum = 50.0
const ValueMaximum = 350.0

type IShop interface {
	Name() string

	CanFetch() bool
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
	Code  string `json:"code"`
	Name  string `json:"name"`
	Model string `json:"model"`

	RetailPrice float32 `json:"oldPrice"`
	Price       float32 `json:"price"`
	Savings     float32 `json:"savings"`
	Discount    float32 `json:"discount"`

	Quantity int `json:"quantity"`

	URL string `json:"link"`
}

type parseFn func(s IShop) *[]*Product

var _skips = map[string]bool{
	"ACER":            true,
	"ALCATEL":         true,
	"ALIGATOR":        true,
	"AMPLICOMMS":      true,
	"AURO":            true,
	"ARTFONE":         true,
	"BEAFON":          true,
	"BLABLOO":         true,
	"BLACKBERRY":      true,
	"BLAUPUNKT":       true,
	"BRONDI":          true,
	"CARBON":          true,
	"CAT":             true,
	"CATERPILLAR":     true,
	"CELLULARE":       true,
	"CROSSCALL":       true,
	"CUBOT":           true, // EXCLUDE
	"CYRUS":           true,
	"DENVER":          true,
	"DEUTSCHE":        true,
	"DOOGEE":          true, // EXCLUDE
	"DOPOD":           true,
	"DORO":            true,
	"DREAME":          true,
	"EL":              true,
	"EMPORIA":         true,
	"EMPORIAEUPHORIA": true,
	"ENERGIZER":       true,
	"EVOLVEO":         true,
	"FELLOWES":        true,
	"FOLIA":           true,
	"FUNKE":           true,
	"HAGENUK":         true,
	"HAMMER":          true,
	"HOP":             true,
	"HP":              true,
	"I.SAFE":          true,
	"IGET":            true,
	"IIIF150":         true,
	"JABLOCOM":        true,
	"KERKMANN":        true,
	"KONROW":          true,
	"KONTAKT":         true,
	"KRÜGER&MATZ":     true,
	"KXD":             true, // EXCLUDE
	"LEIOA":           true,
	"LENOVO":          true,
	"LEXIBOOK":        true,
	"LG":              true, // EXCLUDE
	"LOGICOM":         true,
	"LUMIA":           true,
	"MAGNETOPLAN":     true,
	"MAUL":            true,
	"MAXCOM":          true,
	"MEDIACOM":        true,
	"MICROSOFT":       true, // EXCLUDE
	"MITEL":           true,
	"MP":              true,
	"MYPHONE":         true,
	"OGO":             true,
	"OLYMPIA":         true,
	"ORDISSIMO":       true,
	"OSCAL":           true, // EXCLUDE
	"OUKITEL":         true, // EXCLUDE
	"PALM":            true,
	"PANASONIC":       true,
	"PANTECH":         true,
	"PEAQ":            true,
	"POLAROID":        true,
	"POWERVISION":     true,
	"PRIMO":           true,
	"PUNKT.":          true,
	"QUBO":            true,
	"RUG":             true,
	"RUGGEAR":         true,
	"SAGEM":           true,
	"SGIN":            true,
	"SGW":             true,
	"SIGEL":           true,
	"SIMVALLEY":       true,
	"SONIM":           true,
	"SP":              true,
	"SPC":             true,
	"STOTZ":           true,
	"STYRO":           true,
	"SUNSTECH":        true,
	"SUUNTO":          true,
	"SWISSTONE":       true,
	"SWISSVOICE":      true,
	"SWITEL":          true,
	"TCL":             true,
	"TECNO":           true,
	"TELEFON":         true,
	"TELEFONAS":       true,
	"TELEFUNKEN":      true,
	"TELEKOM":         true,
	"TELME":           true,
	"TREVI":           true,
	"TTFONE":          true,
	"ULEFONE":         true, // EXCLUDE
	"ULEWAY":          true,
	"UMI":             true,
	"UMIDIGI":         true, // EXCLUDE
	"UNBEKANNT":       true,
	"UNIFY":           true,
	"WEDO":            true,
	"XGODY":           true, // EXCLUDE
	"XS13":            true,
	"ZANCO":           true,
	"ZEBRA":           true,
}

func Skip(brand string) bool {
	// fmt.Println("** SKIP: " + brand)
	if s := strings.Split(brand, " "); len(s) > 0 {
		_brand := strings.ToUpper(strings.ReplaceAll(s[0], "-", ""))

		if _brand == "NOKIA" && len(s) > 1 {
			if _, err := strconv.Atoi(s[1]); err == nil {
				return true
			}
		}

		return _skips[_brand]
	}

	return false
}

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

func (s shop) CanFetch() bool {
	return s.parseFn != nil
}

func (s shop) Fetch() *[]*Product {
	return s.parseFn(s)
}

func (s shop) IsWorth(product *Product) bool {
	return (product.Price > 0 && product.Price < ValueMaximum) || (product.Discount >= ValueDiscount)
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
		// fmt.Printf("ATTR [%v|%v] [%v|%v]\n", a.Key, key, a.Val, val)
		if a.Key == key {
			if val == "" || a.Val == val {
				return true
			} else {
				matched, _ := regexp.Match(`\b`+strings.ReplaceAll(val, "-", "_")+`\b`, []byte(strings.ReplaceAll(a.Val, "-", "_")))
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
