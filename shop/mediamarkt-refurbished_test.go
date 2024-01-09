package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var mediamarktRefurbishedNames = []string{
	"Apple iPhone 11 128GB, Grün",
	"Apple iPhone 11 128GB, Rot",
	"Apple iPhone 11 128GB, Schwarz",
	"Apple iPhone 11 128GB, Weiss",
	"Apple iPhone 11 64GB, Rot",
	"Apple iPhone 11 64GB, Violett",
	"Apple iPhone 11 Pro 64GB, Gold",
	"Apple iPhone 11 Pro 64GB, Silber",
	"Apple iPhone 11 Pro 64GB, Spacegrau",
	"Apple iPhone 11 Pro Max 64GB, Grün",
	"Apple iPhone 11 Pro Max 64GB, Spacegrau",
	"Apple iPhone 12 Mini 64GB, Rot",
	"Apple iPhone 12 Mini 64GB, Schwarz",
	"Apple iPhone 8 256GB, Spacegrau",
	"Apple iPhone 8 64GB, Spacegrau",
	"Apple iPhone 8 Plus 64GB, Silber",
	"Apple iPhone 8 Plus 64GB, Spacegrau",
	"Apple iPhone SE 2020 128GB, Schwarz",
	"Apple iPhone SE 2020 64GB, Schwarz",
	"Apple iPhone X 64GB, Silber",
	"Apple iPhone XR 64GB, Blau",
	"Apple iPhone XR 64GB, Schwarz",
	"Apple iPhone XS 256GB, Gold",
	"Apple iPhone XS 64GB, Silber",
	"Apple iPhone XS 64GB, Spacegrau",
	"Apple iPhone XS Max 256GB, Gold",
	"Apple iPhone XS Max 256GB, Silber",
	"Apple iPhone XS Max 512GB, Gold",
	"Apple iPhone XS Max 512GB, Silber",
	"Apple iPhone XS Max 64GB, Gold",
	"Apple iPhone XS Max 64GB, Silber",
	"Huawei Mate 10 Pro (dual sim) 128GB, Blau",
	"Huawei Mate 10 Pro (dual sim) 128 GB Schwarz - refurbished Smartphone",
	"Huawei Mate 20 128GB, Blau",
	"Huawei Mate 20 128GB, Schwarz",
	"Huawei P20 Pro (dual sim) 128GB, Blau",
	"Huawei P20 Pro (dual sim) 128GB, Violett",
	"Huawei P30 Lite (dual sim) 128GB, Schwarz",
	"Huawei P30 Pro (dual sim) 128GB, Grün",
	"Huawei P30 Pro (dual sim) 128GB, Schwarz",
	"Huawei P30 Pro (dual sim) 128GB, Violett",
	"Samsung Galaxy A40 64GB, Weiss",
	"Samsung Galaxy A51 (dual sim) 128GB, Blau",
	"Samsung Galaxy A51 (dual sim) 128GB, Schwarz",
	"Samsung Galaxy A51 (dual sim) 128GB, Weiss",
	"Samsung Galaxy A51 (dual sim) 64GB, Schwarz",
	"Samsung Galaxy A8 (2018) dual sim 32GB, Schwarz",
	"Samsung Galaxy Note 10 (dual sim) 128GB, Blau",
	"Samsung Galaxy Note 10+ (dual sim) 256GB, Schwarz",
	"Samsung Galaxy Note 10+ (dual sim) 256GB, Violett",
	"Samsung Galaxy S10 (dual sim) 128GB, Blau",
	"Samsung Galaxy S10 (dual sim) 128GB, Grün",
	"Samsung Galaxy S10 (dual sim) 128GB, Rot",
	"Samsung Galaxy S10 (dual sim) 128GB, Schwarz",
	"Samsung Galaxy S10 (dual sim) 128GB, Weiss",
	"Samsung Galaxy S10+ (dual sim) 128GB, Blau",
	"Samsung Galaxy S10+ (dual sim) 128GB, Grün",
	"Samsung Galaxy S10+ (dual sim) 128GB, Schwarz",
	"Samsung Galaxy S10+ (dual sim) 128GB, Weiss",
	"Samsung Galaxy S10+ (dual sim) 512GB, Schwarz",
	"Samsung Galaxy S10+ (dual sim) 512GB, Weiss",
	"Samsung Galaxy S10e (dual sim) 128GB, Schwarz",
	"Samsung Galaxy S10e (dual sim) 128 GB Schwarz - refurbished Smartphone",
	"Samsung Galaxy S10e Dual sim, 128GB, Weiss",
	"Samsung Galaxy S20 5G (dual sim) 128GB, Blau",
	"Samsung Galaxy S20 5G (dual sim) 128GB, Grau",
	"Samsung Galaxy S20 5G (dual sim) 128GB, Rosé",
	"Samsung Galaxy S20 5G (dual sim) 128GB, Weiss",
	"Samsung Galaxy S20 FE 5G 128GB, Blau",
	"Samsung Galaxy S20 FE 5G 128GB, Grün",
	"Samsung Galaxy S20 FE 5G 128GB, Rot",
	"Samsung Galaxy S20 Ultra 5G (mono sim) 128GB, Grau",
	"Samsung Galaxy S20 Ultra 5G (mono sim) 128GB, Schwarz",
	"Samsung Galaxy S20 Ultra 5G (mono sim) 128GB, Weiss",
	"Samsung Galaxy S20+ (dual sim) 128GB, Grau",
	"Samsung Galaxy S20+ (dual sim) 128GB, Schwarz",
	"Samsung Galaxy S20+ 5G (dual sim) 128GB, Blau",
	"Samsung Galaxy S20+ 5G (dual sim) 128GB, Grau",
	"Samsung Galaxy S20+ 5G (dual sim) 128GB, Schwarz",
	"Samsung Galaxy S20+ 5G (dual sim) 128GB, Weiss",
	"Samsung Galaxy S21 5G (dual sim) 128GB, Grau",
	"Samsung Galaxy S21 5G (dual sim) 128GB, Rosé",
	"Samsung Galaxy S21 5G (dual sim) 128GB, Violett",
	"Samsung Galaxy S21 5G (dual sim) 128GB, Weiss",
	"Samsung Galaxy S21 5G (dual sim) 256GB, Grau",
	"Samsung Galaxy S21 5G (dual sim) 256GB, Rosé",
	"Samsung Galaxy S21 5G (dual sim) 256GB, Violett",
	"Samsung Galaxy S21 5G (dual sim) 256GB, Weiss",
	"Samsung Galaxy S21 5G (mono sim) 256GB, Grau",
	"Samsung Galaxy S21 FE 5G (dual sim) 128GB, Schwarz",
	"Samsung Galaxy S21+ 5G (dual sim) 128GB, Schwarz",
	"Samsung Galaxy S21+ 5G (dual sim) 128GB, Violett",
	"Samsung Galaxy S21+ 5G (dual sim) 128GB, Weiss",
	"Samsung Galaxy S21+ 5G (mono sim) 128GB, Schwarz",
	"Samsung Galaxy S21+ 5G (mono sim) 128GB, Weiss",
	"Samsung Galaxy S8 64GB, Schwarz",
	"Samsung Galaxy S8 64GB, Silber",
	"Samsung Galaxy S9 (dual sim) 64GB, Schwarz",
	"Samsung Galaxy S9 (mono sim) 64GB, Blau",
	"Samsung Galaxy S9 (mono sim) 64GB, Schwarz",
	"Samsung Galaxy S9+ (dual sim) 64GB, Violett",
	"Samsung Galaxy S9+ (mono sim) 64GB, Blau",
	"Samsung Galaxy S9+ (mono sim) 64GB, Schwarz",
}

