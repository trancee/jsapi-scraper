package shop

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	helpers "jsapi-scraper/helpers"
)

var AmazonRegex3 = regexp.MustCompile(`(Android \d{1,2}( Go)?|Quad Core |Telekom |Telefon |All Carriers | WiFi| Octa Core|Neu |EU | SMD|Xia |MOVIL |Smartfon |Marke Modell |Samsung Galaxy Milwaukee|Sam |Cellulare |Unlocked |Senior | Master Edition|\(Hybrid SIM\))\s*|(-?4G )?(Simlockfreie |Lockfreie |Outdoor |Android[ -]|SIM Free |Schlankes und Elegantes )?(Handys?|Smartphones?)( [Oo]hne [Vv]ertr?agy?,?)?( Günstig,?)?|(\W*\d+(GB)?\s*\+\s*\d+\s*GB\W*)|\W*\d+([,.]\d+)?\s*(cm|\"|''|')|(Dual|DUAL)\W+(SIM|Sim)|\d+MP(\+8MP)?\W+(AI\W*)?(Dual\W+|Quad\W+|Unterwasser)?Kamera|\(?5G|\d{4,5}mAh( Akku)?|\d+\.\d+\s*mAh|Buds|Cinemagic|Dark Silver| \+ 5G|\d+,\d Zoll| LTE`)
var AmazonRegex4 = regexp.MustCompile(`\s*-?\(?\d+([+/]\d+)?\s*(GB|TB|gb)|\d\+\d+G|\W*[45][Gg](\s+|$)?| DS| EU| NFC| -|,|\s+\(\d{2}| 32-2| 2\+32| 3\+64| 4\+64| 64-3| 64/4| 4\+128| 128-4| 128-6(-5G)?| 128-8`)
var AmazonExclusionRegex = regexp.MustCompile(`(?i)Abdeckung|Adapter|AirTag|Armband|Band|CABLE|Charger|Ch?inch|Christbaum|Clamshell|^Core|\bCover\b|Display|Earphones|Etui|Halterung|Handschuhe|HARDCASE|Headset|Hülle|Kopfhörer|Ladegerät|Ladestation|Lautsprecher|Magnet|Majestic|Netzkabel|Objektiv|Protector|Reiselader|S Pen|Saugnapf|Schutzfolie|SmartTag|Smartwatch|Stand|Ständer|Stativ|Stylus|Tastatur|Virtual-Reality|Wasserdicht(es)?|Weihnachtsbaum`)

