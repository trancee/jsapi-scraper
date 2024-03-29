package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var alternateNames = []string{
	"Apple iPhone SE (2022) 64GB, Handy",
	"ASUS Zenfone 8 128GB Obsidian Black",
	"ASUS Zenfone 8 256GB Obsidian Black",
	"ASUS Zenfone 9 128GB inkl. Chromebook Bundle Midnight Black",
	"ASUS Zenfone 9 128GB Midnight Black",
	"ASUS Zenfone 9 128GB Moonlight White",
	"ASUS Zenfone 9 128GB Starry Blue",
	"ASUS Zenfone 9 256GB inkl. Chromebook Bundle Midnight Black",
	"ASUS Zenfone 9 256GB Midnight Black",
	"ASUS Zenfone 9 256GB Moonlight White",
	"Google Pixel 6 128GB Sorta Seafoam",
	"Google Pixel 6 128GB Stormy Black",
	"Google Pixel 6 Pro 128GB Cloudy White",
	"Google Pixel 6 Pro 128GB Stormy Black",
	"Google Pixel 6 Pro 256GB Stormy Black",
	"Google Pixel 6a 128GB Chalk",
	"Google Pixel 6a 128GB Charcoal",
	"Google Pixel 6a 128GB Sage",
	"Google Pixel 7 128GB Lemongrass",
	"Google Pixel 7 128GB Obsidian",
	"Google Pixel 7 Pro 128GB Obsidian",
	"Google Pixel 7 Pro 128GB Snow",
	"Honor 50 256GB Midnight Black",
	"Motorola Defy (2021) 64GB Schwarz",
	"Motorola Edge 20 128GB Frost Grau",
	"Motorola Edge 30 Neo 128GB Very Peri",
	"Motorola Edge 30 pro 256GB Cosmos Blue",
	"Motorola Edge 30 Ultra 256GB Interstellar Black",
	"Motorola Moto e20 32GB Graphite Gray",
	"Motorola Moto e22 32GB Astro Black",
	"Motorola Moto e22 64GB Astro Black",
	"Motorola Moto e22i 32GB Winter White",
	"Motorola Moto G22 64GB Cosmic Black",
	"Motorola Moto G22 64GB Iceberg Blue",
	"Motorola Moto G31 64GB Baby Blue",
	"Motorola Moto G31 64GB Mineral Grey",
	"Motorola Moto G42 64GB Atlantic Green",
	"Motorola Moto G42 64GB Metallic Rosé",
	"Motorola Moto G52 128GB Charcoal Grey",
	"Motorola Moto G60 128GB Dynamic Gray",
	"Motorola Moto G71 5G 128GB Meteorite Gray",
	"Motorola Moto G72 128GB Meteorite Gray",
	"Motorola Moto G82 128GB Meteorite Gray",
	"Motorola Moto G82 128GB White Lily",
	"Nokia G11 32GB Charcoal",
	"Nokia G11 32GB Eis",
	"Nokia G11 64GB Charcoal",
	"Nokia G21 64GB Dusk",
	"Nokia G60 5G 128GB Schwarz",
	"Nokia X10 128GB Forest",
	"Nokia X30 5G 128GB Cloudy Blue",
	"Nokia X30 5G 128GB Ice White",
	"Nokia X30 5G 256GB Ice White",
	"Nokia XR20 64GB Granite",
	"Nokia XR20 64GB Ultra Blue",
	"OnePlus 10T 5G 128GB Jade Green",
	"OnePlus 10T 5G 256GB Jade Green",
	"OnePlus 9 Pro 128GB Stellar Black",
	"Oppo A76 128GB Glowing Blue",
	"Oppo A96 128GB Starry Black",
	"Oppo Find X2 Lite 128GB Pearl White",
	"Oppo Find X2 Neo 256GB Starry Blue",
	"Oppo Find X3 Lite 128GB Starry Black",
	"Oppo Find X5 Lite 256GB Starry Black",
	"Oppo Reno5 5G 128GB Azure Blue",
	"Oppo Reno6 5G 128GB Stellar Black",
	"realme 9 5G 128GB Stargazer White",
	"realme 9i 128GB Prism Black",
	"realme 9i 64GB Prism Black",
	"realme C25Y 128GB Glacier Blue",
	"realme C25Y 128GB Metal Grey",
	"realme C31 64GB Dark Green",
	"realme GT2 Pro 128GB Paper Green",
	"realme GT2 Pro 128GB Paper White",
	"realme narzo 50A Prime 64GB Flash Black",
	"SAMSUNG Galaxy A04s 32GB Black",
	"SAMSUNG Galaxy A04s 32GB Green",
	"SAMSUNG Galaxy A13 (SM-A137) 128GB Black",
	"SAMSUNG Galaxy A13 (SM-A137) 128GB Light Blue",
	"SAMSUNG Galaxy A13 (SM-A137) 128GB White",
	"SAMSUNG Galaxy A13 (SM-A137) 32GB Black",
	"SAMSUNG Galaxy A13 (SM-A137) 32GB Light Blue",
	"SAMSUNG Galaxy A13 (SM-A137) 32GB White",
	"SAMSUNG Galaxy A13 (SM-A137) 64GB Black",
	"SAMSUNG Galaxy A13 (SM-A137) 64GB Light Blue",
	"SAMSUNG Galaxy A13 5G 128GB Awesome Black",
	"SAMSUNG Galaxy A13 5G 128GB White",
	"SAMSUNG Galaxy A13 5G 64GB Awesome Black",
	"SAMSUNG Galaxy A13 5G 64GB Silver Blue",
	"SAMSUNG Galaxy A13 5G 64GB White",
	"SAMSUNG Galaxy A22 5G 64GB White",
	"SAMSUNG Galaxy A23 5G 128GB Awesome Black",
	"SAMSUNG Galaxy A23 5G 128GB Light Blue",
	"SAMSUNG Galaxy A23 5G 128GB White",
	"SAMSUNG Galaxy A23 5G 64GB Awesome Black",
	"SAMSUNG Galaxy A23 5G 64GB Light Blue",
	"SAMSUNG Galaxy A23 5G 64GB White",
	"SAMSUNG Galaxy A32 Enterprise Edition 128GB, Handy",
	"SAMSUNG Galaxy A33 5G 128GB Awesome Blue",
	"SAMSUNG Galaxy A33 5G 128GB Awesome White",
	"SAMSUNG Galaxy A52s 5G 128GB Awesome Black",
	"SAMSUNG Galaxy A52s 5G Enterprise Edition 128GB Awesome Black",
	"SAMSUNG Galaxy A53 5G 128GB Awesome Black",
	"SAMSUNG Galaxy A53 5G 128GB Awesome Peach",
	"SAMSUNG Galaxy A53 5G 256GB Awesome Peach",
	"SAMSUNG Galaxy A53 5G 256GB Awesome White",
	"SAMSUNG Galaxy M13 64GB Light Blue",
	"SAMSUNG Galaxy M23 5G 128GB  Deep Green",
	"SAMSUNG Galaxy M23 5G 128GB orange-copper",
	"SAMSUNG Galaxy S20 FE 5G 128GB Cloud Mint",
	"SAMSUNG Galaxy S20 FE 5G 128GB Cloud Orange",
	"SAMSUNG Galaxy S20 FE 5G 128GB Cloud White",
	"SAMSUNG Galaxy S20 FE 5G 256GB Cloud Navy",
	"SAMSUNG Galaxy S21 FE 5G 128GB Graphite",
	"SAMSUNG Galaxy S21 FE 5G 128GB Lavender",
	"SAMSUNG Galaxy S21 FE 5G 128GB Olive",
	"SAMSUNG Galaxy S21 FE 5G 128GB White",
	"SAMSUNG Galaxy S21 FE 5G 256GB Graphite",
	"SAMSUNG Galaxy S21 FE 5G 256GB Lavender",
	"SAMSUNG Galaxy S22 128GB Green",
	"SAMSUNG Galaxy S22 128GB Phantom Black",
	"SAMSUNG Galaxy S22 128GB Phantom White",
	"SAMSUNG Galaxy S22 128GB Pink Gold",
	"SAMSUNG Galaxy S22 256GB Phantom White",
	"SAMSUNG Galaxy S22 Ultra 128GB Green",
	"SAMSUNG Galaxy S22 Ultra 128GB Phantom Black",
	"SAMSUNG Galaxy S22 Ultra 128GB Phantom White",
	"SAMSUNG Galaxy S23 128GB Cream",
	"SAMSUNG Galaxy S23 128GB Green",
	"SAMSUNG Galaxy S23 128GB Lavender",
	"SAMSUNG Galaxy S23 128GB Phantom Black",
	"SAMSUNG Galaxy S23 256GB Phantom Black",
	"SAMSUNG Galaxy XCover 5 64GB Black",
	"SAMSUNG Galaxy XCover6 Pro 128GB Black",
	"SAMSUNG Galaxy Z Flip3 5G 128GB Cream",
	"SAMSUNG Galaxy Z Flip3 5G 256GB Cream",
	"SAMSUNG Galaxy Z Flip4 128GB Blue",
	"SAMSUNG Galaxy Z Flip4 128GB Bora Purple",
	"SAMSUNG Galaxy Z Flip4 128GB Graphite",
	"SAMSUNG Galaxy Z Flip4 128GB Pink Gold",
	"SAMSUNG Galaxy Z Flip4 256GB Blue",
	"SAMSUNG Galaxy Z Flip4 256GB Graphite",
	"Xiaomi 11T 128GB Moonlight White",
	"Xiaomi 11T Pro 128GB Celestial Blue",
	"Xiaomi 11T Pro 128GB Meteorite Gray",
	"Xiaomi 11T Pro 128GB Moonlight White",
	"Xiaomi 11T Pro 256GB Celestial Blue",
	"Xiaomi 11T Pro 256GB Meteorite Gray",
	"Xiaomi 12 256GB Blue",
	"Xiaomi 12 256GB Gray",
	"Xiaomi 12 Lite 128GB Black",
	"Xiaomi 12 Lite 128GB Lite Green",
	"Xiaomi 12 Pro 256GB Blau",
	"Xiaomi 12 Pro 256GB Grau",
	"Xiaomi 12T 128GB Blau",
	"Xiaomi 12T 128GB Schwarz",
	"Xiaomi 12T 128GB Silber",
	"Xiaomi 12T 256GB Schwarz",
	"Xiaomi 12T Pro 256GB Blau",
	"Xiaomi 12T Pro 256GB Schwarz",
	"Xiaomi 12X 256GB Blue",
	"Xiaomi 13 256GB Schwarz",
	"Xiaomi 13 Lite 128GB Black",
	"Xiaomi 13 Lite 128GB Lite Blue",
	"Xiaomi 13 Lite 128GB Lite Pink",
	"Xiaomi Mi 11 Lite 5G NE 128GB Bubblegum Blue",
	"Xiaomi Mi 11 Lite 5G NE 128GB Mint Green",
	"Xiaomi Mi 11 Lite 5G NE 128GB Truffle Black",
	"Xiaomi Poco C40 32GB POCO Yellow",
	"Xiaomi Poco C40 32GB Power Black",
	"Xiaomi Poco C40 64GB Coral Green",
	"Xiaomi Poco C40 64GB POCO Yellow",
	"Xiaomi Poco C40 64GB Power Black",
	"Xiaomi Poco F4 128GB Night Black",
	"Xiaomi Poco F4 256GB Night Black",
	"Xiaomi Poco F4 GT 128GB Stealth Black",
	"Xiaomi Poco F4 GT 256GB Knight Silver",
	"Xiaomi Poco M5 128GB Black",
	"Xiaomi Poco M5 128GB Yellow",
	"Xiaomi Poco M5 64GB Black",
	"Xiaomi Poco M5 64GB Green",
	"Xiaomi Poco M5 64GB Yellow",
	"Xiaomi Poco M5s 64GB Grey",
	"Xiaomi Poco X5 5G 128GB Black",
	"Xiaomi Poco X5 5G 256GB Blue",
	"Xiaomi Poco X5 5G 256GB Green",
	"Xiaomi Poco X5 Pro 5G 128GB Blue",
	"Xiaomi Redmi 10  64GB Pebble White",
	"Xiaomi Redmi 10 (2022) 64GB Carbon Gray",
	"Xiaomi Redmi 10 (2022) 64GB Sea Blue",
	"Xiaomi Redmi 10 128GB Pebble White",
	"Xiaomi Redmi 10 5G 64GB Graphite Gray",
	"Xiaomi Redmi 10A 128GB Graphite Gray",
	"Xiaomi Redmi 10A 32GB Chrome Silver",
	"Xiaomi Redmi 10A 32GB Graphite Gray",
	"Xiaomi Redmi 10A 32GB Sky Blue",
	"Xiaomi Redmi 10A 64GB Chrome Silver",
	"Xiaomi Redmi 10A 64GB Graphite Gray",
	"Xiaomi Redmi 10A 64GB Sky Blue",
	"Xiaomi Redmi 10C 128GB Graphite Grey",
	"Xiaomi Redmi 10C 128GB Mint Green",
	"Xiaomi Redmi 10C 128GB Ocean Blue",
	"Xiaomi Redmi 10C 64GB Graphite Gray",
	"Xiaomi Redmi 10C 64GB Graphite Grey",
	"Xiaomi Redmi 10C 64GB Ocean Blue",
	"Xiaomi Redmi 9A 32GB Midnight Grey",
	"Xiaomi Redmi 9AT 32GB Glacial Blue",
	"Xiaomi Redmi 9AT 32GB Midnight Grey",
	"Xiaomi Redmi 9C 128GB Midnight Grey",
	"Xiaomi Redmi A1 32GB Black",
	"Xiaomi Redmi A1 32GB Light Green",
	"Xiaomi Redmi Note 10 Pro 128GB Glacier Blue",
	"Xiaomi Redmi Note 10 Pro 128GB Gradient Bronze",
	"Xiaomi Redmi Note 10 Pro 128GB Onyx Gray",
	"Xiaomi Redmi Note 10 Pro 64GB Glacier Blue",
	"Xiaomi Redmi Note 10 Pro 64GB Onyx Gray",
	"Xiaomi Redmi Note 10S 64GB Ocean Blue",
	"Xiaomi Redmi Note 10S 64GB Onyx Gray",
	"Xiaomi Redmi Note 11 128GB Twilight Blue",
	"Xiaomi Redmi Note 11 64GB Star Blue",
	"Xiaomi Redmi Note 11 64GB Twilight Blue",
	"Xiaomi Redmi Note 11 Pro 5G 128GB Mirage Blue",
	"Xiaomi Redmi Note 11 Pro 5G 128GB Phantom White",
	"Xiaomi Redmi Note 11 Pro 5G 128GB Stealth Black",
	"Xiaomi Redmi Note 11 Pro+ 5G 128GB Forest Green",
	"Xiaomi Redmi Note 11 Pro+ 5G 128GB Graphite Gray",
	"Xiaomi Redmi Note 11 Pro+ 5G 128GB Star Blue",
	"Xiaomi Redmi Note 11 Pro+ 5G 256GB Forest Green",
	"Xiaomi Redmi Note 11 Pro+ 5G 256GB Graphite Gray",
	"Xiaomi Redmi Note 11 Pro+ 5G 256GB Star Blue",
	"Xiaomi Redmi Note 11S 128GB Graphite Grey",
}

