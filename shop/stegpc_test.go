package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var stegNames = []string{
	"Blackview BV4900 Dual SIM (3/32GB, schwarz)",
	"Blackview BV4900 Dual SIM (3/32GB, schwarz/orange)",
	"Blackview BV6300 Pro Dual SIM (6/128GB, schwarz)",
	"Honor Magic 4 Lite 5G Dual SIM (6/128GB, blau)",
	"Honor Magic 4 Lite 5G Dual SIM (6/128GB, schwarz)",
	"Huawei Nova 9 SE Dual SIM (8/128GB, blau)",
	"Motorola Defy Dual SIM (4/64GB, schwarz)",
	"Motorola E20 Dual SIM (2/32GB, grau)",
	"Motorola Edge 20 Dual SIM (8/128GB, grau)",
	"Motorola G22 Dual SIM (4/64GB, blau)",
	"Motorola G22 Dual SIM (4/64GB, schwarz)",
	"Motorola G41 Dual SIM (6/128GB, schwarz)",
	"Motorola Moto C Plus Dual SIM - 16GB - schwarz",
	"Motorola Moto E22 Dual SIM (3/32GB, schwarz)",
	"Motorola Moto E22i Dual SIM (2/32GB, grau)",
	"Motorola Moto E32s Dual SIM (3/32GB, grau)",
	"Motorola Moto E40 Dual SIM (4/64GB, dunkelgrau)",
	"Motorola Moto G30 Dual SIM (4/128GB, schwarz)",
	"Motorola Moto G31 Dual SIM (4/128GB, grau)",
	"Motorola Moto G31 Dual SIM (4/64GB, blau)",
	"Motorola Moto G31 Dual SIM (4/64GB, grau)",
	"Motorola Moto G32 Dual SIM (4/128GB, schwarz)",
	"Motorola Moto G42 Dual SIM (4/64GB, grün)",
	"Motorola Moto G42 Dual SIM (4/64GB, pink)",
	"Motorola Moto G52 Dual SIM (4/128GB, grau)",
	"Motorola Moto G72 Dual SIM (6/128GB, blau)",
	"Motorola Moto G72 Dual SIM (6/128GB, grau)",
	"Motorola Moto G82 5G Dual SIM (6/128GB, weiss)",
	"Nokia C21 Dual SIM (2/32GB, blau)",
	"Nokia C21 Plus Dual SIM (2/32GB, blau)",
	"Nokia G11 Dual SIM (3/32GB, blau)",
	"Nokia G11 Dual SIM (3/32GB, braun)",
	"Nokia G21 Dual SIM (4/128GB, blau)",
	"Nokia G21 Dual SIM (4/128GB, braun)",
	"Nokia G21 Dual SIM (4/64GB, blau)",
	"Nokia G21 Dual SIM (4/64GB, braun)",
	"Nokia G22 Dual SIM (4/64GB, blau)",
	"Nokia G22 Dual SIM (4/64GB, grau)",
	"Nokia G50 Dual SIM (4/128GB, blau)",
	"Nokia G50 Dual SIM (4/128GB, pink/gold)",
	"Nokia G60 5G Dual SIM (4/128GB, grau)",
	"Nokia G60 5G Dual SIM (4/128GB, schwarz)",
	"Nokia X10 Dual SIM (4/128, grün)",
	"Oppo A54s Dual SIM (4/128GB, schwarz)",
	"Oppo A57s Dual SIM (4/128GB, blau)",
	"Oppo A57s Dual SIM (4/128GB, schwarz)",
	"Oppo A96 Dual SIM (8/128GB, blau)",
	"Oppo A96 Dual SIM (8/128GB, schwarz)",
	"Realme 9 5G Dual SIM (4/128GB, schwarz)",
	"Realme 9 5G Dual SIM (4/128GB, weiss)",
	"Realme 9 Dual SIM (8/128GB, gold)",
	"Realme 9 Dual SIM (8/128GB, schwarz)",
	"Realme 9 Dual SIM (8/128GB, weiss)",
	"Realme 9i Dual SIM (4/128GB, blau)",
	"Realme 9i Dual SIM (4/128GB, schwarz)",
	"Realme C25Y Dual SIM (4/128GB, blau)",
	"Realme C25Y Dual SIM (4/128GB, grau)",
	"Realme C30 Dual SIM (3/32GB, blau)",
	"Realme C30 Dual SIM (3/32GB, grün)",
	"Realme C30 Dual SIM (3/32GB, schwarz)",
	"Realme C31 (3/32GB, argenté)",
	"Realme C31 (4/64GB, grün)",
	"Realme C31 (4/64GB, silber)",
	"Realme C33 Dual SIM (4/64GB, dunkelblau)",
	"Realme C33 Dual SIM (4/64GB, gold)",
	"Realme C35 Dual SIM (4/128GB, schwarz)",
	"Renewd Apple iPhone 8 (2/64GB, gold)",
	"Renewd Apple iPhone 8 (2/64GB, grau)",
	"Renewd Apple iPhone 8 Plus (3/64GB, grau)",
	"Renewd Apple iPhone X (3/64GB, silber)",
	"Renewd iPhone 8 Plus (3/64GB, gold)",
	"Renewd iPhone 8 Plus (3/64GB, silber)",
	"Samsung Galaxy A04s A047 Dual SIM (3/32GB, grün)",
	"Samsung Galaxy A04s A047 Dual SIM (3/32GB, schwarz)",
	"Samsung Galaxy A04s A047 Dual SIM (3/32GB, weiss)",
	"Samsung Galaxy A12 A127F Dual SIM (3/32GB, schwarz) - EU Modell",
	"Samsung Galaxy A13 128 GB CH Blue",
	"Samsung Galaxy A13 5G A136 Dual SIM (4/128GB, blau)",
	"Samsung Galaxy A13 5G A136 Dual SIM (4/128GB, schwarz)",
	"Samsung Galaxy A13 5G A136 Dual SIM (4/128GB, weiss)",
	"Samsung Galaxy A13 5G A136 Dual SIM (4/64GB, blau)",
	"Samsung Galaxy A13 5G A136 Dual SIM (4/64GB, schwarz)",
	"Samsung Galaxy A13 5G A136 Dual SIM (4/64GB, weiss)",
	"Samsung Galaxy A13 A135F Dual SIM (4/128GB, schwarz) - EU Modell",
	"Samsung Galaxy A13 A137F Dual SIM (4/128GB, blau)",
	"Samsung Galaxy A13 A137F Dual SIM (4/128GB, schwarz)",
	"Samsung Galaxy A13 A137F Dual SIM (4/128GB, weiss)",
	"Samsung Galaxy A13 A137F Dual SIM (4/64GB, blau)",
	"Samsung Galaxy A13 A137F Dual SIM (4/64GB, schwarz)",
	"Samsung Galaxy A13 A137F Dual SIM (4/64GB, weiss)",
	"Samsung Galaxy A13 Dual SIM (3/32GB, blau)",
	"Samsung Galaxy A13 Dual SIM (3/32GB, schwarz)",
	"Samsung Galaxy A13 Dual SIM (3/32GB, weiss)",
	"Samsung Galaxy A13 Dual SIM (4/128GB, blau)",
	"Samsung Galaxy A13 Dual SIM (4/128GB, weiss)",
	"Samsung Galaxy A13 Dual SIM (4/64GB, blau)",
	"Samsung Galaxy A13 Dual SIM (4/64GB, schwarz)",
	"Samsung Galaxy A13 Dual SIM (4/64GB, weiss)",
	"Samsung Galaxy A22 5G Dual SIM (4/64GB, grau)",
	"Samsung Galaxy A22 5G Dual SIM (4/64GB, weiss)",
	"Samsung Galaxy A23 5G A236 Dual SIM (4/128GB, blau)",
	"Samsung Galaxy A23 5G A236 Dual SIM (4/128GB, schwarz)",
	"Samsung Galaxy A23 5G A236 Dual SIM (4/128GB, weiss)",
	"Samsung Galaxy A23 5G A236 Dual SIM (4/64GB, blau)",
	"Samsung Galaxy A23 5G A236 Dual SIM (4/64GB, schwarz)",
	"Samsung Galaxy A23 5G A236 Dual SIM (4/64GB, weiss)",
	"Samsung Galaxy A32 Dual SIM (4/128GB, schwarz)",
	"Samsung Galaxy A32 Dual SIM Enterprise Edition (4/128GB, schwarz)",
	"Samsung Galaxy A33 5G Dual SIM (6/128GB, blau)",
	"Samsung Galaxy A33 5G Dual SIM (6/128GB, schwarz)",
	"Samsung Galaxy A33 5G Dual SIM (6/128GB, weiss)",
	"Samsung Galaxy A33 5G Dual SIM Enterprise Edition (6/128GB, schwarz)",
	"Samsung Galaxy M23 5G Dual SIM (4/128GB, blau)",
	"Samsung Galaxy M23 5G Dual SIM (4/128GB, grün)",
	"Samsung Galaxy M33 M336 5G Dual SIM (6/128GB, blau)",
	"Samsung Galaxy M33 M336 5G Dual SIM (6/128GB, braun)",
	"Samsung Galaxy M33 M336 5G Dual SIM (6/128GB, grün)",
	"Samsung Galaxy Xcover 5 Dual SIM (4/64GB, schwarz)",
	"Samsung Galaxy Xcover 5 Dual SIM Enterprise Edition (4/64GB, schwarz)",
	"Vivo Y01 Dual SIM (2/32GB, schwarz)",
	"Vivo Y21 Dual SIM (4/64GB, blau)",
	"Vivo Y21 Dual SIM (4/64GB, weiss)",
	"Vivo Y21s Dual SIM (4/128GB, blau)",
	"Vivo Y21s Dual SIM (4/128GB, violett)",
	"Vivo Y33s Dual SIM (8/128GB, blau)",
	"Xiaomi Poco M5 Dual SIM (4/128GB, schwarz)",
	"Xiaomi Poco M5 Dual SIM (4/64GB, schwarz)",
	"Xiaomi Poco M5s Dual SIM (4/128GB, weiss)",
	"Xiaomi Poco M5s Dual SIM (4/64GB, weiss)",
	"Xiaomi Poco X5 5G Dual SIM (6/128GB, blau)",
	"Xiaomi Poco X5 5G Dual SIM (6/128GB, schwarz)",
	"Xiaomi Redmi 10 2022 Dual SIM (4/128GB, grau)",
	"Xiaomi Redmi 10 2022 Dual SIM (4/64GB, blau)",
	"Xiaomi Redmi 10 2022 Dual SIM (4/64GB, grau)",
	"Xiaomi Redmi 10 2022 Dual SIM (4/64GB, weiss)",
	"Xiaomi Redmi 10C Dual SIM (4/128GB, blau)",
	"Xiaomi Redmi 10C Dual SIM (4/128GB, grau)",
	"Xiaomi Redmi 10C Dual SIM (4/128GB, grün)",
	"Xiaomi Redmi 10C Dual SIM (4/64GB, grau)",
	"Xiaomi Redmi 9A Dual SIM (2/32GB, blau)",
	"Xiaomi Redmi 9A Dual SIM (2/32GB, grau)",
	"Xiaomi Redmi 9A Dual SIM (2/32GB, grün)",
	"Xiaomi Redmi 9AT Dual SIM (2/32GB, blau)",
	"Xiaomi Redmi 9AT Dual SIM (2/32GB, grau)",
	"Xiaomi Redmi 9C Dual SIM (4/128GB, grau)",
	"Xiaomi Redmi Note 10 Pro Dual SIM (6/128GB, blau)",
	"Xiaomi Redmi Note 10 Pro Dual SIM (6/128GB, bronze)",
	"Xiaomi Redmi Note 10 Pro Dual SIM (6/128GB, grau)",
	"Xiaomi Redmi Note 10 Pro Dual SIM (6/64GB, blau)",
	"Xiaomi Redmi Note 10 Pro Dual SIM (6/64GB, bronze)",
	"Xiaomi Redmi Note 10 Pro Dual SIM (6/64GB, grau)",
	"Xiaomi Redmi Note 10S Dual SIM (6/128GB, blau) - Retoure",
	"Xiaomi Redmi Note 11 Dual SIM (4/128GB, blau)",
	"Xiaomi Redmi Note 11 Dual SIM (4/128GB, grau)",
	"Xiaomi Redmi Note 11 Dual SIM (4/128GB, hellblau)",
	"Xiaomi Redmi Note 11 Dual SIM (4/64GB, blau)",
	"Xiaomi Redmi Note 11 Dual SIM (4/64GB, grau)",
	"Xiaomi Redmi Note 11 Dual SIM (4/64GB, hellblau)",
	"Xiaomi Redmi Note 11 Pro 5G Dual SIM (6/128GB, grau)",
	"Xiaomi Redmi Note 11 Pro 5G Dual SIM (6/128GB, weiss)",
	"Xiaomi Redmi Note 11 Pro Dual SIM (6/128GB, blau)",
	"Xiaomi Redmi Note 11 Pro Dual SIM (6/128GB, grau)",
	"Xiaomi Redmi Note 11S Dual SIM (6/128GB, blau)",
	"Xiaomi Redmi Note 11S Dual SIM (6/128GB, grau)",
	"Xiaomi Redmi Note 11S Dual SIM (6/64GB, blau)",
	"Xiaomi Redmi Note 11S Dual SIM (6/64GB, grau)",
	"ZTE Axon 11 5G Dual SIM (6/128GB, schwarz)",
	"ZTE Blade A31 Dual SIM (2/32GB, grau)",
	"ZTE Blade A31 Lite Dual SIM (1/32GB, grau)",
	"ZTE Blade A52 Dual SIM (2/64GB, blau)",
	"ZTE Blade A52 Dual SIM (2/64GB, grau)",
	"ZTE Blade A72 5G Dual SIM (4/64GB, blau)",
	"ZTE Blade A72 5G Dual SIM (4/64GB, grau)",
	"ZTE Blade A72 Dual SIM (4/64GB, blau)",
	"ZTE Blade A72 Dual SIM (4/64GB, grau)",
	"ZTE Blade V30 Dual SIM (4/128GB, schwarz)",
	"ZTE Blade V30 Vita Dual SIM (3/128GB, blau)",
	"ZTE Blade V30 Vita Dual SIM (3/128GB, grau)",
	"ZTE Blade V30 Vita Dual SIM (3/128GB, grau) + Buds SE",
	"ZTE Blade V30 Vita Dual SIM (4/64GB, blau)",
	"ZTE Blade V40 Vita Dual SIM (3/128GB, grün)",
	"ZTE Blade V40 Vita Dual SIM (3/128GB, schwarz)",
	"ZTE Blade V40S Dual SIM (4/128GB, blau)",
	"ZTE Blade V40S Dual SIM (4/128GB, schwarz)",
}

