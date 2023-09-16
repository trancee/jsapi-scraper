package lint

import (
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var nameMapping = map[string]string{
	"2Nd": "2nd",
	"2ND": "2nd",
	// "3A":   "3a",
	// "4A":   "4a",
	// "5A":   "5a",
	// "6A":   "6a",
	// "7A":   "7a",
	// "8A":   "8a",
	"9At":  "9AT",
	"9I":   "9i",
	"Ce":   "CE",
	"Ds":   "DS",
	"Ee":   "EE",
	"El":   "EL",
	"Fe":   "FE",
	"FLIP": "Flip",
	"Gt":   "GT",
	"Hd":   "HD",
	"Htc":  "HTC",
	// "Huawei":  "HUAWEI",
	"IIif150": "IIIF150",
	"IIi":     "III",
	"Iii":     "III",
	"Ii":      "II",
	"Iphone":  "iPhone",
	"Iv":      "IV",
	"Jvc":     "JVC",
	"Kxd":     "KXD",
	"Lg":      "LG",
	"Mini":    "mini",
	"Mm":      "MM",
	"NEO2":    "Neo2",
	"Oneplus": "OnePlus",
	// "Poco":    "POCO",
	// "Realme":  "realme",
	"Se":      "SE",
	"Tcl":     "TCL",
	"Tecno":   "TECNO",
	"Umidigi": "UMIDIGI",
	// "Vivo":    "vivo",
	"Xa":     "XA",
	"Xcover": "XCover",
	"Xgody":  "XGODY",
	"Xl":     "XL",
	"Xr":     "XR",
	"Xs":     "XS",
	"Xz":     "XZ",
	"Xzs":    "XZs",
	"Zte":    "ZTE",
}

var nameRegex = maps.Keys(nameMapping)

var r = regexp.MustCompile(`\W*(` + strings.Join(nameRegex, "|") + `)\W*`)

func Title(name string) string {
	return cases.Title(language.Und, cases.NoLower).String(name)
}

func Lint(name string) string {
	name = Title(strings.ToLower(regexp.MustCompile(`\s+`).ReplaceAllString(strings.TrimSpace(name), " ")))

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

	// fmt.Printf("%v\n[%s]\n", s, name)
	// defer func() { fmt.Printf("[%s]\n\n", name) }()

	if s[0] == "iPhone" {
		name = "Apple" + " " + name

		s = strings.Split(name, " ")
	}

	if s[0] == "Apple" {
		name = strings.Split(name, "/")[0]

		name = regexp.MustCompile(`\s+\(?(2016|2020|2022)\)?`).ReplaceAllString(name, " ($1)")
		name = regexp.MustCompile(`(?i)\(?1(\.|st)\s*Gen(eration)?\.?\)?`).ReplaceAllString(name, "(2016)")
		name = regexp.MustCompile(`(?i)\(?2(\.|nd)\s*Gen(eration)?\.?\)?`).ReplaceAllString(name, "(2020)")
		name = regexp.MustCompile(`(?i)\(?3(\.|rd)\s*Gen(eration)?\.?\)?`).ReplaceAllString(name, "(2022)")
		name = regexp.MustCompile(`(?i)\s*(\d+)\s*(S)`).ReplaceAllString(name, " $1$2")

		name = strings.ReplaceAll(name, "11 (2020)", "11")
	}

	if s[0] == "Asus" {
		name = regexp.MustCompile(`(?i)ASUS`).ReplaceAllStringFunc(name, strings.ToUpper)
		name = regexp.MustCompile(`(?i)ROG`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Beafon" {
		name = regexp.MustCompile(`(?i)^bea-?fon`).ReplaceAllString(name, "beafon")

		name = regexp.MustCompile(`(?i)-EU\d{3}[B]\s*$`).ReplaceAllString(name, "")

		name = regexp.MustCompile(`(?i)MX\s*(\d+)`).ReplaceAllString(name, "MX$1")
	}

	if s[0] == "Blackview" {
		name = regexp.MustCompile(`(?i)BL\s*(\d+)`).ReplaceAllString(name, "BL$1")
		name = regexp.MustCompile(`(?i)BV\s*(\d+)`).ReplaceAllString(name, "BV$1")
	}

	if s[0] == "Cat" {
		name = regexp.MustCompile(`(?i)CAT`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Cyrus" {
		name = regexp.MustCompile(`(?i)CS\d{2}XA`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Doogee" {
		name = regexp.MustCompile(`(?i)DOOGEE`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Doro" {
		name = regexp.MustCompile(`(?i)PLUS`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Emporia" || len(name) > 7 && name[:7] == "Emporia" {
		name = strings.ToUpper(name)
		name = regexp.MustCompile(`(?i)^emporia\s*`).ReplaceAllString(name, "emporia")
		name = regexp.MustCompile(`(?i)basic|glam|mini|smart$`).ReplaceAllStringFunc(name, strings.ToLower)
	}

	if s[0] == "Gigaset" {
		name = strings.ReplaceAll(name, " Ls", "LS")

		name = regexp.MustCompile(`(?i)G[LSX]\d+(LS)?`).ReplaceAllStringFunc(name, strings.ToUpper)
		name = regexp.MustCompile(`(?i)senior$`).ReplaceAllStringFunc(name, strings.ToLower)
		name = regexp.MustCompile(`(?i)LITE$`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Google" {
		name = regexp.MustCompile(`(?i)\d+[A]$`).ReplaceAllStringFunc(name, strings.ToLower)
	}

	if s[0] == "Honor" {
		name = regexp.MustCompile(`(?i)HONOR`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`(?i)Magic\s*(\d)\s*(\w)?`).ReplaceAllString(name, "Magic$1 $2")
		name = regexp.MustCompile(`(?i)HONOR\s*(\d+)\s*(\w)?`).ReplaceAllString(name, "HONOR $1 $2")
	}

	if s[0] == "Huawei" {
		name = regexp.MustCompile(`(?i)HUAWEI`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`(?i)Ascend\s*([Y])\s*(\d+)`).ReplaceAllString(name, "Ascend $1$2")
		name = regexp.MustCompile(`(?i)Magic\s*(\d+)`).ReplaceAllString(name, "Magic$1")
		name = regexp.MustCompile(`(?i)Mate\s*(\d+)`).ReplaceAllString(name, "Mate $1")
		name = regexp.MustCompile(`(?i)nova\s*([Y]?\d+)?`).ReplaceAllString(name, "nova $1")
		name = regexp.MustCompile(`(?i)P\s*(smart[+]?)\s*(\+?)`).ReplaceAllString(name, "P smart$2 $3")
		name = regexp.MustCompile(`(?i)P-?\s*(\d+)`).ReplaceAllString(name, "P$1")

		name = regexp.MustCompile(`(?i)lite`).ReplaceAllString(name, "lite")

		name = regexp.MustCompile(`(?i)\d+i$`).ReplaceAllStringFunc(name, strings.ToLower)
	}

	if s[0] == "I.safe" {
		name = strings.ToUpper(name)
		name = regexp.MustCompile(`(?i)^i.safe`).ReplaceAllStringFunc(name, strings.ToLower)
	}

	if s[0] == "IIif150" || s[0] == "IIIf150" {
		name = regexp.MustCompile(`(?i)^IIIF150`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Infinix" {
		name = strings.ToUpper(name)
		name = strings.ReplaceAll(name, "INFINIX", "Infinix")
		name = regexp.MustCompile(`(?i)\d+i\b`).ReplaceAllStringFunc(name, strings.ToLower)
	}

	if s[0] == "Inoi" {
		name = regexp.MustCompile(`(?i) A\s*(\d+)`).ReplaceAllString(name, " A$1")
	}

	if s[0] == "Moto" {
		name = "Motorola" + " " + name

		s = strings.Split(name, " ")
	}

	if s[0] == "Motorola" {
		name = strings.ToLower(name)

		// name = regexp.MustCompile(`(?i)motorola`).ReplaceAllStringFunc(name, strings.ToLower)

		if len(s) > 1 && s[1] == "Moto" && s[2] == "Edge" {
			name = strings.ReplaceAll(name, "moto ", "")
		}

		name = strings.ReplaceAll(name, "2nd gen", "???")

		name = regexp.MustCompile(`(?i)edge\s*(\d+)\s*(\w*)`).ReplaceAllString(name, "edge $1 $2")
		name = regexp.MustCompile(`(?i)^motorola\s*(moto\s*)?([eg]\s*)?([eg])\s*(\d+)`).ReplaceAllString(name, "motorola moto $3$4")

		name = strings.ReplaceAll(name, "???", "2nd gen")

		name = strings.ReplaceAll(name, "defy (2021)", "defy")
	}

	if s[0] == "Nothing" {
		name = regexp.MustCompile(`-?\(?(\d+)\)?`).ReplaceAllString(name, "($1)")
	}

	if s[0] == "Oukitel" {
		name = regexp.MustCompile(`(?i)OUKITEL`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`(?i)WP\s*(\d+)`).ReplaceAllString(name, "WP$1")
	}

	if s[0] == "OnePlus" {
		name = regexp.MustCompile(`(?i)\s*CPH\d{4}`).ReplaceAllString(name, "")
		name = regexp.MustCompile(`(?i)CE\s*(\d+)`).ReplaceAllString(name, "CE $1")
	}

	if s[0] == "Oppo" {
		name = regexp.MustCompile(`(?i)OPPO`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`(?i)\s*CPH\d{4}`).ReplaceAllString(name, "")

		name = regexp.MustCompile(`(?i)([A])\s*(\d+)\s*([ks])?\s*\(?(\d{4})?\)?`).ReplaceAllString(name, "$1$2$3 $4")
		name = regexp.MustCompile(`[KS]\s*$`).ReplaceAllStringFunc(name, strings.ToLower)

		name = regexp.MustCompile(`(?i)Reno\s*(\d+)\s*(\w)?`).ReplaceAllString(name, "Reno$1 $2")
		name = regexp.MustCompile(`(?i)OPPO\s*(\d+)\s*(\w)?`).ReplaceAllString(name, "OPPO Reno$1 $2")
	}

	if s[0] == "Oscal" {
		name = regexp.MustCompile(`(?i)OSCAL`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Realme" {
		name = regexp.MustCompile(`(?i)realme`).ReplaceAllStringFunc(name, strings.ToLower)

		name = regexp.MustCompile(`(?i)narzo`).ReplaceAllStringFunc(name, strings.ToLower)

		name = regexp.MustCompile(`(?i)GT\s*(\d+)`).ReplaceAllString(name, "GT $1")

		name = regexp.MustCompile(`(?i)\d+i`).ReplaceAllStringFunc(name, strings.ToLower)
		name = regexp.MustCompile(`(?i)\d+[y]`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Galaxy" {
		name = "Samsung" + " " + name

		s = strings.Split(name, " ")
	}

	if s[0] == "Samsung" {
		name = regexp.MustCompile(`[^FE][E]$`).ReplaceAllStringFunc(name, strings.ToLower)

		name = regexp.MustCompile(`(?i)\s+(SM-)?[AFMS]\d{3}[BFR]?[N]?(\/DSN?)?`).ReplaceAllString(name, "")

		name = regexp.MustCompile(`Note\s*(\d+)`).ReplaceAllString(name, "Note $1")
		name = regexp.MustCompile(`(?i)( Galaxy)? (Tab )?([AFMS])\s*(\d+| Duos)`).ReplaceAllString(name, " Galaxy $2$3$4")

		name = regexp.MustCompile(`(?i)FieldPro`).ReplaceAllString(name, "Field Pro")

		name = regexp.MustCompile(`\(\d{4}\)`).ReplaceAllString(name, "") // remove year annotation
	}

	if s[0] == "Sony" {
		name = regexp.MustCompile(`[DF]\d{4}`).ReplaceAllString(name, "")

		name = regexp.MustCompile(`(?i)\s*I+$`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Vivo" {
		name = regexp.MustCompile(`(?i)vivo`).ReplaceAllStringFunc(name, strings.ToLower)

		name = regexp.MustCompile(`(?i)NEX`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`\d+[CEIS]\b`).ReplaceAllStringFunc(name, strings.ToLower)
	}

	if s[0] == "Wiko" {
		name = regexp.MustCompile(`(?i)View\s*(\d+)`).ReplaceAllString(name, "View $1")
	}

	if s[0] == "Poco" || s[0] == "Redmi" {
		name = "Xiaomi" + " " + name

		s = strings.Split(name, " ")
	}

	if s[0] == "Xiaomi" {
		name = regexp.MustCompile(`^Xiaomi\s*(\d+)\s*\(?(\d{4})\)?`).ReplaceAllString(name, "Xiaomi Redmi $1 ($2)")
		name = regexp.MustCompile(`(?i)(POCO\s*|Pocophone\s*)?([CFMX]\d+)`).ReplaceAllString(name, "POCO $2")
		name = regexp.MustCompile(`(Redmi\s*)?Note\s*(\d+)`).ReplaceAllString(name, "Redmi Note $2")
		// name = regexp.MustCompile(`Note\s*(\d+)`).ReplaceAllString(name, "Note $1")
		name = regexp.MustCompile(`(?i)Redmi\s*(\d+)\s*([ABC])`).ReplaceAllString(name, "Redmi $1$2")

		// name = regexp.MustCompile(`(?i)\d+[A]$`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`\s*\(?(\d{4})\)?`).ReplaceAllString(name, " $1")

		name = regexp.MustCompile(`(?i)Mi(\d+)`).ReplaceAllString(name, "Mi $1")
	}

	if s[0] == "ZTE" {
		name = regexp.MustCompile(`(?i)nubia`).ReplaceAllString(name, "nubia")

		name = regexp.MustCompile(`(Blade\s*)?([AV])\s*(\d+)(\s*([Ss]))?`).ReplaceAllString(name, "Blade $2$3$5")
		name = regexp.MustCompile(`[S]\s*$`).ReplaceAllStringFunc(name, strings.ToLower)
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
