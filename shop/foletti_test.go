package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var folettiNames = []string{
	"Blackview BV4900 Pro 14.5 cm (5.7) Dual SIM Android 10.0 4G Micro-USB 4 GB 64 GB 5580 mAh Black",
	"Blackview BV5200 5180 mAh 4/32 GB Green smartphone",
	"Motorola Moto E 13",
	"Motorola Moto E 22",
	"Motorola Moto E e32s 16.5 cm (6.5) Dual SIM Android 12 4G USB Type-C 4 GB 64 GB 5000 mAh Silver",
	"Motorola Moto E20 grau",
	"MOTOROLA MOTO E22 4/64GB ASTRO BLACK",
	"Motorola moto g22",
	"Motorola Moto G31 sterling blue",
	"Motorola XT2173-3 Moto G31 128GB Grey 6.4 EU (4GB) Android",
	"Motorola XT2173-3 moto g31 Dual Sim 4+64GB mineral grey DE",
	"Motorola XT2229-2 moto e32s Dual Sim 3+32GB  slate grey DE",
	"Motorola XT2239-7 moto e22 Dual Sim 3+32GB  astro black DE",
	"Nokia G11 Charcoal, 3+32GB",
	"Nokia G11 Ice, 3+32GB",
	"Nokia G21 Blue, 4+64GB",
	"Realme C30 bamboo green              3+32GB",
	"Realme C30 denim black               3+32GB",
	"Realme C30 lake blue                 3+32GB",
	"Realme C33 night sea 4+64GB",
	"Realme C33 sandy gold 4+64GB",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A04s SM-A047F/DSN",
	"Samsung Galaxy A13 32GB Black",
	"Samsung Galaxy A13 SM-A137FZWUEUE smartphone",
	"Samsung SM-A137F Galaxy A13 Dual Sim 3+32GB black EU",
	"vivo Y01 elegant black",
	"Xiaomi Redmi 10 5G aurora green 4GB+64GB",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C 17 cm (6.71) Dual SIM Android 11 4G USB Type-C 4 GB 128 GB 5000 mAh Blue",
	"Xiaomi Redmi 10C Dual Sim 4+128GB graphite grey DE",
	"Xiaomi Redmi 10C Dual Sim 4+128GB mint green DE",
	"Xiaomi Redmi 10C Dual Sim 4+128GB ocean blue DE",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9AT 32GB DS Grey 6.5 EU",
	"Xiaomi Redmi 9AT Dual Sim 2+32GB glacial blue DE",
	"Xiaomi Redmi 9C NFC 16.6 cm (6.53) Hybrid Dual SIM 4G Micro-USB 3 GB 64 GB 5000 mAh Green",
	"Xiaomi Redmi A1",
	"ZTE Blade A31 Lite grau",
	"ZTE Blade A72 blau",
	"ZTE Blade A72 grau",
}

var folettiNamesExpected = []string{
	"Blackview BV4900 Pro",
	"Blackview BV5200",
	"Motorola Moto E13",
	"Motorola Moto E22",
	"Motorola Moto e32s",
	"Motorola Moto E20",
	"MOTOROLA MOTO E22",
	"Motorola moto g22",
	"Motorola Moto G31",
	"Motorola Moto G31",
	"Motorola moto g31",
	"Motorola moto e32s",
	"Motorola moto e22",
	"Nokia G11",
	"Nokia G11",
	"Nokia G21",
	"Realme C30",
	"Realme C30",
	"Realme C30",
	"Realme C33",
	"Realme C33",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A04s",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"vivo Y01",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi A1",
	"ZTE Blade A31 Lite",
	"ZTE Blade A72",
	"ZTE Blade A72",
}

func TestFolettiClean(t *testing.T) {
	for i, name := range folettiNames {
		if _name := shop.FolettiCleanFn(name); _name != folettiNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, folettiNamesExpected[i], name)
		}
	}
}
