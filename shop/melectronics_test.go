package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var melectronicsNames = []string{
	"Apple iPhone 11 64GB (2021) Black",
	"Apple iPhone 12 64 GB Purple",
	"Apple iPhone 12 64GB (PRODUCT)RED",
	"Apple iPhone 13 128GB Blue",
	"Apple iPhone SE 3rd Gen 128GB (PRODUCT)RED",
	"Apple iPhone SE 3rd Gen 128GB Midnight",
	"Apple iPhone SE 3rd Gen 64GB (PRODUCT)RED",
	"Apple iPhone SE 3rd Gen 64GB Midnight",
	"Apple iPhone SE 3th 128GB Starlight",
	"Apple iPhone SE 3th 256GB Midnight",
	"Apple iPhone SE 3th 256GB Starlight",
	"Apple iPhone SE 3th 64GB Starlight",
	"Fairphone 4 5G 128 GB Grau",
	"Fairphone 4 5G 256 GB Grau",
	"Google Pixel 6a 128 GB - Charcoal",
	"Google Pixel 7 128 GB - Snow",
	"Google Pixel 7 128 GB Obsidian",
	"Huawei Nova 9 - Black",
	"Huawei Nova 9 - Starry Blue",
	"Huawei Nova 9 SE 128GB Midnight Black",
	"Huawei Y6P 64 GB Dual SIM Midnight Black (ohne Google Mobile Services)",
	"Motorola Edge 30 Fusion 5G 128GB Quartz Black",
	"Motorola Edge 30 Neo 5G 128GB Meteor Grey",
	"Motorola moto e22i 32GB Graphite Grey",
	"Motorola moto g22 64GB Cosmic Black",
	"Nokia G11 32GB Charcoal",
	"Nokia G21 128 GB Dusk",
	"Nokia G21 128 GB Nordic Blue",
	"Nokia G22 DS 64GB - Grey",
	"NOKIA G22 DS 64GB - lagoon blue",
	"Nokia G60 (5G) 128GB - black",
	"Nokia G60 (5G) 128GB - Grey",
	"Nokia X30 (5G) 128GB - Blue",
	"Nokia X30 (5G) 256GB - Blue",
	"Nokia XR20 Granite",
	"Nothing Phone (1) 8GB 128GB - black",
	"Nothing Phone (1) 8GB 256GB - black",
	"Oppo A53 s Electric Black",
	"Oppo A54s (4G) 128GB Crystal Black",
	"Oppo A57s 128GB - sky blue",
	"Oppo A57s 128GB - starry black",
	"Oppo A94 5G 128 GB cosmo blue",
	"Oppo A94 5G 128 GB fluid black",
	"Oppo A96 (4G) 128GB - starry black",
	"Oppo A96 (4G) 128GB - sunset blue",
	"Oppo Find X3 Lite 128GB astral blue",
	"Oppo Find X3 Neo 256 GB starlight black",
	"Oppo Find X5 Lite 5G 256GB Startrails Blue",
	"Oppo Reno 6 128GB Arctic Blue",
	"Oppo Reno 6 Pro 256GB Arctic Blue",
	"Oppo Reno 8 5G 256 GB - Shimmer Gold",
	"Oppo Reno 8 Lite (5G) 128GB - cosmic black",
	"Oppo Reno 8 Lite (5G) 128GB - rainbow spectrum",
	"Oppo Reno2 Luminous Ocean Blue",
	"Oppo Reno8 5G 256GB - Shimmer Black",
	"Redmi Note 12 5G 128GB - onyx gray",
	"Samsung Galaxy A12 64 GB black",
	"Samsung Galaxy A13 128GB Black",
	"Samsung Galaxy A13 128GB White",
	"Samsung Galaxy A14 5G Black 128GB",
	"Samsung Galaxy A14 5G Green 128GB",
	"Samsung Galaxy A14 5G Silver 128GB",
	"Samsung Galaxy A14 Black 128GB",
	"Samsung Galaxy A14 Green 128GB",
	"Samsung Galaxy A14 Silver 128GB",
	"Samsung Galaxy A33 5G 128GB Awesome Black",
	"Samsung Galaxy A33 5G 128GB Awesome Blue",
	"Samsung Galaxy A33 5G 128GB Awesome Peach",
	"Samsung Galaxy A33 5G 128GB Awesome White",
	"Samsung Galaxy A34 5G Awesome Graphite 128GB",
	"Samsung Galaxy A34 5G Awesome Lime 128GB",
	"Samsung Galaxy A34 5G Awesome Silver 128GB",
	"Samsung Galaxy A34 5G Awesome Violet 128GB",
	"Samsung Galaxy A52 Awesome white",
	"Samsung Galaxy A53 5G 128GB Awesome Black",
	"Samsung Galaxy A53 5G 128GB Awesome Blue",
	"Samsung Galaxy A53 5G 128GB Awesome Peach",
	"Samsung Galaxy A53 5G 128GB Awesome White",
	"Samsung Galaxy A54 5G Awesome Graphite 128GB",
	"Samsung Galaxy A54 5G Awesome Lime 128GB",
	"Samsung Galaxy A54 5G Awesome Violet 128GB",
	"Samsung Galaxy A54 5G Awesome White 128GB",
	"Samsung Galaxy S21 FE 5G 128GB Graphite",
	"Samsung Galaxy S21 FE 5G 128GB Lavender",
	"Samsung Galaxy S21 FE 5G 128GB Olive",
	"Samsung Galaxy S21 FE 5G 128GB White",
	"Samsung Galaxy XCover 5 Enterprise Edition",
	"Samsung Galaxy Z Flip3 5G 128 GB Phantom Black",
	"xiaomi 11 Lite 5G NE 128GB Truffle Black",
	"xiaomi 11T 5G 128GB Meteorite Gray",
	"Xiaomi 13 Lite 128GB - black",
	"Xiaomi 13 Lite 128GB - blue",
	"Xiaomi 13 Lite 128GB - pink",
	"xiaomi Mi 11 Lite 128 GB Boba Black",
	"xiaomi Redmi 9A 32 GB Aurora Green",
	"xiaomi Redmi 9C 128 GB Midnight Gray",
	"xiaomi Redmi Note 11 128GB Graphite Gray",
	"xiaomi Redmi Note 11 Pro 5G 128GB",
}

