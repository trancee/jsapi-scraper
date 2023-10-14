package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var folettiNames = []string{
	"Blackview BV4900 Pro 14.5 cm (5.7) Dual SIM Android 10.0 4G Micro-USB 4 GB 64 GB 5580 mAh Black",
	"Blackview BV5200 5180 mAh 4/32 GB Green smartphone",
	"Blackview BV7100 6/128 Black",
	"Blackview BV9300 12/25BGB Green",
	"Gigaset GL390 titan/silber",
	"Gigaset GS5 Lite dark titanium grey",
	"Gigaset GX290 plus titanium grey 4+64GB",
	"Gigaset GX4 Outdoor",
	"Gigaset GX4 petrol",
	"Google Pixel 6a sage",
	"Huawei Nova 10SE 8/128GB Mint Green",
	"Moto e20 32GB, Handy",
	"Moto G42 64GB, Handy",
	"Motorola Mobility Edge 20 128GB, Handy",
	"Motorola edge 20 frosted grey",
	"Motorola edge 20 lite electric graphite",
	"Motorola Edge 30 Neo black onyx 8+128GB",
	"Motorola G42 atlantic green",
	"Motorola G42 metallic rose",
	"Motorola Mobility Motorola XT2173-3 moto g31 Dual Sim 4+64GB mineral grey",
	"Motorola Moto E 13",
	"Motorola Moto E 22",
	"Motorola Moto E e32s 16.5 cm (6.5) Dual SIM Android 12 4G USB Type-C 4 GB 64 GB 5000 mAh Silver",
	"Motorola Moto E13 cosmic black",
	"Motorola Moto E20 grau",
	"Motorola Moto E22 astro black",
	"MOTOROLA MOTO E22 4/64GB ASTRO BLACK",
	"Motorola Moto E32s gravity grau",
	"Motorola Moto G13 lavender blue",
	"Motorola moto G14 sky blue",
	"Motorola moto G14 steel grey",
	"Motorola moto g22",
	"Motorola Moto G23 matte charcoal",
	"Motorola Moto G23 steel blue",
	"Motorola Moto G31 sterling blue",
	"Motorola Moto G41 meteorite black",
	"Motorola Moto G52 charcoal grey",
	"Motorola Moto G72 meteorite grey",
	"Motorola Moto G72 polar blue",
	"Motorola Moto G 42 16.3 cm (6.4) Dual SIM Android 12 4G USB Type-C 4 GB 64 GB 5000 mAh Pink",
	"Motorola Moto G G51 5G 17.3 cm (6.8) Dual SIM Android 11 USB Type-C 4 GB 64 GB 5000 mAh Blue",
	"Motorola Solutions Moto E32S Grey",
	"Motorola XT2173-3 Moto G31 128GB Grey 6.4 EU (4GB) Android",
	"Motorola XT2173-3 moto g31 Dual Sim 4+64GB mineral grey DE",
	"Motorola XT2229-2 moto e32s Dual Sim 3+32GB  slate grey DE",
	"Motorola XT2239-7 moto e22 Dual Sim 3+32GB  astro black DE",
	"Nokia G11 Charcoal, 3+32GB",
	"Nokia G11 Ice, 3+32GB",
	"Nokia G21 Blue, 4+64GB",
	"Nokia G50 blue 4+128GB",
	"Nothing Phone 1 8/256GB Black",
	"Nothing Phone (1) 16.6 cm (6.55) Dual SIM 5G USB Type-C 8 GB 256 GB 4500 mAh White",
	"Nothing Phones Phone (1) 8 GB / 128 GB - Bildschirmdiagonale: 6.55 - Betriebssystem: Android - Detailfarbe: Schwarz - Speicherkapazität total: 128 GB - Verbauter Arbeitsspeicher: 8 GB - Induktionsladung: Ja",
	"OnePlus CPH2399 Nord 2T Dual Sim 8+128GB gray shadow DE",
	"OPPO A76 16.7 cm (6.56) Dual SIM Android 11 4G USB Type-C 4 GB 128 GB 5000 mAh Black",
	"OPPO A76 glowing blue",
	"OPPO Reno 8 Lite 16.3 cm (6.43) Dual SIM Android 11 5G USB Type-C 8 GB 128 GB 4500 mAh Black",
	"Realme 9 meteor black 8+128GB",
	"Realme 9 stargaze white 8+128GB",
	"Realme 9 sunburst gold 8+128GB",
	"Realme 9 Pro Midnight Black 8+128GB",
	"Realme 9i Prism Black 4+128GB",
	"Realme 9i Prism Blue 4+128GB",
	"Realme C30 bamboo green              3+32GB",
	"Realme C30 denim black               3+32GB",
	"Realme C30 lake blue                 3+32GB",
	"Realme C33 night sea 4+64GB",
	"Realme C33 sandy gold 4+64GB",
	"Realme C35 Glowing Black 128GB",
	"Realme C35 Glowing Green 128GB",
	"Redmi 10 (2022) 128GB, Handy",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A04s SM-A047F/DSN",
	"Samsung Galaxy A13 32GB Black",
	"Samsung Galaxy A13 Black 64GB",
	"Samsung Galaxy A13 NE Black 32GB",
	"Samsung Galaxy A13 SM-A137FZWUEUE smartphone",
	"Samsung SM-A137F Galaxy A13 Dual Sim 3+32GB black EU",
	"Samsung SM-M336B/DS 16.8 cm (6.6) Dual SIM 5G USB Type-C 6 GB 128 GB 5000 mAh Blue",
	"Samsung Galaxy A53 ( A536B) 5G 128GB Enterprise Edition Black",
	"Samsung Galaxy SM-M336B/DS 16.8 cm (6.6) Dual SIM 5G USB Type-C 6 GB 128 GB 5000 mAh Brown",
	"Samsung Galaxy Xcover 5 EE 64GB 4RAM 4G DE black",
	"Samsung Galaxy XCover 5 Enterprise Edition CH - Bildschirmdiagonale: 5.3 - Betriebssystem: Android - Detailfarbe: Schwarz - Speicherkapazität total: 64 GB - Verbauter Arbeitsspeicher: 4 GB - Induktionsladung: Nein",
	"Samsung Galaxy Xcover 5 Enterprise 64GB Black 5.3 EU Model Android",
	"Samsung Galaxy Xcover 5 (G525F) 64GB Black",
	"Samsung Galaxy XCover 5 SM-G525F/DS 13.5 cm (5.3) Dual SIM Android 11 4G USB Type-C 4 GB 64 GB 3000 mAh Black",
	"Samsung Galaxy XCover 5 -- Enterprise Edition - 64GB Black",
	"Smartfon OnePlus Nord CE3 Lite 8/128GB 5G Szary",
	"Smartfon Samsung Galaxy A14 (A146P) 5G ds 4/64GB Silver",
	"Smartfon  Xiaomi Redmi 10C 3/64GB Zielony",
	"SMARTPHONE MOTOROLA MOTO G53 5G 4/128 ARCTIC SILVER",
	"TCL 3189 HIMALAYA GREY .",
	"vivo Y01 elegant black",
	"vivo Y21s midnight blue",
	"Xiaomi 23021RAA2Y",
	"Xiaomi Note12 Pro 5G 6/128GB Midnight Black",
	"Xiaomi Note 12 Pro 5G 6/128GB Polar White",
	"Xiaomi Redmi 10 5G aurora green 4GB+64GB",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C 17 cm (6.71) Dual SIM Android 11 4G USB Type-C 4 GB 128 GB 5000 mAh Blue",
	"Xiaomi Redmi 10C Dual Sim 4+128GB graphite grey DE",
	"Xiaomi Redmi 10C Dual Sim 4+128GB mint green DE",
	"Xiaomi Redmi 10C Dual Sim 4+128GB ocean blue DE",
	"Xiaomi Redmi 10C Graphite Gray 4+128GB",
	"Xiaomi Redmi 10C Mint Green 4+128GB",
	"Xiaomi Redmi 10C Ocean Blue 4+128GB",
	"Xiaomi Redmi 12C graphite grey 4GB+128GB",
	"Xiaomi Redmi 12C mint green 4GB+128GB",
	"Xiaomi Redmi 12C ocean blue 4GB+128GB",
	"Xiaomi Redmi 9a 16.6 cm (6.53) Hybrid Dual SIM Android 11 4G Micro-USB 2 GB 32 GB 5000 mAh Grey",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9AT 32GB DS Grey 6.5 EU",
	"Xiaomi Redmi 9AT Dual Sim 2+32GB glacial blue DE",
	"Xiaomi Redmi 9C NFC 16.6 cm (6.53) Hybrid Dual SIM 4G Micro-USB 3 GB 64 GB 5000 mAh Green",
	"Xiaomi Redmi A1",
	"Xiaomi Redmi Note 12 ice blue 4GB+128GB",
	"Xiaomi Redmi Note 12 onyx gray 4GB+128GB",
	"TIM Xiaomi Redmi 10C 17 cm (6.71) Dual SIM Android 11 4G USB Type-C 4 GB 64 GB 5000 mAh Grey",
	"ZTE Blade A31 Lite grau",
	"ZTE Blade V40 Vita inkl. Buds zeus black",
	"ZTE Blade V50 Vita icy blue",
	"ZTE Blade A72 blau",
	"ZTE Blade A72 grau",
}

