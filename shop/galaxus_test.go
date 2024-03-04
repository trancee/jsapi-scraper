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
	"Blackview MOBILE PHONE A70/BLUE",
	"Blackview MOBILE PHONE BV9300/LASER BLACK",
	"Blackview MOBILE PHONE N6000/8/256 BLACK",
	"Blackview MOBILE PHONE N6000/8/256 ORANGE",
	"Blackview BL5000 5G 8",
	"Blackview BV4900 Pro (5.7 Zoll) Dual-SIM Android 10.0 Mikro-USB",
	"Blackview BV5200 Pro 5180 mAh 4/64 GB Schwarz (schwarz) Smartphone",
	"Blackview BV7100 13000 mAh 6/128 GB Grünes Smartphone",
	"Blackview Smartfon Blackview BV9300 12/256GB Pomarańczowy",
	"Blackview Smartfon BV5300 czarny",
	"Blackview Smartfon BV5300 zielony",
	"Blackview Smartfon BV5300 PRO pomarańczowy",
	"Blackview Smartphone Blackview Smartphone BV9900E 6/128 Silver",
	"Blackview Smartphone Blackview BV9200 8/256 Orange",
	"Emporia Super Easy",
	"Fairphone 2",
	"Gigaset GX4 Outdoor",
	"Google Pixel 3a",
	"Google Pixel 4a",
	"Google Pixel 4a 5G",
	"Google Pixel 6a",
	"Honor 70",
	"Honor 70 Lite Black UC",
	"Honor 8",
	"Honor 90 LITE 8G/256G MIDNIGHT BLACK",
	"Honor HON DS X7 4+128 TIM BLK",
	"Honor Magic 4 Lite 5G",
	"Honor Telefon MAGIC5 LITE",
	"Honor Magic 6 Lite Smaragdgrün",
	"Honor Telefonas HONOR X7A 4",
	"HTC Desire 22 Pro",
	"Huawei Mate S Force Touch",
	"Huawei Nexus 6P",
	"Huawei Nova 10 SE",
	"Huawei Nova 8I",
	"Huawei nova 9",
	"Huawei Nova 9 SE",
	"Huawei Nova Y70",
	"Huawei Nova Y90",
	"Huawei P40 Lite",
	"Huawei P40 Lite + Huawei Bluetooth Speaker",
	"Huawei P8",
	"Huawei Porsche Design Mate 10",
	"Huawei Y5p",
	"Inapa Tecno Spark 10 5G 4GB/64GB Android 13 Smartphone Meta Blue",
	"Infinix Hot 11 X689F",
	"Infinix HOT 11S 17 2 cm (6.78 Zoll) Dual-SIM Android 11 4G USB Typ-C 6 GB 128 GB 5000 mAh Schwarz",
	"Infinix HOT205G4128BL smartphone (6.6\") Dual SIM Android 12 USB Type-C Silver",
	"Infinix HOT205G4128BK smartphone (6.6\") Dual SIM Android 12 Black",
	"Infinix INFINIX Smart 6 hellblau",
	"Infinix NOTE 12 PRO (6.7\") Dual SIM Android 12 USB Type-C Black",
	"Infinix Smartphone Infinix Hot 11S 4/64GB Dual SIM Grün (X6812B4GW)",
	"Infinix Telefon HOT 12I/4/64GB RACING Schwarz INFINIX",
	"Infinix Telefon SMART 6 HD/2/32GB FORCE Schwarz INFINIX",
	"Maxcom Telefonas MM 827 4G VoLTE",
	"Motorola 30 Neo",
	"Motorola 22",
	"Motorola 41",
	"Motorola Defy",
	"Motorola e22",
	"Motorola E22i",
	"Motorola Edge 20",
	"Motorola EDGE 20 LITE 8 / 128GB LAGOON Green",
	"Motorola Edge 30",
	"Motorola Edge 30 5G OLED Dual SIM 256/8GB Aurora Green",
	"Motorola Edge 30 Neo (6,28\") Dual-SIM Android 12 5G USB Typ-C 8 GB 128 GB 4020 mAh MOONLESS NIGHT Juodas",
	"Motorola G42",
	"Motorola Google Nexus 6",
	"Motorola GRADE-C Motorola One Macro 64GB Space Blau",
	"Motorola Moto E Moto E6 play - 14 cm (5.5 Zoll) - 2 GB - 32 GB - 13 MP - Android",
	"Motorola Moto E moto e7 (6.5 Zoll) Dual-SIM Android 10.0 USB Typ-C",
	"Motorola Moto E 13 16 5 cm (6.5 Zoll) Dual-SIM Android 13 Go edition 4G USB Typ-C 2 GB 64 GB 5000 mA",
	"Motorola Moto E13 16,5 cm (6,5\") Dual-SIM Android 13 Go Edition 4G USB Type-C 2 GB 64 GB 5000 mAh Kapazität",
	"Motorola Moto E13 -puhelin, 128/8 Gt, Cosmic Black",
	"Motorola Moto E20 2 ”",
	"Motorola Moto E32",
	"Motorola Moto E32s Dual SIM, 4GB RAM, 64GB, Misty Silver",
	"Motorola Moto E32s Gravity Grey",
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
	"Motorola Moto G73 5G Phone, 256/8 GB, Midnight Blue",
	"Motorola moto g84 5G 12/256 Midnight Blue",
	"Motorola moto g84 5G 12/256 Viva Magenta",
	"Motorola Motorola Moto G53",
	"Motorola Smartfon Motorola Moto G13 4/128GB Matte Charcoal",
	"Motorola SMARTPHONE MOTOROLA MOTO G53 5G 4/128 ARCTIC Silber",
	"Motorola XT2143-1 Edge 20 Grey EU Android",
	"Nokia 3.4",
	"Nokia 7.1",
	"Nokia C02 TA-1460 DS 2/32 BNLFRI CHARCOAL",
	"Nokia G11",
	"Nokia G20",
	"Nokia G21",
	"Nokia G22 Blue",
	"Nokia G50",
	"Nokia Handy lyra 110 4g ds ta-1386 juoda",
	"Nokia Lumia 930",
	"Nokia Smartphone G11 DS 3",
	"Nokia X20 (6GB)",
	"Nokia XR20",
	"OnePlus 3T",
	"OnePlus 5T",
	"OnePlus 6",
	"OnePlus 6T (8GB)",
	"OnePlus 6T McLaren Edition",
	"OnePlus 7T (8GB)",
	"OnePlus 8T (12GB)",
	"OnePlus 8 Pro Limited (12GB)",
	"OnePlus Nord 2",
	"OnePlus Nord 2 (12GB)",
	"OnePlus Nord 2 (8GB)",
	"OnePlus Nord 2T",
	"OnePlus Nord CE",
	"OnePlus Nord CE 2 Lite 5G",
	"OnePlus Nord CE3 Lite 5G Pastel Lime",
	"OPPO 8 Lite",
	"OPPO A57s",
	"OPPO A74",
	"OPPO A74 5G Space Silver",
	"OPPO Find X2 Lite",
	"OPPO Find X3 Lite",
	"OPPO Find X3 Neo",
	"OPPO Find X5 Lite",
	"OPPO OPP DS RENO8 LITE 5G 8+128 OEM EU BLK",
	"OPPO Reno 7 Lite",
	"OPPO Reno 8 Lite",
	"OPPO Reno4 Pro",
	"OPPO Reno4 Z",
	"OPPO Reno6",
	"OPPO Smartphone A57s 4",
	"OPPO Smartphone oppo reno 8 lite 8/128gb schwarz",
	"Poco M4 Pro",
	"POCO MOBILE PHONE X5 5G/6/128GB BLACK MZB0D5REU",
	"realme 53",
	"realme 7 Pro",
	"realme C21Y (3GB+32GB) cross black",
	"realme C30 bamboo green 3+32GB",
	"realme GT Master",
	"realme GT Master Edition",
	"realme Narzo 50 4g 128 GB Speed Blue",
	"realme SM REALME C21 BLACK 6,5\" 3+64GB DS ITA",
	"realme Telefon C21-Y DS 4 GB RAM 64 GB (Juodas)",
	"Renewd iPhone SE2020",
	"Rephone Rephone",
	"Samsung Galaxy A12 (2021)",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13 Android",
	"Samsung Galaxy A13 EU",
	"Samsung Galaxy A14 LTE",
	"Samsung Galaxy A21s DE",
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
	"Samsung Galaxy M13 . Bildschirmdiagonale: 16,8 cm (6.6\" ), Bildschirmauflösung: 2408 x 1080 Pixel, Display-T",
	"Samsung Galaxy Note 20 Ultra Snapdragon",
	"Samsung Galaxy Note 4",
	"Samsung Galaxy Note 9",
	"Samsung Galaxy S20 FE 5G",
	"Samsung Galaxy S20 FE (Snapdragon) EU",
	"Samsung Galaxy XCover 4s",
	"Samsung Galaxy XCover 5 EE Enterprise Edition",
	"Samsung Galaxy XCover Pro EE Enterprise Edition",
	"Samsung Galaxy Z Flip3 5G EU",
	"Samsung I9195I Galaxy S4 Mini Value Edition",
	"Sony Xperia 1",
	"Sony Xperia 10 III",
	"Sony Xperia XZ",
	"Sony Xperia Z3 Compact, Schwarz",
	"Sony Xperia Z3+",
	"Sony Xperia Z5 Compact - The Bond Phone",
	"Sony Xperia Z5 Premium",
	"TCL 30E (6.52 Zoll) Dual-SIM Android 12 USB Typ-C",
	"TCL 40 40SE 17,1 cm (6.75\" ) Dual-SIM Android 13 4G USB Typ-C 6 GB 256 GB 5010 mAh Violett",
	"TCL 403 15,2 cm (6\" ) Dual-SIM Android 12 Go Edition 4G Mikro-USB 2 GB 32 GB 3000 mAh Malve",
	"TCL TCT 30 17 cm (6.7\" ) Dual-SIM Android 12 4G USB Typ-C 4 GB 64 GB 5010 mAh Blau",
	"TE Connectivity ZTE Blade A71",
	"TIM HONOR Magic5 Lite",
	"TIM moto g53 5g",
	"TIM SAM DS A236B GAL A23 5G 4+128 TIM BLK",
	"Vivo NEX S",
	"Vivo X60 Pro 5G",
	"Vivo Y01 Elegant Black",
	"Vivo Y21 Pearl White",
	"Vivo Y21s midday dream",
	"Vivo Y21s midnight blue",
	"Vivo Y72 5G (V2041) dream glow",
	"Vivo Y76 5G Cosmic Aurora",
	"Vivo Y76 5G Midnight Space",
	"Vodafone Samsung Galaxy A33",
	"Xiaomi 11 Lite 5G NE",
	"Xiaomi 11T",
	"Xiaomi 11T 5G",
	"Xiaomi 11T Pro",
	"Xiaomi Honor Magic5 Lite",
	"Xiaomi M5 6/128GB GELB",
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
	"Xiaomi Redmi10 2022",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10 Dual-Sim (2022) 64GB, Sea Blue",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 12C (Ocean Blue) DS 6.71“ IPS LCD 720x1650/2.0GHz&1.8GHz/32GB/3GB RAM/MIUI 13/microSDXC",
	"Xiaomi Redmi 13C Clover Green, 6/128Gb",
	"Xiaomi Redmi 13C Navy Blue, 6/128Gb",
	"Xiaomi Redmi 9",
	"Xiaomi Redmi 9A",
	"Xiaomi redmi 9a 32gb handy",
	"Xiaomi Redmi 9C NFC",
	"Xiaomi Redmi A1 Plus",
	"Xiaomi Redmi A1+ Dual Sim 32GB, 2GB RAM, Blue",
	"Xiaomi Redmi A2 (Juodas) Dual-SIM 6,52 IPS LCD 7200x1600",
	"Xiaomi Redmi A2 2/32 GB blue EU",
	"Xiaomi Redmi Note 1",
	"Xiaomi REDMI NOTE 1 - Smartphone - 256 GB - Grau",
	"Xiaomi Redmi Note 10",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro 5G (6GB)",
	"Xiaomi Redmi Note 11 Pro+ 5G",
	"Xiaomi Redmi Note 11s 5G Dual Sim 4GB RAM 128GB - Twilight Blue EU",
	"Xiaomi Redmi Note 11S 5G modrá star/6,6´´ AMOLED/90HZ/FullHD+/2GHz OC/4GB/128GB/SD/2xSIM/50+8+2MPx/5000mAh",
	"Xiaomi Redmi Note 12 OEM",
	"Xiaomi Redmi Note 12 Pro 5G Sky Blue, 8/256Gb",
	"Xiaomi Redmi Note 8",
	"Xiaomi Redmi Note 8 Pro",
	"Xiaomi Redmi Note 9",
	"Xiaomi Redmi Note 9 Pro",
	"Xiaomi Samsung Galaxy A14 5G",
	"Xiaomi Telefon M5/64GB Gelb MZB0C9FEU",
	"Xiaomi X5 6/128GB blue EU",
	"Xiaomi Xia Poco C40 32-3-4G-ye Xiaomi Poco C40 LTE 32/3GB Yellow",
	"Xiaomi Xia Redmi 10 (2022) 128-4-4G-wh Xia Redmi 10 2022 LTE 128/4 Pebble White",
	"Xiaomi Xiaomi 12 Lite",
	"ZTE Blade A31",
	"ZTE Blade A31 Lite",
	"ZTE Blade A31 Lite 1GB RAM, 32GB, Blau",
	"ZTE Blade A73 blue",
	"ZTE Blade V40 pro dark green",
	"ZTE Blade V40 Vita 128G+4G black(inkl. Buds",
	"ZTE Blade V40S 4G Black",
	"ZTE Blade A73 black",
	"ZTE Blade A73 5G grey",
	"ZTE Graues Smartphone Blade V40 Design + schwarze Smartwatch Watch Live",
	"ZTE Smartfon ZTE Blade A31 2/32GB Szary",
	"ZTE Smartphone A310 1/8GB Graphit (A31232/GY)",
	"ZTE Smartphone Blade A72s 3",
	"ZTE SMARTPHONE BLADE A51",
	"ZTE Supplier did not provide product name",
}