var melectronicsNamesExpected = []string{
	"Apple iPhone 11",
	"Apple iPhone 12",
	"Apple iPhone 12",
	"Apple iPhone 13",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Apple iPhone SE (2022)",
	"Fairphone 4",
	"Fairphone 4",
	"Google Pixel 6a",
	"Google Pixel 7",
	"Google Pixel 7",
	"Huawei Nova 9",
	"Huawei Nova 9",
	"Huawei Nova 9 SE",
	"Huawei Y6P",
	"Motorola Edge 30 Fusion",
	"Motorola Edge 30 Neo",
	"Motorola moto e22i",
	"Motorola moto g22",
	"Nokia G11",
	"Nokia G21",
	"Nokia G21",
	"Nokia G22",
	"NOKIA G22",
	"Nokia G60",
	"Nokia G60",
	"Nokia X30",
	"Nokia X30",
	"Nokia XR20",
	"Nothing Phone (1)",
	"Nothing Phone (1)",
	"Oppo A53s",
	"Oppo A54s",
	"Oppo A57s",
	"Oppo A57s",
	"Oppo A94",
	"Oppo A94",
	"Oppo A96",
	"Oppo A96",
	"Oppo Find X3 Lite",
	"Oppo Find X3 Neo",
	"Oppo Find X5 Lite",
	"Oppo Reno6",
	"Oppo Reno6 Pro",
	"Oppo Reno8",
	"Oppo Reno8 Lite",
	"Oppo Reno8 Lite",
	"Oppo Reno2",
	"Oppo Reno8",
	"Xiaomi Redmi Note 12",
	"Samsung Galaxy A12",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A52",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy Z Flip3",
	"xiaomi 11 Lite",
	"xiaomi 11T",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Lite",
	"xiaomi Mi 11 Lite",
	"xiaomi Redmi 9A",
	"xiaomi Redmi 9C",
	"xiaomi Redmi Note 11",
	"xiaomi Redmi Note 11 Pro",
}

func TestMelectronicsClean(t *testing.T) {
	for i, name := range melectronicsNames {
		if _name := shop.MelectronicsCleanFn(name); _name != melectronicsNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, melectronicsNamesExpected[i], name)
		}
	}
}
