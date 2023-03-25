package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var mediamarktNames = []string{
	"MOTOROLA Moto E20 - Smartphone (6.5 \", 32 GB, Graphite Grey)",
	"MOTOROLA Moto E22i - Smartphone (6.5 \", 32 GB, Graphite Grey)",
	"MOTOROLA Moto E32 - Smartphone (6.5 \", 64 GB, Misty Silver)",
	"MOTOROLA Moto E32 - Smartphone (6.5 \", 64 GB, Slate Grey)",
	"NOKIA C12 - Smartphone (6.3 \", 64 GB, Charcoal)",
	"NOKIA C21 Plus - Smartphone (6.517 \", 32 GB, Dark Cyan)",
	"NOKIA G11 - Smartphone (6.5 \", 32 GB, Charcoal)",
	"NOKIA G21 - Smartphone (6.5 \", 128 GB, Dusk)",
	"NOKIA G22 - Smartphone (6.52 \", 64 GB, Meteor Grey)",
	"OPPO A57s - Smartphone (6.56 \", 128 GB, Sky Blue)",
	"OPPO A57s - Smartphone (6.56 \", 128 GB, Starry Black)",
	"SAMSUNG Galaxy A13 4G (mit MediaTek CPU) - Smartphone (6.6 \", 128 GB, Black)",
	"SAMSUNG Galaxy A13 4G (mit MediaTek CPU) - Smartphone (6.6 \", 128 GB, Blue)",
	"SAMSUNG Galaxy A13 4G (mit MediaTek CPU) - Smartphone (6.6 \", 128 GB, White)",
	"WIKO POWER U10 - Smartphone (6.82 \", 32 GB, Carbon Blue)",
	"WIKO Y52 - Smartphone (5 \", 16 GB, Deep Blue)",
	"WIKO Y52 - Smartphone (5 \", 16 GB, Grau)",
	"XIAOMI Redmi 9A - Smartphone (6.53 \", 32 GB, Granite Grey)",
	"XIAOMI Redmi 9C - Smartphone (6.53 \", 128 GB, Midnight Grey)",
}

var mediamarktNamesExpected = []string{
	"MOTOROLA Moto E20",
	"MOTOROLA Moto E22i",
	"MOTOROLA Moto E32",
	"MOTOROLA Moto E32",
	"NOKIA C12",
	"NOKIA C21 Plus",
	"NOKIA G11",
	"NOKIA G21",
	"NOKIA G22",
	"OPPO A57s",
	"OPPO A57s",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"WIKO POWER U10",
	"WIKO Y52",
	"WIKO Y52",
	"XIAOMI Redmi 9A",
	"XIAOMI Redmi 9C",
}

func TestMediamarktClean(t *testing.T) {
	for i, name := range mediamarktNames {
		if _name := shop.MediamarktCleanFn(name); _name != mediamarktNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, mediamarktNamesExpected[i], name)
		}
	}
}
