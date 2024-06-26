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

	"github.com/sugawarayuuta/sonnet"

	"golang.org/x/net/html"

	helpers "jsapi-scraper/helpers"
)

var TuttiRegex = regexp.MustCompile(`(?i)(RAM )?[,-]? ?[\[(]?\b((2|3|4|6|8|12|16)[/]?)?(2|3|4|6|8|12|16|32|54|64|65|125|128|250|256|265|512)(\.0)?[ ./]?([KMG][BG]|BG|G|B\b|\+)[)\]]?( RAM)?|(\. )?\b(6|8|12|128|256|512)(GB)?/(6|8|12|128|256)(GB)?\b|GB 64|/6\s+|\(\d\.\d{1,2} "| \d(\.\d{1,2})?"|\d+ mm|-?[ (/,]?[45] ?G[)]?| gb$| GSM| LTE| inkl\. | mit | und | and more| plié| [*|] | \([^\d]| /|, |/ | -|;`)
var TuttiExclusionRegex = regexp.MustCompile(`(?i)^(emporia|ericsson|htc)|galaxy (s8|s7|s5|s4|s3|s|j\d+|gt)|iph?one? ?(3gs|3g|3|s4|4s|4|5s|5c|5|6s|6|7|8)|motorola (v8|razr)|nokia|orange|samsung (galaxy (young|s|note ii|note 2|j3|ace)|mini|rex|s7|s8|s9)|sonn?y ?(err?ics?son)|swisscom|adapter|akku|alt|astuccio|atrappe|audio|bastler|beschädigt|bilschirmschutz|bootloop|case|cloudlocked|cover|charger|charging stand|custodia|defec?kt|display|folie|gehäuse|gesperrt|gigaset|hülle|kabel|kameraschutz|kinder|klapp|mainboard|nostalgie|(nur )?verpackung|panzerglas|poco mode 1|riss|rott[io]|sammlung|scambio|scatola|screen protector|senior| set|siemens|silikon|skin|sperre|teile|vecchio|vintage|voip|zersplittert`)
var TuttiInclusionRegex = regexp.MustCompile(`(?i)^(apple (iphone (x|se|\d{2}))|asus (zenfone|rog)|blackview (bv\d+|bl\d+|a\d+)|fairphone|google (pixel)|honor (x\d+|magic|\d+)|huawei (y\d+|p[ -]?\d+|p smart|nova|mate)|infinix|inoi (note|a\d+)|motorola (moto|edge|defy)|nothing|oneplus (nord|\d+)|oppo (reno|find|a\d+)|realme (narzo|c\d+|\d+)|samsung (galaxy [amnsxz])|sony (xperia)|vivo (y\d+|v\d+)|wiko (y\d+|view|sunny|power|fever)|xiaomi (redmi |poco|mi|\d+)|zte (blade|axon))`)

