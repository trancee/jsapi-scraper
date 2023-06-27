package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var manorNames = []string{
	"Blackview  BV4900 14,5 cm (5.7 Zoll) Dual-SIM Android 10.0 4G Mikro-USB 3 GB 32 GB 5580 mAh Schwarz",
	"Blackview  BV6300 Pro 14,5 cm (5.7 Zoll) Dual-SIM Android 10.0 4G 6 GB 128 GB 4380 mAh Schwarz",
	"HUAWEI  nova 9 SE 17,2 cm (6.78 Zoll) Dual-SIM EMUI 12.0 4G USB Typ-C 8 GB 128 GB 4000 mAh Blau",
	"MOTOROLA Moto E E22i 16,5 cm (6.5 Zoll) Dual-SIM Android 12 Go Edition 4G USB Typ-C 2 GB 32 GB 4020 mAh Grau",
	"MOTOROLA  Moto G g32 16,5 cm (6.5 Zoll) Dual-SIM Android 12 4G USB Typ-C 4 GB 128 GB 5000 mAh Grau",
	"MOTOROLA  Moto G51 XT2171-2 Dual 5G 128 GB Horizontblau (4 GB)",
	"NOKIA   G60 5G 16,7 cm (6.58 Zoll) Dual-SIM Android 12 USB Typ-C 4 GB 128 GB Schwarz",
	"NOKIA  G21 16,5 cm (6.5 Zoll) Dual-SIM Android 11 4G USB Typ-C 4 GB 64 GB 5050 mAh Blau",
	"NOKIA  NOK G21 4+64GB purple 16,5 cm (6.5 Zoll) Dual-SIM Android 11 4G USB Typ-C 4 GB 5050 mAh Violett",
	"NOKIA G11, 6.51'' Smartphone",
	"NOKIA G21, 6.51'' Smartphone",
	"NOKIA G50 5G, 6.82'' Smartphone",
	"OPPO   A54 6.5&quot; Dual SIM 64GB 5G Smartphone Lila",
	"OPPO   A96 16,7 cm (6.59 Zoll) Dual-SIM Android 11 4G USB Typ-C 8 GB 128 GB 5000 mAh Schwarz",
	"OPPO  A57s 16,7 cm (6.56 Zoll) Dual-SIM Android 12 4G USB Typ-C 4 GB 128 GB 5000 mAh Blau",
	"OPPO  A57s 16,7 cm (6.56 Zoll) Dual-SIM Android 12 4G USB Typ-C 4 GB 128 GB 5000 mAh Schwarz",
	"OPPO OPPO Reno 8 Lite 16,3 cm (6.43 Zoll) Dual-SIM Android 11 5G USB Typ-C 8 GB 128 GB 4500 mAh Schwarz",
	"OPPO Oppo A74 6,5&quot; Dual-SIM 128GB 5G Smartphone Schwarz",
	"SAMSUNG   Galaxy A02 Dual A022fd 32GB Rot (3GB)",
	"SAMSUNG   Galaxy A02 Dual A022fd 64GB Grau (3GB)",
	"SAMSUNG   Galaxy A02 Dual A022GD 32GB Blau (2GB)",
	"SAMSUNG   Galaxy A02 Dual A022GD 32GB Schwarz (2GB)",
	"SAMSUNG   Galaxy A03 Dual A035fd 64GB Rot (4GB)",
	"SAMSUNG   Galaxy A03 Dual A035fd 64GB Schwarz (4GB)",
	"SAMSUNG   Galaxy A03S Dual A037fd 64GB Blau (4GB)",
	"SAMSUNG   Galaxy A03S Dual A037fd 64GB Schwarz (4GB)",
	"SAMSUNG   Galaxy A12 Dual A127FD 128GB Weiß (6GB)",
	"SAMSUNG   Galaxy A12 Dual A127FD 64GB Schwarz (4GB)",
	"SAMSUNG   Galaxy A13 Dual A135FD 128GB  (4GB)",
	"SAMSUNG   Galaxy A13 Dual A135FD 128GB Blau (4GB)",
	"SAMSUNG   Galaxy A13 Dual A135FD 128GB Weiß (4GB)",
	"SAMSUNG   Galaxy A13 Dual A135FD 64GB Blau (4GB)",
	"SAMSUNG   Galaxy A22 Dual A225F 4G 128GB Violett (4GB)",
	"SAMSUNG   Galaxy A23 Dual A235FD 128GB Blau (4GB)",
	"SAMSUNG   Galaxy A23 Dual A235FD 128GB Schwarz (4GB)",
	"SAMSUNG   Galaxy A23 Dual A235FD 128GB Schwarz (6GB)",
	"SAMSUNG   Galaxy A23 Dual A235FD 128GB Weiß (4GB)",
	"SAMSUNG   Galaxy A32 Dual A325FD 4G 128 GB Schwarz (6 GB)",
	"SAMSUNG   Galaxy M32 Dual M325FD 4G 128GB Schwarz (6GB)",
	"SAMSUNG  Galaxy A13 16,8 cm (6.6 Zoll) Dual-SIM Android 12 4G USB Typ-C 4 GB 128 GB 5000 mAh Hellblau",
	"SAMSUNG  Galaxy A13 16,8 cm (6.6 Zoll) Dual-SIM Android 12 4G USB Typ-C 4 GB 128 GB 5000 mAh Schwarz",
	"SAMSUNG  Galaxy A13 16,8 cm (6.6 Zoll) Dual-SIM Android 12 4G USB Typ-C 4 GB 128 GB 5000 mAh Weiß",
	"SAMSUNG  Galaxy A13 16,8 cm (6.6 Zoll) Hybride Dual-SIM 4G USB Typ-C 4 GB 128 GB 5000 mAh Hellblau",
	"SAMSUNG  Galaxy A13 16,8 cm (6.6 Zoll) Hybride Dual-SIM 4G USB Typ-C 4 GB 64 GB 5000 mAh Schwarz",
	"SAMSUNG  Galaxy A13 SM-A135F 16,8 cm (6.6 Zoll) Dual-SIM Android 12 4G USB Typ-C 4 GB 128 GB 5000 mAh Blau",
	"SAMSUNG  Galaxy A13 SM-A135F 16,8 cm (6.6 Zoll) Hybride Dual-SIM Android 12 4G USB Typ-C 4 GB 128 GB 5000 mAh Schwarz",
	"SAMSUNG  Galaxy A32 4G Enterprise Edition 16,3 cm (6.4 Zoll) Android 11 USB Typ-C 4 GB 128 GB 5000 mAh Schwarz",
	"SAMSUNG  Galaxy A32 4G SM-A325F 16,3 cm (6.4 Zoll) Dual-SIM Android 11 USB Typ-C 4 GB 128 GB 5000 mAh Schwarz",
	"SAMSUNG  Galaxy XCover 5 Enterprise Edition 13,5 cm (5.3 Zoll) Android 11 4G 4 GB 64 GB 3000 mAh Schwarz",
	"SAMSUNG  Galaxy XCover 5 SM-G525F 13,5 cm (5.3 Zoll) Dual-SIM 4G USB Typ-C 4 GB 64 GB 3000 mAh Schwarz",
	"SAMSUNG Smartphone Bundle Galaxy A33 6.4&quot; 5G Dual SIM 128 GB Schwarz + JBL Tune 510 BT Wireless On-Ear KopfhÃ¶rer Schwarz",
	"SAMSUNG  Smartphone Pack  Galaxy A13 6.6 &quot;Dual SIM 64 GB Schwarz + Transparente -HÃ¼lle",
	"SAMSUNG Galaxy A13, 6.6'' Smartphone",
	"SAMSUNG Galaxy Xcover 5 EE 5.3'' Smartphone",
	"Vivo Pack Smartphone Vivo Y21s 6.51\" Dual SIM 128 GB Mitternachtsblau + Graue SchutzhÃ¼lle + Transparente Schutzfolie",
	"XIAOMI   Redmi 10 2022 Dual 128GB Weiß (6GB)",
	"XIAOMI   Redmi 10C Dual 128GB Grau (4GB)",
	"XIAOMI   Redmi 10C Dual 128GB O.Blau (4GB)",
	"XIAOMI   Redmi Note 10S 128GB Grau (8GB)",
	"XIAOMI  Redmi 10 16,5 cm (6.5 Zoll) Dual-SIM Android 11 4G USB Typ-C 4 GB 64 GB 5000 mAh Blau",
	"XIAOMI  Redmi 10 16,5 cm (6.5 Zoll) Dual-SIM Android 11 4G USB Typ-C 4 GB 64 GB 5000 mAh Grau",
	"XIAOMI  Redmi 10 16,5 cm (6.5 Zoll) Dual-SIM Android 11 4G USB Typ-C 4 GB 64 GB 5000 mAh Weiß",
	"XIAOMI  Redmi 10 2022 16,5 cm (6.5 Zoll) Hybride Dual-SIM Android 11 4G USB Typ-C 4 GB 128 GB 5000 mAh Grau",
	"XIAOMI  Redmi 9a 16,6 cm (6.53 Zoll) Hybride Dual-SIM Android 11 4G Mikro-USB 2 GB 32 GB 5000 mAh Grau",
	"XIAOMI  Redmi 9C 16,6 cm (6.53 Zoll) Hybride Dual-SIM 4G Mikro-USB 4 GB 128 GB 5000 mAh Grau",
	"XIAOMI  Redmi Note 10 Pro 16,9 cm (6.67 Zoll) Dual-SIM Android 11 4G USB Typ-C 6 GB 128 GB 5020 mAh Grau",
	"XIAOMI  Redmi Note 11 16,3 cm (6.43 Zoll) Dual-SIM Android 11 4G USB Typ-C 4 GB 128 GB 5000 mAh Grau",
	"ZTE  Blade A31 Lite 12,7 cm (5 Zoll) Dual-SIM Android 11 4G Mikro-USB 1 GB 32 GB 2000 mAh Grau",
	"ZTE  Blade V30 vita 17,3 cm (6.82 Zoll) Hybride Dual-SIM Android 11 4G 4 GB 64 GB 5000 mAh Blau",
	"ZTE  Blade V40s 16,9 cm (6.67 Zoll) Hybride Dual-SIM Android 12 4G USB Typ-C 4 GB 128 GB 4500 mAh Blau",
}

