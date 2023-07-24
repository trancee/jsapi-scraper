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
	"Google Pixel 3a mit 64 GB Speicher und Android 12",
	"Handy leicht gebraucht WIKO Y50",
	"Honor 10 / Dual-SIM / 64GB / LTE / USB-C / Earphone Jack",
	"Huawei Ascend Y 530 Black",
	"Huawei P smart + zu verkaufen",
	"Huawei y5, inkl. Original Ladegerät und Handbuch",
	"Huawei Y5p Smartphone NEU und OVP",
	"I Phone 5 S",
	"I Phone SE 1. Generation 128 MB",
	"iPhone  6",
	"iPhone Apple 6S 32GB Gold - Guter Zustand",
	"iPhone X con vetro da sostituire",
	"Iphone SE 2016 (intakt, nur Touch Screen zersprungen)",
	"Iphone SE 2020/2022",
	"IPhone SE 64GB Silber mit GARANTIE",
	"iPhone SE in gutem Zustand",
	"iPhone Se prima generazione",
	"Iphone SE RED",
	"iPhone SE Roségold",
	"iphone se rose",
	"iPhone SE Rose Gold 64 GB 1. Generation",
	"iPhone SE mit Gebrauchsspuren",
	"iphone XR mit OVP",
	"IPHONE XS 64 BG",
	"Iphone XSMax",
	"ipohne 6 S, 64 GB",
	"Mi Xiaomi 10 Lite 5G",
	"Moto G5G - wie neu, unbenutztes Zubehör",
	"Motorola e 20",
	"Motorola Moto G 6 mit Dual-SIM",
	"Motorolla V8 RAZR2 OVP",
	"Neuwertiges Handy Redmi 9a",
	"One plus 8 pro 256GB 16GB",
	"Oppo A 1k",
	"ORIGINAL HUAWEI P40 MIT GOOGLE SERVICES",
	"Samsung A5-6",
	"Samsung A 51",
	"Samsung A13 garandieschein",
	"Samsung A3,A5,A7,A8 Läuft Einwandfrei",
	"Samsung A7 (2018)",
	"Samsung galaxy a 20 e",
	"Samsung Galaxy A3 in Lederetui mit Eingabestift",
	"Samsung Galaxy A3 2017 Occasion",
	"Sansung Galaxy A3 / 2017 goldig",
	"Samsung Galaxy A3/6 2016 schwarz",
	"Samsung Galaxy A32 5 G",
	"Samsung Galaxy a6 duas 32 GB",
	"samsung Galaxy a8  Duos 32 GB mit Packung Ladekabel",
	"SAMSUNG Galaxy FE20 5G",
	"Samsung Galaxy J1 J100H White Android Smartphone",
	"Samsung Galaxy Tablet A 6",
	"Samsung Galxy S3 mini 8GB",
	"Samsung GT-C3530 Chrome Silver - OVP - Ladegerät",
	"Samsung Mobile Phone - Galaxy J5 - Batteria rimovibile!!",
	"samsung s 10",
	"Samsung S duos dual sim",
	"Samsung S10 und Huawei P8",
	"Samsung s3 wie neu",
	"Samsung S6e Occassion Top Zustand",
	"Samsung S20 FE 5G",
	"Smartphone XIAOMI Redmi 9A ancora in Garanzia",
	"Sony XPERIA Handy Telefon",
	"Sony Xperia xz F8331",
	"Wiko Fever 4G (16Go/3Go) (neuf/neu)",
	"Wiko View3 , 64GB",
	"xiaomi mi 9 Smartphone",
	"Xiaomi mobile redmi 10 blau chrome handy",
	"XIAOMI Redmi 9A HD+ Smartphone Android 32GB 13MP Nero",
	"Xiaomi Redmi 9a in OVP und noch verschweisst",
	"Xiaomi Redmi Note5",
	"Zu verkaufen Honor 9Lite",
}

var tuttiNamesExpected = []string{
	"Apple iPhone 6 Plus",
	"Apple iPhone 6S",
	"Apple iPhone X",
	"Samsung Galaxy A40",
	"Google Pixel 2",
	"Google Pixel 3a",
	"WIKO Y50",
	"Honor 10",
	"Huawei Ascend Y530",
	"Huawei P smart+",
	"Huawei y5",
	"Huawei Y5p",
	"Apple iPhone 5S",
	"Apple iPhone SE (2016)",
	"Apple iPhone 6",
	"Apple iPhone 6S",
	"Apple iPhone X",
	"Apple iPhone SE",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone Se (2016)",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone se",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone XR",
	"Apple iPhone XS",
	"Apple iPhone XS Max",
	"Apple iPhone 6S",
	"Xiaomi Mi 10 Lite",
	"MOTOROLA Moto G",
	"Motorola Moto e20",
	"Motorola Moto G6",
	"Motorola V8 RAZR2",
	"Xiaomi Redmi 9a",
	"OnePlus 8 pro",
	"Oppo A1k",
	"HUAWEI P40",
	"Samsung Galaxy A5",
	"Samsung Galaxy A51",
	"Samsung Galaxy A13",
	"Samsung Galaxy A3",
	"Samsung Galaxy A7",
	"Samsung Galaxy a20e",
	"Samsung Galaxy A3",
	"Samsung Galaxy A3",
	"Samsung Galaxy A3",
	"Samsung Galaxy A3",
	"Samsung Galaxy A32",
	"Samsung Galaxy a6 duos",
	"samsung Galaxy a8 Duos",
	"SAMSUNG Galaxy S20 FE",
	"Samsung Galaxy J1",
	"Samsung Galaxy Tab A6",
	"Samsung Galaxy S3 mini",
	"Samsung GT-C3530",
	"Samsung Galaxy J5",
	"samsung Galaxy s10",
	"Samsung Galaxy S duos",
	"Samsung Galaxy S10",
	"Samsung Galaxy s3",
	"Samsung Galaxy S6e",
	"Samsung Galaxy S20 FE",
	"XIAOMI Redmi 9A",
	"Sony XPERIA",
	"Sony Xperia xz",
	"Wiko Fever",
	"Wiko View 3",
	"xiaomi mi 9",
	"Xiaomi redmi 10",
	"XIAOMI Redmi 9A",
	"Xiaomi Redmi 9a",
	"Xiaomi Redmi Note 5",
	"Honor 9 Lite",
}

func TestTuttiClean(t *testing.T) {
	for i, name := range tuttiNames {
		if _name := shop.TuttiCleanFn(name); _name != tuttiNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, tuttiNamesExpected[i], name)
		}
	}
}
