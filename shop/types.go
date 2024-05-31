package shop

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/net/html"
)

var EUR_CHF = 1.0

const ValueDiscount = 75.0
const ValueWorth = 100.0

const ValueMinimum = 0.0
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
	EURPrice    float32 `json:"eurPrice"`
	Savings     float32 `json:"savings"`
	Discount    float32 `json:"discount"`

	Quantity int `json:"quantity"`

	URL string `json:"link"`

	Counter          int64 `json:"counter"`
	CreationDate     int64 `json:"createdAt"`
	ModificationDate int64 `json:"modifiedAt"`
	NotificationDate int64 `json:"notifiedAt"`
}

type parseFn func(s IShop) *[]*Product

var _skips = map[string]bool{
	"'MASTER":         true,
	"30":              true,
	"ACER":            true,
	"AGM":             true,
	"ALCATEL":         true,
	"ALIGATOR":        true,
	"AMPLICOMMS":      true,
	"ARCHOS":          true,
	"AURO":            true,
	"ARTFONE":         true,
	"ATHESI":          true,
	"BEAFON":          true,
	"BEGHELLI":        true,
	"BLABLOO":         true,
	"BLACKBERRY":      true,
	"BLAUPUNKT":       true,
	"BOOKSENSE":       true,
	"BRONDI":          true,
	"BULLIT":          true,
	"CARBON":          true,
	"CAT":             true,
	"CATERPILLAR":     true,
	"CELLULARE":       true,
	"CLEMENTONI":      true,
	"CROSSCALL":       true,
	"CUBOT":           true, // EXCLUDE
	"CUSTOM":          true,
	"CYRUS":           true,
	"DENSO":           true,
	"DENVER":          true,
	"DEUTSCHE":        true,
	"DOOGEE":          true, // EXCLUDE
	"DOPOD":           true,
	"DORO":            true,
	"DREAME":          true,
	"EL":              true,
	"EMPORIA":         true,
	"EMPORIAEUPHORIA": true,
	"EMPORIASMART":    true,
	"EMPORIATOUCH":    true,
	"ENERGIZER":       true,
	"EVOLVEO":         true,
	"FELLOWES":        true,
	"FOLIA":           true,
	"FUNKE":           true,
	"FYSIC":           true,
	"HAGENUK":         true,
	"HAMMER":          true,
	"HISENSE":         true,
	"HOP":             true,
	"HOTWAV":          true,
	"HP":              true,
	"I.SAFE":          true,
	"IBASSO":          true,
	"IGET":            true,
	"IIIF150":         true,
	"INAB":            true,
	"JABLOCOM":        true,
	"KERKMANN":        true,
	"KONROW":          true,
	"KONTAKT":         true,
	"KRUGER":          true,
	"KRUGER&MATZ":     true,
	"KRÜGER&MATZ":     true,
	"KXD":             true, // EXCLUDE
	"LANCOM":          true,
	"LEIOA":           true,
	"LENOVO":          true,
	"LEXIBOOK":        true,
	"LG":              true, // EXCLUDE
	"LOGICOM":         true,
	"LUMIA":           true,
	"MAGNETOPLAN":     true,
	"MARSHALL":        true,
	"MAUL":            true,
	"MAXCOM":          true,
	"MEDIACOM":        true,
	"MFOX":            true,
	"MICROSOFT":       true, // EXCLUDE
	"MIGNON":          true,
	"MITEL":           true,
	"MOBISTEL":        true,
	"MP":              true,
	"MYPHONE":         true,
	"NEFFOS":          true,
	"NGM":             true,
	"NOMU":            true,
	"OEM":             true,
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
	"PREMIER":         true,
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
	"SKROSS":          true,
	"SONIM":           true,
	"SP":              true,
	"SPARK":           true,
	"SPC":             true,
	"STOTZ":           true,
	"STYRO":           true,
	"SUNSTECH":        true,
	"SUUNTO":          true,
	"SWISSTONE":       true,
	"SWISSVOICE":      true,
	"SWITEL":          true,
	"SYCO":            true,
	"TCL":             true,
	"TECNO":           true,
	"TELECOM":         true,
	"TELEFON":         true,
	"TELEFONAS":       true,
	"TELEFUNKEN":      true,
	"TELEKOM":         true,
	"TELME":           true,
	"TOPCO":           true,
	"TREVI":           true,
	"TTFONE":          true,
	"ULEFONE":         true, // EXCLUDE
	"ULEWAY":          true,
	"UMI":             true,
	"UMIDIGI":         true, // EXCLUDE
	"UNBEKANNT":       true,
	"UNIFY":           true,
	"WEDO":            true,
	"WOLDER":          true,
	"XGODY":           true, // EXCLUDE
	"XS13":            true,
	"YEZZ":            true,
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
	if n != nil && n.Type == html.TextNode {
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
				matched, _ := regexp.Match(`\b`+strings.ReplaceAll(val, "-", "_")+`\b`, StringToBytes(strings.ReplaceAll(a.Val, "-", "_")))
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

// https://josestg.medium.com/140x-faster-string-to-byte-and-byte-to-string-conversions-with-zero-allocation-in-go-200b4d7105fc
func BytesToString(b []byte) string {
	// Ignore if your IDE shows an error here; it's a false positive.
	p := unsafe.SliceData(b)
	return unsafe.String(p, len(b))
}

// https://josestg.medium.com/140x-faster-string-to-byte-and-byte-to-string-conversions-with-zero-allocation-in-go-200b4d7105fc
func StringToBytes(s string) []byte {
	p := unsafe.StringData(s)
	b := unsafe.Slice(p, len(s))
	return b
}
