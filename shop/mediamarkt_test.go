package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var mediamarktNames = []string{
	"APPLE iPhone 11 (2020) - Smartphone (6.1 \", 64 GB, Black)",
	"APPLE iPhone 11 (2020) - Smartphone (6.1 \", 64 GB, White)",
	"MOTOROLA Edge 20 Lite - Smartphone (6.7 \", 128 GB, Lagoon Green)",
	"MOTOROLA Edge 30 Neo - Smartphone (6.28 \", 128 GB, Onyx Black)",
	"MOTOROLA Moto E20 - Smartphone (6.5 \", 32 GB, Graphite Grey)",
	"MOTOROLA Moto E22i - Smartphone (6.5 \", 32 GB, Graphite Grey)",
	"MOTOROLA Moto E32 - Smartphone (6.5 \", 64 GB, Misty Silver)",
	"MOTOROLA Moto E32 - Smartphone (6.5 \", 64 GB, Slate Grey)",
	"MOTOROLA Moto G 5G - Smartphone (6.7 \", 64 GB, Frosted Silver)",
	"MOTOROLA Moto G 5G - Smartphone (6.7 \", 64 GB, Volcanic Grey)",
	"MOTOROLA Moto G 5G Plus - Smartphone (6.7 \", 64 GB, Surfing Blue)",
	"MOTOROLA Moto G22 - Smartphone (6.5 \", 64 GB, Cosmic Black)",
	"MOTOROLA Moto G30 - Smartphone (6.5 \", 128 GB, Dark Pearl)",
	"MOTOROLA Moto G31 - Smartphone (6.4 \", 128 GB, Mineral Grey)",
	"MOTOROLA Moto G32 - Smartphone (6.5 \", 128 GB, Mineral Grey)",
	"MOTOROLA Moto G41 - Smartphone (6.4 \", 128 GB, Meteorite Black)",
	"MOTOROLA Moto G42 - Smartphone (6.4 \", 128 GB, Atlantic Green)",
	"MOTOROLA Moto G52 - Smartphone (6.6 \", 128 GB, Charcoal Grey)",
	"MOTOROLA Moto G52 - Smartphone (6.6 \", 128 GB, Porcelain White)",
	"MOTOROLA Moto G71 5G - Smartphone (6.4 \", 128 GB, Iron Black)",
	"MOTOROLA Moto G82 5G - Smartphone (6.6 \", 128 GB, Meteorite Gray)",
	"NOKIA C12 - Smartphone (6.3 \", 64 GB, Charcoal)",
	"NOKIA C21 Plus - Smartphone (6.517 \", 32 GB, Dark Cyan)",
	"NOKIA G11 - Smartphone (6.5 \", 32 GB, Charcoal)",
	"NOKIA G21 - Smartphone (6.5 \", 128 GB, Dusk)",
	"NOKIA G22 - Smartphone (6.52 \", 64 GB, Meteor Grey)",
	"NOKIA G50 - Smartphone (6.82 \", 128 GB, Midnight Sun)",
	"NOKIA G60 5G - Smartphone (6.58 \", 128 GB, Grau)",
	"NOKIA G60 5G - Smartphone (6.58 \", 128 GB, Schwarz)",
	"NOTHING phone (1) - Smartphone (6.55 \", 128 GB, Schwarz)",
	"ONE PLUS Nord CE 2 5G - Smartphone (6.43 \", 128 GB, Bahama Blue)",
	"ONE PLUS Nord CE 2 5G - Smartphone (6.43 \", 128 GB, Grey Mirror)",
	"OPPO A54 5G - Smartphone (6.5 \", 64 GB, Fantastic Purple)",
	"OPPO A57s - Smartphone (6.56 \", 128 GB, Sky Blue)",
	"OPPO A57s - Smartphone (6.56 \", 128 GB, Starry Black)",
	"OPPO A96 - Smartphone (6.59 \", 128 GB, Starry Black)",
	"OPPO A96 - Smartphone (6.59 \", 128 GB, Sunset Blue)",
	"OPPO Find X3 Neo - Smartphone (6.5 \", 256 GB, Galactic Silver)",
	"OPPO Find X3 Neo - Smartphone (6.5 \", 256 GB, Starlight Black)",
	"OPPO Reno8 5G - Smartphone (6.4 \", 256 GB, Shimmer Black)",
	"OPPO Reno8 5G - Smartphone (6.4 \", 256 GB, Shimmer Gold)",
	"OPPO Reno8 Lite 5G - Smartphone (6.43 \", 128 GB, Cosmic Black)",
	"OPPO Reno8 Lite 5G - Smartphone (6.43 \", 128 GB, Rainbow Spectrum)",
	"SAMSUNG Galaxy A13 4G (mit MediaTek CPU) - Smartphone (6.6 \", 128 GB, Black)",
	"SAMSUNG Galaxy A13 4G (mit MediaTek CPU) - Smartphone (6.6 \", 128 GB, Blue)",
	"SAMSUNG Galaxy A13 4G (mit MediaTek CPU) - Smartphone (6.6 \", 128 GB, White)",
	"SAMSUNG Galaxy A14 4G - Smartphone (6.6 \", 128 GB, Lime Green)",
	"SAMSUNG Galaxy A14 4G - Smartphone (6.6 \", 128 GB, Schwarz)",
	"SAMSUNG Galaxy A14 4G - Smartphone (6.6 \", 128 GB, Silber)",
	"SAMSUNG Galaxy A14 5G - Smartphone (6.6 \", 128 GB, Lime Green)",
	"SAMSUNG Galaxy A14 5G - Smartphone (6.6 \", 128 GB, Schwarz)",
	"SAMSUNG Galaxy A14 5G - Smartphone (6.6 \", 128 GB, Silber)",
	"SAMSUNG Galaxy A33 5G (EU) - Smartphone (6.4 \", 128 GB, Awesome Black)",
	"SAMSUNG Galaxy A33 5G (EU) - Smartphone (6.4 \", 128 GB, Awesome Blue)",
	"SAMSUNG Galaxy A33 5G - Smartphone (6.4 \", 128 GB, Awesome Black)",
	"SAMSUNG Galaxy A33 5G - Smartphone (6.4 \", 128 GB, Awesome Blue)",
	"SAMSUNG Galaxy A33 5G - Smartphone (6.4 \", 128 GB, Awesome Peach)",
	"SAMSUNG Galaxy A33 5G - Smartphone (6.4 \", 128 GB, Awesome White)",
	"SAMSUNG Galaxy A34 5G - Smartphone (6.6 \", 128 GB, Awesome Graphite)",
	"SAMSUNG Galaxy A34 5G - Smartphone (6.6 \", 128 GB, Awesome Lime)",
	"SAMSUNG Galaxy A34 5G - Smartphone (6.6 \", 128 GB, Awesome Silver)",
	"SAMSUNG Galaxy A34 5G - Smartphone (6.6 \", 128 GB, Awesome Violet)",
	"SAMSUNG Galaxy A34 5G - Smartphone (6.6 \", 256 GB, Awesome Graphite)",
	"SAMSUNG Galaxy A34 5G - Smartphone (6.6 \", 256 GB, Awesome Lime)",
	"SAMSUNG Galaxy A34 5G - Smartphone (6.6 \", 256 GB, Awesome Silver)",
	"SAMSUNG Galaxy A34 5G - Smartphone (6.6 \", 256 GB, Awesome Violet)",
	"SAMSUNG Galaxy A52s 5G - Smartphone (6.5 \", 128 GB, Awesome Violet)",
	"SAMSUNG Galaxy A52s 5G - Smartphone (6.5 \", 128 GB, Awesome White)",
	"SAMSUNG Galaxy A53 5G - Smartphone (6.5 \", 128 GB, Awesome Black)",
	"SAMSUNG Galaxy A53 5G - Smartphone (6.5 \", 128 GB, Awesome Blue)",
	"SAMSUNG Galaxy A53 5G - Smartphone (6.5 \", 128 GB, Awesome Peach)",
	"SAMSUNG Galaxy A53 5G - Smartphone (6.5 \", 128 GB, Awesome White)",
	"SAMSUNG Galaxy XCover 5 Enterprise Edition - Smartphone (5.3 \", 64 GB, Schwarz)",
	"WIKO POWER U10 - Smartphone (6.82 \", 32 GB, Carbon Blue)",
	"WIKO Y52 - Smartphone (5 \", 16 GB, Deep Blue)",
	"WIKO Y52 - Smartphone (5 \", 16 GB, Grau)",
	"XIAOMI 11 Lite 5G NE - Smartphone (6.55 \", 128 GB, Peach Pink)",
	"XIAOMI 11 Lite 5G NE - Smartphone (6.55 \", 128 GB, Truffle Black)",
	"XIAOMI 12 Lite - Smartphone (6.55 \", 128 GB, Black)",
	"XIAOMI 12 Lite - Smartphone (6.55 \", 128 GB, Lite Green)",
	"XIAOMI POCO M5 - Smartphone (6.58 \", 128 GB, Gelb)",
	"XIAOMI POCO M5 - Smartphone (6.58 \", 128 GB, Grün)",
	"XIAOMI POCO M5 - Smartphone (6.58 \", 128 GB, Schwarz)",
	"XIAOMI POCO M5 - Smartphone (6.58 \", 64 GB, Gelb)",
	"XIAOMI POCO M5 - Smartphone (6.58 \", 64 GB, Grün)",
	"XIAOMI POCO M5 - Smartphone (6.58 \", 64 GB, Schwarz)",
	"XIAOMI POCO M5s - Smartphone (6.43 \", 128 GB, Blau)",
	"XIAOMI POCO M5s - Smartphone (6.43 \", 128 GB, Grau)",
	"XIAOMI POCO M5s - Smartphone (6.43 \", 128 GB, Weiss)",
	"XIAOMI POCO M5s - Smartphone (6.43 \", 64 GB, Blau)",
	"XIAOMI POCO M5s - Smartphone (6.43 \", 64 GB, Grau)",
	"XIAOMI POCO M5s - Smartphone (6.43 \", 64 GB, Weiss)",
	"XIAOMI Redmi 10 2022 - Smartphone (6.5 \", 128 GB, Carbon Grey)",
	"XIAOMI Redmi 9A - Smartphone (6.53 \", 32 GB, Granite Grey)",
	"XIAOMI Redmi 9C - Smartphone (6.53 \", 128 GB, Midnight Grey)",
	"XIAOMI Redmi Note 10S - Smartphone (6.43 \", 128 GB, Onyx Grey)",
	"XIAOMI Redmi Note 11 - Smartphone (6.43 \", 128 GB, Graphite Grey)",
	"XIAOMI Redmi Note 11 Pro 4G - Smartphone (6.67 \", 128 GB, Graphite Grey)",
	"XIAOMI Redmi Note 11 Pro 4G - Smartphone (6.67 \", 128 GB, Polar White)",
	"XIAOMI Redmi Note 11 Pro 4G - Smartphone (6.67 \", 128 GB, Star Blue)",
	"XIAOMI Redmi Note 11 Pro 5G - Smartphone (6.67 \", 128 GB, Atlantic Blue)",
	"XIAOMI Redmi Note 11 Pro 5G - Smartphone (6.67 \", 128 GB, Graphite Grey)",
	"XIAOMI Redmi Note 11 Pro 5G - Smartphone (6.67 \", 128 GB, Polar White)",
	"XIAOMI Redmi Note 11S 4G - Smartphone (6.43 \", 128 GB, Graphite Grey)",
	"XIAOMI Redmi Note 11S 4G - Smartphone (6.43 \", 128 GB, Pearl White)",
	"XIAOMI Redmi Note 11S 4G - Smartphone (6.43 \", 128 GB, Twilight Blue)",
	"XIAOMI Redmi Note 12 4G - Smartphone (6.67 \", 128 GB, Ice Blue)",
	"XIAOMI Redmi Note 12 4G - Smartphone (6.67 \", 128 GB, Onyx Gray)",
	"XIAOMI Redmi Note 12 5G - Smartphone (6.67 \", 128 GB, Forest Green)",
	"XIAOMI Redmi Note 12 5G - Smartphone (6.67 \", 128 GB, Ice Blue)",
	"XIAOMI Redmi Note 12 Pro 5G - Smartphone (6.67 \", 128 GB, Midnight Black)",
}