var folettiNamesExpected = []string{
	"Blackview BV4900 Pro",
	"Blackview BV5200",
	"Blackview BV7100",
	"Blackview BV9300",
	"Gigaset GL390",
	"Gigaset GS5 LITE",
	"Gigaset GX290 Plus",
	"Gigaset GX4",
	"Gigaset GX4",
	"Google Pixel 6a",
	"HUAWEI nova 10 SE",
	"motorola moto e20",
	"motorola moto g42",
	"motorola edge 20",
	"motorola edge 20",
	"motorola edge 20 lite",
	"motorola edge 30 neo",
	"motorola moto g42",
	"motorola moto g42",
	"motorola moto g31",
	"motorola moto e13",
	"motorola moto e22",
	"motorola moto e32s",
	"motorola moto e13",
	"motorola moto e20",
	"motorola moto e22",
	"motorola moto e22",
	"motorola moto e32s",
	"motorola moto g13",
	"motorola moto g14",
	"motorola moto g14",
	"motorola moto g22",
	"motorola moto g23",
	"motorola moto g23",
	"motorola moto g31",
	"motorola moto g41",
	"motorola moto g52",
	"motorola moto g72",
	"motorola moto g72",
	"motorola moto g42",
	"motorola moto g51",
	"motorola moto e32s",
	"motorola moto g31",
	"motorola moto g31",
	"motorola moto e32s",
	"motorola moto e22",
	"Nokia G11",
	"Nokia G11",
	"Nokia G21",
	"Nokia G50",
	"Nothing Phone (1)",
	"Nothing Phone (1)",
	"Nothing Phone (1)",
	"OnePlus Nord 2T",
	"OPPO A76",
	"OPPO A76",
	"OPPO Reno8 Lite",
	"realme 9",
	"realme 9",
	"realme 9",
	"realme 9 Pro",
	"realme 9i",
	"realme 9i",
	"realme C30",
	"realme C30",
	"realme C30",
	"realme C33",
	"realme C33",
	"realme C35",
	"realme C35",
	"Xiaomi Redmi 10",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy M33",
	"Samsung Galaxy A53",
	"Samsung Galaxy M33",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover 5",
	"OnePlus Nord CE 3 Lite",
	"Samsung Galaxy A14",
	"Xiaomi Redmi 10C",
	"motorola moto g53",
	"TCL 3189",
	"vivo Y01",
	"vivo Y21s",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 12C",
	"Xiaomi Redmi 12C",
	"Xiaomi Redmi 12C",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi A1",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi 10C",
	"ZTE Blade A31 Lite",
	"ZTE Blade V40 Vita",
	"ZTE Blade V50 Vita",
	"ZTE Blade A72",
	"ZTE Blade A72",
}

func TestFolettiClean(t *testing.T) {
	for i, name := range folettiNames {
		if _name := shop.FolettiCleanFn(name); _name != folettiNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, folettiNamesExpected[i], name)
		}
	}
}
