package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var brackNames = []string{
	"Recommerce Switzerland SA iPhone 8 64GB Space Grau – refurbished",
	"Apple iPhone 11 128GB Schwarz",
	"Apple iPhone 11 64GB Schwarz",
	"Apple iPhone 12 64GB Schwarz",
	"Apple iPhone 13 128GB Grün",
	"Apple iPhone 13 128GB Mitternacht",
	"Apple iPhone 13 128GB Rosé",
	"Apple iPhone 13 256GB Blau",
	"Apple iPhone 13 256GB Polarstern",
	"Apple iPhone 13 256GB PRODUCT(RED)",
	"Apple iPhone 13 mini 128GB PRODUCT(RED)",
	"Apple iPhone 13 mini 256GB Grün",
	"Apple iPhone 13 mini 256GB Polarstern",
	"Apple iPhone 14 128 GB Gelb",
	"Apple iPhone 14 128 GB Mitternacht",
	"Apple iPhone 14 128 GB Violett",
	"Apple iPhone 14 256 GB Blau",
	"Apple iPhone 14 256 GB Polarstern",
	"Apple iPhone 14 256 GB PRODUCT(RED)",
	"Apple iPhone 14 Plus 128 GB Blau",
	"Apple iPhone 14 Plus 256 GB Gelb",
	"Apple iPhone 14 Plus 256 GB Mitternacht",
	"Apple iPhone 14 Plus 256 GB Polarstern",
	"Apple iPhone 14 Plus 256 GB PRODUCT(RED)",
	"Apple iPhone 14 Plus 256 GB Violett",
	"Apple iPhone 14 Pro 128 GB Dunkellila",
	"Apple iPhone 14 Pro 128 GB Gold",
	"Apple iPhone 14 Pro 256 GB Silber",
	"Apple iPhone 14 Pro 256 GB Space Schwarz",
	"Apple iPhone 14 Pro Max 128 GB Dunkellila",
	"Apple iPhone 14 Pro Max 128 GB Silber",
	"Apple iPhone 14 Pro Max 256 GB Gold",
	"Apple iPhone 14 Pro Max 256 GB Space Schwarz",
	"Apple iPhone SE 3. Gen. 128 GB PRODUCT(RED)",
	"Apple iPhone SE 3. Gen. 64 GB Mitternacht",
	"Apple iPhone SE 3. Gen. 64 GB Polarstern",
	"ASUS ROG Phone 6 256 GB / 12 GB Phantom Schwarz",
	"ASUS ROG Phone 6 256 GB / 12 GB Sturm Weiss",
	"ASUS ROG Phone 6 512 GB / 16 GB Phantom Schwarz",
	"ASUS ROG Phone 6 521 GB / 16 GB Sturm Weiss",
	"Fairphone 4 5G 128 GB Grau",
	"Fairphone 4 5G 256 GB Grau",
	"Fairphone 4 5G 256 GB Grün",
	"Google Pixel 6a 128 GB Charcoal",
	"Google Pixel 7 128 GB Snow",
	"Google Pixel 7 256 GB Lemongrass",
	"Google Pixel 7 256 GB Obsidian",
	"Google Pixel 7 Pro 128 GB Hazel",
	"Google Pixel 7 Pro 128 GB Snow",
	"Google Pixel 7 Pro 256 GB Obsidian",
	"Huawei Mate 50 Pro 256 GB Schwarz",
	"Huawei Mate 50 Pro 256 GB Silber",
	"Motorola Razr 5G 256GB",
	"Nokia C21 32 GB Blau",
	"Nokia G21 128 GB Dusk",
	"Nokia X30 5G 128 GB Cloudy blue",
	"Nokia X30 5G 128 GB Ice white",
	"Nokia X30 5G 256 GB Cloudy blue",
	"Nothing Phones Phone (1) 8 GB / 128 GB",
	"Nothing Phones Phone (1) 8 GB / 256 GB",
	"OnePlus 11 5G 128 GB Titan Black",
	"OnePlus 11 5G 256 GB Titan Black",
	"OnePlus 8 Pro 256GB Glacial Green",
	"OPPO A54s 128 GB Crystal Black",
	"OPPO A57s 128 GB Sky Blue",
	"OPPO A57s 128 GB Starry Black",
	"OPPO A94 5G 128 GB Cosmo Blue",
	"OPPO A94 5G 128 GB Fluid Black",
	"OPPO A96 128 GB Sunset Blue",
	"OPPO Find N2 Flip 256 GB Astral Black",
	"OPPO Find N2 Flip 256 GB Moonlit Purple",
	"OPPO Find X3 Lite 128 GB Black",
	"OPPO Find X3 Lite 128 GB Blue",
	"OPPO Find X3 Pro 256 GB Black",
	"OPPO Find X3 Pro 256 GB Blue",
	"OPPO Find X5 256 GB Schwarz",
	"OPPO Find X5 256 GB Weiss",
	"OPPO Find X5 Lite 256 GB Hellblau",
	"OPPO Find X5 Lite 256 GB Schwarz",
	"OPPO Find X5 Pro 256 GB Schwarz",
	"OPPO Find X5 Pro 256 GB Weiss",
	"OPPO Reno2 256GB Ocean Blue",
	"OPPO Reno6 Pro 5G Arctic Blue",
	"OPPO Reno8 256 GB Shimmer Black",
	"OPPO Reno8 256 GB Shimmer Gold",
	"OPPO Reno8 Lite 128 GB Cosmic Black",
	"OPPO Reno8 Lite 128 GB Rainbow Spectrum",
	"OPPO Reno8 Pro 256 GB Glazed Black",
	"OPPO Reno8 Pro 256 GB Glazed Green",
	"Realme GT2 5G 128 GB Steel Black",
	"Samsung Galaxy A13 128 GB CH Black",
	"Samsung Galaxy A14 128 GB CH Black",
	"Samsung Galaxy A14 128 GB CH Lime Green",
	"Samsung Galaxy A14 128 GB CH Silver",
	"Samsung Galaxy A14 5G 128 GB CH Black",
	"Samsung Galaxy A14 5G 128 GB CH Lime Green",
	"Samsung Galaxy A14 5G 128 GB CH Silver",
	"Samsung Galaxy A33 5G 128 GB CH Awesome Blue",
	"Samsung Galaxy A33 5G 128 GB CH Awesome White",
	"Samsung Galaxy A34 5G 128 GB CH Awesome Graphite",
	"Samsung Galaxy A34 5G 128 GB CH Awesome Lime",
	"Samsung Galaxy A34 5G 128 GB CH Enterprise Edition Awesome Graphite",
	"Samsung Galaxy A34 5G 256 GB CH Awesome Silver",
	"Samsung Galaxy A34 5G 256 GB CH Awesome Violet",
	"Samsung Galaxy A54 5G 128 GB CH Awesome Violet",
	"Samsung Galaxy A54 5G 128 GB CH Awesome White",
	"Samsung Galaxy A54 5G 128 GB CH Enterprise Edition Awesome Graphite",
	"Samsung Galaxy A54 5G 256 GB CH Awesome Graphite",
	"Samsung Galaxy A54 5G 256 GB CH Awesome Lime",
	"Samsung Galaxy S22 5G 128 GB CH Phantom Black",
	"Samsung Galaxy S23 128 GB CH Green",
	"Samsung Galaxy S23 128 GB CH Lavender",
	"Samsung Galaxy S23 128 GB CH Phantom Black",
	"Samsung Galaxy S23 256 GB CH Cream",
	"Samsung Galaxy S23 Ultra 256 GB CH Cream",
	"Samsung Galaxy S23 Ultra 256 GB CH Lavender",
	"Samsung Galaxy S23 Ultra 256 GB CH Phantom Black",
	"Samsung Galaxy S23 Ultra 512 GB CH Green",
	"Samsung Galaxy S23+ 256 GB CH Cream",
	"Samsung Galaxy S23+ 256 GB CH Green",
	"Samsung Galaxy S23+ 256 GB CH Lavender",
	"Samsung Galaxy S23+ 256 GB CH Phantom Black",
	"Samsung Galaxy S23+ 512 GB CH Green",
	"Samsung Galaxy S23+ 512 GB CH Phantom Black",
	"Samsung Galaxy XCover 5 Enterprise Edition CH",
	"Samsung Galaxy XCover 6 Pro Enterprise Edition CH",
	"Samsung Galaxy Z Flip4 5G 128 GB CH Bora Purple",
	"Samsung Galaxy Z Flip4 5G 256 GB CH Pink Gold",
	"Samsung Galaxy Z Flip4 5G 512 GB CH Graphite",
	"Samsung Galaxy Z Fold4 5G 256 GB CH Beige",
	"Samsung Galaxy Z Fold4 5G 512 GB CH Graygreen",
	"Samsung Galaxy Z Fold4 5G 512 GB CH Phantom Black",
	"Xiaomi 12 5G 128 GB Blau",
	"Xiaomi 12T 256 GB Silber",
	"Xiaomi 13 256 GB Schwarz",
	"Xiaomi 13 256 GB Weiss",
	"Xiaomi 13 Lite 128 GB Blau",
	"Xiaomi 13 Lite 128 GB Pink",
	"Xiaomi 13 Lite 128 GB Schwarz",
	"Xiaomi 13 Pro 256 GB Ceramic Black",
	"Xiaomi Poco F3 256 GB Ocean Blue",
	"Xiaomi Pocophone F2 Pro 128GB Violett",
	"Xiaomi Redmi 10 2022 128 GB Carbon Gray",
	"Xiaomi Redmi 12C 128 GB Gray",
	"Xiaomi Redmi 9A 32 GB Aurora Green",
	"Xiaomi Redmi 9A 32 GB Glacial Blue",
	"Xiaomi Redmi 9A 32 GB Granite Gray",
	"Xiaomi Redmi 9C 128 GB Midnight Grey",
	"Xiaomi Redmi Note 11 128 GB Grau",
	"Xiaomi Redmi Note 11 Pro 5G 128 GB Gray",
	"Xiaomi Redmi Note 11S 128 GB Twilight Blue",
	"Xiaomi Redmi Note 12 128 GB Blau",
	"Xiaomi Redmi Note 12 128 GB Grün",
	"Xiaomi Redmi Note 12 128 GB Schwarz",
	"Xiaomi Redmi Note 12 PRO 5G 128 GB Midnight Black",
	"Xiaomi Redmi Note 12 PRO 5G 128 GB Polar White",
	"Xiaomi Redmi Note 12 PRO 5G 128 GB Sky Blue",
	"Xiaomi Redmi Note 12 PRO+ 5G 256 GB Midnight Black",
	"Xiaomi Redmi Note 12 PRO+ 5G 256 GB Polar White",
	"Xiaomi Redmi Note 12 PRO+ 5G 256 GB Sky Blue",
}