var mediamarktNamesExpected = []string{
	"Apple iPhone 11",
	"Apple iPhone 11",
	"motorola edge 20 lite",
	"motorola edge 30 neo",
	"motorola moto e20",
	"motorola moto e22i",
	"motorola moto e32",
	"motorola moto e32",
	"motorola moto g",
	"motorola moto g",
	"motorola moto g",
	"motorola moto g22",
	"motorola moto g30",
	"motorola moto g31",
	"motorola moto g32",
	"motorola moto g41",
	"motorola moto g42",
	"motorola moto g52",
	"motorola moto g52",
	"motorola moto g71",
	"motorola moto g82",
	"Nokia C12",
	"Nokia C21 Plus",
	"Nokia G11",
	"Nokia G21",
	"Nokia G22",
	"Nokia G50",
	"Nokia G60",
	"Nokia G60",
	"Nothing Phone (1)",
	"OnePlus Nord CE 2",
	"OnePlus Nord CE 2",
	"OPPO A54",
	"OPPO A57s",
	"OPPO A57s",
	"OPPO A96",
	"OPPO A96",
	"OPPO Find X3 Neo",
	"OPPO Find X3 Neo",
	"OPPO Reno8",
	"OPPO Reno8",
	"OPPO Reno8 Lite",
	"OPPO Reno8 Lite",
	"Samsung Galaxy A13",
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
	"Samsung Galaxy A33",
	"Samsung Galaxy A33",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A34",
	"Samsung Galaxy A52s",
	"Samsung Galaxy A52s",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy XCover 5",
	"Wiko Power U10",
	"Wiko Y52",
	"Wiko Y52",
	"Xiaomi 11 Lite",
	"Xiaomi 11 Lite",
	"Xiaomi 12 Lite",
	"Xiaomi 12 Lite",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5",
	"Xiaomi POCO M5s",
	"Xiaomi POCO M5s",
	"Xiaomi POCO M5s",
	"Xiaomi POCO M5s",
	"Xiaomi POCO M5s",
	"Xiaomi POCO M5s",
	"Xiaomi Redmi 10 2022",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12 Pro",
}

func TestMediamarktClean(t *testing.T) {
	for i, name := range mediamarktNames {
		if _name := shop.MediamarktCleanFn(name); _name != mediamarktNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, mediamarktNamesExpected[i], name)
		}
	}
}
