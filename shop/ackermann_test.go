package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var ackermannNames = []string{
	" Nothing Phone -2 12 GB / 25",
	" Nothing Phone -2 12 GB / 51",
	"Apple iPhone 11, 4G",
	"Apple iPhone 12, 64 GB",
	"Apple iPhone 13 mini, 128 GB",
	"Apple iPhone 13, 128 GB",
	"Apple iPhone 14 Plus, 128 GB",
	"Apple iPhone 14 Pro Max, 128 GB",
	"Apple iPhone 14 Pro, 128 GB",
	"Apple iPhone SE 3. Gen., 128 GB",
	"Apple XR",
	"Fairphone 4 5G 128 GB",
	"Fairphone 4 5G 256 GB",
	"Google 7 128 GB Snow",
	"Google Pixel 6a 128GB black",
	"Google Pixel 7 Pro 128GB Obsidian",
	"Google Pixel 7a 128 GB Schwarz",
	"Huawei Mate 50 Pro black",
	"Huawei Mate 50 Pro silver",
	"Huawei Pro 256 GB Pearl",
	"Huawei Pro 256 GB Schwarz",
	"Motorola edge20 Pro",
	"Motorola Moto 454",
	"Motorola Moto G 5G",
	"Motorola Moto G100",
	"Motorola moto g32",
	"Motorola razr 40",
	"Motorola razr 40 Ultra",
	"Nokia 128 GB Black",
	"Nokia 5G 128GB/6GB RAM, DS, cloudy blue",
	"Nokia 5G 256GB/8GB RAM, DS, cloudy blue",
	"Nokia G22 64GB Lagoon Blue",
	"Nokia G42 128 GB Grey",
	"Nokia G42 128 GB Purple",
	"Nokia X30 5G 128 GB/6 RAM, DS, ice white",
	"OnePlus 11 5G 128 GB Titan Black",
	"OnePlus Nord CE 3 128 GB Lite Chrom",
	"OnePlus Nord CE 3 128 GB Lite Paste",
	"Oppo 128 GB Cosmo Blue",
	"Oppo 128 GB Fluid Black",
	"Oppo 128 GB Sunset Blue",
	"Oppo 256GB Ocean Blue",
	"Oppo Pro 5G Arctic Blue",
	"Oppo Reno 8 Lite 128 GB Cosmic",
	"Oppo Reno10 256 GB Ice Blue",
	"Oppo Reno10 256 GB Silvery Grey",
	"Oppo Reno8 Pro Glazed Black",
	"Oppo Reno8 Pro Glazed Green",
	"Oppo X3 Pro 256 GB Black",
	"Oppo X3 Pro 256 GB Blue",
	"Oppo X5 Lite 256 GB Hellblau",
	"Oppo X5 Lite 256 GB Schwarz",
	"Samsung A13 128 GB CH Black",
	"Samsung A13 128 GB CH Blue",
	"Samsung A14 128 GB CH Black",
	"Samsung A14 128 GB CH Lime",
	"Samsung A14 128 GB CH Silver",
	"Samsung A14 5G 128 GB CH",
	"Samsung A14 5G 128 GB CH Lim",
	"Samsung A34 5G 128 GB CH Awe",
	"Samsung A34 5G 128 GB CH Ent",
	"Samsung A34 5G 256 GB CH Awe",
	"Samsung A53 5G 128 GB CH Awe",
	"Samsung A54 5G 128 GB CH",
	"Samsung A54 5G 128 GB CH Awe",
	"Samsung A54 5G 256 GB CH Awe",
	"Samsung Galaxy A13",
	"Samsung GALAXY A14 LTE 64GB",
	"Samsung Galaxy A34 5G",
	"Samsung Galaxy A53 5G",
	"Samsung Galaxy A54 5G",
	"Samsung Galaxy S21 FE 5G",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22 Ultra",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23 Plus",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy Xcover",
	"Samsung Galaxy Xcover Pro",
	"Samsung S23 128 GB CH Cream",
	"Samsung S23 128 GB CH Green",
	"Samsung S23 128 GB CH Lavend",
	"Samsung S23 128 GB CH Phanto",
	"Samsung S23 256 GB CH Cream",
	"Samsung S23 256 GB CH Enterp",
	"Samsung S23 256 GB CH Green",
	"Samsung S23 256 GB CH Phanto",
	"Samsung S23 Ultra 256 GB CH",
	"Samsung S23 Ultra 512 GB CH",
	"Samsung S23+ 256 GB CH Cream",
	"Samsung S23+ 256 GB CH Green",
	"Samsung S23+ 256 GB CH Laven",
	"Samsung S23+ 256 GB CH Phant",
	"Samsung S23+ 512 GB CH Cream",
	"Samsung S23+ 512 GB CH Green",
	"Samsung S23+ 512 GB CH Laven",
	"Samsung S23+ 512 GB CH Phant",
	"Xiaomi 10 2022 128 GB Carbon",
	"Xiaomi 13 Lite 128 GB Black",
	"Xiaomi 13 Lite 128 GB Blue",
	"Xiaomi 13 Lite 128 GB Pink",
	"Xiaomi F3 256 GB Ocean Bl",
	"Xiaomi Note 11 128 GB Grau",
	"Xiaomi Note 11 Pro 5G 128 GB",
	"Xiaomi Redmi Note 12 PRO 5G 128GB black",
	"Xiaomi Redmi Note 12 PRO 5G 128GB blue",
	"Xiaomi Redmi Note 12 PRO+ 5G 256GB black",
}