var stegNamesExpected = []string{
	"Blackview BV4900",
	"Blackview BV4900",
	"Blackview BV6300 Pro",
	"Honor Magic 4 Lite",
	"Honor Magic 4 Lite",
	"Huawei Nova 9 SE",
	"Motorola Defy",
	"Motorola Moto E20",
	"Motorola Edge 20",
	"Motorola Moto G22",
	"Motorola Moto G22",
	"Motorola Moto G41",
	"Motorola Moto C Plus",
	"Motorola Moto E22",
	"Motorola Moto E22i",
	"Motorola Moto E32s",
	"Motorola Moto E40",
	"Motorola Moto G30",
	"Motorola Moto G31",
	"Motorola Moto G31",
	"Motorola Moto G31",
	"Motorola Moto G32",
	"Motorola Moto G42",
	"Motorola Moto G42",
	"Motorola Moto G52",
	"Motorola Moto G72",
	"Motorola Moto G72",
	"Motorola Moto G82",
	"Nokia C21",
	"Nokia C21 Plus",
	"Nokia G11",
	"Nokia G11",
	"Nokia G21",
	"Nokia G21",
	"Nokia G21",
	"Nokia G21",
	"Nokia G22",
	"Nokia G22",
	"Nokia G50",
	"Nokia G50",
	"Nokia G60",
	"Nokia G60",
	"Nokia X10",
	"Oppo A54s",
	"Oppo A57s",
	"Oppo A57s",
	"Oppo A96",
	"Oppo A96",
	"Realme 9",
	"Realme 9",
	"Realme 9",
	"Realme 9",
	"Realme 9",
	"Realme 9i",
	"Realme 9i",
	"Realme C25Y",
	"Realme C25Y",
	"Realme C30",
	"Realme C30",
	"Realme C30",
	"Realme C31",
	"Realme C31",
	"Realme C31",
	"Realme C33",
	"Realme C33",
	"Realme C35",
	"Apple iPhone 8",
	"Apple iPhone 8",
	"Apple iPhone 8 Plus",
	"Apple iPhone X",
	"Apple iPhone 8 Plus",
	"Apple iPhone 8 Plus",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A12",
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
	"Samsung Galaxy A22",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A32",
	"Samsung Galaxy A32",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy M23",
	"Samsung Galaxy M23",
	"Samsung Galaxy M33",
	"Samsung Galaxy M33",
	"Samsung Galaxy M33",
	"Samsung Galaxy Xcover 5",
	"Samsung Galaxy Xcover 5",
	"Vivo Y01",
	"Vivo Y21",
	"Vivo Y21",
	"Vivo Y21s",
	"Vivo Y21s",
	"Vivo Y33s",
	"Xiaomi Poco M5",
	"Xiaomi Poco M5",
	"Xiaomi Poco M5s",
	"Xiaomi Poco M5s",
	"Xiaomi Poco X5",
	"Xiaomi Poco X5",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 11S",
	"ZTE Axon 11",
	"ZTE Blade A31",
	"ZTE Blade A31 Lite",
	"ZTE Blade A52",
	"ZTE Blade A52",
	"ZTE Blade A72",
	"ZTE Blade A72",
	"ZTE Blade A72",
	"ZTE Blade A72",
	"ZTE Blade V30",
	"ZTE Blade V30 Vita",
	"ZTE Blade V30 Vita",
	"ZTE Blade V30 Vita",
	"ZTE Blade V30 Vita",
	"ZTE Blade V40 Vita",
	"ZTE Blade V40 Vita",
	"ZTE Blade V40S",
	"ZTE Blade V40S",
}

func TestStegClean(t *testing.T) {
	for i, name := range stegNames {
		if _name := shop.StegCleanFn(name); _name != stegNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, stegNamesExpected[i], name)
		}
	}
}