var AmazonCleanFn = func(name string) string {
	name = regexp.MustCompile(`\d{5}[A-Z]{3}|RM-\d{4}|TA-\d{4}|SIPP5 |/Motorola PA4N0106IT|MOBILE PHONE |Prefix : | Exclusive|(XIA|REA) DS | SLP|^Brodos |^Milwaukee |^KERDEJAR | was-LX1|3\. Generation|Flashlight|all carriers ,|^[-0] `).ReplaceAllString(name, "")
	name = strings.NewReplacer(" ", " ", "，", ",", "（", "(", "）", ")", "–", "|", "‎", "", "Kingkong", "King Kong", "KXD Handy,", "KXD", "Mobile Phone", "", "TELEFONO MOVIL", "", "Telefonas ", " ", "Mobility", "", "Galaxy-A", "Galaxy A", "Galaxy-Xcover", "Galaxy XCover", "Enterprise Edition", "EE", " A 90", " A90", " M5/", " M5|", "8GBRAM128GBROM", " ", "Black Smartphone", " ", "XIA DS", "Xiaomi").Replace(name)
	name = regexp.MustCompile(`^A34`).ReplaceAllString(name, "Samsung Galaxy A34")
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

	if s[0] == "Honor" || s[0] == "HONOR" {
		name = strings.ReplaceAll(name, "HonorMagic", "Magic")
	}

	if s[0] == "Huawei" || s[0] == "HUAWEI" {
		name = regexp.MustCompile(`(?i)[A-Z]{3}-L\d{2}`).ReplaceAllString(name, "")
	}

	if s[0] == "Moto" || s[0] == "moto" {
		name = "Motorola " + name
		s[0] = "Motorola"
	}
	if s[0] == "Motorola" || len(s) > 1 && s[1] == "Motorola" {
		// XT2159-3
		name = regexp.MustCompile(`(?i)[X][T]\d{4}-\d`).ReplaceAllString(name, "")
		name = strings.NewReplacer("Moto Edge", "Edge", "Light", "Lite", " E ", " E", "Moto One", "one", "moto one", "one").Replace(name)
	}

	if s[0] == "Nothing" {
		if s[1] != "Phone" {
			name = strings.ReplaceAll(name, "Nothing", "Nothing Phone")
		}
	}

	if s[0] == "Pixel" {
		name = "Google " + name
	}

	if s[0] == "Samsung" {
		// SM-A057GLGU
		name = regexp.MustCompile(`SM-A057G([LZ][GKS][UV])?(/DSN)?`).ReplaceAllString(name, "A05s")

		name = regexp.MustCompile(`\s+(SM-)?[AFMS]\d{3}[BFPR]?(/DSN?)?`).ReplaceAllString(name, "")
		name = strings.ReplaceAll(name, "Samsung A", "Samsung Galaxy A")
		name = strings.ReplaceAll(name, "Samsung M", "Samsung Galaxy M")
		name = strings.ReplaceAll(name, "Samsung G ", "Samsung Galaxy ")
	}

	if loc := AmazonRegex4.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%s\t%s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	if s[0] == "Apple" {
		var year string
		name = regexp.MustCompile(`\s+\(?(2016|2020|2022)\)?`).ReplaceAllStringFunc(name, func(s string) string {
			year = s
			return ""
		})
		name += year
	}

	if s[0] == "Blackview" || s[0] == "DOOGEE" || s[0] == "OSCAL" {
		name = regexp.MustCompile(`\s*\(?20[12]\d\)?`).ReplaceAllString(name, "")
	}

	if s[0] == "Xiaomi" {
		name = strings.ReplaceAll(name, "Xiaomi Smart ", "")
	}

	if s[0] == "ZTE" {
		name = strings.ReplaceAll(name, "Plues", "Plus")
	}

	name = regexp.MustCompile(`\s*\($`).ReplaceAllString(name, "")

	return helpers.Lint(name)
}

type loggingTransport struct{}