var TuttiCleanFn = func(name string) string {
	name = regexp.MustCompile(`(?i)^Original | - neuwertiger Zustand| excellent etat|^\*\*\*\*|^\*\*|^\+\+ | \+\+$| \+ Airpods| Entsperrt|^Für Einsteiger: |^Neu: |^Zu verkaufen: |^NEW |Supper | Occass?ion|Gebrauchtes | Schnäppchen| GÜNSTIG( zu vergeben)?| MAKELLOS|>BILLIG: |AM ?OLED Bildschirm|\. FestPreis|Hochwertiges |Cellulate |funktioniert (100%)?|bisschen kaputt aber geht noch gut| \d+%( Battery| di batteria)?|Batteriezustand|Batteriekapazität( \d+%)?|COMPLETO FUNZIONA BENISSIMO|(Perfetto |Ricondizionato |Rigenerato )?Garanzia( 12Mesi)?( n\.\d{3})?|RESERVIERT|(guter?|in (einem )?sehr gute[nm]|in exzellente[nm]|neuwertiger|Perfekter|(in )?(einem )?Top|Super)[ -]Zustand( \+ Garantie)?|ohne Kratzer|(voll )?funktionstüchtig|im sehr guten Zustand|mit Box|( - )?sehr guter Zustand|Hammer Geil| in ottimo stato|semplificato |Mobile Phone( - )?|Mobile?telefon (von )?| (Smart(fon|phone)|Handy)( - |:)?|(Satelliten)?Telefon(ino|o)?( per anziani)?|mobile |(leicht)? gebraucht( & gelöscht)?|(Micro-|Neuwertiges |Komfort-|Natel/)?(Handy|Natel) (von |/ |- )?|Neuwertiges |Handy\s*/\s*(Natel|Smart(fon|phone))( - |:)?|(zu )?verkaufen? ?(ein )?|Verkauf von |vendo | vendesi|preis verhandelbar|\(Tausche/Verhandelbar\) |originalverpackt|neu & ungeöffnet|(in|inklusive|mit|NEU und| und)? OVP| ungeöffnet|( - )?(fast |(fast )?wie |Fabrik)?neu(es?|wertig)?|24 Monate Garantie| New Version| und noch verschweisst| nie benutzt| in gute[mn] Zustand| mit Gebrauchsspuren|einwandfrei|renoviert|4 Farben|Gratis ?versand|in Lederetui|mit Eingabestift|läuft einwandfrei| MIT GOOGLE SERVICES|Original OS|CalyxOS|GrapheneOS|(günstiges |neues |Android )?Smart(fon|phon)e?( -|:)? |Burnerphone |Neuwertiges |in usato|(Usato ma )?(Come |Praticamente )?Nuovo( |!)?(batteria|Sigillato)?| garandieschein| anno| année| con vetro da sostituire| \+ accessori| Glasschaden| HD\+|( - )?(dual|Dueo)[ -]?sim( free)?|(\d\.)?\d Zoll|miui|Firmengerät| Apple| ESR| MagSafe| Original Taptic Engine| Original Front Kamera Module| Original Kamera Module| 4K Kamera| inkl\. Ladegerät| inkl Zubehoer| Gehäuse Original|Google Sperre|SIMLOCKED|kaputter SIM Slot| Speicher| di Xiaomi|ottimo condizione( senza graffi)?|condizioni ottime| modell|cellulare |\[DANNEGGIATO\]|4/128(GB)?|(12|16|32|64|128|256)\+(2|3|4|6|8|12|16)GB|(2|3|4|6|8|12|16)\s*ram|(16|32|64|128|256)\s*rom|512,$|- Graphit|Breathing Crystal|Red Product| RED$|Smaragdgrün| 1TB| Etui| Cleav| goldfarben| \+ viel Zubehör|Bitte lesen!|\s*!*$| ((mit )?Android|EU)$`).ReplaceAllString(name, "")

	if loc := TuttiRegex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	name = regexp.MustCompile(`(?i)i[ -]?P(hi|ho|hp|oh|o)ne?\s*`).ReplaceAllString(name, "iPhone ")
	name = regexp.MustCompile(`(?i)One ?Plus`).ReplaceAllString(name, "OnePlus")
	name = regexp.MustCompile(`(?i)Xioa?mi`).ReplaceAllString(name, "Xiaomi")
	name = regexp.MustCompile(`(?i)Mi Xiaomi`).ReplaceAllString(name, "Xiaomi Mi")
	name = regexp.MustCompile(`(?i)Huawaii`).ReplaceAllString(name, "Huawai")
	name = regexp.MustCompile(`(?i)Blackwew`).ReplaceAllString(name, "Blackview")
	name = strings.NewReplacer("Appel", "Apple", "max pro", "Pro Max", "Xsmax", "XS Max", "prima generazione", "1. Gen.", "1Gen  Rigenerato", "1. Gen.", "1 Generation", "1. Gen.", "1st. Generation", "1. Gen.", "1Gen", "1. Gen.", "1gen", "1. Gen.", "1 Gen", "1. Gen.", "1 gen", "1. Gen.", "Gen 1", "1. Gen.", "primoo", "1. Gen.", "SE1", "SE 1. Gen.", "SE 3", "SE 3. Gen.", "3°Generazione", "3. Gen.", "Redmy", "Redmi", " G5G", " G", " 2GB", "", "10 e", "10e", "20 e", "20e", "FE20", "S20 FE", "S21FE", "S21 FE", "A5-6", "A5", "Galxy", "Galaxy", "XSMax", "XS Max", "Mate-20", "Mate 20", "Motorolla", "Motorola", "Sansung", "Samsung", "SAMSUG", "SAMSUNG", "Galaxie", "Galaxy", "S10 DUOS S10 NORMAL", "S10 DUOS", " Tablet", " Tab", "Xcovet", "XCover", "Enterprise Edition", "EE", "2 Stück", "", "2 Stuck", "", "Android", "", "n.201", "", "Litle", "Lite", "ß", "ss", "|", " ", "  ", " ").Replace(name)

	s := strings.Split(name, " ")

	if s[len(s)-1] == "Samsung" {
		name = strings.ReplaceAll(name, "Samsung", "")
	}

	if s[0] == "Google" {
		name = regexp.MustCompile(`(?i)(\d)\s+(a)`).ReplaceAllString(name, "$1$2")
	}

	if s[0] == "Honor" {
		name = regexp.MustCompile(`(?i)PLK-[ATU]?L\d{2}[H]?`).ReplaceAllString(name, "")
	}

	if s[0] == "Huawei" || s[0] == "HUAWEI" {
		name = strings.ReplaceAll(name, " leit", " lite")

		name = regexp.MustCompile(`(?i)[A-Z]{3}-L\d{2}`).ReplaceAllString(name, "")

		name = regexp.MustCompile(`(\s*)[Pp]ro`).ReplaceAllString(name, " Pro")

		name = regexp.MustCompile(`(?i)New Edition`).ReplaceAllString(name, "NE")
	}

	if s[0] == "iPhone" || s[0] == "Apple" {
		name = strings.ReplaceAll(name, "**", " ")
		name = strings.ReplaceAll(name, " con", " ")
		name = strings.ReplaceAll(name, " pto", " Pro")
		name = strings.ReplaceAll(name, " 1387", " 4S")
		name = strings.ReplaceAll(name, " (Corail)", "")
		name = strings.ReplaceAll(name, " 4.7?", "")
		name = strings.ReplaceAll(name, " 10 x", " X")
		name = strings.ReplaceAll(name, " 10", " X")
		name = strings.ReplaceAll(name, "10", "")
		name = strings.ReplaceAll(name, " (2017", " (2016)")
		name = strings.ReplaceAll(name, " 2017", " 2016")
		name = strings.ReplaceAll(name, ". 2021", " 2016")
		name = strings.ReplaceAll(name, " 256", "")
		name = strings.ReplaceAll(name, " 128", "")
		name = strings.ReplaceAll(name, " 64", "")

		name = regexp.MustCompile(`(XS|XR|X|11|12|13|14|15)\s*([Pp]ro|PRO|[Mm]ini)`).ReplaceAllString(name, "$1 $2")

		name = regexp.MustCompile(` 1$`).ReplaceAllString(name, " (2016)")

		name = regexp.MustCompile(`[.!]+\s*$`).ReplaceAllString(name, "")
	}

	if s[0] == "Oppo" || s[0] == "OPPO" {
		name = regexp.MustCompile(`(?i)([X])\s+(\d+)`).ReplaceAllString(name, "$1$2")

		name = strings.ReplaceAll(name, " Champions league edition", "")
	}

	if s[0] == "Pocophone" || s[0] == "POCOPHONE" {
		name = strings.ReplaceAll(name, "Pocophone", "Xiaomi POCO")
		name = strings.ReplaceAll(name, "POCOPHONE", "Xiaomi POCO")
	}

	if s[0] == "Samsung" || s[0] == "samsung" || s[0] == "Galaxy" || s[0] == "galaxy" {
		name = strings.NewReplacer("duas", "duos", "GT 19070", "", " 64", "", "s21Ultra", "S21 Ultra", "S21Ultra", "S21 Ultra", "Note S20", "Note20", "Z 3 flip", "Z Flip3").Replace(name)

		name = regexp.MustCompile(`\($`).ReplaceAllString(name, "")
		name = regexp.MustCompile(`\(?20\d{2}\)?`).ReplaceAllString(name, "")
		name = regexp.MustCompile(`[J]\d{3}[H]`).ReplaceAllString(name, "")
		name = regexp.MustCompile(`(A\d{2})\s+[Ss]\b`).ReplaceAllString(name, "${1}s")

		name = strings.Split(name, ",")[0]
		name = strings.Split(name, " - ")[0]

		name = regexp.MustCompile(`(?i)\b(A)\s+(\d+)\s+([S])\b`).ReplaceAllString(name, "$1$2$3")
	}

	if s[0] == "Sony" {
		name = regexp.MustCompile(`[E]\d{4}`).ReplaceAllString(name, "")

		name = regexp.MustCompile(`\blll\b`).ReplaceAllString(name, "III")
		name = regexp.MustCompile(`\bll\b`).ReplaceAllString(name, "II")
	}

	if s[0] == "Redmi" {
		// name = strings.ReplaceAll(name, "(5G)", "")
		name = strings.ReplaceAll(name, " C9", " 9C")
		name = strings.ReplaceAll(name, " Not ", " Note ")

		name = regexp.MustCompile(`(?i)\b(\d{1,2})\s+([C])\b`).ReplaceAllString(name, "$1$2")
	}
	if s[0] == "Xiaomi" {
		if len(s) > 1 && s[1] == "9C" {
			name = strings.ReplaceAll(name, "9C", "Redmi 9C")
		}

		name = strings.ReplaceAll(name, "Xiaomi Redmi 9 Note", "Xiaomi Redmi Note 9")

		name = regexp.MustCompile(`(?i)\b(\d{1,2})\s+([CT])\b`).ReplaceAllString(name, "$1$2")
	}

	name = regexp.MustCompile(`A\d{4}|\.$`).ReplaceAllString(name, "")

	name = strings.ReplaceAll(name, "Mi 0", "Mi 10")
	name = strings.ReplaceAll(name, "Uktra", "Ultra")

	name = strings.ReplaceAll(name, "Xiaomi 11 Ultra", "Xiaomi Mi 11 Ultra")

	return helpers.Lint(name)
}