var galaxusNamesExpected = []string{
	"Apple iPhone 11 Pro",
	"Apple iPhone 6",
	"Apple iPhone 7",
	"Apple iPhone 8",
	"Apple iPhone SE (2020)",
	"Apple iPhone XS",
	"Blackview A100",
	"Blackview A55",
	"Blackview A70",
	"Blackview BV9300",
	"Blackview N6000",
	"Blackview N6000",
	"Blackview BL5000",
	"Blackview BV4900 Pro",
	"Blackview BV5200 Pro",
	"Blackview BV7100",
	"Blackview BV9300",
	"Blackview BV5300",
	"Blackview BV5300",
	"Blackview BV5300 Pro",
	"Blackview BV9900e",
	"Blackview BV9200",
	"emporiaSUPEREASY",
	"Fairphone 2",
	"Gigaset GX4",
	"Google Pixel 3a",
	"Google Pixel 4a",
	"Google Pixel 4a",
	"Google Pixel 6a",
	"HONOR 70",
	"HONOR 70 Lite",
	"HONOR 8",
	"HONOR 90 Lite",
	"HONOR X7",
	"HONOR Magic4 Lite",
	"HONOR Magic5 Lite",
	"HONOR Magic6 Lite",
	"HONOR X7a",
	"HTC Desire 22 Pro",
	"HUAWEI Mate S",
	"HUAWEI Nexus 6P",
	"HUAWEI nova 10 SE",
	"HUAWEI nova 8i",
	"HUAWEI nova 9",
	"HUAWEI nova 9 SE",
	"HUAWEI nova Y70",
	"HUAWEI nova Y90",
	"HUAWEI P40 lite",
	"HUAWEI P40 lite",
	"HUAWEI P8",
	"HUAWEI Mate 10",
	"HUAWEI Y5p",
	"TECNO Spark 10",
	"Infinix HOT 11",
	"Infinix HOT 11S",
	"Infinix HOT 20",
	"Infinix HOT 20",
	"Infinix SMART 6",
	"Infinix NOTE 12 PRO",
	"Infinix HOT 11S",
	"Infinix HOT 12i",
	"Infinix SMART 6 HD",
	"Maxcom MM 827",
	"motorola edge 30 neo",
	"motorola moto g22",
	"motorola moto g41",
	"motorola defy",
	"motorola moto e22",
	"motorola moto e22i",
	"motorola edge 20",
	"motorola edge 20 lite",
	"motorola edge 30",
	"motorola edge 30",
	"motorola edge 30 neo",
	"motorola moto g42",
	"motorola google nexus 6",
	"motorola one macro",
	"motorola moto e6 play",
	"motorola moto e7",
	"motorola moto e13",
	"motorola moto e13",
	"motorola moto e13",
	"motorola moto e20",
	"motorola moto e32",
	"motorola moto e32s",
	"motorola moto e32s",
	"motorola edge 30 neo",
	"motorola edge 30 neo",
	"motorola moto g",
	"motorola moto g plus",
	"motorola moto g10",
	"motorola moto g200",
	"motorola moto g22",
	"motorola moto g30",
	"motorola moto g31",
	"motorola moto g41",
	"motorola moto g42",
	"motorola moto g5 plus",
	"motorola moto g52",
	"motorola moto g5s plus",
	"motorola moto g6 play",
	"motorola moto g62",
	"motorola moto g73",
	"motorola moto g84",
	"motorola moto g84",
	"motorola moto g53",
	"motorola moto g13",
	"motorola moto g53",
	"motorola edge 20",
	"Nokia 3.4",
	"Nokia 7.1",
	"Nokia C02",
	"Nokia G11",
	"Nokia G20",
	"Nokia G21",
	"Nokia G22",
	"Nokia G50",
	"Nokia Lyra 110",
	"Nokia Lumia 930",
	"Nokia G11",
	"Nokia X20",
	"Nokia XR20",
	"OnePlus 3T",
	"OnePlus 5T",
	"OnePlus 6",
	"OnePlus 6T",
	"OnePlus 6T McLaren Edition",
	"OnePlus 7T",
	"OnePlus 8T",
	"OnePlus 8 Pro",
	"OnePlus Nord 2",
	"OnePlus Nord 2",
	"OnePlus Nord 2",
	"OnePlus Nord 2T",
	"OnePlus Nord CE",
	"OnePlus Nord CE 2 Lite",
	"OnePlus Nord CE 3 Lite",
	"OPPO Reno8 Lite",
	"OPPO A57s",
	"OPPO A74",
	"OPPO A74",
	"OPPO Find X2 Lite",
	"OPPO Find X3 Lite",
	"OPPO Find X3 Neo",
	"OPPO Find X5 Lite",
	"OPPO Reno8 Lite",
	"OPPO Reno7 Lite",
	"OPPO Reno8 Lite",
	"OPPO Reno4 Pro",
	"OPPO Reno4 Z",
	"OPPO Reno6",
	"OPPO A57s",
	"OPPO Reno8 Lite",
	"Xiaomi POCO M4 Pro",
	"Xiaomi POCO X5",
	"realme C53",
	"realme 7 Pro",
	"realme C21Y",
	"realme C30",
	"realme GT",
	"realme GT",
	"realme narzo 50",
	"realme C21",
	"realme C21Y",
	"Apple iPhone SE (2020)",
	"Rephone",
	"Samsung Galaxy A12",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A14",
	"Samsung Galaxy A21s",
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
	"Samsung Galaxy M13",
	"Samsung Galaxy Note20 Ultra",
	"Samsung Galaxy Note4",
	"Samsung Galaxy Note9",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy XCover 4S",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover Pro",
	"Samsung Galaxy Z Flip3",
	"Samsung Galaxy S4 mini",
	"Sony Xperia 1",
	"Sony Xperia 10 III",
	"Sony Xperia XZ",
	"Sony Xperia Z3 Compact",
	"Sony Xperia Z3+",
	"Sony Xperia Z5 Compact",
	"Sony Xperia Z5 Premium",
	"TCL 30E",
	"TCL 40SE",
	"TCL 403",
	"TCL 30",
	"ZTE Blade A71",
	"HONOR Magic5 Lite",
	"motorola moto g53",
	"Samsung Galaxy A23",
	"vivo NEX S",
	"vivo X60 Pro",
	"vivo Y01",
	"vivo Y21",
	"vivo Y21s",
	"vivo Y21s",
	"vivo Y72",
	"vivo Y76",
	"vivo Y76",
	"Samsung Galaxy A33",
	"Xiaomi 11 Lite",
	"Xiaomi 11T",
	"Xiaomi 11T",
	"Xiaomi 11T Pro",
	"HONOR Magic5 Lite",
	"Xiaomi POCO M5",
	"Xiaomi Mi 10 Lite",
	"Xiaomi Mi 10T Pro",
	"Xiaomi Mi 11 Lite",
	"Xiaomi Mi 11 Lite",
	"Xiaomi Mi 9",
	"Xiaomi POCO M5s",
	"Xiaomi POCO F3",
	"Xiaomi POCO M3",
	"Xiaomi POCO M4",
	"Xiaomi POCO X3",
	"Xiaomi POCO X4 Pro",
	"Xiaomi POCO X4 Pro",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 12C",
	"Xiaomi Redmi 13C",
	"Xiaomi Redmi 13C",
	"Xiaomi Redmi 9",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi A1+",
	"Xiaomi Redmi A1+",
	"Xiaomi Redmi A2",
	"Xiaomi Redmi A2",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 10",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi Redmi Note 8",
	"Xiaomi Redmi Note 8 Pro",
	"Xiaomi Redmi Note 9",
	"Xiaomi Redmi Note 9 Pro",
	"Samsung Galaxy A14",
	"Xiaomi POCO M5",
	"Xiaomi POCO X5",
	"Xiaomi POCO C40",
	"Xiaomi Redmi 10",
	"Xiaomi 12 Lite",
	"ZTE Blade A31",
	"ZTE Blade A31 Lite",
	"ZTE Blade A31 Lite",
	"ZTE Blade A73",
	"ZTE Blade V40 Pro",
	"ZTE Blade V40 Vita",
	"ZTE Blade V40s",
	"ZTE Blade A73",
	"ZTE Blade A73",
	"ZTE Blade V40",
	"ZTE Blade A31",
	"ZTE Blade A310",
	"ZTE Blade A72s",
	"ZTE Blade A51",
	"ZTE",
}

func TestGalaxusClean(t *testing.T) {
	for i, name := range galaxusNames {
		if _name := shop.GalaxusCleanFn(name); _name != galaxusNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, galaxusNamesExpected[i], name)
		}
	}
}
