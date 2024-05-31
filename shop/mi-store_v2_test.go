package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var miStoreV2Names = []string{
	"*** BUNDLE ***  Xiaomi 13T Pro 16GB/1024GB inkl. Xiaomi Watch S1 Active",
	"*** BUNDLE *** Xiaomi 13T 8GB/256GB inkl. Redmi Watch 2 Lite GL",
	"*** BUNDLE *** Xiaomi 13T Pro 12GB/512GB inkl. Xiaomi Watch S1 Active",
	"***BUNDLE*** Xiaomi 13T 8GB/256GB inkl. Redmi Watch 2 Lite GL",
	"Redmi 13 | 6GB/128GB | Schwarz",
	"Xiaomi 12 Lite 8GB/128GB Smartphone",
	"Xiaomi 13 Lite 8GB/128GB Smartphone",
	"Xiaomi POCO F4 8GB/256GB Smartphone",
	"Xiaomi POCO M5 4GB/128GB Smartphone",
	"Xiaomi POCO M5 4GB/64GB Smartphone",
	"Xiaomi POCO M5s 4GB/64GB Smartphone",
	"Xiaomi POCO X5 5G 6GB/128GB Smartphone",
	"Xiaomi POCO X5 5G 8GB/256GB Smartphone",
	"Xiaomi Redmi 10 2022 4GB/128GB Smartphone",
	"Xiaomi Redmi 12 4G 4GB/128GB Smartphone",
	"Xiaomi Redmi 12 4G 8GB/256GB Smartphone",
	"Xiaomi Redmi 12C 3GB/64GB Smartphone",
	"Xiaomi Redmi 12C 4GB/128GB Smartphone",
	"Xiaomi Redmi 9A 2GB/32GB Smartphone",
	"Xiaomi Redmi A2 2GB/32GB Smartphone",
	"Xiaomi Redmi Note 11 4GB/128GB Smartphone",
	"Xiaomi Redmi Note 12 4G 4GB/128GB Smartphone",
	"Xiaomi Redmi Note 12 4G 4GB/64GB Smartphone",
	"Xiaomi Redmi Note 12 5G 4GB/128GB Smartphone",
	"Xiaomi Redmi Note 12 Pro 6GB/128GB Smartphone",
	"Xiaomi Redmi Note 12 Pro 8GB/256GB Smartphone",
}

var miStoreV2NamesExpected = []string{
	"Xiaomi 13T Pro",
	"Xiaomi 13T",
	"Xiaomi 13T Pro",
	"Xiaomi 13T",
	"Xiaomi Redmi 13",
	"Xiaomi 12 Lite",
	"Xiaomi 13 Lite",
	"Xiaomi POCO F4",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5s",
	"Xiaomi POCO X5",
	"Xiaomi POCO X5",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 12",
	"Xiaomi Redmi 12",
	"Xiaomi Redmi 12C",
	"Xiaomi Redmi 12C",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi A2",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi Redmi Note 12 Pro",
}

func TestMiStoreV2Clean(t *testing.T) {
	for i, name := range miStoreV2Names {
		if _name := shop.MiStoreV2CleanFn(name); _name != miStoreV2NamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, miStoreV2NamesExpected[i], name)
		}
	}
}
