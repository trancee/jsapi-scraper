package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var tuttiNames = []string{
	"2 Stuck  Iphone  7  128 GB",
	"Android Handy: INOI A63 (Neupreis 99CHF)",
	"Android Wiko Y80",
	"Apple iPhone SE | 32 GB | Space Gray | New Battery - Wie Neu",
	"Apple I Phone 6 Plus Gold 16GB",
	"Apple I Phone 6S Space gray 32GB",
	"Apple I Phone X 64GB silber",
	"Blackwew A80 2GB RAM 16GB ROM",
	"Fairphone 1 2 Stück",
	"Fairphone 3 in guten Zustand zu verkaufen",
	"Galaxy A40 (Sehr guter Zustand",
	"gebrauchtes Xiaomi 9C",
	"Gebrauchtes Xiaomi Redmi 7A",
	"Google Pixel 2 5\" 64GB",
	"Google Pixel 3a mit 64 GB Speicher und Android 12",
	"Handy leicht gebraucht WIKO Y50",
	"Handy Redmi Note 10 Onyx Gray- Achtung! mit Glassbruch",
	"Handy - Oppo A53s electric black, 128 GB",
	"Handy / Smartphone Samsung Galaxy A3",
	"Handy / Smartphone: Samsung GT-E3210",
	"Handy/ Smartphone Xiaomi Redmi A1",
	"Handy/Natel Huawei P20",
	"Handy XIAOMI Redmi 9 - Dual-SIM 6.5 Zoll. 4GB RAM / 64GM,OVP",
	"Honor 10 / Dual-SIM / 64GB / LTE / USB-C / Earphone Jack",
	"Honor 7 PLK-L01 16GB",
	"Huawei Ascend Y 530 Black",
	"Huawei P smart + zu verkaufen",
	"HUAWEI P20 LITE HANDY/ GOOGLE FÄHIG",
	"Huawei P20 Perfetto Garanzia",
	"HUAWEI P30 PRO GOOGLE SPERRE",
	"Huawei P-30 Lite 128 GB",
	"huawei P 30 lite",
	"HUAWEI P 40 LITE ,128 GB",
	"HUAWEI P9 FUNKTIONIERT EINWANDFREI.",
	"Huawei y5, inkl. Original Ladegerät und Handbuch",
	"Huawei Y5p Smartphone NEU und OVP",
	"Huawei Y6 - 2018 - Dual (2) SIM 16GB ATU-L21,",
	"I Phone 5 S",
	"I Phone 5 16G, Black",
	"I Phone 6 16 G Gold",
	"I Phone 8 64 G, silber",
	"I Phone SE 1. Generation 128 MB",
	"iPhone  6",
	"iPhone Apple 6S 32GB Gold - Guter Zustand",
	"iPhone X con vetro da sostituire",
	"Iphone X funktionstüchtig / Jedoch mit Streifen",
	"iPhone SE 1 Gen 64GB SpaceGray",
	"Iphone SE 1Gen  Rigenerato Garanzia",
	"Iphone SE 1 Generation 32GB",
	"iPhone SE 1. GEN",
	"iPhone SE 1. gen 64gb",
	"iPhone SE 2. Generation (2020) Rot",
	"Iphone SE 2016 (intakt, nur Touch Screen zersprungen)",
	"iPhone SE 2017 32GB",
	"Iphone SE 2020 Red Edition Rot !! Gebraucht !!",
	"Iphone SE 2020/2022",
	"IPhone SE 54 GB gold",
	"IPhone SE 64GB Silber mit GARANTIE",
	"iPhone SE A1723 2016, Space Gray, 128GB",
	"iPhone SE in gutem Zustand",
	"iPhone SE im sehr guten Zustand mit Box",
	"iPhone Se prima generazione",
	"Iphone SE RED",
	"iPhone SE Rosa 32GB",
	"iPhone SE Roségold",
	"iphone se rose",
	"iPhone SE Rose Gold 64 GB 1. Generation",
	"Iphone SE Top Zustand RESERVIERT",
	"iPhone SE mit Gebrauchsspuren",
	"iphone x 265 GB nur Verpackung",
	"iphone XR mit OVP",
	"iphone xr red",
	"Iphone XR rot 64 GB",
	"IPHONE XS 64 BG",
	"iphone xs SIMLOCKED",
	"Iphone XSMax",
	"IPhone 10 64 GB",
	"iPhone 11 Speicher 64 GB",
	"iPhone 12 Pro Max Gehäuse Original",
	"iPhone 12 Pro Max Original Front Kamera Module",
	"iPhone 12 Pro Max Original Kamera Module",
	"iPhone 12 Pro Max Original Taptic Engine mit Lautsprecher",
	"iPhone 15 Pro Max ESR MagSafe, Armorite Glas, AirTag Loops",
	"ipohne 6 S, 64 GB",
	"Mi Xiaomi 10 Lite 5G",
	"Mobiltelefon Emporia 4G",
	"Moto G5G - wie neu, unbenutztes Zubehör",
	"Motorola e 20",
	"Motorola Moto G 6 mit Dual-SIM",
	"Motorolla V8 RAZR2 OVP",
	"Neues Galaxy a 13 Samsung GÜNSTIG",
	"Neuwertiges Handy Redmi 9a",
	"One plus 8 pro 256GB 16GB",
	"Oppo A 1k",
	"Oppo A16 S in ottimo stato",
	"OPPO A54 Neu & ungeöffnet",
	"Oppo reno 8 Lite Glasschaden",
	"ORIGINAL HUAWEI P40 MIT GOOGLE SERVICES",
	"Pocophone F1 6GB/128GB 4G (mit neuer 4000mAh Batterie)",
	"Redmi 9C. 32 gb",
	"Redmi C9 128 GB",
	"Samsung A5-6",
	"Samsung A 51",
	"Samsung A13 garandieschein",
	"Samsung A2 Core Garanzia n.201",
	"Samsung A3,A5,A7,A8 Läuft Einwandfrei",
	"Samsung A7 (2018)",
	"Samsung galaxy a 20 e",
	"Samsung Galaxy A13 Neuwertiger zustand (Preis auf Anfrage)",
	"Samsung Galaxy A3 in Lederetui mit Eingabestift",
	"Samsung Galaxy A3 2017 Occasion",
	"Sansung Galaxy A3 / 2017 goldig",
	"Samsung Galaxy A3/6 2016 schwarz",
	"Samsung Galaxy A32 5 G",
	"Samsung Galaxy A40 * black * 64GB * OVP",
	"Samsung Galaxy a6 duas 32 GB",
	"Samsung Galaxy A7 Smartphone SM-A750FN/DS",
	"samsung Galaxy a8  Duos 32 GB mit Packung Ladekabel",
	"SAMSUNG Galaxy FE20 5G",
	"Samsung Galaxy J1 J100H White Android Smartphone",
	"Samsung Galaxy Tablet A 6",
	"Samsung Galaxy Xcovet 5",
	"Samsung Galxy S3 mini 8GB",
	"Samsung GT-C3530 Chrome Silver - OVP - Ladegerät",
	"Samsung Mobile Phone - Galaxy J5 - Batteria rimovibile!!",
	"samsung s 10",
	"Samsung Galaxy S Advance GT 19070",
	"Samsung S duos dual sim",
	"Samsung s10 - dual sim - 128 GB",
	"Samsung S10 und Huawei P8",
	"Samsung s3 wie neu",
	"Samsung S6e Occassion Top Zustand",
	"Samsung S20 FE 5G",
	"Smartfon Huawei P8 lite 5 Zoll",
	"Smartphone - Honor 7X",
	"Smartphone Google Pixel - sehr guter Zustand",
	"Smartphone Inoi a 151",
	"Smartphone XIAOMI Redmi 9A ancora in Garanzia",
	"Sony Xperia Handy einwandfrei",
	"Sony XPERIA Handy Telefon",
	"sony xperia modell f5321",
	"Sony Xperia xz F8331",
	"Sony XPERIA XZ2 4k kamera",
	"Sony Xperia Z3 Compact renoviert 4 Farben Gratisversand",
	"Telefono per anziani Doro 7080",
	"Vendo iphone 11 rosso",
	"VENDO SAMSUNG GALAXY NUOVO A14 4G, LUGANO",
	"Wiko Sunny 2 Dual SIM Neu nie benutzt",
	"Wiko Y60 Handy neuwertig",
	"Wiko Fever 4G (16Go/3Go) (neuf/neu)",
	"Wiko View3 , 64GB",
	"WIKO VIEW 4 LITE COMPLETO FUNZIONA BENISSIMO",
	"Xiaomi Mi9",
	"Xiaomi MI9T 128GB, GB Ram, 48Mpx",
	"Xiaomi Mi 9 Litle",
	"xiaomi mi 9 Smartphone",
	"Xiaomi mobile redmi 10 blau chrome handy",
	"Xiaomi Redmi 9 a originalverpackt neu und OVP",
	"Xiaomi Redmi 9 a originalverpackt und OVP",
	"Xiaomi Redmi 9A 2GB 32 GB Grau 2 Jahre Garantie",
	"XIAOMI Redmi 9A HD+ Smartphone Android 32GB 13MP Nero",
	"Xiaomi Redmi 9a in OVP und noch verschweisst",
	"XIAOMI Redmi 9AT NUOVO!",
	"Xiaomi Redmi Note 8 Pro Pearl White",
	"Xiaomi Redmi Note5",
	"Zu verkaufen Honor 9Lite",
}

