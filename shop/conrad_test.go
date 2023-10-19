package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var conradNames = []string{
	"5G Smartphone Nokia G42 5G 128 GB 16.7 cm (6.56 Zoll) Grau Android™ 13 Single-SIM",
	"Apple iPhone 12 (PRODUCT) RED™ 128 GB",
	"Apple iPhone 12 (PRODUCT) RED™ 64 GB",
	"Apple iPhone 12 Blue 128 GB",
	"Apple iPhone 12 Blue 64 GB",
	"Apple iPhone 12 Grün 128 GB",
	"Apple iPhone 12 Schwarz 128 GB",
	"Apple iPhone 12 Schwarz 256 GB",
	"Apple iPhone 12 Schwarz 64 GB",
	"Apple iPhone 12 Weiß 128 GB",
	"Apple iPhone 12 Weiß 256 GB",
	"Apple iPhone 12 Weiß 64 GB",
	"Apple iPhone 13 (PRODUCT) RED™ 128 GB",
	"Apple iPhone 13 Blau 128 GB",
	"Apple iPhone 13 Grün 128 GB",
	"Apple iPhone 13 Grün 256 GB",
	"Apple iPhone 13 Grün 512 GB",
	"Apple iPhone 13 iPhone 256 GB (PRODUCT) RED™",
	"Apple iPhone 13 iPhone 256 GB Blau",
	"Apple iPhone 13 iPhone 256 GB Rose",
	"Apple iPhone 13 iPhone 256 GB Weiß",
	"Apple iPhone 13 mini Grün 128 GB",
	"Apple iPhone 13 mini Grün 512 GB",
	"Apple iPhone 13 Mini iPhone  Schwarz",
	"Apple iPhone 13 Mini iPhone 128 GB (PRODUCT) RED™",
	"Apple iPhone 13 Mini iPhone 128 GB Blau",
	"Apple iPhone 13 Mini iPhone 128 GB Rose",
	"Apple iPhone 13 Mini iPhone 256 GB (PRODUCT) RED™",
	"Apple iPhone 13 Mini iPhone 256 GB Blau",
	"Apple iPhone 13 Mini iPhone 256 GB Rose",
	"Apple iPhone 13 Mini iPhone 256 GB Weiß",
	"Apple iPhone 13 Mini iPhone 512 GB (PRODUCT) RED™",
	"Apple iPhone 13 Mini iPhone 512 GB Rose",
	"Apple iPhone 13 Mini iPhone 512 GB Weiß",
	"Apple iPhone 13 Mini Mitternacht 128 GB",
	"Apple iPhone 13 Mini Mitternacht 256 GB",
	"Apple iPhone 13 Mini Polarstern 128 GB",
	"Apple iPhone 13 Mitternacht 128 GB",
	"Apple iPhone 13 Mitternacht 256 GB",
	"Apple iPhone 13 Polarstern 128 GB",
	"Apple iPhone 13 Polarstern 256 GB",
	"Apple iPhone 13 Rose 128 GB",
	"Apple iPhone 14 (PRODUCT) RED™ 128 GB",
	"Apple iPhone 14 Blau 128 GB",
	"Apple iPhone 14 Gelb 128 GB",
	"Apple iPhone 14 Mitternacht 128 GB",
	"Apple iPhone 14 Mitternacht 256 GB",
	"Apple iPhone 14 Plus (PRODUCT) RED™ 128 GB",
	"Apple iPhone 14 Plus Blau 128 GB",
	"Apple iPhone 14 Plus Gelb 128 GB",
	"Apple iPhone 14 Plus Mitternacht 128 GB",
	"Apple iPhone 14 Plus Polarstern 128 GB",
	"Apple iPhone 14 Plus Polarstern 256 GB",
	"Apple iPhone 14 Plus Violett 128 GB",
	"Apple iPhone 14 Polarstern 128 GB",
	"Apple iPhone 14 Pro Dunkellila 128 GB",
	"Apple iPhone 14 Pro Silber 128 GB",
	"Apple iPhone 14 Violett 128 GB",
	"Apple iPhone SE 128GB Midnight Mitternacht 128 GB",
	"Apple iPhone SE 128GB Starlight Polarstern 128 GB",
	"Apple iPhone SE 256GB Midnight Mitternacht 256 GB",
	"Apple iPhone SE 256GB Starlight Polarstern 256 GB",
	"Apple iPhone SE 64GB Midnight Mitternacht 64 GB",
	"Apple iPhone SE 64GB Starlight Polarstern 64 GB",
	"Apple iPhone SE Rot 128 GB",
	"Apple iPhone SE Rot 256 GB",
	"Apple iPhone SE Rot 64 GB",
	"Apple refurbished iPhone SE Refurbished (sehr gut) 64 GB 4.7 Zoll (11.9 cm) iOS 16 12 Megapixel Schwarz",
	"Asus ROG Phone 6 Smartphone 256 GB Schwarz",
	"Asus ROG Phone 6 Smartphone 256 GB Weiß",
	"Asus ROG Phone 6 Smartphone 512 GB Weiß",
	"Asus ROG Phone 6D 5G Smartphone 256 GB Space Grau",
	"Asus Zenfone 8 Smartphone 128 GB Schwarz",
	"Asus Zenfone 8 Smartphone 256 GB Schwarz",
	"Asus Zenfone 9 5G Smartphone 128 GB Blau",
	"Asus Zenfone 9 5G Smartphone 128 GB Weiß",
	"Asus Zenfone 9 5G Smartphone 256 GB Schwarz",
	"Asus Zenfone 9 5G Smartphone 256 GB Weiß",
	"beafon M6s Senioren-Smartphone 32 GB Schwarz",
	"beafon MX1-EU001B Outdoor Smartphone 128 GB Schwarz",
	"Blackview BV7200 Outdoor Smartphone 128 GB Schwarz",
	"CAT S42 H+ (Version 2022) Outdoor Smartphone 32 GB Schwarz",
	"CAT S42 H+ Outdoor Smartphone 32 GB Schwarz",
	"CAT S53 5G Smartphone 128 GB Schwarz",
	"CAT S62 Pro (Version 2022) Smartphone 128 GB Schwarz",
	"CAT S75 Satellite 5G Smartphone 128 GB Schwarz",
	"Cyrus CS22XA Outdoor Smartphone 16 GB Schwarz",
	"Cyrus CS45XA Outdoor Smartphone 64 GB Schwarz",
	"doro 8050 Plus Senioren-Smartphone  Grau (transparent) ",
	"doro 8050 Senioren-Smartphone 16 GB Graphit",
	"Emporia SMART.4 Smartphone 32 GB Schwarz",
	"Emporia SMART.5 Senioren-Smartphone 32 GB Schwarz",
	"Emporia SUPEREASY Senioren-Smartphone 32 GB Schwarz/Silber",
	"Fairphone 3+ Smartphone 64 GB Schwarz",
	"Gigaset Gigaset GS5 senior Smartphone 64 GB Schwarz",
	"Gigaset Gigaset GX6, Titanium Black Outdoor Smartphone 128 GB Schwarz",
	"Gigaset GS5 LITE Smartphone 64 GB Dunkelgrau",
	"Gigaset GS5 LITE Smartphone 64 GB Perlweiß",
	"Gigaset GX4 Outdoor Smartphone 64 GB Petrol",
	"Gigaset GX4 Outdoor Smartphone 64 GB Schwarz",
	"Gigaset GX6 Outdoor Smartphone 128 GB Grau",
	"Google Pixel 6 Pro Smartphone 128 GB Schwarz",
	"Google Pixel 6 Smartphone 128 GB Schwarz",
	"Google Pixel 6a 5G Smartphone 128 GB Charcoal ",
	"Google Pixel 7 5G Smartphone 128 GB Schwarz ",
	"Google Pixel 7 Pro 5G Smartphone 128 GB Schwarz ",
	"i.safe MOBILE IS120.2 Ex-geschütztes Handy 16 GB Schwarz ",
	"i.safe MOBILE IS655.RG Industrie Smartphone 32 GB ",
	"iPhone 8 Refurbished (sehr gut) 64 GB 4.7 Zoll (11.9 cm) iOS 11 12 Megapixel Spacegrau",
	"Motorola Edge20 5G Smartphone 128 GB 17 cm (6.7 Zoll) Schwarz Android™ 11 Dual-SIM",
	"Motorola Edge 30 Fusion Holiday Edition 5G Smartphone 128 GB Rot",
	"Motorola Edge 30 Fusion Smartphone 128 GB Grau",
	"Motorola Edge 30 Neo Smartphone 128 GB Schwarz",
	"Motorola Edge 30 Neo Smartphone 128 GB Violett",
	"Motorola Edge 30 Ultra Smartphone 256 GB Schwarz",
	"Motorola Edge20 5G Smartphone 128 GB Schwarz",
	"Motorola G31 Smartphone 64 GB Grau",
	"Motorola G82 5G Smartphone 128 GB Grau",
	"Motorola moto e22 Smartphone 32 GB Schwarz",
	"Motorola Moto E40 Smartphone 64 GB Dunkelgrau",
	"Motorola moto g22 Smartphone 64 GB Eisblau",
	"Motorola moto g22 Smartphone 64 GB Schwarz",
	"Motorola Moto G31 Smartphone 64 GB Blau",
	"Motorola moto G42 Smartphone 64 GB Grün",
	"Motorola moto G42 Smartphone 64 GB Metallic, Rose",
	"Motorola Moto G72 Smartphone 128 GB Blau",
	"Motorola Moto G72 Smartphone 128 GB Schwarz",
	"Nokia C21 Plus, 32GB Smartphone 32 GB Blau",
	"Nokia C21 Plus, 32GB Smartphone 32 GB Mitternacht",
	"Nokia G11 Plus Smartphone 32 Grau",
	"Nokia G11 Smartphone 32 GB Charcoal",
	"Nokia G11 Smartphone 32 GB Ice",
	"Nokia G21, 4+64 Smartphone 64 GB Blue",
	"Nokia G21, 4+64 Smartphone 64 GB Purple",
	"Nokia G22 Smartphone 64 GB Blau",
	"Nokia G22 Smartphone 64 GB Grau",
	"Nokia G50 5G Smartphone 128 GB Ocean Blue",
	"Nokia G60 5G Smartphone 128 GB Grau",
	"Nokia G60 5G Smartphone 128 GB Schwarz",
	"Nokia X10 Smartphone 128 GB Grün",
	"Nokia X30 5G Smartphone 128 GB Blau",
	"Nokia X30 5G Smartphone 128 GB Weiß",
	"Nothing Phone (1) 5G Smartphone 256 GB Schwarz",
	"Nothing Phone (1) 5G Smartphone 256 GB Weiß",
	"OPPO A96 Smartphone 128 GB Blau",
	"OPPO Reno8 Lite Smartphone 128 GB Mehrfarbig",
	"OPPO Reno8 Lite Smartphone 128 GB Schwarz",
	"OPPO Reno8 Pro Smartphone 256 GB Hellgrün",
	"OPPO Reno8 Pro Smartphone 256 GB Schwarz",
	"OPPO Reno8 Smartphone 256 GB Gold",
	"OPPO Reno8 Smartphone 256 GB Schwarz",
	"Realme 8i Smartphone 128 GB Lila",
	"Realme 8i Smartphone 64 GB Lila",
	"Realme 9 5G 5G Smartphone 64 GB Weiß",
	"Realme C21 Smartphone 64 GB Blau",
	"Realme C31 Smartphone 32 GB Silber",
	"Realme C35 Smartphone 128 GB Schwarz",
	"Realme GT 2 Smartphone 128 GB Grün",
	"Renewd® iPhone 8 Refurbished (sehr gut) 64 GB 4.7 Zoll (11.9 cm) iOS 11 12 Megapixel Spacegrau",
	"Renewd® iPhone SE (2020) Renewd® (Grade A) 64 GB 4.7 Zoll (11.9 cm) iOS 14 12 Megapixel Schwarz",
	"Renewd® iPhone SE (2. Generation) Renewd® (Grade A) 64 GB 4.7 Zoll (11.9 cm) iOS 14 12 Megapixel Schwarz",
	"Samsung Galaxy A04s Smartphone 32 GB Schwarz",
	"Samsung Galaxy A23 5G Smartphone 64 GB Hellblau",
	"Samsung Galaxy A23 5G Smartphone 64 GB Schwarz",
	"Samsung Galaxy A23 5G Smartphone 64 GB Weiß",
	"Samsung Galaxy A33 5G 5G Smartphone 128 GB Schwarz",
	"Samsung Galaxy A33 5G Enterprise Edition 5G Smartphone 128 GB Schwarz",
	"Samsung Galaxy A33 5G Smartphone 128 GB Hellblau",
	"Samsung Galaxy A33 5G Smartphone 128 GB Weiß",
	"Samsung Galaxy A33 EU 5G Smartphone 128 GB 16.3 cm (6.4 Zoll) Schwarz Android™ 12 Hybrid-Slot",
	"Samsung Galaxy A52s 5G (A528B) 5G Smartphone 128 GB Schwarz",
	"Samsung Galaxy A53 5G Enterprise Edition 5G Smartphone 128 GB Schwarz",
	"Samsung Galaxy A53 5G Smartphone 128 GB Hellblau",
	"Samsung Galaxy A53 5G Smartphone 128 GB Schwarz",
	"Samsung Galaxy A53 5G Smartphone 128 GB Weiß",
	"Samsung Galaxy S20 FE 5G 5G Smartphone 128 GB 6.5 Zoll (16.5 cm) Blau",
	"Samsung Galaxy S21 5G Enterprise Edition 5G Smartphone 128 GB Grau",
	"Samsung Galaxy S21 FE 5G 5G Smartphone 128 GB Graphite",
	"Samsung Galaxy S22 5G Smartphone 128 GB Grün",
	"Samsung Galaxy S22 5G Smartphone 128 GB Roségold",
	"Samsung Galaxy S22 5G Smartphone 128 GB Schwarz",
	"Samsung Galaxy S22 5G Smartphone 128 GB Weiß",
	"Samsung Galaxy S22 5G Smartphone 256 GB Weiß",
	"Samsung Galaxy S22 Enterprise Edition 5G Smartphone 128 GB Schwarz",
	"Samsung Galaxy S22 Ultra 5G Smartphone 128 GB Schwarz",
	"Samsung Galaxy S22+ 5G Smartphone 128 GB Weiß",
	"Samsung Galaxy S23 256 GB CH Cream 5G Smartphone 256 GB Cream",
	"Samsung Galaxy S23 256 GB CH Green Smartphone 256 GB Grün",
	"Samsung Galaxy S23 256 GB CH Phantom Black Smartphone 256 GB Schwarz",
	"Samsung Galaxy Xcover FieldPro Smartphone 64 GB Schwarz",
	"Samsung Galaxy XCover Pro Enterprise Edition Smartphone 64 GB Schwarz",
	"Samsung Galaxy Xcover6 Pro Enterprise Edition Smartphone 128 GB Schwarz",
	"Samsung Galaxy Z Flip4 5G Enterprise Edition 5G Smartphone 128 GB Graphit",
	"Samsung XCover 5 Enterprise Edition Outdoor Smartphone 64 GB Schwarz",
	"Samsung XCover 5 Enterprise Edition Outdoor Smartphone 64 GB 13.5 cm (5.3 Zoll) Schwarz Android™ 11 Dual-SIM",
	"Sony Xperia 5 III 5G Smartphone 128 GB Schwarz",
	"Xiaomi 11T Pro 5G Smartphone 256 GB Blau",
	"Xiaomi 11T Pro 5G Smartphone 256 GB Grau",
	"Xiaomi Redmi 9A Smartphone 32 GB Grau",
	"Xiaomi Redmi 9AT Smartphone 32 GB Gletscherblau",
	"Xiaomi Redmi 9AT Smartphone 32 GB Grau",
	"Xiaomi Redmi Note 10 Pro Smartphone 128 GB Grau",
	"Xiaomi Redmi Note 11 Pro 5G Smartphone 128 GB Graphitgrau",
	"ZTE Blade V30 Smartphone 128 GB Schwarz",
	"ZTE Blade V30 vita + Buds white Smartphone 128 GB Grau",
	"ZTE Blade V30 Vita Smartphone 64 GB Blau",
	"ZTE Blade V40 Vita Smartphone 128 GB Schwarz",
	"ZTE Blade V40S Smartphone 128 GB Schwarz",
}

