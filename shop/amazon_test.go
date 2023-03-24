package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var amazonNames = []string{
	"Apple 2022 iPhone SE (128 GB) - Polarstern (3. Generation)",
	"Apple iPhone 12 (128 GB) - Violett",
	"Apple iPhone 13 (128 GB) - (Product) RED",
	"Apple iPhone 13 (128 GB) - Mitternacht",
	"Apple iPhone 13 Mini (128 GB) - Blau",
	"Apple iPhone 14 (128 GB) - Mitternachtsblau",
	"Apple iPhone 14 (512 GB) - Violett",
	"Apple iPhone 14 Plus (256 GB) - Blau",
	"Apple iPhone 14 Plus (256 GB) - Gelb",
	"Apple iPhone 14 Plus (512 GB) - Gelb",
	"Apple iPhone 14 Pro (128 GB) - Dunkellila",
	"Apple iPhone 14 Pro (128 GB) - Space Schwarz",
	"Apple iPhone 14 Pro (256 GB) - Dunkellila",
	"Apple iPhone 14 Pro (256 GB) - Space Schwarz",
	"Apple iPhone 14 Pro (1 TB) - Dunkellila",
	"Apple iPhone 14 Pro (1 TB) - Gold",
	"Apple iPhone 14 Pro (1 TB) - Silber",
	"Apple iPhone 14 Pro Max (128 GB) - Silber",
	"Apple iPhone 14 Pro Max (128 GB) - Space Schwarz",
	"Apple iPhone 14 Pro Max (256 GB) - Dunkellila",
	"Apple iPhone 14 Pro Max (256 GB) - Space Schwarz",
	"Apple iPhone 14 Pro Max (512 GB) - Gold",
	"Bewinner 4G Smartphone Ohne Vertrag Günstig, S22 Ultra Android 11 Handy mit 6.52 Zoll HD Screen und 4000mAh Akku, 4GB+64GB...",
	"Bewinner 4G Smartphone Ohne Vertrag Günstig, S22 Ultra Android 11 Handy mit 6.52 Zoll HD Screen und 4000mAh Akku, 4GB+64GB...",
	"Blackview 5200 PRO Outdoor Handy, 7GB+64GB/1TB Erweiterbar Octa Core Android 12 Smartphone Ohne Vertrag,6.1'' Dual SIM 4G ...",
	"Blackview A100 Smartphone ohne Vertrag, 4G Android Handy 6GB + 128GB 512GB Erweitern, 18W 4680mAh Akku, Octo-Core Prozesso...",
	"Blackview A52(2023) Android 12 Smartphone Ohne Vertrag Günstig, Octa Core 3+32GB/1TB Erweiterbar Handy Ohne Vertrag, 13MP ...",
	"Blackview A55 PRO 4G Handy ohne Vertrag, Helio P22 Octa-Core 4GB+64GB, 5MP+13MP Kamera, 6,53’’ HD+ IPS, 4780mAh Akku, Andr...",
	"Blackview A55 Smartphone Ohne Vertrag Günstig, 4G Android 11 Günstige Handys, 6.5 Zoll HD + Display/Dual SIM/4780mAh/3GB+1...",
	"Blackview A55 Smartphone Ohne Vertrag Günstig, 4G Android 11 Günstige Handys, 6.5 Zoll HD + Display/Dual SIM/4780mAh/3GB+1...",
	"Blackview A55 Smartphone ohne Vertrag, Android 11 Dual SIM Handy, 6,52\" HD+ Display, Quad-Core 3GB + 16GB, 4780mAh Akku, 5...",
	"Blackview A85 Handy Ohne Vertrag, 50MP+8MP Kamera, 8GB+128GB, 6.5\" HD+ 90Hz, 18W Schnellladung, Dual SIM Android 12 Smartp...",
	"Blackview A95 Smartphone Ohne Vertrag, 20MP+8MP Kamera, Helio P70 8GB+128GB, 6,53\" HD+ Bildschirm, 4380mAh 18W Schnellladu...",
	"Blackview BV5200 Android 12 Wasserdichit Outdoor Smartphone, Octa Core 7GB+32GB/1TB Erweiterbar Outdoor Handy Ohne Vertra...",
	"Blackview BV5200 Android 12 Wasserdichit Outdoor Smartphone, Octa Core 7GB+32GB/1TB Erweiterbar Outdoor Handy Ohne Vertra...",
	"Blackview BV7100 Outdoor Handys ohne Vertrag 2023, 10GB+128GB+1TB Erweiterbar,13000 mAh 33W Aufladen, IP68 Android 12 6,5...",
	"Blackview BV7200 Outdoor Handy ohne Vertrag, 50MP Dual Kamera, Helio G85 6GB+128GB, IP68 Android 12 Robust Smartphone, 6....",
	"Blackview Outdoor Handy Ohne Vertrag BV5200, ArcSoft® 13MP+5MP, 4GB+32GB(1TB Erweiterung), Android 12 DUAL SIM IP68 MIL-ST...",
	"CUBOT J10 Smartphone ohne Vertrag, 4 Zoll Touch Bildschirm, Android 11, Quad Prozessor, 32GB ROM, 5MP Kamera, Face ID/GPS/...",
	"CUBOT Max 3 Smartphone ohne Vertrag 4 + 64GB, 6,95 Zoll HD Display, 5000mAh Akku, 48MP AI DREI Kamera, 4G LET Dual SIM Han...",
	"CUBOT Note 30 (2023) Smartphone Ohne Vertrag, Android 12 Günstig Handy, 4GB + 64GB/256GB Erweiterbar, 20MP Kamera, 6.52\" D...",
	"CUBOT P60 Smartphone Ohne Vertrag Günstig Android 12 Handy, 20MP+8MP Dual SIM 4G Handys, 6.52 Zoll HD+, 5000mAh Akku Octa ...",
	"DOOGEE Android 12 Outdoor Handy Ohne Vertrag S41, Ouad Core 3GB+16GB (1TB Erweiterbar), 6300mAh Akku, 8MP Dreifachkamera, ...",
	"DOOGEE S41 Pro（2023） Outdoor Handy Android 12, 7GB+32GB/1TB Outdoor Smartphone Ohne Vertrag, 6300mAh, 5.5\" Bildschirm, 13M...",
	"DOOGEE S51 Outdoor Smartphone, 4GB RAM 64GB ROM, Android 12 4G Simlockfreie Handys Robustes Smartphones, 6,0-Zoll, 12MP + ...",
	"DOOGEE S88 Plus(8GB+128GB) 10000mAh Akku Outdoor Smartphone Ohne Vertrag, 48MP Quad-Kamera, Octa-Core Android 10, 6,3’’ FH...",
	"DOOGEE X95(T) Handy ohne Vertrag Günstig, 4G Smartphone ohne Vertrag Android 10, 13MP+5MP Kamera, 4350mAh Akuu 6,52 Zoll H...",
	"DOOGEE X95(T) Smartphone ohne Vertrag Günstig, 4G Handy ohne Vertrag Android 10 mit 13MP+5MP Kamera, 4350mAh Akuu 6,52 Zol...",
	"DOOGEE X97 (2022) 4G Smartphone ohne vertrag - 6,0 Zoll DH+ Android 12 Handys Simlockfreie, 3GB + 16GB + 256GB Erweiterun...",
	"DOOGEE X97 Pro Smartphone ohne Vertrag Android 12 4GB +64GB Octa-Core Processor Handy Günstig 6,0 Zoll HD 12MP+5MP Kamera,...",
	"DOOGEE X98 Pro Android 12 Smartphone ohne Vertrag Günstig, 9GB/64GB 1TB Erweiterbar Handy Helio G25 Octa-Core 6.52 Zoll H...",
	"DOOGEE X98 Pro Android 12 Smartphone ohne Vertrag, 9GB/64GB 1TB Erweiterbar Handy Günstig Helio G25 Octa-Core 6.52 Zoll H...",
	"EL Smartphone ohne Vertrag Günstig D68 Dual SIM 4G LTE Handy, Android 10 Phone, 6,088 Zoll HD, 3GB RAM+32GB ROM, 4000 mAh ...",
	"Google Pixel 7 Pro – Entsperrtes Android-Smartphone mit Tele- und Weitwinkelobjektiv – 128GB - Hazel",
	"Google Pixel 7 Pro – Entsperrtes Android-Smartphone mit Tele- und Weitwinkelobjektiv – 128GB - Snow",
	"Google Pixel 7 Pro – Entsperrtes Android-Smartphone mit Tele- und Weitwinkelobjektiv – 256GB - Obsidian",
	"Google Pixel 7 Pro – Entsperrtes Android-Smartphone mit Tele- und Weitwinkelobjektiv – 256GB - Snow",
	"Google Pixel 7 – Entsperrtes Android-Smartphone mit Weitwinkelobjektiv – 256GB - Obsidian",
	"Handy ohne Vertrag Günstig,OUKITEL C25 Android 11 6.52 Zoll HD+ 5000mAh Batterie Smartphone ohne Vertrag,256GB Erweiterbar...",
	"HUAWEI P40 lite 5G Dual-SIM Smartphone BUNDLE (16,51cm(6,5 Zoll), 128 GB ROM, 6 GB RAM, Android 10.0 AOSP ohne Google Play...",
	"HUAWEI P40 lite 5G Dual-SIM Smartphone BUNDLE (16,51cm(6,5 Zoll), 128 GB ROM, 6 GB RAM, Android 10.0 AOSP ohne Google Play...",
	"IIIF150 R2022 Outdoor Smartphone, 6.78\" FHD+ 90Hz, 8GB + 128GB Outdoor Handy Ohne Vertrag, 64MP + 20MP Nachtsicht Kamera, ...",
	"KXD Handy, A1 SIM Free Smartphone Entsperrt, 5,71 Zoll Vollbild, 1GB RAM 16G ROM 128G Erweiterung, Android 8.1 Handys, 5MP Dual Rückfahrkameras, Dual SIM, Triple Kartensteckplätze",
	"Motorola edge30 neo Smartphone (6,3\"-FHD+-Display, 64-MP-Kamera, 8/128 GB, 4020 mAh, Android 12), Very Peri, inkl. Schutzc...",
	"Motorola edge30 ultra Smartphone (6,7\"-FHD+-Display, 200-MP-Kamera, 12/256 GB, 4610 mAh, Android 12), Interstellar Black, ...",
	"Motorola Moto g72 Smartphone (6,6\"-FHD+-Display,108-MP-Kamera,6/128 GB,5000 mAh, Android 12), Polar Blue, inkl. Schutzcove...",
	"Nokia X30 5G 6,43\" Smartphone mit AMOLED PureDisplay, FHD+, 6/128 GB, Gorilla Glass Victus, 3 Jahre Garantie, 50MP PureVie...",
	"Nothing Phone (1) – 8 GB RAM + 128 GB, Glyph Interface, 50-MP-Dualkamera, OS, OLED-Display (6,55 Zoll, 120 Hz), schwarz, A063",
	"Nothing Telekom Phone (1) 8 GB 256 GB juodas",
	"OSCAL C20 PRO Android 11 4G Smartphone Ohne Vertrag, Dual-SIM Handy, 6.08\" HD+ Display, Unibody-Design, Quad-Core 2GB RAM+...",
	"OSCAL C60 Smartphone Ohne Vertrag, 4G Android 11 Handy Günstig mit 6.5 Zoll HD+ Display, 2.0 GHz Processor 4GB RAM 128GB E...",
	"OSCAL C60 Smartphone Ohne Vertrag, 4G Android 11 Handy Günstig mit 6.5 Zoll HD+ Display, 2.0 GHz Processor 4GB RAM 128GB E...",
	"OSCAL C60 Smartphone Ohne Vertrag, 4G Android 11 Handy Günstig mit 6.5 Zoll HD+ Display, 2.0 GHz Processor 4GB RAM 128GB E...",
	"OSCAL C60 Smartphone Ohne Vertrag, 4G Android 11 Handy Günstig mit 6.5 Zoll HD+ Display, 2.0 GHz Processor 4GB RAM 128GB E...",
	"OSCAL C60 Smartphone Ohne Vertrag, 4G Android 11 Handy Günstig mit 6.5 Zoll HD+ Display, 2.0 GHz Processor 4GB RAM 128GB E...",
	"OSCAL C80 Android 12 Smartphone Ohne Vertrag, 8GB+128GB, 50MP+8MP Dual Kamera, 6.5\" 90Hz HD+, Dual SIM Handy Schlankes Des...",
	"OSCAL Handy Ohne Vertrag(14+128GB), C80 Android Smartphone 6,5 Zoll 90Hz Display, Octa-Core Prozessor, 50MP Kamera, 4G Dua...",
	"OSCAL Handy Ohne Vertrag(14+128GB), C80 Android Smartphone 6,5 Zoll 90Hz Display, Octa-Core Prozessor, 50MP Kamera, 4G Dua...",
	"OSCAL Outdoor Handy S80(2023) 13.000mAh 33W Laden, 10GB+128GB, Android 12 Octa-Core Handy Ohne Vertrag, 6,58\" FHD+ Display...",
	"OSCAL Smartphone ohne Vertrag, C20 Pro Android Handy 32GB 128GB Erweiter, Handy Günstig 6.1 Zoll HD+ Display, Octa Core Pr...",
	"OSCAL Smartphone Ohne Vertrag Günstig Neu C20(2022), Android 11 Go Handy 6.1 Zoll Wassertropfen Bildschirm 32GB ROM (128GB...",
	"OUKITEL Android 12 Smartphone ohne Vertrag C31 PRO, 4GB+64GB, DUAL SIM+SD (3 Kartensteckplatz), Helio P22 Qcta-Core Günsti...",
	"OUKITEL C21 Smartphone ohne Vertrag, Android 10 Handy, 4 GB + 64 GB, 6,4 Zoll FHD+, AI-Quad-Kamera, Helio P60 Dual SIM, 40...",
	"OUKITEL C25 Handy ohne Vertrag,6.52 Zoll HD+ 5000mAh Batterie Smartphone Günstiges,13MP Triple-Kamera/8MP Frontkamera,4GB+32GB, Quad-Core,Android 11, Dual SIM 4G Handy,Face ID/Fingerabdruck/-Grün",
	"OUKITEL C31 Pro Smartphone ohne Vertrag Günstig Android 12 Handy, Octa Core 4GB+64GB/256GB Erweiterbar, 5150mAh Akku, Dual...",
	"OUKITEL C31 Pro Smartphone ohne Vertrag Günstig Android 12 Handy, Octa Core 4GB+64GB/256GB Erweiterbar, 5150mAh Akku, Dual...",
	"OUKITEL C31 Pro Smartphone ohne Vertrag Günstig Android 12 Handy, Octa Core 4GB+64GB/256GB Erweiterbar, 5150mAh Akku, Dual...",
	"OUKITEL C32 Smartphone Ohne Vertrag, Android 12 Handy mit 8GB(13GB)+128GB/SD 1TB Octa-Core Prozessor, 6,52\" HD+ Display, 5...",
	"OUKITEL C32 Smartphone Ohne Vertrag, Android 12 Handy mit 8GB(13GB)+128GB/SD 1TB Octa-Core Prozessor, 6,52\" HD+ Display, 5...",
	"OUKITEL C32 Smartphone Ohne Vertrag, Android 12 Handy mit 8GB(13GB)+128GB/SD 1TB Octa-Core Prozessor, 6,52\" HD+ Display, 5...",
	"OUKITEL Outdoor Smartphone Ohne Vertrag WP18, 12500mAh Großer Akku 5.93\" Outdoor Handy, 4GB + 32GB Android 11 IP68 Smartph...",
	"OUKITEL Outdoor Smartphone Ohne Vertrag WP18, 12500mAh Großer Akku 5.93\" Outdoor Handy, 4GB + 32GB Android 11 IP68 Smartph...",
	"OUKITEL WP12 Outdoor Smartphone Android 11 Handy Ohne Vertrag, Dual SIM IP68 Wasserdichter, 5,5 Zoll 4000mAh Akku 13MP Kam...",
	"OUKITEL WP18 (2022) Outdoor Smartphone, 12500mAh Großer Akku 5.93\" Outdoor Handy Ohne Vertrag, 4GB + 32GB Android 11 Handy...",
	"OUKITEL WP5 Outdoor Smartphone Ohne Vertrag, 8000mAh Akku Outdoor Handy, 7GB 32GB, 1TB Erweiterbar, Android 11, IP68 Wasse...",
	"OUKITEL WP5 Outdoor Smartphone Ohne Vertrag, 8000mAh Akku Outdoor Handy, 7GB 32GB, 1TB Erweiterbar, Android 11, IP68 Wasse...",
	"realme 8 Smartphone ohne Vertrag, 64MP AI-Quad-Kamera Android Handy, 6,4 Zoll Super AMOLED Display, 30W Dart Charge, Stark...",
	"realme 9 5G - 4+128GB Smartphone, Snapdragon 695 5G-Prozessor, Ultraflüssiges 120-Hz-Display,50 MP KI-Dreifach-Kamera ,Sta...",
	"realme 9 5G - 4+64GB Smartphone, Snapdragon 695 5G-Prozessor, Ultraflüssiges 120-Hz-Display,50 MP KI-Dreifach-Kamera ,Star...",
	"realme 9 Pro+ 5G Smartphone ohne Vertragy,Sony IMX766 Flaggschiff-Kamera,MediaTek Dimensity 920 5G Prozessor,60 W SuperDar...",
	"realme 9 Pro+ 5G Smartphone ohne Vertragy,Sony IMX766 Flaggschiff-Kamera,MediaTek Dimensity 920 5G Prozessor,60 W SuperDar...",
	"realme Narzo 50 5G-4+64GB Smartphone ohne Vertragy, Starker 5000 mAh-Akku, Dimensity 810 5G-Prozessor Android Handy, 33W D...",
	"realme Narzo 50A Prime - 4+64GB Smartphone 16,7 cm (6,6'') FHD+-Vollbildschirm, 50 MP KI-Dreifach-Kamera, Starker 5000-mAh-Akku, Starker Unisoc T612-Prozessor, Flash Black, ohne Netzteil",
	"Redmi 9A Smartphone 2GB 32GB 6.53\" HD+ DotDrop Display 5000mAh (typ) 13 MP AI Rear Camera [Global Version] Green",
	"Samsung F936B Galaxy Z Fold 4 256GB/12GB Dual-SIM graygreen",
	"Samsung Galaxy M52 5G Smartphone Android 128 GB Weiß",
	"Samsung Galaxy S20 FE Cloud Navy G780F Dual-SIM 128GB Android 10.0 Smartphone SM-G780FZBDEUB",
	"Samsung Galaxy S22 SM-S901B 15.5 cm (6.1) Dual SIM Android 12 5G USB Type-C 8 GB 128 GB 3700 mAh Black",
	"Samsung Galaxy S22 S908 Ultra EU 128GB, Android, phantom white",
	"Samsung Galaxy S22+, Android Smartphone, 6,6 Zoll Dynamic AMOLED Display, 4.500 mAh Akku, 128 GB/8 GB RAM, Handy in Phanto...",
	"Samsung Galaxy XCover6 Pro, robustes Android Smartphone ohne Vertrag, 16,72 cm/ 6,6 Zoll Display, 4.050 mAh Akku, 128 GB/6...",
	"Samsung Galaxy XCover6 Pro, robustes Android Smartphone ohne Vertrag, 16,72 cm/ 6,6 Zoll Display, 4.050 mAh Akku, 128 GB/6...",
	"Samsung Galaxy-A33 5G - 128GB Enterprise Edition (30 Monate Garantie), Awesome Black",
	"Smartphone ohne Vertrag OUKITEL C25 Android 11 5000mAh Akku Handy ohne Vertrag Günstig 4GB/32GB 256GB Erweiterbar 6.52 Zol...",
	"Smartphone ohne Vertrag, OSCAL C20 Pro Android Handy 32GB 128GB Erweiter, Handy Günstig 6.1 Zoll HD+ Display, Octa Core Pr...",
	"Sony Xperia 10 IV (5G Smartphone, 6 Zoll, OLED-Display , Dreifach-Kamera, 3,5-mm-Audioanschluss, 5.000mAh Akku, Dual SIM h...",
	"Sony Xperia 10 IV (5G Smartphone, 6 Zoll, OLED-Display , Dreifach-Kamera, 3,5-mm-Audioanschluss, 5.000mAh Akku, Dual SIM h...",
	"Sony Xperia 10 IV (5G Smartphone, 6 Zoll, OLED-Display , Dreifach-Kamera, 3,5-mm-Audioanschluss, 5.000mAh Akku, Dual SIM h...",
	"Ulefone 4G Handy Günstig, Note 8P Smartphone ohne Vertrag Dual SIM 5,5-Zoll 16 GB ROM 128GB Erweiterbar 8MP kameras 3 Slot...",
	"Ulefone 4G Smartphone ohne Vertrag, Note 6(P), 8,5mm Ultra Dünn DUAL-SIM-Handy, 6,1'' HD+ Bildschirm, 3-Karten Slot Desig...",
	"Ulefone 4G Smartphone ohne Vertrag, Note 6(P), 8,5mm Ultra Dünn DUAL-SIM-Handy, 6,1'' HD+ Bildschirm, 3-Karten Slot Desig...",
	"Ulefone 4G Smartphone ohne Vertrag, Note 6(P), 8,5mm Ultra Dünn DUAL-SIM-Handy, 6,1'' HD+ Bildschirm, 3-Karten Slot Desig...",
	"Ulefone Smartphone ohne Vertrag Günstig, Note 6P 4G Handy Android 11 Go Octa-Core 32GB ROM 128GB Erweiterbar mit 8MP Kamera 6,1 Zoll 3 Slots/Face ID/GPS/WiFi/FM/Große Schrift Dual SIM Violett",
	"Ulefone 7700mAh Note 12P Handy ohne Vertrag 2022 OTG Reverse Charge 6,82-Zoll Android 11 OS Smartphone, Octa-Core 4+64GB/S...",
	"Ulefone 7700mAh Note 12P Handy ohne Vertrag 2022 OTG Reverse Charge 6,82-Zoll Android 11 OS Smartphone, Octa-Core 4+64GB/S...",
	"Ulefone Android 12 4G Smartphone ohne Vertrag Note 14, 3-Karten-Steckplatz, Helio A22 Quad Core 3GB+16GB, DUAL-SIM Handy, ...",
	"Ulefone Armor 8 Pro Smartphones Wasserdicht Android 11 Outdoor Handy, Helio P60 Octa-Core 5580mAh Akku 8GB+128GB/ 1TB Erwe...",
	"Ulefone Armor X3 Outdoor Smartphones Ohne Vertrag, 5,5-Zoll Handys IP68/IP69K Wasserdicht, 32GB ROM 128GB Erweiterbar, 500...",
	"Ulefone Android 12 4G Outdoor Smartphone Ohne Vertrag, Armor X6 Pro, Quad-Core 4GB+32GB, 5.0\" IP68 Robust Handy, Dual SIM,...",
	"Ulefone Armor X10 PRO Outdoor Smartphone Ohne Vertrag, 20MP Unterwasserkamera,Helio P22 4GB+64GB, 5,45\" IP68 Wasserdicht H...",
	"Ulefone Note 14 Pro Smartphone Ohne Vertrag, Android 12 Handy 4GB/64GB+128GB(SD) 4500mAh mit 13MP+5MP Kamera 6,52 Zoll HD+...",
	"Ulefone Note 14 Pro Smartphone Ohne Vertrag, Android 12 Handy 4GB/64GB+128GB(SD) 4500mAh mit 13MP+5MP Kamera 6,52 Zoll HD+...",
	"Ulefone Note 14 Pro Smartphone Ohne Vertrag, Android 12 Handy 4GB/64GB+128GB(SD) 4500mAh mit 13MP+5MP Kamera 6,52 Zoll HD+...",
	"Ulefone Power Armor 13 Outdoor Handy ohne Vertrag，13200mAh Akku,Infrarot Entfernungsmessung, 6,81'' 48MP Kamera, IP68 Wass...",
	"Ulefone Power Armor 14 Pro (8GB+128GB) Outdoor Handy Ohne Vertrag, Android 12 Smartphone 10000mAh IP68 Wasserdicht Octa-Co...",
	"Ulefone Power Armor X11 PRO (2023) 4G Simlockfreie Handys, 8150mAh Akku, 64GB Speicher 4GB RAM, 16MP + 5MP Smartphones, An...",
	"UMIDIGI Power 7S Smartphone Ohne Vertrag, Großer Akku 6150mAh, Dual SIM Android 11 Handy mit 6.7 Zoll HD+ Display 4GB + 64...",
	"UMIDIGI G1 MAX Smartphone Ohne Vertrag Günstig,6GB+128GB Handy Günstig,Android 12 Handy Ohne Vertrag,6.52 Zoll HD+ Display,50MP AI Kamera,5150mAh Akku,Octa Core/4G Dual SIM/OTG(Schwarz)",
	"UMIDIGI G1 MAX Smartphone Ohne Vertrag Günstig,6GB + 128GB Handy Günstig,Android 12 Go Handy Ohne Vertrag,6.52 Zoll HD+ Display,50MP AI Kamera,5150mAh Akku,Octa Core/4G Dual SIM/OTG(Blau)",
	"VIVO y76 5G Smartphone, 4.100 (typisch) mAh+ 44 W FlashCharge 50 MP Hauptkamera 8 GB RAM + 4 GB erweiterter RAM+128 GB RO...",
	"XGODY Y13 Quad Core Android 9.0 Smartphone Ohne Vertrag 6.0'' 4G Dual SIM günstiges Handys, 1GB RAM 8GB ROM, 256GB Erweiterbar, 3 Kartensteckplätze, 5MP Kamera, Face ID, GPS, Deutsche (Grün)",
	"Xiaomi Redmi Note 11S, Smartphone + Kopfhörer, 6 + 64 GB Handy, 6,43'' 90 Hz FHD + AMOLED DotDisplay, MediaTek Helio G96, ...",
	"ZTE Smartphone Blade V40 S 4G (16,94cm (6,67 Zoll) FHD+ Display, 4G LTE, 4GB RAM und 128GB interner Speicher, 50MP Hauptka...",
	"ZTE Smartphone Blade V40 (16,94cm (6,67 Zoll) FHD+ Display, 4G LTE, 4GB RAM und 128GB interner Speicher, 48MP Hauptkamera und 8MP Frontkamera, Dual-SIM, Android 11) schwarz 123401201022",
}