var tuttiNamesExpected = []string{
	"Apple iPhone 7",
	"Inoi A63",
	"Wiko Y80",
	"Apple iPhone SE",
	"Apple iPhone 6 Plus",
	"Apple iPhone 6S",
	"Apple iPhone X",
	"Blackview A80",
	"Fairphone 1",
	"Fairphone 3",
	"Samsung Galaxy A40",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi 7A",
	"Google Pixel 2",
	"Google Pixel 3a",
	"Wiko Y50",
	"Xiaomi Redmi Note 10",
	"OPPO A53s",
	"Samsung Galaxy A3",
	"Samsung GT-E3210",
	"Xiaomi Redmi A1",
	"HUAWEI P20",
	"Xiaomi Redmi 9",
	"HONOR 10",
	"HONOR 7",
	"HUAWEI Ascend Y530",
	"HUAWEI P smart+",
	"HUAWEI P20 lite",
	"HUAWEI P20",
	"HUAWEI P30 Pro",
	"HUAWEI P30 lite",
	"HUAWEI P30 lite",
	"HUAWEI P40 lite",
	"HUAWEI P9",
	"HUAWEI Y5",
	"HUAWEI Y5p",
	"HUAWEI Y6",
	"Apple iPhone 5S",
	"Apple iPhone 5",
	"Apple iPhone 6",
	"Apple iPhone 8",
	"Apple iPhone SE (2016)",
	"Apple iPhone 6",
	"Apple iPhone 6S",
	"Apple iPhone X",
	"Apple iPhone X",
	"Apple iPhone SE (2016)",
	"Apple iPhone SE (2016)",
	"Apple iPhone SE (2016)",
	"Apple iPhone SE (2016)",
	"Apple iPhone SE (2016)",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE (2016)",
	"Apple iPhone SE (2016)",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE (2016)",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE (2016)",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone SE",
	"Apple iPhone X",
	"Apple iPhone XR",
	"Apple iPhone XR",
	"Apple iPhone XR",
	"Apple iPhone XS",
	"Apple iPhone XS",
	"Apple iPhone XS Max",
	"Apple iPhone X",
	"Apple iPhone 11",
	"Apple iPhone 12 Pro Max",
	"Apple iPhone 12 Pro Max",
	"Apple iPhone 12 Pro Max",
	"Apple iPhone 12 Pro Max",
	"Apple iPhone 15 Pro Max",
	"Apple iPhone 6S",
	"Xiaomi Mi 10 Lite",
	"emporia",
	"motorola moto g",
	"motorola moto e20",
	"motorola moto g6",
	"motorola v8 razr2",
	"Samsung Galaxy A13",
	"Xiaomi Redmi 9A",
	"OnePlus 8 Pro",
	"OPPO A1k",
	"OPPO A16s",
	"OPPO A54",
	"OPPO Reno8 Lite",
	"HUAWEI P40",
	"Xiaomi POCO F1",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi 9C",
	"Samsung Galaxy A5",
	"Samsung Galaxy A51",
	"Samsung Galaxy A13",
	"Samsung Galaxy A2 Core",
	"Samsung Galaxy A3",
	"Samsung Galaxy A7",
	"Samsung Galaxy A20e",
	"Samsung Galaxy A13",
	"Samsung Galaxy A3",
	"Samsung Galaxy A3",
	"Samsung Galaxy A3",
	"Samsung Galaxy A3",
	"Samsung Galaxy A32",
	"Samsung Galaxy A40",
	"Samsung Galaxy A6 Duos",
	"Samsung Galaxy A7",
	"Samsung Galaxy A8 Duos",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy J1",
	"Samsung Galaxy Tab A6",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy S3 mini",
	"Samsung GT-C3530",
	"Samsung Galaxy J5",
	"Samsung Galaxy S10",
	"Samsung Galaxy S Advance",
	"Samsung Galaxy S Duos",
	"Samsung Galaxy S10",
	"Samsung Galaxy S10",
	"Samsung Galaxy S3",
	"Samsung Galaxy S6e",
	"Samsung Galaxy S20 FE",
	"HUAWEI P8 lite",
	"HONOR 7X",
	"Google Pixel",
	"Inoi A151",
	"Xiaomi Redmi 9A",
	"Sony Xperia",
	"Sony Xperia",
	"Sony Xperia",
	"Sony Xperia XZ",
	"Sony Xperia XZ2",
	"Sony Xperia Z3 Compact",
	"Doro 7080",
	"Apple iPhone 11",
	"Samsung Galaxy A14",
	"Wiko Sunny 2",
	"Wiko Y60",
	"Wiko Fever",
	"Wiko View 3",
	"Wiko View 4 Lite",
	"Xiaomi Mi 9",
	"Xiaomi Mi 9T",
	"Xiaomi Mi 9 Lite",
	"Xiaomi Mi 9",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9AT",
	"Xiaomi Redmi Note 8 Pro",
	"Xiaomi Redmi Note 5",
	"HONOR 9 Lite",
}

func TestTuttiClean(t *testing.T) {
	for i, name := range tuttiNames {
		if _name := shop.TuttiCleanFn(name); _name != tuttiNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, tuttiNamesExpected[i], name)
		}
	}
}