func XXX_tutti(isDryRun bool) IShop {
	const _name = "Tutti"
	const _url = "https://www.tutti.ch/api/v10/graphql"

	const _tests = false

	testCases := map[string]string{}

	type _Result struct {
		Code string `json:"listingID"`
		Name string `json:"title"`

		Description string `json:"body"`

		FormattedPrice string `json:"formattedPrice"`

		SEO struct {
			Slug string `json:"deSlug"`
		} `json:"seoInformation"`

		Title string
		Model string
		Price float32
	}

	type _Node struct {
		Node _Result `json:"node"`
	}

	type _Response struct {
		Data struct {
			SearchListingsByQuery struct {
				Listings struct {
					Edges      []_Node `json:"edges"`
					TotalCount int     `json:"totalCount"`
				} `json:"listings"`
			} `json:"searchListingsByQuery"`
		} `json:"data"`
	}

	var _result _Response
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	var _results []_Result
	_count := 0

	first := true
	for p := 1; p <= 10; p++ {
		fn := fmt.Sprintf("shop/tutti.%d.json", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			// "query": "Apple iPhone | Google Pixel | Huawei | Motorola | Nothing | OnePlus | OPPO | realme | Samsung Galaxy | Vivo | Xiaomi | ZTE",
			var jsonData = StringToBytes(
				fmt.Sprintf(
					`{
						"query": "query SearchListingsByConstraints($query: String, $constraints: ListingSearchConstraints, $category: ID, $first: Int!, $offset: Int!, $sort: ListingSortMode!, $direction: SortDirection!) {\n  searchListingsByQuery(\n    query: $query\n    constraints: $constraints\n    category: $category\n  ) {\n    ...searchResultFields\n  }\n}\n\nfragment searchResultFields on ListingSearchResult {\n  listings(first: $first, offset: $offset, sort: $sort, direction: $direction) {\n    ...listingsConnectionFields\n  }\n  galleryListings(first: 3) {\n    ...listingFields\n  }\n  filters {\n    ...filterFields\n  }\n  suggestedCategories {\n    ...suggestedCategoryFields\n  }\n  selectedCategory {\n    ...selectedCategoryFields\n  }\n  seoInformation {\n    seoIndexable\n    deQuerySlug: querySlug(language: DE)\n    frQuerySlug: querySlug(language: FR)\n    itQuerySlug: querySlug(language: IT)\n    bottomSEOLinks {\n      label\n      slug\n      searchToken\n    }\n  }\n  searchToken\n  query\n}\n\nfragment selectedCategoryFields on Category {\n  categoryID\n  label\n  ...categoryParentFields\n}\n\nfragment categoryParentFields on Category {\n  parent {\n    categoryID\n    label\n    parent {\n      categoryID\n      label\n      parent {\n        categoryID\n        label\n      }\n    }\n  }\n}\n\nfragment suggestedCategoryFields on Category {\n  categoryID\n  label\n  searchToken\n  mainImage {\n    rendition(width: 300) {\n      src\n    }\n  }\n}\n\nfragment filterFields on ListingFilter {\n  __typename\n  ...filterDescriptionFields\n  ... on ListingIntervalFilter {\n    ...intervalFilterFields\n  }\n  ... on ListingSingleSelectFilter {\n    ...singleSelectFilterFields\n  }\n  ... on ListingMultiSelectFilter {\n    ...multiSelectFilterFields\n  }\n  ... on ListingPricingFilter {\n    ...pricingFilterFields\n  }\n  ... on ListingLocationFilter {\n    ...locationFilterFields\n  }\n}\n\nfragment filterDescriptionFields on ListingsFilterDescription {\n  name\n  label\n  disabled\n}\n\nfragment intervalFilterFields on ListingIntervalFilter {\n  ...filterDescriptionFields\n  intervalType {\n    __typename\n    ... on ListingIntervalTypeText {\n      ...intervalTypeTextFields\n    }\n    ... on ListingIntervalTypeSlider {\n      ...intervalTypeSliderFields\n    }\n  }\n  intervalValue: value {\n    min\n    max\n  }\n  step\n  unit\n  minField {\n    placeholder\n  }\n  maxField {\n    placeholder\n  }\n}\n\nfragment intervalTypeTextFields on ListingIntervalTypeText {\n  minLimit\n  maxLimit\n}\n\nfragment intervalTypeSliderFields on ListingIntervalTypeSlider {\n  sliderStart: minLimit\n  sliderEnd: maxLimit\n}\n\nfragment singleSelectFilterFields on ListingSingleSelectFilter {\n  ...filterDescriptionFields\n  ...selectFilterFields\n  selectedOption: value\n}\n\nfragment selectFilterFields on ListingSelectFilter {\n  options {\n    ...selectOptionFields\n  }\n  placeholder\n  inline\n}\n\nfragment selectOptionFields on ListingSelectOption {\n  value\n  label\n}\n\nfragment multiSelectFilterFields on ListingMultiSelectFilter {\n  ...filterDescriptionFields\n  ...selectFilterFields\n  selectedOptions: values\n}\n\nfragment pricingFilterFields on ListingPricingFilter {\n  ...filterDescriptionFields\n  pricingValue: value {\n    min\n    max\n    freeOnly\n  }\n  minField {\n    placeholder\n  }\n  maxField {\n    placeholder\n  }\n}\n\nfragment locationFilterFields on ListingLocationFilter {\n  ...filterDescriptionFields\n  value {\n    radius\n    selectedLocalities {\n      ...localityFields\n    }\n  }\n}\n\nfragment localityFields on Locality {\n  localityID\n  name\n  localityType\n}\n\nfragment listingFields on Listing {\n  listingID\n  title\n  body\n  postcodeInformation {\n    postcode\n    locationName\n    canton {\n      shortName\n      name\n    }\n  }\n  timestamp\n  formattedPrice\n  formattedSource\n  highlighted\n  primaryCategory {\n    categoryID\n  }\n  sellerInfo {\n    alias\n    logo {\n      rendition {\n        src\n      }\n    }\n    subscriptionInfo {\n      subscriptionBadge {\n        src(format: SVG)\n      }\n    }\n  }\n  images(first: 15) {\n    __typename\n  }\n  thumbnail {\n    normalRendition: rendition(width: 235, height: 167) {\n      src\n    }\n    retinaRendition: rendition(width: 470, height: 334) {\n      src\n    }\n  }\n  seoInformation {\n    deSlug: slug(language: DE)\n    frSlug: slug(language: FR)\n    itSlug: slug(language: IT)\n  }\n}\n\nfragment listingsConnectionFields on ListingsConnection {\n  totalCount\n  edges {\n    node {\n      ...listingFields\n    }\n  }\n  placements {\n    keyValues {\n      key\n      value\n    }\n    pageName\n    pagePath\n    positions {\n      adUnitID\n      mobile\n      position\n      positionType\n    }\n    afs {\n      customChannelID\n      styleID\n      adUnits {\n        adUnitID\n        mobile\n      }\n    }\n  }\n}",
						"variables": {
							"constraints": {
								"strings": [{
									"key": "organic",
									"value": ["tutti"]
								}],
								"prices": [{
									"key": "price",
									"min": %.f,
									"max": %.f,
									"freeOnly": false
								}],
								"intervals": null,
								"locations": null
							},
							"category": "cellPhones",
							"status": "pendingUpdateWithoutToken",
							"first": 100,
							"offset": %d,
							"direction": "DESCENDING",
							"sort": "TIMESTAMP"
						}
					}`,
					ValueMinimum,
					ValueMaximum,
					_count,
				),
			)

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
			req.Header.Set("Origin", "https://www.tutti.ch")
			// req.Header.Set("X-Csrf-Token", "f77e43ecb46452eadd7a9c33af762565086815d653388d7d6208502169648ba0")
			req.Header.Set("X-Tutti-Client-Identifier", "web/1.0.0+env-live.git-4936d92")
			req.Header.Set("X-Tutti-Hash", "312015e0-4495-437f-9ac5-65144c7e8bb2")
			req.Header.Set("X-Tutti-Source", "web r1.0-2023-07-04-14-33")

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
		// fmt.Println(BytesToString(_body))

		if err := sonnet.Unmarshal(_body, &_result); err != nil {
			panic(err)
		}
		// fmt.Println(_result.Data.SearchListingsByQuery.Listings.Edges)

		for _, edge := range _result.Data.SearchListingsByQuery.Listings.Edges {
			_title := html.UnescapeString(edge.Node.Name)

			if TuttiExclusionRegex.MatchString(_title) {
				continue
			}

			_model := TuttiCleanFn(_title)

			if Skip(_model) {
				continue
			}

			if !TuttiInclusionRegex.MatchString(_model) {
				if first {
					fmt.Println()
					first = false
				}

				fmt.Printf("*** %-50s %s\n", _model, _title)
				continue
			}

			// _link := fmt.Sprintf("https://www.tutti.ch/de/vi/%s/%s", edge.Node.SEO.Slug, edge.Node.Code)
			// fmt.Printf("%s\n%s\n\t\t%s\n%s\n--\n", _model, _title, strings.TrimSpace(strings.ReplaceAll(html.UnescapeString(edge.Node.Description), "\n", "\n\t\t")), _link)

			_price := edge.Node.FormattedPrice
			_price = strings.NewReplacer(".-", ".00", "'", "", "Gratis", "0.00").Replace(_price)

			edge.Node.Model = _model

			if _price, err := strconv.ParseFloat(_price, 32); err != nil {
				panic(err)
			} else {
				edge.Node.Price = float32(_price)
			}

			_results = append(_results, edge.Node)
		}

		_count += len(_result.Data.SearchListingsByQuery.Listings.Edges)

		if _count >= _result.Data.SearchListingsByQuery.Listings.TotalCount {
			break
		}

		fmt.Println()
	}

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_results))
		for _, result := range _results {
			_title := html.UnescapeString(result.Name)
			_model := result.Model

			if _tests {
				testCases[_title] = _model
			}

			var _retailPrice float32
			var _price float32
			var _savings float32
			var _discount float32

			_price = result.Price
			_retailPrice = _price

			if _savings == 0 {
				_savings = _price - _retailPrice
			}
			_discount = 100 - ((100 / _retailPrice) * _price)

			_productUrl := fmt.Sprintf("https://www.tutti.ch/de/vi/%s/%s", result.SEO.Slug, result.Code)

			{
				code := result.Code
				product := &Product{
					Code:  _name + "//" + code,
					Name:  _title,
					Model: _model,

					RetailPrice: _retailPrice,
					Price:       _price,
					Savings:     _savings,
					Discount:    _discount,

					URL: _productUrl,
				}
				// fmt.Println(product)

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