var manorNamesExpected = []string{
	"Blackview BV4900",
	"Blackview BV6300 Pro",
	"HUAWEI nova 9 SE",
	"MOTOROLA Moto E22i",
	"MOTOROLA Moto g32",
	"MOTOROLA Moto G51",
	"NOKIA G60",
	"NOKIA G21",
	"NOKIA G21",
	"NOKIA G11",
	"NOKIA G21",
	"NOKIA G50",
	"OPPO A54",
	"OPPO A96",
	"OPPO A57s",
	"OPPO A57s",
	"OPPO Reno8 Lite",
	"OPPO A74",
	"SAMSUNG Galaxy A02",
	"SAMSUNG Galaxy A02",
	"SAMSUNG Galaxy A02",
	"SAMSUNG Galaxy A02",
	"SAMSUNG Galaxy A03",
	"SAMSUNG Galaxy A03",
	"SAMSUNG Galaxy A03S",
	"SAMSUNG Galaxy A03S",
	"SAMSUNG Galaxy A12",
	"SAMSUNG Galaxy A12",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A22",
	"SAMSUNG Galaxy A23",
	"SAMSUNG Galaxy A23",
	"SAMSUNG Galaxy A23",
	"SAMSUNG Galaxy A23",
	"SAMSUNG Galaxy A32",
	"SAMSUNG Galaxy M32",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A32",
	"SAMSUNG Galaxy A32",
	"SAMSUNG Galaxy XCover 5",
	"SAMSUNG Galaxy XCover 5",
	"SAMSUNG Galaxy A33",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy A13",
	"SAMSUNG Galaxy Xcover 5",
	"Vivo Y21s",
	"XIAOMI Redmi 10",
	"XIAOMI Redmi 10C",
	"XIAOMI Redmi 10C",
	"XIAOMI Redmi Note 10S",
	"XIAOMI Redmi 10",
	"XIAOMI Redmi 10",
	"XIAOMI Redmi 10",
	"XIAOMI Redmi 10",
	"XIAOMI Redmi 9a",
	"XIAOMI Redmi 9C",
	"XIAOMI Redmi Note 10 Pro",
	"XIAOMI Redmi Note 11",
	"ZTE Blade A31 Lite",
	"ZTE Blade V30 vita",
	"ZTE Blade V40s",
}

func TestManorClean(t *testing.T) {
	for i, name := range manorNames {
		if _name := shop.ManorCleanFn(name); _name != manorNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, manorNamesExpected[i], name)
		}
	}
}
