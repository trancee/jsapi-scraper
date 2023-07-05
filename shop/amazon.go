package shop

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// r := regexp.MustCompile(`(?i)\W*((Handys?|(4G )?Smartphones?)( mit)?|ohne Vertragy?,?|Outdoor|(\+\W*)Kopfhörer|Günstige?,?|Telekom|Wasserdichit|50MP\+8MP Kamera,|OTG Reverse Charge|(\d{4,5}|\d{1,3}\.\d{3})mAh(\W*(Großer )?Akku)?|\(?\s*202\d\)?|\(\d+\+\d+GB\),|\d+ \+ \d+\s*GB|Android \d+( 4GB?)?)`)
// r := regexp.MustCompile(`(?i)\W*((Handys?|(4G )?Smartphones?)( mit)?|ohne Vertragy?,?|Outdoor|(\+\W*)Kopfhörer|Günstige?,?|Telekom|Wasserdichit|50MP\+8MP Kamera,|OTG Reverse Charge|Erweiterbar|Octa\W*Core(\W*Pro[cz]essor)?|(Starker )?(\d{4,5}|\d{1,3}\.\d{3})\s*mAh(\W*(Großer )?Akku)?|\(?\s*202\d\)?|\d+(GB)?\s*\+\s*\d+\s*GB(\/\d+[GT]B)?\)?,?|Android \d+)`)
// r := regexp.MustCompile(`(?i)\s*([ ，]|(Handys?|(4G )?Smartphones?)( mit)?|ohne Vertragy?,?|Outdoor|(\+\W*)Kopfhörer|Günstige?,?|Telekom|Wasserdichit|50MP\+8MP (Dual )?Kamera,|OTG Reverse Charge|Erweiterbar|Octa\W*Core(\W*Pro[cz]essor)?|(Starker )?(\d{4,5}|\d{1,3}\.\d{3})\s*mAh(\W*(Großer )?Akku)?|[(（]?\s*202\d[)）]?|\W*\d+(GB)?\s*\+\s*\d+\s*GB(\/\d+[GT]B)?\)?,?|Android \d+)`)
// var AmazonRegex = regexp.MustCompile(`(?i)\s*([ ，]|(Handys?|(4G )?Smartphones?)( mit)?|ohne Vertragy?,?|Outdoor|(\+\W*)Kopfhörer|Günstige?,?|Telekom|Wasserdichit|50MP\+8MP (Dual )?Kamera,|OTG Reverse Charge|Erweiterbar|Octa\W*Core(\W*Pro[cz]essor)?|(Starker )?(\d{4,5}|\d{1,3}\.\d{3})\s*mAh(\W*(Großer )?Akku)?|[(（]?\s*202\d[)）]?|\W*\d+(GB)?\s*\+\s*\d+\s*GB(\/\d+[GT]B)?\)?,?|Android \d+)`)
// var AmazonRegex = regexp.MustCompile(`(?i)\s*([ ，]|(Handys?|(4G )?Smartphones?)( mit)?|ohne Vertragy?,?|(4G )?Outdoor|(\+\W*)Kopfhörer|Günstige?,?|Telekom|Wasserdichi?t|50MP\+8MP (Dual )?Kamera,|OTG Reverse Charge|Erweiterbar|Octa\W*Core(\W*Pro[cz]essor)?|(Starker )?(\d{4,5}|\d{1,3}\.\d{3})\s*mAh(\W*(Großer )?Akku)?|[（]?\s*20[12]\d[）]?|\W*\d+(GB)?\s*\+\s*\d+\s*GB(\/\d+[GT]B)?\)?,?|Android \d+)`)
// var AmazonRegex = regexp.MustCompile(`(?i)\s*(((4G |Lockfreie )?(Handys?|Smartphones?))( mit)?|ohne Vertragy?,?(\d\.\d+'*( Zoll HD\+)?)?|(4G )?Outdoor|(\+\W*)Kopfhörer|Günstig(,|es|e)?|Neu|Telekom|(IP\d+\s+)?Wasserdichi?t(er)?|\d+MP(\+8MP)?\W+(AI\W*)?(Dual\W+|Quad\W+|Unterwasser)?Kamera|Dual\W+SIM(\+SD \(.*?\))?|\d Zoll Touch Bildschirm,|EU 128GB|OTG Reverse Charge|Cloud Navy|Erweiterbar|Octa\W*Core(\W*Pro[cz]essor)?|(Großer?|Größten) Akku|(Starker )?(\d{4,5}|\d{1,3}\.\d{3})\s*mAh(\W*(Großer )?(Akku|Batterie))?|\b20[12]\d|\W*\d+(GB)?\s*\+\s*\d+\s*GB([\/+]\d+[GT]B)?\)?,?|Android \d+(\.\d)?( Go)?|(SM )?[SG]\d{3}[A-Z]*)`)
// var AmazonRegex2 = regexp.MustCompile(`(?i)^(.*?)(\s+\(?\dG\W*|\s*\d+\W*([GT]B|W)|\W\d+[,.]\d+|\s*–\s*|\s*Android| Helio | mit | Octa |,)`)
var AmazonRegex3 = regexp.MustCompile(`(Android \d{1,2}( Go)?|Quad Core |Telekom |Telefon |All Carriers |Neu |EU |Xia |MOVIL |Smartfon |Marke Modell |Sam |Cellulare |Unlocked |Senior | Master Edition|\(Hybrid SIM\))\s*|(-?4G )?(Simlockfreie |Lockfreie |Outdoor |Android[ -]|SIM Free )?(Handys?|Smartphones?)( [Oo]hne [Vv]ertragy?,?)?( Günstig,?)?|(\W*\d+(GB)?\s*\+\s*\d+\s*GB\W*)|\W*\d+([,.]\d+)?\s*(cm|\"|''|')|(Dual|DUAL)\W+(SIM|Sim)|\d+MP(\+8MP)?\W+(AI\W*)?(Dual\W+|Quad\W+|Unterwasser)?Kamera|\(?5G|\d{4,5}mAh( Akku)?|Blau|Blue|Buds|Cinemagic|Cloud Navy|Dark Silver|Glacier Blue|Gradient Bronze|Graphite Gray|Grau|Midnight Gray|Mint Green|\(?Ocean Blue\)?|Onyx Gray|Oro|Pebble White|Polar White|Sunrise Orange|White|\/Black|\/?BLUE|\/?GREEN|\/?ORANGE|GRIS|\W+\(?20[12]\d\)?| \+ 5G| Enterprise Edition`)
var AmazonRegex4 = regexp.MustCompile(`\s*-?\(?\d+([+\/]\d+)?\s*(GB|TB|gb)|\d\+\d+G|\W*[45][Gg](\s+|$)?| DS| EU| NFC| -|,|\s+\(\d{2}| 32-2| 2\+32| 4\+64| 4\+128| 128-4`)