var brackNamesExpected = []string{
	"Apple iPhone 8",
	"Apple iPhone 11",
	"Apple iPhone 11",
	"Apple iPhone 12",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 14",
	"Apple iPhone 14",
	"Apple iPhone 14",
	"Apple iPhone 14",
	"Apple iPhone 14",
	"Apple iPhone 14",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro Max",
	"Apple iPhone 14 Pro Max",
	"Apple iPhone 14 Pro Max",
	"Apple iPhone 14 Pro Max",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"ASUS ROG Phone 6",
	"ASUS ROG Phone 6",
	"ASUS ROG Phone 6",
	"ASUS ROG Phone 6",
	"Fairphone 4",
	"Fairphone 4",
	"Fairphone 4",
	"Google Pixel 6a",
	"Google Pixel 7",
	"Google Pixel 7",
	"Google Pixel 7",
	"Google Pixel 7 Pro",
	"Google Pixel 7 Pro",
	"Google Pixel 7 Pro",
	"HUAWEI Mate 50 Pro",
	"HUAWEI Mate 50 Pro",
	"motorola razr",
	"Nokia C21",
	"Nokia G21",
	"Nokia X30",
	"Nokia X30",
	"Nokia X30",
	"Nothing Phone (1)",
	"Nothing Phone (1)",
	"OnePlus 11",
	"OnePlus 11",
	"OnePlus 8 Pro",
	"OPPO A54s",
	"OPPO A57s",
	"OPPO A57s",
	"OPPO A94",
	"OPPO A94",
	"OPPO A96",
	"OPPO Find N2 Flip",
	"OPPO Find N2 Flip",
	"OPPO Find X3 Lite",
	"OPPO Find X3 Lite",
	"OPPO Find X3 Pro",
	"OPPO Find X3 Pro",
	"OPPO Find X5",
	"OPPO Find X5",
	"OPPO Find X5 Lite",
	"OPPO Find X5 Lite",
	"OPPO Find X5 Pro",
	"OPPO Find X5 Pro",
	"OPPO Reno2",
	"OPPO Reno6 Pro",
	"OPPO Reno8",
	"OPPO Reno8",
	"OPPO Reno8 Lite",
	"OPPO Reno8 Lite",
	"OPPO Reno8 Pro",
	"OPPO Reno8 Pro",
	"realme GT 2",
	"Samsung Galaxy A13",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy S22",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy XCover 5 EE",
	"Samsung Galaxy XCover 6 Pro EE",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Fold 4",
	"Samsung Galaxy Z Fold 4",
	"Samsung Galaxy Z Fold 4",
	"Xiaomi 12",
	"Xiaomi 12T",
	"Xiaomi 13",
	"Xiaomi 13",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Pro",
	"Xiaomi POCO F3",
	"Xiaomi POCO F2 Pro",
	"Xiaomi Redmi 10 2022",
	"Xiaomi Redmi 12C",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi Redmi Note 12 Pro+",
	"Xiaomi Redmi Note 12 Pro+",
	"Xiaomi Redmi Note 12 Pro+",
}

func TestBrackClean(t *testing.T) {
	for i, name := range brackNames {
		if _name := shop.BrackCleanFn(name); _name != brackNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, brackNamesExpected[i], name)
		}
	}
}
