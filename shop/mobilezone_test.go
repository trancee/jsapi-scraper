package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var mobilezoneNames = []string{
	"Apple iPhone 11 64GB Black",
	"Apple iPhone 12 5G 64GB Black",
	"Apple iPhone 12 mini 5G 256GB Purple",
	"Apple iPhone 13 5G 128GB Midnight",
	"Apple iPhone 13 mini 5G 128GB Midnight",
	"Apple iPhone 14 5G 128GB Midnight",
	"Apple iPhone 14 Plus 5G 128GB Midnight",
	"Apple iPhone 14 Pro 5G 128GB Space Black",
	"Apple iPhone 14 Pro Max 5G 128GB Space Black",
	"Apple iPhone SE 2020 128GB Red",
	"Apple iPhone SE 5G 2022 64GB Starlight",
	"Fairphone 4 128GB 5G Grey Dual-SIM",
	"Google Pixel 7 5G 128GB Obsidian",
	"Google Pixel 7 Pro 5G 128GB Obsidian",
	"Motorola EDGE20 256GB 5G Frosted Grey",
	"Motorola EDGE20 Lite 128GB 5G Electric Graphite",
	"Motorola EDGE20 Pro 256GB 5G Midnight Blue",
	"Motorola moto e20 32GB Graphite Grey Dual-SIM",
	"Motorola moto g22 64GB Black Dual-SIM",
	"Motorola moto g31 128GB Grey Dual-SIM",
	"Motorola moto g52 128GB Charcoal Grey Dual-SIM",
	"Nokia 8.3 128GB 5G Polar Night Dual-Sim",
	"Nokia X20 128GB Blue Dual-SIM",
	"Nokia X30 128GB Cloudy Blue Dual-SIM",
	"Nokia XR20 128GB 5G Granite Grey",
	"Nothing Phone 1 A063 256GB Black",
	"OnePlus 11 5G 128GB Titan Black",
	"Oppo A96 128GB Starry Black Dual-SIM",
	"Oppo Find N2 Flip 5G 256GB Astral Black Dual-SIM",
	"Oppo Find X3 Lite 5G 128GB Starry Black Dual-SIM",
	"Oppo Find X3 Neo 5G 256GB Starlight Black Dual-SIM",
	"Oppo Find X3 Pro 5G 256GB Gloss Black Dual-SIM",
	"Oppo Find X5 5G 256GB Black",
	"Oppo Find X5 Lite 5G 256GB Starry Black",
	"Oppo Find X5 Pro 5G 256GB Glaze Black Dual-SIM",
	"OPPO Reno4 Pro 256GB 5G Space Black Dual-SIM",
	"Oppo Reno6 5G 128GB Stellar Black Dual-SIM",
	"Oppo Reno6 Pro 5G 256GB Lunar Grey Dual-SIM",
	"Oppo Reno8 5G 256GB Shimmer Black Dual-SIM",
	"Oppo Reno8 Lite 5G 128GB Cosmic Black Dual-SIM",
	"Oppo Reno8 Pro 5G 256GB Glazed Black Dual-SIM",
	"Samsung Galaxy A13 4G 128GB White Dual-SIM",
	"Samsung Galaxy A14 4G 128GB Black",
	"Samsung Galaxy A33 5G 128GB Awesome Black Dual-SIM",
	"Samsung Galaxy A34 5G 128GB Awesome Graphite",
	"Samsung Galaxy A53 5G 128GB Awesome Black Dual-SIM",
	"Samsung Galaxy A54 5G 128GB Awesome Graphite",
	"Samsung Galaxy Note 10 256GB Aura Glow Dual-SIM",
	"Samsung Galaxy S21 5G 256GB Phantom Gray Dual-SIM",
	"Samsung Galaxy S21 FE 5G 128GB Graphite Dual-SIM",
	"Samsung Galaxy S22 5G 256GB Black Dual-SIM",
	"Samsung Galaxy S22 Ultra 5G 256GB Black Dual-SIM",
	"Samsung Galaxy S22+ 5G 256GB Black Dual-SIM",
	"Samsung Galaxy S23 5G 128GB Phantom Black Dual-SIM",
	"Samsung Galaxy S23 Ultra 5G 256GB Phantom Black Dual-SIM",
	"Samsung Galaxy S23+ 5G 256GB Phantom Black Dual-SIM",
	"Samsung Galaxy Xcover5 64GB Black Dual-SIM EE",
	"Samsung Galaxy Z Flip4 5G 128GB Graphite",
	"Samsung Galaxy Z Fold3 512GB 5G Silver",
	"Samsung Galaxy Z Fold4 5G 256GB Phantom Black",
	"Wiko Power U20 64GB Navy Blue",
	"Xiaomi 12 256GB 5G Gray Dual-SIM",
	"Xiaomi 12 Pro 256GB 5G Gray Dual-SIM",
	"Xiaomi 12T 5G 256GB Black Dual-SIM",
	"Xiaomi 12T Pro 5G 256GB Black Dual-SIM",
	"Xiaomi 13 256GB 5G Dual-SIM Black",
	"Xiaomi 13 Lite 128GB 5G Dual-SIM Black",
	"Xiaomi 13 Pro 256GB 5G Dual-SIM Ceramic Black",
	"Xiaomi Redmi Note 11 128GB Graphite Gray Dual-SIM",
	"Xiaomi Redmi Note 12 Pro 5G 128GB Black Dual-SIM",
}

var mobilezoneNamesExpected = []string{
	"Apple iPhone 11",
	"Apple iPhone 12",
	"Apple iPhone 12 mini",
	"Apple iPhone 13",
	"Apple iPhone 13 mini",
	"Apple iPhone 14",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro Max",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE (2022)",
	"Fairphone 4",
	"Google Pixel 7",
	"Google Pixel 7 Pro",
	"motorola edge 20",
	"motorola edge 20 lite",
	"motorola edge 20 pro",
	"motorola moto e20",
	"motorola moto g22",
	"motorola moto g31",
	"motorola moto g52",
	"Nokia 8.3",
	"Nokia X20",
	"Nokia X30",
	"Nokia XR20",
	"Nothing Phone (1)",
	"OnePlus 11",
	"OPPO A96",
	"OPPO Find N2 Flip",
	"OPPO Find X3 Lite",
	"OPPO Find X3 Neo",
	"OPPO Find X3 Pro",
	"OPPO Find X5",
	"OPPO Find X5 Lite",
	"OPPO Find X5 Pro",
	"OPPO Reno4 Pro",
	"OPPO Reno6",
	"OPPO Reno6 Pro",
	"OPPO Reno8",
	"OPPO Reno8 Lite",
	"OPPO Reno8 Pro",
	"Samsung Galaxy A13",
	"Samsung Galaxy A14",
	"Samsung Galaxy A33",
	"Samsung Galaxy A34",
	"Samsung Galaxy A53",
	"Samsung Galaxy A54",
	"Samsung Galaxy Note 10",
	"Samsung Galaxy S21",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22 Ultra",
	"Samsung Galaxy S22+",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23+",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy Z Flip 4",
	"Samsung Galaxy Z Fold 3",
	"Samsung Galaxy Z Fold 4",
	"Wiko Power U20",
	"Xiaomi 12",
	"Xiaomi 12 Pro",
	"Xiaomi 12T",
	"Xiaomi 12T Pro",
	"Xiaomi 13",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Pro",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 12 Pro",
}

func TestMobileZoneClean(t *testing.T) {
	for i, name := range mobilezoneNames {
		if _name := shop.MobileZoneCleanFn(name); _name != mobilezoneNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, mobilezoneNamesExpected[i], name)
		}
	}
}