var AmazonCleanFn = func(name string) string {
	name = regexp.MustCompile(`\d{5}[A-Z]{3}|RM-\d{4}|SIPP5 |\/Motorola PA4N0106IT|MOBILE PHONE |(XIA|REA) DS |3\. Generation|all carriers ,|^[-0] `).ReplaceAllString(name, "")
	name = strings.NewReplacer(" ", " ", "，", ",", "（", "(", "）", ")", "–", "|", "‎", "", "Kingkong", "King Kong", "KXD Handy,", "KXD", "Mobile Phone", "", "TELEFONO MOVIL", "", "Telefonas ", " ", "Mobility", "", "Galaxy-A", "Galaxy A", " A 90", " A90", " M5/", " M5|", "8GBRAM128GBROM", " ").Replace(name)
	name = AmazonRegex3.ReplaceAllString(name, "|")

	if s := strings.Split(name, "|"); len(s) > 0 {
		_name := strings.TrimSpace(s[0])
		if len(strings.Split(_name, " ")) == 1 {
			for i := 1; i < len(s); i++ {
				if name := strings.TrimSpace(s[i]); name != "" {
					// fmt.Println("[" + name + "]")
					_name += " " + name
					break
				}
			}
		}
		name = strings.TrimSpace(_name)
	}

	s := strings.Split(name, " ")

	if s[0] == "HUAWEI" || s[0] == "Huawei" {
		name = strings.ReplaceAll(name, "Mate10", "Mate 10")
		name = strings.ReplaceAll(name, "Mate20", "Mate 20")
	}

	if s[0] == "Honor" || s[0] == "HONOR" {
		name = regexp.MustCompile(`Magic\s*(\d)\s*(\w)?`).ReplaceAllString(name, "Magic$1 $2")
	}

	if s[0] == "moto" || s[0] == "Moto" {
		name = "Motorola " + name
	}
	if s[0] == "Motorola" {
		name = strings.NewReplacer("Light", "Lite", " E ", " E").Replace(name)

		if len(s) > 1 && (s[1] == "Moto" || s[1] == "moto") && (s[2] == "Edge" || s[2] == "edge") {
			name = strings.ReplaceAll(name, "Moto ", "")
		}
		if len(s[1]) > 1 && (s[1][0:1] == "e" || s[1][0:1] == "E" || s[1][0:1] == "g" || s[1][0:1] == "G") && s[1][1:2] >= "0" && s[1][1:2] <= "9" {
			name = strings.ReplaceAll(name, "Motorola ", "Motorola Moto ")
		}

		name = regexp.MustCompile(`(?i)edge\s*(\d+)\s*(\w*)`).ReplaceAllString(name, "edge $1 $2")
	}

	if s[0] == "OPPO" || s[0] == "Oppo" {
		name = regexp.MustCompile(`Reno\s*(\d)\s*(\w)?`).ReplaceAllString(name, "Reno$1 $2")
	}

	if s[0] == "Galaxy" {
		name = "Samsung " + name
	}

	if s[0] == "Redmi" || s[0] == "REDMI" || s[0] == "POCO" {
		name = "Xiaomi " + name
	}
	if s[0] == "Samsung" {
		name = regexp.MustCompile(`\s+(SM-)?[AFMS]\d{3}[BFR]?(\/DSN?)?`).ReplaceAllString(name, "")
		name = strings.ReplaceAll(name, "Samsung A", "Samsung Galaxy A")
		name = strings.ReplaceAll(name, "Samsung M", "Samsung Galaxy M")
	}
	if s[0] == "Xiaomi" {
		name = strings.ReplaceAll(name, "Xiaomi 10", "Xiaomi Redmi 10")
		name = strings.ReplaceAll(name, "Xiaomi M5", "Xiaomi POCO M5")
		name = strings.ReplaceAll(name, "Xiaomi Note ", "Xiaomi Redmi Note ")
	}
	if s[0] == "ZTE" {
		name = strings.ReplaceAll(name, "V40 S", "V40S")

		if strings.HasSuffix(name, "Vita") && s[1] != "Blade" {
			name = strings.ReplaceAll(name, "ZTE", "ZTE Blade")
		}
	}

	if loc := AmazonRegex4.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%s\t%s\n", loc, name[:loc[0]], name)
		_model := name[:loc[0]]
		// // fmt.Printf("%v\t%s\t%s\n", loc, name[loc[2]:loc[3]], name)
		// _model := name[loc[2]:loc[3]]
		// fmt.Println(_model)
		// _model = strings.TrimSpace(strings.Trim(_model, "("))
		// fmt.Println(_model)
		// _model = strings.TrimSpace(strings.Split(strings.ReplaceAll(_model, "()", "|"), "|")[0])
		// fmt.Println(_model)
		return regexp.MustCompile(`\s+`).ReplaceAllString(strings.TrimSpace(_model), " ")
	}

	return regexp.MustCompile(`\s+`).ReplaceAllString(strings.TrimSpace(name), " ")
}

