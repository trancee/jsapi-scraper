package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var galaxusNames = []string{
	"Apple iPhone 11 Pro",
	"Apple iPhone 6",
	"Apple iPhone 7",
	"Apple iPhone 8",
	"Apple iPhone SE (2nd Gen)",
	"Apple iPhone XS",
	"Blackview A100 DUAL SIM 128GB Blue",
	"Blackview A55GREEN Phone A55 / Green",
	"Blackview BL5000 5G 8",
	"Blackview BV4900 Pro (5.7 Zoll) Dual-SIM Android 10.0 Mikro-USB",
	"Blackview BV5200 Pro 5180 mAh 4/64 GB Schwarz (schwarz) Smartphone",
	"Blackview Smartphone Blackview Smartphone BV9900E 6/128 Silver",
	"Fairphone 2",
	"Google Pixel 3a",
	"Google Pixel 4a",
	"Google Pixel 4a 5G",
	"Google Pixel 6a",
	"Honor 70",
	"Honor 8",
	"Honor Magic 4 Lite 5G",
	"HTC Desire 22 Pro",
	"Huawei Mate S Force Touch",
	"Huawei Nova 10 SE",
	"Huawei Nova 8I",
	"Huawei nova 9",
	"Huawei Nova 9 SE",
	"Huawei Nova Y70",
	"Huawei Nova Y90",
	"Huawei P40 Lite",
	"Huawei P8",
	"Huawei Y5p",
	"Infinix Hot 11 X689F",
	"Infinix HOT 11S 17 2 cm (6.78 Zoll) Dual-SIM Android 11 4G USB Typ-C 6 GB 128 GB 5000 mAh Schwarz",
	"Motorola Defy",
	"Motorola e22",
	"Motorola E22i",
	"Motorola Edge 20",
	"Motorola Edge 30",
	"Motorola G42",
	"Motorola Google Nexus 6",
	"Motorola Moto E32",
	"Motorola Moto E32s Dual SIM, 4GB RAM, 64GB, Misty Silver",
	"Motorola Moto Edge 30 Neo",
	"Motorola Moto Edge30 Neo",
	"Motorola Moto G 5G",
	"Motorola Moto G 5G Plus",
	"Motorola Moto G10",
	"Motorola Moto G200",
	"Motorola moto g22 16.5 cm (6.5\") Dual SIM Android 12 4G USB Type-C 4 GB 64 GB 5000 mAh Juodas",
	"Motorola Moto G30",
	"Motorola Moto G31 4",
	"Motorola Moto G41 (4GB)",
	"Motorola Moto G42 4",
	"Motorola Moto G5 Plus",
	"Motorola Moto G52",
	"Motorola Moto G5s Plus",
	"Motorola Moto G6 Play",
	"Motorola Moto G62 5G",
	"Motorola XT2143-1 Edge 20 Grey EU Android",
	"Nokia 3.4",
	"Nokia 7.1",
	"Nokia G11",
	"Nokia G20",
	"Nokia G21",
	"Nokia G50",
	"Nokia Lumia 930",
	"Nokia Smartphone G11 DS 3",
	"Nokia X20 (6GB)",
	"Nokia XR20",
	"OnePlus 3T",
	"OnePlus 5T",
	"OnePlus 6",
	"OnePlus 6T (8GB)",
	"OnePlus 7T (8GB)",
	"OnePlus 8T (12GB)",
	"OnePlus Nord 2",
	"OnePlus Nord 2 (12GB)",
	"OnePlus Nord 2 (8GB)",
	"OnePlus Nord 2T",
	"OnePlus Nord CE",
	"OnePlus Nord CE 2 Lite 5G",
	"OPPO A57s",
	"OPPO A74",
	"OPPO Find X2 Lite",
	"OPPO Find X3 Lite",
	"OPPO Find X3 Neo",
	"OPPO Find X5 Lite",
	"OPPO Reno 7 Lite",
	"OPPO Reno 8 Lite",
	"OPPO Reno4 Pro",
	"OPPO Reno4 Z",
	"OPPO Reno6",
	"realme 7 Pro",
	"realme C30 bamboo green 3+32GB",
	"realme GT Master",
	"realme Narzo 50 4g 128 GB Speed Blue",
	"Rephone Rephone",
	"Samsung Galaxy A12 (2021)",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13 EU",
	"Samsung Galaxy A14 LTE",
	"Samsung Galaxy A32",
	"Samsung Galaxy A32 5G",
	"Samsung Galaxy A32 EU",
	"Samsung Galaxy A33 5G",
	"Samsung Galaxy A33 5G EU",
	"Samsung Galaxy A52",
	"Samsung Galaxy A52s",
	"Samsung Galaxy A52s EU",
	"Samsung Galaxy A53 5G",
	"Samsung Galaxy A53 5G EU",
	"Samsung Galaxy A72",
	"Samsung Galaxy Note 20 Ultra Snapdragon",
	"Samsung Galaxy Note 4",
	"Samsung Galaxy Note 9",
	"Samsung Galaxy S20 FE 5G",
	"Samsung Galaxy XCover 4s",
	"Samsung Galaxy XCover 5 EE Enterprise Edition",
	"Samsung Galaxy XCover Pro EE Enterprise Edition",
	"Sony Xperia 1",
	"Sony Xperia 10 III",
	"Sony Xperia XZ",
	"Sony Xperia Z3 Compact, Schwarz",
	"Sony Xperia Z3+",
	"Sony Xperia Z5 Compact - The Bond Phone",
	"Sony Xperia Z5 Premium",
	"TCL 30E (6.52 Zoll) Dual-SIM Android 12 USB Typ-C",
	"Vivo NEX S",
	"Vivo X60 Pro 5G",
	"Vivo Y01 Elegant Black",
	"Vivo Y21s midday dream",
	"Vivo Y21s midnight blue",
	"Vivo Y72 5G (V2041) dream glow",
	"Xiaomi 11 Lite 5G NE",
	"Xiaomi 11T",
	"Xiaomi 11T 5G",
	"Xiaomi 11T Pro",
	"Xiaomi Mi 10 Lite 5G",
	"Xiaomi Mi 10T Pro",
	"Xiaomi Mi 11 Lite",
	"Xiaomi Mi 11 Lite 5G (8GB)",
	"Xiaomi Mi 9",
	"Xiaomi M5s",
	"Xiaomi Poco F3",
	"Xiaomi Poco M3",
	"Xiaomi POCO M4 5G",
	"Xiaomi Poco X3",
	"Xiaomi Poco X4 Pro",
	"Xiaomi Poco X4 Pro 5G",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10 Dual-Sim (2022) 64GB, Sea Blue",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 9",
	"Xiaomi Redmi 9C NFC",
	"Xiaomi Redmi A1 Plus",
	"Xiaomi Redmi A1+ Dual Sim 32GB, 2GB RAM, Blue",
	"Xiaomi Redmi A2 2/32 GB blue EU",
	"Xiaomi Redmi Note 10",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro 5G (6GB)",
	"Xiaomi Redmi Note 11 Pro+ 5G",
	"Xiaomi Redmi Note 11s 5G Dual Sim 4GB RAM 128GB - Twilight Blue EU",
	"Xiaomi Redmi Note 8",
	"Xiaomi Redmi Note 8 Pro",
	"Xiaomi Redmi Note 9",
	"Xiaomi Redmi Note 9 Pro",
	"ZTE Blade A31",
	"ZTE Blade A31 Lite",
	"ZTE Blade V40S 4G Black",
}

