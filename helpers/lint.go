package lint

import (
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var colorRegex = regexp.MustCompile(`(?i)(in )?([/(,]?(Ancora|Aqua|Arctic|Astral|Astro|Atlantic|Aurora|Awesome|Azul|Bamboo|Bianco|Black|(Hell ?|Pazifik)?Blau|Bleu|Blue?|Burgund|Butter|Champagne|Charcoal|Chrome|Cloudy?|Clover|Copper|Cosmic|Cosmo|Dark|Denim|Diamond|Dusk|Electric|Elegant|Frost(ed)?|Galactic|Gelb|Glacier|Glazed|Glowing|Gold(ig)?|Gradient|Granite|Graphite|Gr[ae]y|Green|Grau|Gravity|Gris|Grün|Himalaya|Holunderweiss|Ic[ey]|Interstellar|Lagoon|Lake|Lavende[lr]|Light|(Dunkel)?Lila|Luminous|Magenta|Marine|Matte|Metallic|Meteorite|Meteor|Midday|Midnight|Mint(green)?|Misty?|Mitternacht(sschwarz)?|Moonlight|Mystic|Nachtgrün|Navy|Nero|Night|Noir|Ocean|Onyx|Orange| Oro|Pacific|Pastel|Peacock|Pearl|Pebble|Pepper|Perlmutweiss|Petrol|Piano|Pink(gold)?|Polar(stern)?|Prism|Purple|Red( Edition)?|Rosa|Rose|Ros[aée]gold|Rosso|Rot|Sage|Sakura|Salbeigrün|Sandy|Schwarz|Shadow|Silber|Silver|Sky|Space(grey)?|Stargaze|Starlight|Starry|Star|Steel|Sterling|Sternenblau|Sunburst|Sunrise|Sunset|Titanium|Titan|Türkis|Twilight|Violette?|Violet|Waldgrün|Weiss|Weiß|White|Yellow|Zeus)\b[\s/]?)(Azur|Black|Blau|Bleen|Blue|Bronze|Cream|Dream|Gold|Green|Gr[ae]y|Grün|Grau|Lime|Navy|Onyx|Pink|Rose|Schwarz|Silber|Silver|White|Weiss)?[)]?`)

var nameMapping = map[string]string{
	"2Nd":  "2nd",
	"2ND":  "2nd",
	"3Rd":  "3rd",
	"3RD":  "3rd",
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
	"Mclaren": "McLaren",
	"Mini":    "mini",
	"Mm":      "MM",
	"Myphone": "myPhone",
	"NEO2":    "Neo2",
	"Nfc":     "NFC",
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
	if loc := colorRegex.FindStringSubmatchIndex(strings.NewReplacer("ß", "ss", "é", "e").Replace(name)); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

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

		name = regexp.MustCompile(`\s+SE\s*\(?(2016|2020|2022)\)?`).ReplaceAllString(name, " SE ($1)")
		name = regexp.MustCompile(`(?i)\(?1(\.|st)\s*Gen(eration)?\.?\)?`).ReplaceAllString(name, "(2016)")
		name = regexp.MustCompile(`(?i)\(?2(\.|nd)\s*Gen(eration)?\.?\)?`).ReplaceAllString(name, "(2020)")
		name = regexp.MustCompile(`(?i)\(?3(\.|rd)\s*Gen(eration)?\.?\)?`).ReplaceAllString(name, "(2022)")
		name = regexp.MustCompile(`(?i)\s*(\d+)\s*(S)`).ReplaceAllString(name, " $1$2")

		name = strings.ReplaceAll(name, "SE 2", "SE (2020)")
		name = strings.ReplaceAll(name, "SE2", "SE (2020)")
		name = strings.ReplaceAll(name, "11 (2020)", "11")

		name = strings.ReplaceAll(name, "(2016) 2016", "(2016)")
		name = strings.ReplaceAll(name, "(2020) 2020", "(2020)")
		name = strings.ReplaceAll(name, "(2022) 2022", "(2022)")
		name = strings.ReplaceAll(name, "(2016) (2016)", "(2016)")
		name = strings.ReplaceAll(name, "(2020) (2020)", "(2020)")
		name = strings.ReplaceAll(name, "(2020) (2022)", "(2020)")
		name = strings.ReplaceAll(name, "(2022) (2022)", "(2022)")
	}

	if s[0] == "Asus" {
		name = regexp.MustCompile(`(?i)ASUS`).ReplaceAllStringFunc(name, strings.ToUpper)
		name = regexp.MustCompile(`(?i)ROG`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`(?i)Z[ES]\d{3}[KM]L`).ReplaceAllString(name, "")
	}

	if s[0] == "Beafon" {
		name = regexp.MustCompile(`(?i)^bea-?fon`).ReplaceAllString(name, "beafon")

		name = regexp.MustCompile(`(?i)-EU\d{3}[B]\s*$`).ReplaceAllString(name, "")

		name = regexp.MustCompile(`(?i)MX\s*(\d+)`).ReplaceAllString(name, "MX$1")
	}

	if s[0] == "Blackview" {
		name = regexp.MustCompile(`(?i)BL\s*(\d+)`).ReplaceAllString(name, "BL$1")
		name = regexp.MustCompile(`(?i)BV\s*(\d+)`).ReplaceAllString(name, "BV$1")
		// name = regexp.MustCompile(`(?i)(B[LV]\d+)\s*s`).ReplaceAllString(name, "$1s")
		name = regexp.MustCompile(`(?i)\(\d+\)`).ReplaceAllString(name, "")
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
		name = regexp.MustCompile(`(?i)HONOR\s*(\d+[X]?)\s*(\w)?`).ReplaceAllString(name, "HONOR $1 $2")
	}

	if s[0] == "Hmd" || s[0] == "HMD" {
		name = regexp.MustCompile(`(?i)HMD`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`(?i)TA-\d{4}`).ReplaceAllString(name, "")
	}

	if s[0] == "Huawei" {
		name = regexp.MustCompile(`(?i)HUAWEI`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`(?i)Ascend\s*([Y])\s*(\d+)`).ReplaceAllString(name, "Ascend $1$2")
		name = regexp.MustCompile(`(?i)Magic\s*(\d+)`).ReplaceAllString(name, "Magic$1")
		name = regexp.MustCompile(`(?i)Mate\s*(\d+)`).ReplaceAllString(name, "Mate $1")
		name = regexp.MustCompile(`(?i)nova\s*([Y]?\d+)?`).ReplaceAllString(name, "nova $1")
		name = regexp.MustCompile(`(?i)P\s*(smart[+]?)\s*(\+?)`).ReplaceAllString(name, "P smart$2 $3")
		name = regexp.MustCompile(`(?i)P-?\s*(\d+)`).ReplaceAllString(name, "P$1")
		name = regexp.MustCompile(`(?i)\s*SE$`).ReplaceAllString(name, " SE")

		name = regexp.MustCompile(`(?i)lite`).ReplaceAllString(name, "lite")

		name = regexp.MustCompile(`(?i)\d+i$`).ReplaceAllStringFunc(name, strings.ToLower)

		name = regexp.MustCompile(`(?i)NE$`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = strings.ReplaceAll(name, " Gx8", " GX8")
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

		if len(s) > 2 && s[1] == "Moto" && s[2] == "Edge" {
			name = strings.ReplaceAll(name, "moto ", "")
		}

		name = strings.ReplaceAll(name, "2nd gen", "???")

		name = regexp.MustCompile(`(?i)edge\s*(\d+)\s*(\w*)`).ReplaceAllString(name, "edge $1 $2")
		name = regexp.MustCompile(`(?i)^motorola\s*(moto\s*)?([eg]\s*)?([eg])\s*(\d+)`).ReplaceAllString(name, "motorola moto $3$4")

		name = strings.ReplaceAll(name, "???", "2nd gen")

		name = strings.ReplaceAll(name, "defy (2021)", "defy")
	}

	if s[0] == "Nothing" {
		name = regexp.MustCompile(`-?\(?(\d+[Aa]?)\)?`).ReplaceAllString(name, "($1)")
		name = regexp.MustCompile(`\(\d+[Aa]?\)`).ReplaceAllStringFunc(name, strings.ToLower)
	}

	if s[0] == "Oukitel" {
		name = regexp.MustCompile(`(?i)OUKITEL`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`(?i)WP\s*(\d+)`).ReplaceAllString(name, "WP$1")
	}

	if s[0] == "OnePlus" {
		name = regexp.MustCompile(`(?i)\s*CPH\d{4}`).ReplaceAllString(name, "")
		name = regexp.MustCompile(`(?i)CE\s*(\d+)`).ReplaceAllString(name, "CE $1")
		name = regexp.MustCompile(`(?i)Nord\s*(\d+)`).ReplaceAllString(name, "Nord $1")
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
		name = strings.ReplaceAll(name, "Neo ", "NEO ")

		name = regexp.MustCompile(`(?i)\d+i`).ReplaceAllStringFunc(name, strings.ToLower)
		name = regexp.MustCompile(`(?i)\d+[y]`).ReplaceAllStringFunc(name, strings.ToUpper)
	}

	if s[0] == "Galaxy" {
		name = "Samsung" + " " + name

		s = strings.Split(name, " ")
	}

	if s[0] == "Samsung" {
		name = regexp.MustCompile(`[^FE][E]$`).ReplaceAllStringFunc(name, strings.ToLower)
		name = regexp.MustCompile(`(?i)\s*FE$`).ReplaceAllString(name, " FE")

		// SM-A546B/DS-128GB
		name = regexp.MustCompile(`(?i)\s+(SM-)?[AFGMNS]\d{3}[BFPR]?[N]?(/DSN?)?`).ReplaceAllString(name, "")
		name = regexp.MustCompile(`(?i)\s+(GT-)?[N]\d{4}`).ReplaceAllString(name, "")

		name = regexp.MustCompile(`(?i)Z\s*Flip\s*(\d+)`).ReplaceAllString(name, "Z Flip$1")
		name = regexp.MustCompile(`(?i)Z\s*Fold\s*(\d+)`).ReplaceAllString(name, "Z Fold $1")
		name = regexp.MustCompile(`Note\s*(\d+)`).ReplaceAllString(name, "Note$1")
		name = regexp.MustCompile(`(?i)(Note\d+)\s+Plus`).ReplaceAllString(name, "$1+")
		name = regexp.MustCompile(`(?i)( Galaxy)? (Tab )?([AFMS])\s*(\d+| Duos)`).ReplaceAllString(name, " Galaxy $2$3$4")

		name = regexp.MustCompile(`(?i)FieldPro`).ReplaceAllString(name, "Field Pro")

		name = regexp.MustCompile(`(?i)X\s*Cover\s*(\d)`).ReplaceAllString(name, "XCover $1")
		name = regexp.MustCompile(`\s+\d[s]$`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = strings.ReplaceAll(name, "Samsung Note", "Samsung Galaxy Note")

		// name = regexp.MustCompile(`\(\d{4}\)`).ReplaceAllString(name, "") // remove year annotation
	}

	if s[0] == "Sony" {
		name = regexp.MustCompile(`[DF]\d{4}`).ReplaceAllString(name, "")

		name = regexp.MustCompile(`(?i)\s*(III|II|IV|V|VI|VII)$`).ReplaceAllStringFunc(name, strings.ToUpper)
		name = regexp.MustCompile(`(?i)\s*(III|II|IV|V|VI|VII)$`).ReplaceAllString(name, " $1")
	}

	if s[0] == "Vivo" {
		name = regexp.MustCompile(`(?i)vivo`).ReplaceAllStringFunc(name, strings.ToLower)

		name = regexp.MustCompile(`(?i)NEX`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = regexp.MustCompile(`\d+[CEIS]\b`).ReplaceAllStringFunc(name, strings.ToLower)
	}

	if s[0] == "Wiko" {
		name = regexp.MustCompile(`(?i)View\s*(\d+)`).ReplaceAllString(name, "View $1")
	}

	if s[0] == "Poco" || s[0] == "Redmi" || s[0] == "Mi" {
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

		name = regexp.MustCompile(`\d+[I]`).ReplaceAllStringFunc(name, strings.ToLower)
		name = regexp.MustCompile(`\d+[t]`).ReplaceAllStringFunc(name, strings.ToUpper)

		name = strings.ReplaceAll(name, "Mi Redmi ", "Mi ")
	}

	if s[0] == "ZTE" {
		name = regexp.MustCompile(`(?i)nubia`).ReplaceAllString(name, "nubia")

		name = regexp.MustCompile(`(Blade\s*)?([AV])\s*(\d+)(\s*([Ss]))?`).ReplaceAllString(name, "Blade $2$3$5")
		name = regexp.MustCompile(`[S]\s*$`).ReplaceAllStringFunc(name, strings.ToLower)
	}

	return strings.TrimSpace(name)
}