func XXX_amazon(isDryRun bool) IShop {
	const _name = "Amazon"
	// const _url = "https://www.amazon.de/s?k=SIM-Free+&+Unlocked+Mobile+Phones&i=electronics&rh=n:15326400031&s=price-asc-rank&c=ts&qid=1678973932&ts_id=15326400031&ref=sr_st_price-asc-rank&ds=v1:pfeFMDyZ0TLfvHopZLWOdvFgfEilJ0+V3TRCd10npNE"
	// const _url = "https://www.amazon.de/-/en/s?k=SIM-Free+%26+Unlocked+Mobile+Phones&i=electronics&rh=n%3A15326400031%2Cp_n_free_shipping_eligible%3A20943778031%2Cp_n_condition-type%3A776949031%2Cp_n_deal_type%3A26902994031&dc&language=de_CH&currency=CHF&c=ts&qid=1678974685&rnid=26902991031&ts_id=15326400031&ref=sr_nr_p_ru_1&ds=v1%3AqBW5BcRHmqi2sjFHp4vDX0aNYSc4IsvDBroYiSaiqkA"
	// const _url = "https://www.amazon.de/s?k=Simlockfreie+Handys&i=electronics&rh=n:15326400031,p_n_free_shipping_eligible:20943778031,p_n_deal_type:26902994031,p_n_condition-type:776949031&language=de_DE&currency=CHF&dc=&c=ts&qid=1678975426&rnid=776942031&page=2"
	// const _url = "https://www.amazon.de/s?k=Simlockfreie+Handys&i=electronics&rh=n:15326400031,p_n_free_shipping_eligible:20943778031,p_n_deal_type:26902994031,p_n_condition-type:776949031&dc=&c=ts&qid=1678975426&rnid=776942031&s=price-asc-rank&page=%d"
	// const _url = "https://www.amazon.de/s?k=Simlockfreie+Handys&i=electronics&rh=n:15326400031,p_n_free_shipping_eligible:20943778031,p_n_deal_type:26902993031&s=price-asc-rank&dc&c=ts&qid=1679414936&rnid=26902991031&page=%d"
	// const _url = "https://www.amazon.de/s?k=Simlockfreie+Handys&i=electronics&rh=n:15326400031,p_n_free_shipping_eligible:20943778031,p_n_deal_type:26902994031,p_6:A3JWKAKR8XB7XF&s=price-asc-rank&dc&c=ts&qid=1680358125&rnid=26902991031&page=%d"
	_url := fmt.Sprintf("https://www.amazon.de/s?k=Simlockfreie+Handys&i=electronics&rh=n:15326400031,p_n_free_shipping_eligible:20943778031,p_6:A3JWKAKR8XB7XF&dc=&c=ts&qid=1681744980&rnid=389294011&low-price=%.f&high-price=%.f&page=%%d", ValueMinimum, ValueMaximum)

	const _debug = false

	type _Response struct {
		code  string
		title string

		link string

		oldPrice float32
		price    float32
	}

	var _result []_Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	for p := 1; p <= 20; p++ {
		fn := fmt.Sprintf("shop/amazon.%d.html", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			req, err := http.NewRequest("GET", fmt.Sprintf(_url, p), nil)
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
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

			req.AddCookie(
				&http.Cookie{
					Name:   "lc-acbde",
					Value:  "de_DE",
					Path:   "/",
					MaxAge: 3600,
					Secure: true,
				},
			)

			req.AddCookie(
				&http.Cookie{
					Name:   "i18n-prefs",
					Value:  "CHF",
					Path:   "/",
					MaxAge: 3600,
					Secure: true,
				},
			)

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

		doc := parse(string(_body))

		if productList := traverse(doc, "div", "class", "s-search-results"); productList != nil {
			// fmt.Println(productList)

			for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling.NextSibling {
				if !contains(item.Attr, "data-component-type", "s-search-result") {
					continue
				}

				// item := traverse(items, "div", "class", "s-result-item")
				// fmt.Println(item)

				_product := _Response{}

				asin, _ := attr(item.Attr, "data-asin")
				if _debug {
					fmt.Println(asin)
				}
				_product.code = asin

				itemImage := traverse(item, "img", "class", "s-image")
				// fmt.Println(itemImage)

				title, _ := attr(itemImage.Attr, "alt")
				if brand := strings.ToUpper(strings.Split(title, " ")[0]); brand == "POCO" {
					title = "Xiaomi " + title
				}
				if _debug {
					fmt.Println(title)
				}
				_product.title = title

				if strings.Contains(title, "Outdoor") || (len(title) > 7 && title[0:7] == "emporia") {
					continue
				}

				if Skip(title) {
					continue
				}

				imageTitleLink := traverse(item, "a", "class", "a-text-normal")
				// fmt.Println(imageTitleLink)

				link, _ := attr(imageTitleLink.Attr, "href")
				if _debug {
					fmt.Println(link)
				}
				_product.link = link

				if itemPrice := traverse(item, "div", "class", "s-price-instructions-style"); itemPrice != nil {
					// fmt.Println(itemPrice)

					if itemOldPrice := traverse(itemPrice, "span", "class", "a-price-whole"); itemOldPrice != nil {
						// fmt.Println(itemOldPrice)

						oldPrice, _ := text(itemOldPrice)
						oldPrice = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(oldPrice, ".", ""), ",", "."), "CHF", "")
						if _debug {
							fmt.Println(oldPrice)
						}

						if _price, err := strconv.ParseFloat(oldPrice, 32); err != nil {
							panic(err)
						} else {
							_product.price = float32(_price)
						}
					}

					if itemPrice := traverse(itemPrice, "span", "class", "a-text-price"); itemPrice != nil {
						// fmt.Println(currentPrice)

						if currentPrice := traverse(itemPrice, "span", "aria-hidden", "true"); currentPrice != nil {
							// fmt.Println(currentPrice)

							price, _ := text(currentPrice)
							price = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(price, ".", ""), ",", "."), "CHF", "")
							if _debug {
								fmt.Println(price)
							}

							if _price, err := strconv.ParseFloat(price, 32); err != nil {
								panic(err)
							} else {
								_product.oldPrice = float32(_price)
							}
						}
					}
				}

				if _debug {
					fmt.Println()
				}

				_result = append(_result, _product)
			}

			resultInfo := traverse(doc, "span", "data-component-type", "s-result-info-bar")
			results := traverse(resultInfo, "span", "", "")
			if result, ok := text(results); ok {
				if x := regexp.MustCompile(`(\d+)-(\d+) von (\d+)`).FindStringSubmatch(result); x != nil && x[2] == x[3] {
					break
				}
			}
		}
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result))

		// _products := map[string]int{}
		// keys := make([]string, 0, len(_result))
		// for i, k := range _result {
		// 	keys = append(keys, k.title)
		// 	_products[k.title] = i
		// }
		// sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

		// for _, key := range keys {
		// 	product := _result[_products[key]]
		for _, product := range _result {
			// fmt.Println(product)

			_title := product.title
			// fmt.Println(_title)
			// fmt.Println("\"" + strings.ReplaceAll(_title, "\"", "\\\"") + "\",")
			_model := AmazonCleanFn(_title)
			// fmt.Println("\"" + strings.ReplaceAll(_model, "\"", "\\\"") + "\",")

			if Skip(_model) {
				continue
			}

			_retailPrice := product.price
			_price := _retailPrice
			if product.oldPrice > 0 {
				_retailPrice = product.oldPrice
			}

			_savings := _price - _retailPrice
			_discount := 100 - ((100 / _retailPrice) * _price)

			_link := s.ResolveURL(product.link).String()
			_link = strings.Split(_link, "/ref=")[0]

			{
				product := &Product{
					Code:  _name + "//" + product.code,
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

		// fmt.Printf("%#v\n", products)
		return &products
	}

	return NewShop(
		_name,
		_url,

		_parseFn,
	)
}