var galaxusNamesExpected = []string{
	"Apple iPhone 11 Pro",
	"Apple iPhone 6",
	"Apple iPhone 7",
	"Apple iPhone 8",
	"Apple iPhone SE (2nd Gen)",
	"Apple iPhone XS",
	"Blackview A100",
	"Blackview A55",
	"Blackview BL5000",
	"Blackview BV4900 Pro",
	"Blackview BV5200 Pro",
	"Blackview BV9900E",
	"Fairphone 2",
	"Google Pixel 3a",
	"Google Pixel 4a",
	"Google Pixel 4a",
	"Google Pixel 6a",
	"Honor 70",
	"Honor 8",
	"Honor Magic4 Lite",
	"HTC Desire 22 Pro",
	"Huawei Mate S",
	"Huawei Nova 10 SE",
	"Huawei Nova 8I",
	"Huawei nova 9",
	"Huawei Nova 9 SE",
	"Huawei Nova Y70",
	"Huawei Nova Y90",
	"Huawei P40 Lite",
	"Huawei P8",
	"Huawei Y5p",
	"Infinix Hot 11",
	"Infinix HOT 11S",
	"Motorola Defy",
	"Motorola Moto e22",
	"Motorola Moto E22i",
	"Motorola Edge 20",
	"Motorola Edge 30",
	"Motorola Moto G42",
	"Motorola Google Nexus 6",
	"Motorola Moto E32",
	"Motorola Moto E32s",
	"Motorola Edge 30 Neo",
	"Motorola Edge 30 Neo",
	"Motorola Moto G",
	"Motorola Moto G Plus",
	"Motorola Moto G10",
	"Motorola Moto G200",
	"Motorola moto g22",
	"Motorola Moto G30",
	"Motorola Moto G31",
	"Motorola Moto G41",
	"Motorola Moto G42",
	"Motorola Moto G5 Plus",
	"Motorola Moto G52",
	"Motorola Moto G5s Plus",
	"Motorola Moto G6 Play",
	"Motorola Moto G62",
	"Motorola Edge 20",
	"Nokia 3.4",
	"Nokia 7.1",
	"Nokia G11",
	"Nokia G20",
	"Nokia G21",
	"Nokia G50",
	"Nokia Lumia 930",
	"Nokia G11",
	"Nokia X20",
	"Nokia XR20",
	"OnePlus 3T",
	"OnePlus 5T",
	"OnePlus 6",
	"OnePlus 6T",
	"OnePlus 7T",
	"OnePlus 8T",
	"OnePlus Nord 2",
	"OnePlus Nord 2",
	"OnePlus Nord 2",
	"OnePlus Nord 2T",
	"OnePlus Nord CE",
	"OnePlus Nord CE 2 Lite",
	"OPPO A57s",
	"OPPO A74",
	"OPPO Find X2 Lite",
	"OPPO Find X3 Lite",
	"OPPO Find X3 Neo",
	"OPPO Find X5 Lite",
	"OPPO Reno7 Lite",
	"OPPO Reno8 Lite",
	"OPPO Reno4 Pro",
	"OPPO Reno4 Z",
	"OPPO Reno6",
	"realme 7 Pro",
	"realme C30",
	"realme GT Master",
	"realme Narzo 50",
	"Rephone",
	"Samsung Galaxy A12",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A14",
	"Samsung Galaxy A32",
	"Samsung Galaxy A32",
	"Samsung Galaxy A32",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A52",
	"Samsung Galaxy A52s",
	"Samsung Galaxy A52s",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A72",
	"Samsung Galaxy Note 20 Ultra",
	"Samsung Galaxy Note 4",
	"Samsung Galaxy Note 9",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy XCover 4s",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover Pro",
	"Sony Xperia 1",
	"Sony Xperia 10 III",
	"Sony Xperia XZ",
	"Sony Xperia Z3 Compact",
	"Sony Xperia Z3+",
	"Sony Xperia Z5 Compact - The Bond Phone",
	"Sony Xperia Z5 Premium",
	"TCL 30E",
	"Vivo NEX S",
	"Vivo X60 Pro",
	"Vivo Y01",
	"Vivo Y21s",
	"Vivo Y21s",
	"Vivo Y72",
	"Xiaomi 11 Lite",
	"Xiaomi 11T",
	"Xiaomi 11T",
	"Xiaomi 11T Pro",
	"Xiaomi Mi 10 Lite",
	"Xiaomi Mi 10T Pro",
	"Xiaomi Mi 11 Lite",
	"Xiaomi Mi 11 Lite",
	"Xiaomi Mi 9",
	"Xiaomi Poco M5s",
	"Xiaomi Poco F3",
	"Xiaomi Poco M3",
	"Xiaomi POCO M4",
	"Xiaomi Poco X3",
	"Xiaomi Poco X4 Pro",
	"Xiaomi Poco X4 Pro",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 9",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi A1+",
	"Xiaomi Redmi A1+",
	"Xiaomi Redmi A2",
	"Xiaomi Redmi Note 10",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11s",
	"Xiaomi Redmi Note 8",
	"Xiaomi Redmi Note 8 Pro",
	"Xiaomi Redmi Note 9",
	"Xiaomi Redmi Note 9 Pro",
	"ZTE Blade A31",
	"ZTE Blade A31 Lite",
	"ZTE Blade V40S",
}

func TestGalaxusClean(t *testing.T) {
	for i, name := range galaxusNames {
		if _name := shop.GalaxusCleanFn(name); _name != galaxusNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, galaxusNamesExpected[i], name)
		}
	}
}
