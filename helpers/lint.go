package lint

import (
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var nameMapping = map[string]string{
	"2Nd":     "2nd",
	"2ND":     "2nd",
	"3A":      "3a",
	"4A":      "4a",
	"5A":      "5a",
	"6A":      "6a",
	"7A":      "7a",
	"8A":      "8a",
	"9At":     "9AT",
	"9I":      "9i",
	"Ce":      "CE",
	"Ds":      "DS",
	"El":      "EL",
	"Fe":      "FE",
	"FLIP":    "Flip",
	"Gt":      "GT",
	"Hd":      "HD",
	"Htc":     "HTC",
	"Iii":     "III",
	"IIi":     "III",
	"Ii":      "II",
	"IIIF150": "iiiF150",
	"Iphone":  "iPhone",
	"Iv":      "IV",
	"Jvc":     "JVC",
	"Kxd":     "KXD",
	"Lg":      "LG",
	"Mini":    "mini",
	"NEO2":    "Neo2",
	"Oneplus": "OnePlus",
	"Realme":  "realme",
	"Se":      "SE",
	"Tcl":     "TCL",
	"Tecno":   "TECNO",
	"Umidigi": "UMIDIGI",
	"Vivo":    "vivo",
	"Xa":      "XA",
	"Xcover":  "XCover",
	"Xgody":   "XGODY",
	"Xl":      "XL",
	"Xr":      "XR",
	"Xs":      "XS",
	"Xz":      "XZ",
	"Xzs":     "XZs",
	"Zte":     "ZTE",
}

var nameRegex = maps.Keys(nameMapping)

var r = regexp.MustCompile(`\W*(` + strings.Join(nameRegex, "|") + `)\W*`)

func Title(name string) string {
	return cases.Title(language.Und, cases.NoLower).String(name)
}

func Lint(name string) string {
	name = Title(strings.ToLower(strings.TrimSpace(name)))

	if matches := r.FindAllStringSubmatch(name, -1); len(matches) > 0 {
		for _, match := range matches {
			name = strings.ReplaceAll(name, match[0], strings.ReplaceAll(match[0], match[1], nameMapping[match[1]]))
		}

		// fmt.Printf("%-30s %v\n", name, matches)
	}

	name = strings.ReplaceAll(name, "FEver", "Fever")

	return Model(name)
}

func Model(name string) string {
	name = strings.TrimSpace(name)

	s := strings.Split(name, " ")
	// fmt.Println(s)

	if s[0] == "iPhone" {
		name = "Apple" + " " + name

		s[0] = "Apple"
	}

	if s[0] == "Apple" {
		name = strings.Split(name, "/")[0]

		name = regexp.MustCompile(`\s+(2016|2020|2022)`).ReplaceAllString(name, " ($1)")
		name = regexp.MustCompile(`(?i)1(\.|st)\s*Gen(eration)?\.?`).ReplaceAllString(name, "(2016)")
		name = regexp.MustCompile(`(?i)2(\.|nd)\s*Gen(eration)?\.?`).ReplaceAllString(name, "(2020)")
		name = regexp.MustCompile(`(?i)3(\.|rd)\s*Gen(eration)?\.?`).ReplaceAllString(name, "(2022)")
		name = regexp.MustCompile(`(?i)\s*(\d+)\s*(S)`).ReplaceAllString(name, " $1$2")
	}

	if s[0] == "Blackview" {
		name = regexp.MustCompile(`(?i)BV\s*(\d+)`).ReplaceAllString(name, "BV$1")
	}

	if s[0] == "Galaxy" {
		name = "Samsung" + " " + name
	}

	if s[0] == "Honor" {
		name = regexp.MustCompile(`(?i)Magic\s*(\d)\s*(\w)?`).ReplaceAllString(name, "Magic$1 $2")
		name = regexp.MustCompile(`(?i)Honor\s*(\d+)\s*(\w)?`).ReplaceAllString(name, "Honor $1 $2")
	}

	if s[0] == "Huawei" {
		name = regexp.MustCompile(`(?i)Ascend\s*([Y])\s*(\d+)`).ReplaceAllString(name, "Ascend $1$2")
		name = regexp.MustCompile(`(?i)Magic\s*(\d+)`).ReplaceAllString(name, "Magic$1")
		name = regexp.MustCompile(`(?i)nova\s*(\d+)`).ReplaceAllString(name, "nova $1")
		name = regexp.MustCompile(`(?i)P\s*(Smart)\s*(\+?)`).ReplaceAllString(name, "P $1$2")
		name = regexp.MustCompile(`(?i)P-?\s*(\d+)`).ReplaceAllString(name, "P$1")
	}

	if s[0] == "Inoi" {
		name = regexp.MustCompile(`(?i) A\s*(\d+)`).ReplaceAllString(name, " A$1")
	}

	if s[0] == "Moto" {
		name = "Motorola" + " " + name
	}

	if s[0] == "Motorola" {
		// name = strings.NewReplacer("Moto E E", "Moto E", "Moto G g", "Moto g").Replace(name)

		name = regexp.MustCompile(`(?i)edge\s*(\d+)\s*(\w*)`).ReplaceAllString(name, "edge $1 $2")
		name = regexp.MustCompile(`(?i)^Motorola\s*(Moto\s*)?([EG]\s+)?([EG])\s*(\d+)\s*`).ReplaceAllString(name, "Motorola Moto $3$4")
	}

	if s[0] == "Nothing" {
		name = regexp.MustCompile(`-(\d+)`).ReplaceAllString(name, "($1)")
	}

	if s[0] == "OnePlus" {
		name = regexp.MustCompile(`(?i)\s*CPH\d{4}`).ReplaceAllString(name, "")
		name = regexp.MustCompile(`(?i)CE\s*(\d+)`).ReplaceAllString(name, "CE $1")
	}

	if s[0] == "OPPO" || s[0] == "Oppo" {
		name = regexp.MustCompile(`(?i)([A])\s*(\d+)\s*([k])?`).ReplaceAllString(name, "$1$2$3")
		name = regexp.MustCompile(`[KS]$`).ReplaceAllStringFunc(name, strings.ToLower)

		name = regexp.MustCompile(`(?i)Reno\s*(\d+)\s*(\w)?`).ReplaceAllString(name, "Reno$1 $2")
		name = regexp.MustCompile(`(?i)OPPO\s*(\d+)\s*(\w)?`).ReplaceAllString(name, "OPPO Reno$1 $2")
	}

	if s[0] == "Redmi" {
		name = "Xiaomi" + " " + name
	}

	if s[0] == "Samsung" {
		name = regexp.MustCompile(`[^F][E]$`).ReplaceAllStringFunc(name, strings.ToLower)

		name = regexp.MustCompile(`(?i)\s+(SM-)?[AFMS]\d{3}[BFR]?[N]?(\/DSN?)?`).ReplaceAllString(name, "")

		name = regexp.MustCompile(`Note\s*(\d+)`).ReplaceAllString(name, "Note $1")
		name = regexp.MustCompile(`(?i)( Galaxy)? (Tab )?(A|S)\s*(\d+| Duos)`).ReplaceAllString(name, " Galaxy $2$3$4")
	}

	if s[0] == "Sony" {
		name = regexp.MustCompile(`[DF]\d{4}`).ReplaceAllString(name, "")
	}

	if s[0] == "Wiko" {
		name = regexp.MustCompile(`(?i)View\s*(\d+)`).ReplaceAllString(name, "View $1")
	}

	if s[0] == "Xiaomi" {
		name = regexp.MustCompile(`^Xiaomi\s*(\d+)\s*(\d{4})`).ReplaceAllString(name, "Xiaomi Redmi $1 ($2)")
		name = regexp.MustCompile(`(Poco\s*)?([CFMX]\d+)`).ReplaceAllString(name, "Poco $2")
		name = regexp.MustCompile(`(Redmi\s*)?Note\s*(\d+)`).ReplaceAllString(name, "Redmi Note $2")
		// name = regexp.MustCompile(`Note\s*(\d+)`).ReplaceAllString(name, "Note $1")
		name = regexp.MustCompile(`Redmi\s*(\d+)`).ReplaceAllString(name, "Redmi $1")
	}

	if s[0] == "ZTE" {
		name = regexp.MustCompile(`(Blade\s*)?(A\d+)`).ReplaceAllString(name, "Blade $2")
	}

	return strings.TrimSpace(name)

	// var r2 = regexp.MustCompile(`\W*([A-Z][a-z]*)?[0-9]+[A-Za-z]*\W*`)

	// if matches := r2.FindAllStringSubmatch(name, -1); len(matches) > 0 {
	// 	for _, match := range matches {
	// 		if len(match) > 1 && (match[1] == "Reno" || match[1] == "Magic") {
	// 			continue
	// 		}

	// 		oldModel := strings.ToUpper(strings.TrimSpace(match[0]))
	// 		newModel := oldModel
	// 		if n := len(oldModel); n > 1 {
	// 			if s := oldModel[n-1 : n]; s == "I" || s == "E" {
	// 				newModel = oldModel[:n-1] + strings.ToLower(s)
	// 				// fmt.Println("*** MATCH:" + oldModel + "/" + oldModel[n-1:n] + "/" + newModel)
	// 			}
	// 		}

	// 		name = strings.ReplaceAll(name, match[0], strings.ReplaceAll(strings.ToUpper(match[0]), oldModel, newModel))
	// 	}

	// 	// fmt.Printf("%-30s %v\n", name, matches)
	// }

	// return name
}
