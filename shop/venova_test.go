package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var venovaNames = []string{
	"Nokia - 7.2 16 cm Dual-SIM Android 9.0 4G USB Typ-C 4 GB 64 GB 3500 mAh, Schwarz",
	"Samsung - Galaxy A32 5G SM-A326B 16,5 cm 5000 mAh, Blau",
	"Samsung - Galaxy A32 5G SM-A326B 16,5 cm 5000 mAh, Weiss",
	"Samsung - Galaxy XCover 5 Enterprise Edition 13,5 cm 3000 mAh, Schwarz",
	"Samsung Galaxy M13 64GB Deep Green - Dual SIM 4GB RAM Smartphone",
	"Samsung Galaxy M13 64GB Light Blue - Dual SIM 4GB RAM Smartphone",
	"Samsung SM-A137F - Galaxy A13 Dual-SIM, 64GB, Weiss",
	"Samsung SM-A137FLBVEUE - Galaxy A13 Dual-SIM, 64GB, Blau",
	"Samsung SM-A137FZKUeUe - Galaxy A13 32GB, Handy, Schwarz",
	"Xiaomi 40-54-3225 - Redmi A1, 32GB, 2.0GB RAM, Light Green",
	"Xiaomi Poco M5 - Dual SIM 4GB RAM 64GB, grün",
	"Xiaomi Poco M5 - Dual SIM 4GB RAM 64GB, schwarz",
	"Xiaomi Poco M5S - Dual SIM 4GB RAM 128GB, blau",
	"Xiaomi Poco M5S - Dual SIM 4GB RAM 128GB, grau",
	"Xiaomi Redmi 10 C 128GB Graphite Grey - Dual SIM 4GB Smartphone",
	"Xiaomi Redmi 10 C 128GB Ocean Blue - Dual SIM 4GB Smartphone",
	"Xiaomi Redmi 10 C 64GB Graphite Grey - Dual SIM 4GB Smartphone",
	"Xiaomi Redmi 10 C 64GB Ocean Blue - Dual SIM 4GB Smartphone",
	"Xiaomi Redmi 9A 32GB Granite Grey - Dual SIM 2GB RAM Smartphone",
	"Xiaomi Redmi Note 10 128GB Carbon Grey - Dual SIM 4GB Smartphone, 2022",
	"Xiaomi Redmi Note 10 64GB Carbon Grey - Dual SIM 4GB Smartphone, 2022",
}

var venovaNamesExpected = []string{
	"Nokia 7.2",
	"Samsung Galaxy A32",
	"Samsung Galaxy A32",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy M13",
	"Samsung Galaxy M13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Xiaomi Redmi A1",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5s",
	"Xiaomi POCO M5s",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi Note 10",
	"Xiaomi Redmi Note 10",
}

func TestVenovaClean(t *testing.T) {
	for i, name := range venovaNames {
		if _name := shop.VenovaCleanFn(name); _name != venovaNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, venovaNamesExpected[i], name)
		}
	}
}