var alternateNamesExpected = []string{
	"Apple iPhone SE (2022)",
	"ASUS Zenfone 8",
	"ASUS Zenfone 8",
	"ASUS Zenfone 9",
	"ASUS Zenfone 9",
	"ASUS Zenfone 9",
	"ASUS Zenfone 9",
	"ASUS Zenfone 9",
	"ASUS Zenfone 9",
	"ASUS Zenfone 9",
	"Google Pixel 6",
	"Google Pixel 6",
	"Google Pixel 6 Pro",
	"Google Pixel 6 Pro",
	"Google Pixel 6 Pro",
	"Google Pixel 6a",
	"Google Pixel 6a",
	"Google Pixel 6a",
	"Google Pixel 7",
	"Google Pixel 7",
	"Google Pixel 7 Pro",
	"Google Pixel 7 Pro",
	"HONOR 50",
	"motorola defy",
	"motorola edge 20",
	"motorola edge 30 neo",
	"motorola edge 30 pro",
	"motorola edge 30 ultra",
	"motorola moto e20",
	"motorola moto e22",
	"motorola moto e22",
	"motorola moto e22i",
	"motorola moto g22",
	"motorola moto g22",
	"motorola moto g31",
	"motorola moto g31",
	"motorola moto g42",
	"motorola moto g42",
	"motorola moto g52",
	"motorola moto g60",
	"motorola moto g71",
	"motorola moto g72",
	"motorola moto g82",
	"motorola moto g82",
	"Nokia G11",
	"Nokia G11",
	"Nokia G11",
	"Nokia G21",
	"Nokia G60",
	"Nokia X10",
	"Nokia X30",
	"Nokia X30",
	"Nokia X30",
	"Nokia XR20",
	"Nokia XR20",
	"OnePlus 10T",
	"OnePlus 10T",
	"OnePlus 9 Pro",
	"OPPO A76",
	"OPPO A96",
	"OPPO Find X2 Lite",
	"OPPO Find X2 Neo",
	"OPPO Find X3 Lite",
	"OPPO Find X5 Lite",
	"OPPO Reno5",
	"OPPO Reno6",
	"realme 9",
	"realme 9i",
	"realme 9i",
	"realme C25Y",
	"realme C25Y",
	"realme C31",
	"realme GT 2 Pro",
	"realme GT 2 Pro",
	"realme narzo 50A Prime",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A22",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A32 EE",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A52s",
	"Samsung Galaxy A52s",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy M13",
	"Samsung Galaxy M23",
	"Samsung Galaxy M23",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22 Ultra",
	"Samsung Galaxy S22 Ultra",
	"Samsung Galaxy S22 Ultra",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover 6 Pro",
	"Samsung Galaxy Z Flip3",
	"Samsung Galaxy Z Flip3",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Flip4",
	"Xiaomi 11T",
	"Xiaomi 11T Pro",
	"Xiaomi 11T Pro",
	"Xiaomi 11T Pro",
	"Xiaomi 11T Pro",
	"Xiaomi 11T Pro",
	"Xiaomi 12",
	"Xiaomi 12",
	"Xiaomi 12 Lite",
	"Xiaomi 12 Lite",
	"Xiaomi 12 Pro",
	"Xiaomi 12 Pro",
	"Xiaomi 12T",
	"Xiaomi 12T",
	"Xiaomi 12T",
	"Xiaomi 12T",
	"Xiaomi 12T Pro",
	"Xiaomi 12T Pro",
	"Xiaomi 12X",
	"Xiaomi 13",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Lite",
	"Xiaomi Mi 11 Lite",
	"Xiaomi Mi 11 Lite",
	"Xiaomi Mi 11 Lite",
	"Xiaomi POCO C40",
	"Xiaomi POCO C40",
	"Xiaomi POCO C40",
	"Xiaomi POCO C40",
	"Xiaomi POCO C40",
	"Xiaomi POCO F4",
	"Xiaomi POCO F4",
	"Xiaomi POCO F4 GT",
	"Xiaomi POCO F4 GT",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5s",
	"Xiaomi POCO X5",
	"Xiaomi POCO X5",
	"Xiaomi POCO X5",
	"Xiaomi POCO X5 Pro",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10 2022",
	"Xiaomi Redmi 10 2022",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10A",
	"Xiaomi Redmi 10A",
	"Xiaomi Redmi 10A",
	"Xiaomi Redmi 10A",
	"Xiaomi Redmi 10A",
	"Xiaomi Redmi 10A",
	"Xiaomi Redmi 10A",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi A1",
	"Xiaomi Redmi A1",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11S",
}

func TestAlternateClean(t *testing.T) {
	for i, name := range alternateNames {
		if _name := shop.AlternateCleanFn(name); _name != alternateNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, alternateNamesExpected[i], name)
		}
	}
}
