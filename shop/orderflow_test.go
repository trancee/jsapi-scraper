package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var orderflowNames = []string{
	"Apple iPhone 11",
	"Apple iPhone 11 64GB Schwarz, Bildschirmdiagonale: 6.1 \"",
	"Apple iPhone SE 3. Gen. 64 GB Mitternacht, Bildschirmdiagonale",
	"HUAWEI Nova 9 Dual SIM",
	"HUAWEI P SMART 2021 MIDNIGHT BLACK 128GB/HMS/DS LTE/6.67 ANDRD",
	"NOKIA G11",
	"NOKIA G60 5G",
	"OPPO A57s 128 GB Starry Black, Bildschirmdiagonale: 6.56 \"",
	"OPPO A96 128 GB Sunset Blue, Bildschirmdiagonale: 6.59 \"",
	"OPPO Find X5 Lite 256 GB Hellblau, Bildschirmdiagonale: 6.43",
	"OPPO Reno8 256 GB Shimmer Gold, Bildschirmdiagonale: 6.43 \"",
	"OPPO Reno8 Lite 128 GB Cosmic Black, Bildschirmdiagonale: 6.4",
	"OPPO Reno8 Lite 128 GB Rainbow Spectrum, Bildschirmdiagonale",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13 128 GB CH Blue, Bildschirmdiagonale: 6.6",
	"Samsung Galaxy A33 5G 128 GB CH Awesome White",
	"Samsung Galaxy A33 5G 128 GB CH Enterprise Edition",
	"Samsung Galaxy A34 5G 128 GB CH Awesome Violet",
	"Samsung Galaxy A34 5G 256 GB CH Awesome Silver",
	"Samsung Galaxy A34 5G 256 GB CH Awesome Violet",
	"Samsung Galaxy A52s 5G",
	"Samsung Galaxy A54 5G 128 GB CH Awesome Graphite",
	"Samsung Galaxy A54 5G 256 GB CH Awesome Graphite",
	"Xiaomi Redmi 9A 32 GB Aurora Green, Bildschirmdiagonale: 6.53",
	"Xiaomi Redmi 9A 32 GB Granite Gray, Bildschirmdiagonale: 6.53",
	"Xiaomi Redmi 9C 128 GB Midnight Grey, Bildschirmdiagonale: 6.53",
}

var orderflowNamesExpected = []string{
	"Apple iPhone 11",
	"Apple iPhone 11",
	"Apple iPhone SE 3rd Gen",
	"HUAWEI Nova 9",
	"HUAWEI P SMART",
	"NOKIA G11",
	"NOKIA G60",
	"OPPO A57s",
	"OPPO A96",
	"OPPO Find X5 Lite",
	"OPPO Reno8",
	"OPPO Reno8 Lite",
	"OPPO Reno8 Lite",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A52s",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9C",
}

func TestOrderflowClean(t *testing.T) {
	for i, name := range orderflowNames {
		if _name := shop.OrderflowCleanFn(name); _name != orderflowNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, orderflowNamesExpected[i], name)
		}
	}
}