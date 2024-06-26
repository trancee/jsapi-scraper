package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var cashConvertersNames = []string{
	"Apple - IPhone 11",
	"Apple - iPhone 11",
	"Apple - iPhone 11 Rouge",
	"Apple - iPhone 8",
	"Apple iPhone 11 64 gb",
	"Apple IPhone 12 128GB + boîte",
	"Apple iPhone 7 32gb",
	"Apple iPhone 8 64gb",
	"Crosscal Core x4",
	"Huawei P20 Lite 64Go Dual Sim",
	"Ibasso DX160",
	"Iphone 11 64GB",
	"IPhone 11 64GB",
	"iPhone 12 (Blanc) 64gb, 90%",
	"IPhone 12 (rouge) 64gb 88%",
	"iPhone 6S 64Gb Space Gray",
	"iPhone 7 32 Go Black Reconditionné",
	"iPhone 7 32Gb Black",
	"Iphone 7 plus 128gb",
	"iPhone 7 Plus 32GB en l état dans les coins",
	"iPhone 8 256Gb Or Rose Reconditionné",
	"Iphone 8 64 gb",
	"iPhone 8 64 gb",
	"iPhone 8 64GB (Blanc)",
	"iPhone 8 64GB (rose)",
	"IPhone 8 64GB Black",
	"IPhone 8 64GB Pink 100%",
	"Iphone 8 64GB,",
	"iPhone SE (2022) 5G 64Go Midnight",
	"iPhone SE 2020 (Noir) 128GB,",
	"iPhone SE 2022 5G 64Go Midnight Reconditionné",
	"iPhone SE 2022 5G 64Go Red",
	"IPhone SE 2022 64GB",
	"IPhone SE 2022 64GB Bleu Navy",
	"IPhone SE 2022 64GB Red",
	"iPhone SE 3rd 2022 64GB batterie 100%",
	"IPhone SE 64GB 2022",
	"Iphone XR (Corail) 64GB Batt: 99%",
	"Motorola G7 power Dual SIM 64GB",
	"Motorola G7 Power Dual Sim 64Go Noir Reconditionné",
	"Motorola G9 Play 64GB",
	"Motorola Moto G200 5G + Boîte et housse",
	"Nokia G21 + Boîte",
	"One + 10T 128GB",
	"Oppo A54 5G 64GB",
	"Oppo A57 64gb dual sim",
	"Oppo A57S 128GB + Boîte",
	"Oppo A9 2020 128gb",
	"Oppo A91 128gb",
	"Oppo A94 128GG 8GB de ram ,",
	"Oppo A94 5G 128Gb",
	"Oppo A96 128GB",
	"Oppo Find X2 lite",
	"Oppo Find X2 Pro 512GB + Boite + 2 Coques,",
	"Oppo Find X3 Lite 5G 128GB",
	"Oppo Fond x2 Pro 5G 512gb",
	"Oppo Reno 2, 256gb Dual sim",
	"Oppo Reno 2z",
	"Oppo Reno8 5G 256GB - NEUF",
	"Poco M3 Pro 5G 64gb",
	"Redmi 9T Dual Sim 64Go Vert Reconditionné",
	"Redmi Note 10 5G 64GB Dual SIM",
	"Redmi Note 10 Pro 128gb dual sim",
	"Redmi Note 10S 128G Dual SIM",
	"Redmi Note 11S 128gb dual sim",
	"Samsung : Galaxy S7 Edge 32GB",
	"Samsung - A8 2018",
	"Samsung - Note20 5G",
	"Samsung - Note 20 5G",
	"Samsung - S20+ 5G",
	"Samsung - S7 2016",
	"Samsung A12",
	"Samsung Galaxy A03S 32gb",
	"Samsung Galaxy A13 64GB",
	"Samsung galaxy A40 Dual 64GB",
	"SAMSUNG GALAXY A54 - NEUF",
	"Samsung Galaxy A8 (2018) 32Go Dual Sim",
	"Samsung Galaxy A8 2018 32GB Dual",
	"SAMSUNG GALAXY A8 SM-A530F NFC LTE,",
	"Samsung Galaxy Active 2 8’ Wifi + LTE",
	"Samsung Galaxy J5 (2017) 16Go",
	"Samsung Galaxy Note 20 (Black) 256GO",
	"Samsung Galaxy Note 20 128GB",
	"SAMSUNG GALAXY S10PLUS",
	"Samsung Galaxy S10 + 128Go Black Dual Sim",
	"Samsung Galaxy s10 128 GB",
	"Samsung Galaxy S10 128gb",
	"Samsung Galaxy S10 128GB",
	"Samsung Galaxy S10 128Gb + Coque + Boîte",
	"Samsung Galaxy S10 128Go White Reconditionné",
	"Samsung Galaxy S20 FE 5G 128Go",
	"Samsung Galaxy S20 Plus 128gb",
	"Samsung Galaxy S20+ 5G 128Go",
	"Samsung Galaxy S21 5G 256gb",
	"Samsung Galaxy S7 32GB",
	"SAMSUNG GALAXY S7 32GB NFC LTE,",
	"Samsung Galaxy S9 64gb",
	"Samsung Galaxy S9 64GB 273013",
	"Samsung Galaxy S9 64Go Black Reconditionné",
	"Samsung Galaxy S9 64Go Blue Reconditionné",
	"Samsung Galaxy S9 64Go Dual Sim",
	"Samsung Galaxy S9 Plus 64 GB 274209",
	"Samsung Galaxy S9 Plus 64gb",
	"Samsung Galaxy S9, 64gb",
	"Samsung Galaxy Xcover 4 16 Go noir",
	"Samsung Galaxy Xcover 5 64 GB",
	"Samsung Reconditionné - Samsung XCOVER 4S 32GB",
	"SAMSUNG S20FE,",
	"Samsung s20Fe 128Go Bleu navy",
	"Samsung S7 32GB",
	"Téléphone : iPhone 8 64 GB batterie à 78%",
	"Téléphone Galaxy S7 32GB",
	"Téléphone OPPO A15 32GB Dual",
	"Téléphone OPPO A16S 64GB Dual",
	"Téléphone OPPO A54 5G 64GB Dual",
	"Téléphone Oppo A54 5G 64GB,",
	"Téléphone OPPO A74 128GB dual",
	"Téléphone Oppo Find X3 Lite 128GB,",
	"Téléphone Portable MI 10T Lite 128GB,",
	"Téléphone Portable Samsung Galaxy S7 32GB NFC Lite,",
	"Téléphone Portable Xiaomi Mi Mix 2 64GB,",
	"Téléphone Samsung - J6 2018",
	"Téléphone Samsung - Samsung S20FE 5G",
	"Téléphone Samsung Galaxy A02s 32GB",
	"Téléphone Samsung Galaxy A5 2017 32GB",
	"Téléphone Samsung Galaxy AO2S 32B",
	"Téléphone Samsung Galaxy J6 2018 32GB Dual",
	"Téléphone Samsung Galaxy Note 4 32GB",
	"Téléphone Samsung Galaxy S10+ 128GB + Cover et boite",
	"Téléphone Samsung Galaxy S9 Plus 64GB + Boîte",
	"Téléphone Xiaomi Mi 10T Lite (Bleu Azur) 128GB,",
	"Téléphone Xiaomi Mi 10T Lite (Gris) 128GB,",
	"Téléphone Xiaomi MI 5G 256GB + boit",
	"Téléphone Xiaomi MI Lite 5G 64GB",
	"Téléphone Xiaomi POCO M3 64GB Dual",
	"Téléphone Xiaomi Redmi 9 32GB Dual",
	"Téléphone Xiaomi Redmi 9T 64GB",
	"Téléphone Xiaomi Redmi Note 10 128BG Dual",
	"Téléphone Xiaomi Redmi Note 10 128GB Dual",
	"Téléphone Xiaomi Redmi Note 10 5g 64GB Dual",
	"Téléphone Xiaomi Redmi Note 10S 64GB Dual",
	"Téléphone Xiaomi Redmi Note 9S 128GB Dual",
	"Xiaomi - Mi 10 Lite 5G",
	"Xiaomi - Mi 10T Lite",
	"Xiaomi - Redmi Note 11",
	"Xiaomi - Redmi Note 9S",
	"Xiaomi Mi 11 Lite 5G 128Gb Dual",
	"Xiaomi Mi 11 Lite 5G 128GB Dusl Sim",
	"Xiaomi Mi Mix 2 64Go Dual Sim Noir Reconditionné",
	"Xiaomi Redmi Mi 10 Lite 5G 64GB + boite:",
}