var conradNamesExpected = []string{
	"Nokia G42",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13 iPhone",
	"Apple iPhone 13 iPhone",
	"Apple iPhone 13 iPhone",
	"Apple iPhone 13 iPhone",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13 mini",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13",
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
	"Apple iPhone 14 Plus",
	"Apple iPhone 14",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"ASUS ROG Phone 6",
	"ASUS ROG Phone 6",
	"ASUS ROG Phone 6",
	"ASUS ROG Phone 6D",
	"ASUS Zenfone 8",
	"ASUS Zenfone 8",
	"ASUS Zenfone 9",
	"ASUS Zenfone 9",
	"ASUS Zenfone 9",
	"ASUS Zenfone 9",
	"beafon M6s",
	"beafon MX1",
	"Blackview BV7200",
	"CAT S42 H+",
	"CAT S42 H+",
	"CAT S53",
	"CAT S62 Pro",
	"CAT S75",
	"Cyrus CS22XA",
	"Cyrus CS45XA",
	"Doro 8050 PLUS",
	"Doro 8050",
	"emporiaSMART.4",
	"emporiaSMART.5",
	"emporiaSUPEREASY",
	"Fairphone 3+",
	"Gigaset GS5",
	"Gigaset GX6",
	"Gigaset GS5 LITE",
	"Gigaset GS5 LITE",
	"Gigaset GX4",
	"Gigaset GX4",
	"Gigaset GX6",
	"Google Pixel 6 Pro",
	"Google Pixel 6",
	"Google Pixel 6a",
	"Google Pixel 7",
	"Google Pixel 7 Pro",
	"i.safe MOBILE IS120.2",
	"i.safe MOBILE IS655.RG",
	"Apple iPhone 8",
	"motorola edge 20",
	"motorola edge 30",
	"motorola edge 30",
	"motorola edge 30 neo",
	"motorola edge 30 neo",
	"motorola edge 30 ultra",
	"motorola edge 20",
	"motorola moto g31",
	"motorola moto g82",
	"motorola moto e22",
	"motorola moto e40",
	"motorola moto g22",
	"motorola moto g22",
	"motorola moto g31",
	"motorola moto g42",
	"motorola moto g42",
	"motorola moto g72",
	"motorola moto g72",
	"Nokia C21 Plus",
	"Nokia C21 Plus",
	"Nokia G11 Plus",
	"Nokia G11",
	"Nokia G11",
	"Nokia G21",
	"Nokia G21",
	"Nokia G22",
	"Nokia G22",
	"Nokia G50",
	"Nokia G60",
	"Nokia G60",
	"Nokia X10",
	"Nokia X30",
	"Nokia X30",
	"Nothing Phone (1)",
	"Nothing Phone (1)",
	"OPPO A96",
	"OPPO Reno8 Lite",
	"OPPO Reno8 Lite",
	"OPPO Reno8 Pro",
	"OPPO Reno8 Pro",
	"OPPO Reno8",
	"OPPO Reno8",
	"realme 8i",
	"realme 8i",
	"realme 9",
	"realme C21",
	"realme C31",
	"realme C35",
	"realme GT 2",
	"Apple iPhone 8",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE (2020)",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A52s",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22 Ultra",
	"Samsung Galaxy S22+",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy XCover Field Pro",
	"Samsung Galaxy XCover Pro",
	"Samsung Galaxy XCover 6 Pro",
	"Samsung Galaxy Z Flip 4",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover 5",
	"Sony Xperia 5 III",
	"Xiaomi 11T Pro",
	"Xiaomi 11T Pro",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"ZTE Blade V30",
	"ZTE Blade V30 Vita",
	"ZTE Blade V30 Vita",
	"ZTE Blade V40 Vita",
	"ZTE Blade V40s",
}

func TestConradClean(t *testing.T) {
	for i, name := range conradNames {
		if _name := shop.ConradCleanFn(name); _name != conradNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, conradNamesExpected[i], name)
		}
	}
}
