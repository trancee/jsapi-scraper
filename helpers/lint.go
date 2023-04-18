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
	"Ce":      "CE",
	"Ds":      "DS",
	"El":      "EL",
	"Fe":      "FE",
	"Gt":      "GT",
	"Htc":     "HTC",
	"Ii":      "II",
	"IIIF150": "iiiF150",
	"Iphone":  "iPhone",
	"Kxd":     "KXD",
	"Lg":      "LG",
	"Oneplus": "OnePlus",
	"Realme":  "realme",
	"Se":      "SE",
	"Tcl":     "TCL",
	"Tecno":   "TECNO",
	"Umidigi": "UMIDIGI",
	"Xa":      "XA",
	"Xcover":  "XCover",
	"Xgody":   "XGODY",
	"Xr":      "XR",
	"Xs":      "XS",
	"Xzs":     "XZs",
	"Zte":     "ZTE",
}

var nameRegex = maps.Keys(nameMapping)

var r = regexp.MustCompile(`\W*(` + strings.Join(nameRegex, "|") + `)\W*`)
var r2 = regexp.MustCompile(`\W*([A-Z][a-z]*)?[0-9]+[A-Za-z]*\W*`)

func Title(name string) string {
	return cases.Title(language.Und, cases.NoLower).String(name)
}

func Lint(name string) string {
	if matches := r.FindAllStringSubmatch(name, -1); len(matches) > 0 {
		for _, match := range matches {
			name = strings.ReplaceAll(name, match[0], strings.ReplaceAll(match[0], match[1], nameMapping[match[1]]))
		}

		// fmt.Printf("%-30s %v\n", name, matches)
	}

	return name
}

func Model(name string) string {
	if matches := r2.FindAllStringSubmatch(name, -1); len(matches) > 0 {
		for _, match := range matches {
			if len(match) > 1 && (match[1] == "Reno" || match[1] == "Magic") {
				continue
			}

			oldModel := strings.ToUpper(strings.TrimSpace(match[0]))
			newModel := oldModel
			if n := len(oldModel); n > 1 {
				if s := oldModel[n-1 : n]; s == "I" || s == "E" {
					newModel = oldModel[:n-1] + strings.ToLower(s)
					// fmt.Println("*** MATCH:" + oldModel + "/" + oldModel[n-1:n] + "/" + newModel)
				}
			}

			name = strings.ReplaceAll(name, match[0], strings.ReplaceAll(strings.ToUpper(match[0]), oldModel, newModel))
		}

		// fmt.Printf("%-30s %v\n", name, matches)
	}

	return name
}