var cashConvertersNamesExpected = []string{
	"Apple iPhone 11",
	"Apple iPhone 11",
	"Apple iPhone 11",
	"Apple iPhone 8",
	"Apple iPhone 11",
	"Apple iPhone 12",
	"Apple iPhone 7",
	"Apple iPhone 8",
	"Crosscal Core X4",
	"HUAWEI P20 lite",
	"Ibasso Dx160",
	"Apple iPhone 11",
	"Apple iPhone 11",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 6S",
	"Apple iPhone 7",
	"Apple iPhone 7",
	"Apple iPhone 7 Plus",
	"Apple iPhone 7 Plus",
	"Apple iPhone 8",
	"Apple iPhone 8",
	"Apple iPhone 8",
	"Apple iPhone 8",
	"Apple iPhone 8",
	"Apple iPhone 8",
	"Apple iPhone 8",
	"Apple iPhone 8",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE",
	"Apple iPhone XR",
	"motorola moto g7 power",
	"motorola moto g7 power",
	"motorola moto g9 play",
	"motorola moto g200",
	"Nokia G21",
	"OnePlus 10T",
	"OPPO A54",
	"OPPO A57",
	"OPPO A57s",
	"OPPO A9 2020",
	"OPPO A91",
	"OPPO A94",
	"OPPO A94",
	"OPPO A96",
	"OPPO Find X2 Lite",
	"OPPO Find X2 Pro",
	"OPPO Find X3 Lite",
	"OPPO Find X2 Pro",
	"OPPO Reno2",
	"OPPO Reno2 Z",
	"OPPO Reno8",
	"Xiaomi POCO M3 Pro",
	"Xiaomi Redmi 9T",
	"Xiaomi Redmi Note 10",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11S",
	"Samsung Galaxy S7 Edge",
	"Samsung Galaxy A8 2018",
	"Samsung Galaxy Note20",
	"Samsung Galaxy Note20",
	"Samsung Galaxy S20+",
	"Samsung Galaxy S7 2016",
	"Samsung Galaxy A12",
	"Samsung Galaxy A03s",
	"Samsung Galaxy A13",
	"Samsung Galaxy A40",
	"Samsung Galaxy A54",
	"Samsung Galaxy A8 (2018)",
	"Samsung Galaxy A8 2018",
	"Samsung Galaxy A8",
	"Samsung Galaxy Active 2",
	"Samsung Galaxy J5 (2017)",
	"Samsung Galaxy Note20",
	"Samsung Galaxy Note20",
	"Samsung Galaxy S10 Plus",
	"Samsung Galaxy S10+",
	"Samsung Galaxy S10",
	"Samsung Galaxy S10",
	"Samsung Galaxy S10",
	"Samsung Galaxy S10",
	"Samsung Galaxy S10",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 Plus",
	"Samsung Galaxy S20+",
	"Samsung Galaxy S21",
	"Samsung Galaxy S7",
	"Samsung Galaxy S7",
	"Samsung Galaxy S9",
	"Samsung Galaxy S9",
	"Samsung Galaxy S9",
	"Samsung Galaxy S9",
	"Samsung Galaxy S9",
	"Samsung Galaxy S9 Plus",
	"Samsung Galaxy S9 Plus",
	"Samsung Galaxy S9",
	"Samsung Galaxy XCover 4",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover 4S",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S7",
	"Apple iPhone 8",
	"Samsung Galaxy S7",
	"OPPO A15",
	"OPPO A16s",
	"OPPO A54",
	"OPPO A54",
	"OPPO A74",
	"OPPO Find X3 Lite",
	"Xiaomi Mi 10T Lite",
	"Samsung Galaxy S7",
	"Xiaomi Mi Mix 2",
	"Samsung J6 2018",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy A02s",
	"Samsung Galaxy A5 2017",
	"Samsung Galaxy A02s",
	"Samsung Galaxy J6 2018",
	"Samsung Galaxy Note4",
	"Samsung Galaxy S10+",
	"Samsung Galaxy S9 Plus",
	"Xiaomi Mi 10T Lite",
	"Xiaomi Mi 10T Lite",
	"Xiaomi Mi",
	"Xiaomi Mi Lite",
	"Xiaomi POCO M3",
	"Xiaomi Redmi 9",
	"Xiaomi Redmi 9T",
	"Xiaomi Redmi Note 10",
	"Xiaomi Redmi Note 10",
	"Xiaomi Redmi Note 10",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 9S",
	"Xiaomi Mi 10 Lite",
	"Xiaomi Mi 10T Lite",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 9S",
	"Xiaomi Mi 11 Lite",
	"Xiaomi Mi 11 Lite",
	"Xiaomi Mi Mix 2",
	"Xiaomi Redmi Mi 10 Lite",
}

func TestCashConvertersClean(t *testing.T) {
	for i, name := range cashConvertersNames {
		if _name := shop.CashConvertersCleanFn(name); _name != cashConvertersNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, cashConvertersNamesExpected[i], name)
		}
	}
}
