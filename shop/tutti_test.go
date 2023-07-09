package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var tuttiNames = []string{
	"Apple I Phone 6 Plus Gold 16GB",
	"Apple I Phone 6S Space gray 32GB",
	"Apple I Phone X 64GB silber",
	"Galaxy A40 (Sehr guter Zustand",
	"Google Pixel 2 5\" 64GB",
	"Handy leicht gebraucht WIKO Y50",
	"Honor 10 / Dual-SIM / 64GB / LTE / USB-C / Earphone Jack",
	"Huawei Ascend Y 530 Black",
	"Huawei P smart + zu verkaufen",
	"Huawei y5, inkl. Original Ladegerät und Handbuch",
	"Huawei Y5p Smartphone NEU und OVP",
	"I Phone 5 S",
	"iPhone  6",
	"iPhone Apple 6S 32GB Gold - Guter Zustand",
	"Iphone SE 2016 (intakt, nur Touch Screen zersprungen)",
	"Iphone SE 2020/2022",
	"IPhone SE 64GB Silber mit GARANTIE",
	"iPhone SE in gutem Zustand",
	"iphone XR mit OVP",
	"Iphone XSMax",
	"ipohne 6 S, 64 GB",
	"Mi Xiaomi 10 Lite 5G",
	"Motorola e 20",
	"Motorolla V8 RAZR2 OVP",
	"Neuwertiges Handy Redmi 9a",
	"One plus 8 pro 256GB 16GB",
	"Oppo A 1k",
	"ORIGINAL HUAWEI P40 MIT GOOGLE SERVICES",
	"Samsung A5-6",
	"Samsung A 51",
	"Samsung A3,A5,A7,A8 Läuft Einwandfrei",
	"Samsung A7 (2018)",
	"Samsung galaxy a 20 e",
	"Samsung Galaxy A3 in Lederetui mit Eingabestift",
	"Samsung Galaxy A3 2017 Occasion",
	"Sansung Galaxy A3 / 2017 goldig",
	"Samsung Galaxy A3/6 2016 schwarz",
	"Samsung Galaxy a6 duas 32 GB",
	"samsung Galaxy a8  Duos 32 GB mit Packung Ladekabel",
	"SAMSUNG Galaxy FE20 5G",
	"Samsung Galaxy J1 J100H White Android Smartphone",
	"Samsung Galxy S3 mini 8GB",
	"Samsung GT-C3530 Chrome Silver - OVP - Ladegerät",
	"Samsung Mobile Phone - Galaxy J5 - Batteria rimovibile!!",
	"samsung s 10",
	"Samsung S duos dual sim",
	"Samsung S20 FE 5G",
	"Smartphone XIAOMI Redmi 9A ancora in Garanzia",
	"Sony Xperia xz F8331",
	"Wiko Fever 4G (16Go/3Go) (neuf/neu)",
	"Xiaomi mobile redmi 10 blau chrome handy",
	"XIAOMI Redmi 9A HD+ Smartphone Android 32GB 13MP Nero",
}

var tuttiNamesExpected = []string{
	"Apple iPhone 6 Plus",
	"Apple iPhone 6S",
	"Apple iPhone X",
	"Samsung Galaxy A40",
	"Google Pixel 2",
	"WIKO Y50",
	"Honor 10",
	"Huawei Ascend Y530",
	"Huawei P smart+",
	"Huawei y5",
	"Huawei Y5p",
	"Apple iPhone 5S",
	"Apple iPhone 6",
	"Apple iPhone 6S",
	"Apple iPhone SE",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone XR",
	"Apple iPhone XS Max",
	"Apple iPhone 6S",
	"Xiaomi Mi 10 Lite",
	"Motorola Moto e20",
	"Motorola V8 RAZR2",
	"Xiaomi Redmi 9a",
	"OnePlus 8 pro",
	"Oppo A1k",
	"HUAWEI P40",
	"Samsung Galaxy A5",
	"Samsung Galaxy A51",
	"Samsung Galaxy A3",
	"Samsung Galaxy A7",
	"Samsung Galaxy a20e",
	"Samsung Galaxy A3",
	"Samsung Galaxy A3",
	"Samsung Galaxy A3",
	"Samsung Galaxy A3",
	"Samsung Galaxy a6 duos",
	"samsung Galaxy a8 Duos",
	"SAMSUNG Galaxy S20 FE",
	"Samsung Galaxy J1",
	"Samsung Galaxy S3 mini",
	"Samsung GT-C3530",
	"Samsung Galaxy J5",
	"samsung Galaxy s10",
	"Samsung Galaxy S duos",
	"Samsung Galaxy S20 FE",
	"XIAOMI Redmi 9A",
	"Sony Xperia xz",
	"Wiko Fever",
	"Xiaomi redmi 10",
	"XIAOMI Redmi 9A",
}

func TestTuttiClean(t *testing.T) {
	for i, name := range tuttiNames {
		if _name := shop.TuttiCleanFn(name); _name != tuttiNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, tuttiNamesExpected[i], name)
		}
	}
}