var ackermannNamesExpected = []string{
	"Nothing Phone (2)",
	"Nothing Phone (2)",
	"Apple iPhone 11",
	"Apple iPhone 12",
	"Apple iPhone 13 mini",
	"Apple iPhone 13",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Pro Max",
	"Apple iPhone 14 Pro",
	"Apple iPhone SE (2022)",
	"Apple iPhone XR",
	"Fairphone 4",
	"Fairphone 4",
	"Google 7",
	"Google Pixel 6a",
	"Google Pixel 7 Pro",
	"Google Pixel 7a",
	"HUAWEI Mate 50 Pro",
	"HUAWEI Mate 50 Pro",
	"HUAWEI Pro",
	"HUAWEI Pro",
	"motorola edge 20 pro",
	"motorola moto g54",
	"motorola moto g",
	"motorola moto g100",
	"motorola moto g32",
	"motorola razr 40",
	"motorola razr 40 ultra",
	"Nokia",
	"Nokia",
	"Nokia",
	"Nokia G22",
	"Nokia G42",
	"Nokia G42",
	"Nokia X30",
	"OnePlus 11",
	"OnePlus Nord CE 3",
	"OnePlus Nord CE 3",
	"OPPO",
	"OPPO",
	"OPPO",
	"OPPO",
	"OPPO Pro",
	"OPPO Reno8 Lite",
	"OPPO Reno10",
	"OPPO Reno10",
	"OPPO Reno8 Pro",
	"OPPO Reno8 Pro",
	"OPPO X3 Pro",
	"OPPO X3 Pro",
	"OPPO X5 Lite",
	"OPPO X5 Lite",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A53",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A13",
	"Samsung Galaxy A14",
	"Samsung Galaxy A34",
	"Samsung Galaxy A53",
	"Samsung Galaxy A54",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22 Ultra",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23 Plus",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover 6 Pro",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Xiaomi Redmi 10 2022",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Lite",
	"Xiaomi POCO F3",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi Redmi Note 12 Pro+",
}

func TestAckermannClean(t *testing.T) {
	for i, name := range ackermannNames {
		if _name := shop.AckermannCleanFn(name); _name != ackermannNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, ackermannNamesExpected[i], name)
		}
	}
}