func (s *loggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	bytes, _ := httputil.DumpRequestOut(r, true)

	resp, err := http.DefaultTransport.RoundTrip(r)
	// err is returned after dumping the response

	respBytes, _ := httputil.DumpResponse(resp, true)
	bytes = append(bytes, respBytes...)

	fmt.Printf("%s\n", bytes)

	return resp, err
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

		eurPrice float32
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
			url := fmt.Sprintf(_url, p)

			// req, err := http.NewRequest(http.MethodGet, url, nil)
			// if err != nil {
			// 	// panic(err)
			// 	fmt.Printf("[%s] %s (%s)\n", _name, err, url)
			// 	return NewShop(
			// 		_name,
			// 		_url,

			// 		nil,
			// 	)
			// }

			// // req.Header.Set("Host", "www.amazon.de")
			// // req.Header.Set("User-Agent", "XX")
			// req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
			// req.Header.Set("Accept", "*/*")
			// // req.Header.Set("Accept-Encoding", "identity")
			// // req.Header.Set("Connection", "Keep-Alive")

			// // f, err := os.OpenFile(".sslkeylog", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
			// // if err != nil {
			// // 	panic(err)
			// // }
			// // defer f.Close()

			// client := &http.Client{
			// 	Transport: &http2.Transport{
			// 		DisableCompression:         true,
			// 		AllowHTTP:                  true,
			// 		StrictMaxConcurrentStreams: true,

			// 		TLSClientConfig: &tls.Config{
			// 			// KeyLogWriter:  f,
			// 			Renegotiation: tls.RenegotiateFreelyAsClient,
			// 		},
			// 	},
			// 	// Transport: &loggingTransport{},

			// 	// Transport: &http2.Transport{},
			// 	// Timeout:   10 * time.Second,
			// }
			// resp, err := client.Do(req)
			// if err != nil {
			// 	// panic(err)
			// 	fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
			// 	return NewShop(
			// 		_name,
			// 		_url,

			// 		nil,
			// 	)
			// }
			// defer resp.Body.Close()

			// if resp.StatusCode != http.StatusOK {
			// 	// panic(resp.StatusCode)
			// 	fmt.Printf("[%s] %d: %s (%s)\n", _name, resp.StatusCode, resp.Status, resp.Request.URL)
			// 	return NewShop(
			// 		_name,
			// 		_url,

			// 		nil,
			// 	)
			// }

			// if body, err := io.ReadAll(resp.Body); err != nil {
			// 	// panic(err)
			// 	fmt.Printf("[%s] %s (%s)\n", _name, err, resp.Request.URL)
			// 	return NewShop(
			// 		_name,
			// 		_url,

			// 		nil,
			// 	)
			// } else {
			// 	_body = body
			// }

			cmd := exec.Command("curl", "--compressed", url)

			cmdOutput := &bytes.Buffer{}
			cmd.Stdout = cmdOutput

			// if err := cmd.Start(); err != nil {
			// 	panic(err)
			// }
			// cmd.Wait()

			if err := cmd.Run(); err != nil {
				panic(err)
			}

			_body = cmdOutput.Bytes()

			os.WriteFile(path+fn, _body, 0664)
		}
		// fmt.Println(BytesToString(_body))

		doc := parse(BytesToString(_body))

		if productList := traverse(doc, "div", "class", "s-search-results"); productList != nil {
			// fmt.Println(productList)

			for item := productList.FirstChild.NextSibling; item != nil; item = item.NextSibling {
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

				if AmazonExclusionRegex.MatchString(title) {
					continue
				}

				imageTitleLink := traverse(item, "a", "class", "a-text-normal")
				// fmt.Println(imageTitleLink)

				link, _ := attr(imageTitleLink.Attr, "href")
				if _debug {
					fmt.Println(link)
				}
				_product.link = link

				// puis-price-instructions-style
				// s-price-instructions-style
				if itemPrice := traverse(item, "a", "class", "s-no-hover"); itemPrice != nil {
					// fmt.Println(itemPrice)

					if itemOldPrice := traverse(itemPrice, "span", "class", "a-offscreen"); itemOldPrice != nil {
						// fmt.Println(itemOldPrice)

						oldPrice, _ := text(itemOldPrice)
						oldPrice = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(oldPrice, ".", ""), ",", "."), "\u00a0€", "")
						if _debug {
							fmt.Println(oldPrice)
						}

						if _price, err := strconv.ParseFloat(oldPrice, 32); err != nil {
							panic(err)
						} else {
							_product.price = float32(_price * EUR_CHF)
							_product.eurPrice = float32(_price)
							if _debug {
								fmt.Println(_product.price)
							}
						}
					}

					if itemPrice := traverse(itemPrice, "span", "class", "a-text-price"); itemPrice != nil {
						if itemPrice := traverse(itemPrice, "span", "class", "a-offscreen"); itemPrice != nil {
							// fmt.Println(itemPrice)

							price, _ := text(itemPrice)
							price = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(price, ".", ""), ",", "."), "\u00a0€", "")
							if _debug {
								fmt.Println(price)
							}

							if _price, err := strconv.ParseFloat(price, 32); err != nil {
								panic(err)
							} else {
								_product.oldPrice = float32(_price * EUR_CHF)
								if _debug {
									fmt.Println(_product.oldPrice)
								}
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
			_eurPrice := product.eurPrice

			_savings := _price - _retailPrice
			_discount := 100 - ((_price * 100) / _retailPrice)

			_link := s.ResolveURL(product.link).String()
			_link = strings.Split(_link, "/ref=")[0]

			{
				product := &Product{
					Code:  _name + "//" + product.code,
					Name:  _title,
					Model: _model,

					RetailPrice: _retailPrice,
					Price:       _price,
					EURPrice:    _eurPrice,
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