var mediamarktRefurbishedNamesExpected = []string{
	"Apple iPhone 11",
	"Apple iPhone 11",
	"Apple iPhone 11",
	"Apple iPhone 11",
	"Apple iPhone 11",
	"Apple iPhone 11",
	"Apple iPhone 11 Pro",
	"Apple iPhone 11 Pro",
	"Apple iPhone 11 Pro",
	"Apple iPhone 11 Pro Max",
	"Apple iPhone 11 Pro Max",
	"Apple iPhone 12 mini",
	"Apple iPhone 12 mini",
	"Apple iPhone 8",
	"Apple iPhone 8",
	"Apple iPhone 8 Plus",
	"Apple iPhone 8 Plus",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE (2020)",
	"Apple iPhone X",
	"Apple iPhone XR",
	"Apple iPhone XR",
	"Apple iPhone XS",
	"Apple iPhone XS",
	"Apple iPhone XS",
	"Apple iPhone XS Max",
	"Apple iPhone XS Max",
	"Apple iPhone XS Max",
	"Apple iPhone XS Max",
	"Apple iPhone XS Max",
	"Apple iPhone XS Max",
	"HUAWEI Mate 10 Pro",
	"HUAWEI Mate 10 Pro",
	"HUAWEI Mate 20",
	"HUAWEI Mate 20",
	"HUAWEI P20 Pro",
	"HUAWEI P20 Pro",
	"HUAWEI P30 lite",
	"HUAWEI P30 Pro",
	"HUAWEI P30 Pro",
	"HUAWEI P30 Pro",
	"Samsung Galaxy A40",
	"Samsung Galaxy A51",
	"Samsung Galaxy A51",
	"Samsung Galaxy A51",
	"Samsung Galaxy A51",
	"Samsung Galaxy A8 (2018)",
	"Samsung Galaxy Note 10",
	"Samsung Galaxy Note 10+",
	"Samsung Galaxy Note 10+",
	"Samsung Galaxy S10",
	"Samsung Galaxy S10",
	"Samsung Galaxy S10",
	"Samsung Galaxy S10",
	"Samsung Galaxy S10",
	"Samsung Galaxy S10+",
	"Samsung Galaxy S10+",
	"Samsung Galaxy S10+",
	"Samsung Galaxy S10+",
	"Samsung Galaxy S10+",
	"Samsung Galaxy S10+",
	"Samsung Galaxy S10e",
	"Samsung Galaxy S10e",
	"Samsung Galaxy S10e",
	"Samsung Galaxy S20",
	"Samsung Galaxy S20",
	"Samsung Galaxy S20",
	"Samsung Galaxy S20",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 Ultra",
	"Samsung Galaxy S20 Ultra",
	"Samsung Galaxy S20 Ultra",
	"Samsung Galaxy S20+",
	"Samsung Galaxy S20+",
	"Samsung Galaxy S20+",
	"Samsung Galaxy S20+",
	"Samsung Galaxy S20+",
	"Samsung Galaxy S20+",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S21+",
	"Samsung Galaxy S21+",
	"Samsung Galaxy S21+",
	"Samsung Galaxy S21+",
	"Samsung Galaxy S21+",
	"Samsung Galaxy S8",
	"Samsung Galaxy S8",
	"Samsung Galaxy S9",
	"Samsung Galaxy S9",
	"Samsung Galaxy S9",
	"Samsung Galaxy S9+",
	"Samsung Galaxy S9+",
	"Samsung Galaxy S9+",
}

func TestMediamarktRefurbishedClean(t *testing.T) {
	for i, name := range mediamarktRefurbishedNames {
		if _name := shop.MediamarktRefurbishedCleanFn(name); _name != mediamarktRefurbishedNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, mediamarktRefurbishedNamesExpected[i], name)
		}
	}
}