var amazonNamesExpected = []string{
	"Apple iPhone SE",
	"Apple iPhone 12",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13 Mini",
	"Apple iPhone 14",
	"Apple iPhone 14",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Plus",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro",
	"Apple iPhone 14 Pro Max",
	"Apple iPhone 14 Pro Max",
	"Apple iPhone 14 Pro Max",
	"Apple iPhone 14 Pro Max",
	"Apple iPhone 14 Pro Max",
	"Bewinner S22 Ultra",
	"Bewinner S22 Ultra",
	"Blackview 5200 PRO",
	"Blackview A100",
	"Blackview A52",
	"Blackview A55 PRO",
	"Blackview A55",
	"Blackview A55",
	"Blackview A55",
	"Blackview A85",
	"Blackview A95",
	"Blackview BV5200",
	"Blackview BV5200",
	"Blackview BV7100",
	"Blackview BV7200",
	"Blackview BV5200",
	"CUBOT J10",
	"CUBOT Max 3",
	"CUBOT Note 30",
	"CUBOT P60",
	"DOOGEE S41",
	"DOOGEE S41 Pro",
	"DOOGEE S51",
	"DOOGEE S88 Plus",
	"DOOGEE X95(T)",
	"DOOGEE X95(T)",
	"DOOGEE X97",
	"DOOGEE X97 Pro",
	"DOOGEE X98 Pro",
	"DOOGEE X98 Pro",
	"EL D68",
	"Google Pixel 7 Pro",
	"Google Pixel 7 Pro",
	"Google Pixel 7 Pro",
	"Google Pixel 7 Pro",
	"Google Pixel 7",
	"OUKITEL C25",
	"HUAWEI P40 lite",
	"HUAWEI P40 lite",
	"IIIF150 R2022",
	"KXD A1",
	"Motorola edge30 neo",
	"Motorola edge30 ultra",
	"Motorola Moto g72",
	"Nokia X30",
	"Nothing Phone (1)",
	"Nothing Phone (1)",
	"OSCAL C20 PRO",
	"OSCAL C60",
	"OSCAL C60",
	"OSCAL C60",
	"OSCAL C60",
	"OSCAL C60",
	"OSCAL C80",
	"OSCAL C80",
	"OSCAL C80",
	"OSCAL S80",
	"OSCAL C20 Pro",
	"OSCAL C20",
	"OUKITEL C31 PRO",
	"OUKITEL C21",
	"OUKITEL C25",
	"OUKITEL C31 Pro",
	"OUKITEL C31 Pro",
	"OUKITEL C31 Pro",
	"OUKITEL C32",
	"OUKITEL C32",
	"OUKITEL C32",
	"OUKITEL WP18",
	"OUKITEL WP18",
	"OUKITEL WP12",
	"OUKITEL WP18",
	"OUKITEL WP5",
	"OUKITEL WP5",
	"realme 8",
	"realme 9",
	"realme 9",
	"realme 9 Pro+",
	"realme 9 Pro+",
	"realme Narzo 50",
	"realme Narzo 50A Prime",
	"Redmi 9A",
	"Samsung F936B Galaxy Z Fold 4",
	"Samsung Galaxy M52",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22 Ultra",
	"Samsung Galaxy S22+",
	"Samsung Galaxy XCover6 Pro",
	"Samsung Galaxy XCover6 Pro",
	"Samsung Galaxy A33",
	"OUKITEL C25",
	"OSCAL C20 Pro",
	"Sony Xperia 10 IV",
	"Sony Xperia 10 IV",
	"Sony Xperia 10 IV",
	"Ulefone Note 8P",
	"Ulefone Note 6(P)",
	"Ulefone Note 6(P)",
	"Ulefone Note 6(P)",
	"Ulefone Note 6P",
	"Ulefone Note 12P",
	"Ulefone Note 12P",
	"Ulefone Note 14",
	"Ulefone Armor 8 Pro",
	"Ulefone Armor X3",
	"Ulefone Armor X6 Pro",
	"Ulefone Armor X10 PRO",
	"Ulefone Note 14 Pro",
	"Ulefone Note 14 Pro",
	"Ulefone Note 14 Pro",
	"Ulefone Power Armor 13",
	"Ulefone Power Armor 14 Pro",
	"Ulefone Power Armor X11 PRO",
	"UMIDIGI Power 7S",
	"UMIDIGI G1 MAX",
	"UMIDIGI G1 MAX",
	"VIVO y76",
	"XGODY Y13",
	"Xiaomi Redmi Note 11S",
	"ZTE Blade V40 S",
	"ZTE Blade V40",
}

func TestLint(t *testing.T) {
	for i, name := range amazonNames {
		if _name := shop.AmazonCleanFn(name); _name != amazonNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, amazonNamesExpected[i], name)
		}
	}
}
