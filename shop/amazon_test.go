package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var amazonNames = []string{
	"0 Xiaomi Redmi 10C 4GB/128GB Grey Non-NFC EU",
	"- Xiaomi Redmi Note 12 5G 128GB/4GB Dual-SIM Frosted-Green",
	"A34 AWESOME LIME 128GB 5G",
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
	"Apple SLP iPhone 6S 64GB Premium GRIS SIDERAL",
	"Bewinner 4G Smartphone Ohne Vertrag Günstig, S22 Ultra Android 11 Handy mit 6.52 Zoll HD Screen und 4000mAh Akku, 4GB+64GB...",
	"Bewinner 4G Smartphone Ohne Vertrag Günstig, S22 Ultra Android 11 Handy mit 6.52 Zoll HD Screen und 4000mAh Akku, 4GB+64GB...",
	"Blackview 5200 PRO Outdoor Handy, 7GB+64GB/1TB Erweiterbar Octa Core Android 12 Smartphone Ohne Vertrag,6.1'' Dual SIM 4G ...",
	"Blackview A100 Simlockfreie Handys, Helio P70 Octa-Core 6GB RAM 128GB ROM, 6.67\", Android 11 4G Smartphones Ohne Vertrag, 4680mAh Akku, NFC RGB Kompass Dual SIM GPS ID Gesicht Schwarz",
	"Blackview A100 Smartphone ohne Vertrag, 4G Android Handy 6GB + 128GB 512GB Erweitern, 18W 4680mAh Akku, Octo-Core Prozesso...",
	"Blackview A52(2023) Android 12 Smartphone Ohne Vertrag Günstig, Octa Core 3+32GB/1TB Erweiterbar Handy Ohne Vertrag, 13MP ...",
	"Blackview A55 - Smartphone mit Quad-Core Prozessor - 6.528\" HD+ Waterdrop Display - 5MP Front-, 8MP Rückkamera - 3GB RAM+16GB ROM - Android 11 - Leistungsstarker 4780 mAh Akku - Grün",
	"Blackview A55 PRO 4G Handy ohne Vertrag, Helio P22 Octa-Core 4GB+64GB, 5MP+13MP Kamera, 6,53’’ HD+ IPS, 4780mAh Akku, Andr...",
	"Blackview A55 Pro Blau | 4 GB RAM + 64 GB ROM, Betriebssystem Android 11.0",
	"Blackview A55 Smartphone Ohne Vertrag Günstig, 4G Android 11 Günstige Handys, 6.5 Zoll HD + Display/Dual SIM/4780mAh/3GB+1...",
	"Blackview A55 Smartphone Ohne Vertrag Günstig, 4G Android 11 Günstige Handys, 6.5 Zoll HD + Display/Dual SIM/4780mAh/3GB+1...",
	"Blackview A55 Smartphone ohne Vertrag, Android 11 Dual SIM Handy, 6,52\" HD+ Display, Quad-Core 3GB + 16GB, 4780mAh Akku, 5...",
	"Blackview A70 DS 3GB/32GB Black EU",
	"Blackview A85 Handy Ohne Vertrag, 50MP+8MP Kamera, 8GB+128GB, 6.5\" HD+ 90Hz, 18W Schnellladung, Dual SIM Android 12 Smartp...",
	"Blackview A 90 - Schlankes & Leichtes Smartphone - Dual-SIM-Handy - 4 GB RAM + 64 GB ROM - 185g Leichtgewicht - Android 11 - Immersives Gaming mit Schnellem RAM - europäische Version - Schwarz",
	"Blackview A95 Smartphone Ohne Vertrag, 20MP+8MP Kamera, Helio P70 8GB+128GB, 6,53\" HD+ Bildschirm, 4380mAh 18W Schnellladu...",
	"Blackview A958GBRAM128GBROM Handy Negro, BVA958GB128GB-BLK, Black",
	"Blackview BV4900 DUAL SIM 32GB Black Orange, schwarz",
	"Blackview BV4900-4G Smartphone - Ultra Robustes Mobiltelefon mit 5.7\" HD-Display - IP68 Wasserdicht, Staub- und Sturzsicher - 3GB RAM + 32GB ROM - europäische Version - Orange",
	"Blackview BV4900S ‎- Smartphone 32GB, 2GB RAM, Dual SIM, Green",
	"Blackview BV5200 Android 12 Wasserdichit Outdoor Smartphone, Octa Core 7GB+32GB/1TB Erweiterbar Outdoor Handy Ohne Vertra...",
	"Blackview BV5200 Android 12 Wasserdichit Outdoor Smartphone, Octa Core 7GB+32GB/1TB Erweiterbar Outdoor Handy Ohne Vertra...",
	"Blackview BV7100 Outdoor Handys ohne Vertrag 2023, 10GB+128GB+1TB Erweiterbar,13000 mAh 33W Aufladen, IP68 Android 12 6,5...",
	"Blackview BV7200 Outdoor Handy ohne Vertrag, 50MP Dual Kamera, Helio G85 6GB+128GB, IP68 Android 12 Robust Smartphone, 6....",
	"Blackview Mobile Phone A55/BLACK",
	"Blackview Mobile Phone A55/BLUE",
	"Blackview Mobile Phone A55/GREEN",
	"Blackview Mobile Phone BV4900 PRO/ORANGE",
	"Blackview Mobile Phone BV6600 PRO/Black",
	"Blackview Outdoor Handy Ohne Vertrag BV5200, ArcSoft® 13MP+5MP, 4GB+32GB(1TB Erweiterung), Android 12 DUAL SIM IP68 MIL-ST...",
	"Blackview Smartphone A55 Pro 4/64 GB blau + Hydrogel Film für Handy",
	"Brodos WIKO Y81 Smartphone, 6,2 Zoll (15,75 cm), 4G, Dual-SIM, Android 10, Gold [Deutsche Ware], WIKY81WV680GOLSTAM",
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
	"emporiaSMART.4 | Senior Mobile Phone 4G Volte | No-Contract Senior Smartphone | Mobile Phone with Emergency Call Button | 5-inch Display | Android 10 | 13 MP Camera | Black",
	"Galaxy M13",
	"Google Pixel 7 Pro – Entsperrtes Android-Smartphone mit Tele- und Weitwinkelobjektiv – 128GB - Hazel",
	"Google Pixel 7 Pro – Entsperrtes Android-Smartphone mit Tele- und Weitwinkelobjektiv – 128GB - Snow",
	"Google Pixel 7 Pro – Entsperrtes Android-Smartphone mit Tele- und Weitwinkelobjektiv – 256GB - Obsidian",
	"Google Pixel 7 Pro – Entsperrtes Android-Smartphone mit Tele- und Weitwinkelobjektiv – 256GB - Snow",
	"Google Pixel 7 – Entsperrtes Android-Smartphone mit Weitwinkelobjektiv – 256GB - Obsidian",
	"Handy ohne Vertrag Günstig,OUKITEL C25 Android 11 6.52 Zoll HD+ 5000mAh Batterie Smartphone ohne Vertrag,256GB Erweiterbar...",
	"Honor HonorMagic5 Lite 5G 256GB Midnight Black",
	"HONOR Magic 4 Lite, Android Smartphone, 6 + 128 GB Handy, 48-MP-Kamera, 6,81\" 90Hz-LCD, Snapdragon 680, 66W-Schnellladung mit 4800-mAh-Akku, BLAU",
	"HUAWEI Mate10 lite Dual-Sim Smartphone (14.97 cm (5.9 Zoll), 64 GB interner Speicher, 4 GB RAM, 16 MP + 2 MP Kamera, Android 7.0, EMUI 5.1) Aurora Blau",
	"Huawei Mate20 Dual-SIM Smartphone Bundle (6,53 Zoll, 128 GB interner Speicher, 4 GB RAM, Android 9.0, EMUI 9.0) midnight blau + USB Typ-C-Adapter [Exklusiv bei Amazon] - Deutsche Version",
	"HUAWEI 51091CKM P10 lite Dual-SIM Smartphone (13,2 cm (5,2 Zoll) Touch-Display, 32 GB interner Speicher, Android 7.0) Weiß",
	"HUAWEI P40 lite 5G Dual-SIM Smartphone BUNDLE (16,51cm(6,5 Zoll), 128 GB ROM, 6 GB RAM, Android 10.0 AOSP ohne Google Play...",
	"HUAWEI P40 lite 5G Dual-SIM Smartphone BUNDLE (16,51cm(6,5 Zoll), 128 GB ROM, 6 GB RAM, Android 10.0 AOSP ohne Google Play...",
	"Huawei was-LX1 Black Smartphone P10 Lite (Dual SIM 32GB Speicher, Quad Core Prozessor, 13,20 cm (5,2 Zoll)) schwarz",
	"IIIF150 R2022 Outdoor Smartphone, 6.78\" FHD+ 90Hz, 8GB + 128GB Outdoor Handy Ohne Vertrag, 64MP + 20MP Nachtsicht Kamera, ...",
	"KXD Handy, A1 SIM Free Smartphone Entsperrt, 5,71 Zoll Vollbild, 1GB RAM 16G ROM 128G Erweiterung, Android 8.1 Handys, 5MP Dual Rückfahrkameras, Dual SIM, Triple Kartensteckplätze",
	"Lenovo/Motorola PA4N0106IT P2 Smartphone (13.97 cm (5.5 Zoll), 13 MP Kamera, Android 6.0 (Marshmallow)), 32 GB Dunkel Grau",
	"Milwaukee Samsung A047F Galaxy A04s Unlocked 32GB/3GB RAM Dual-SIM grün",
	"MOBILE PHONE POCO X5 5G/6/128GB BLUE MZB0D6UEU POCO",
	"MOBILE PHONE REDMI GO 16GB/BLUE MZB7506EU XIAOMI",
	"moto G 3. Generation Smartphone (12,7 cm (5 Zoll) Touchscreen-Display, 16 GB Speicher, Android 5.1.1) schwarz",
	"Moto One 4+64 White",
	"moto one Smartphone (14,98 cm (5,9 Zoll), 64 GB interner Speicher, 4 GB RAM, Android one ) Weiß, inkl. Schutzcover",
	"Motorola edge30 neo Smartphone (6,3\"-FHD+-Display, 64-MP-Kamera, 8/128 GB, 4020 mAh, Android 12), Very Peri, inkl. Schutzc...",
	"Motorola edge30 ultra Smartphone (6,7\"-FHD+-Display, 200-MP-Kamera, 12/256 GB, 4610 mAh, Android 12), Interstellar Black, ...",
	"Motorola Mobility Edge 20 Light 128GB Handy, grau, Elektrisches Graphit, Dual SIM MOTOEDGE 20LITE",
	"Motorola Moto E 20 16.5 cm (6.5) Dual SIM Android 11 Go Edition USB Type-C 2 GB 32 GB 4000 mAh Blue",
	"Motorola Moto G5 Plus Oro Dual SIM XT1685",
	"Motorola Moto g72 Smartphone (6,6\"-FHD+-Display,108-MP-Kamera,6/128 GB,5000 mAh, Android 12), Polar Blue, inkl. Schutzcove...",
	"MOVIL Smartphone realme X50 PRO 8GB 256GB 5G grün (Moss Green)",
	"Nokia G21 Azul 4+128GB / 6.5' HD+ 90HZ",
	"Nokia RM-1172 Dark Silver All Carriers Handy 230, 7,11 cm (2,8 Zoll) (Dual SIM, MP3 Player, microSD Kartenleser, 1200mAh Akku, Taschenlampe) grau, Standard",
	"Nokia RM-1172 Handy 230, 7,11 cm (2,8 Zoll) (Dual SIM, MP3 Player, microSD Kartenleser, 1200mAh Akku, Taschenlampe) Silber",
	"Nokia X30 5G 6,43\" Smartphone mit AMOLED PureDisplay, FHD+, 6/128 GB, Gorilla Glass Victus, 3 Jahre Garantie, 50MP PureVie...",
	"Nothing Phone (1) – 8 GB RAM + 128 GB, Glyph Interface, 50-MP-Dualkamera, OS, OLED-Display (6,55 Zoll, 120 Hz), schwarz, A063",
	"Nothing Telefon (1) - 8 GB RAM + 256 GB, Glyph-Schnittstelle, 50-MP-Doppelkamera, Nothing Betriebssystem, 6,55-Zoll-OLED-Display mit 120 Hz, Schwarz",
	"Nothing Telekom Phone (1) 8 GB 256 GB juodas",
	"OPPO Reno 4Z - Smartphone 128GB, 8GB RAM, Dual SIM, Ink Black",
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
	"realme 6 6,5 Zoll FHD+ Display DualSIM Smartphone 8GB RAM + 128GB ROM Quad-Kamera Comet White",
	"realme 7 5g Smartphone ohne Vertrag, 6,5 Zoll 120Hz Display, 5000mAh Akku, 48MP+16MP Quad Kamera, 30W Dart Charge, Dual SIM Android Handy, NFC, 6+128GB, Blau",
	"realme 8 Smartphone ohne Vertrag, 64MP AI-Quad-Kamera Android Handy, 6,4 Zoll Super AMOLED Display, 30W Dart Charge, Stark...",
	"realme 9 5G - 4+128GB Smartphone, Snapdragon 695 5G-Prozessor, Ultraflüssiges 120-Hz-Display,50 MP KI-Dreifach-Kamera ,Sta...",
	"realme 9 5G - 4+64GB Smartphone, Snapdragon 695 5G-Prozessor, Ultraflüssiges 120-Hz-Display,50 MP KI-Dreifach-Kamera ,Star...",
	"realme 9 Pro+ 5G Smartphone ohne Vertragy,Sony IMX766 Flaggschiff-Kamera,MediaTek Dimensity 920 5G Prozessor,60 W SuperDar...",
	"realme 9 Pro+ 5G Smartphone ohne Vertragy,Sony IMX766 Flaggschiff-Kamera,MediaTek Dimensity 920 5G Prozessor,60 W SuperDar...",
	"Realme GT Master Edition 6+128 GB, Cosmos Black",
	"realme Narzo 50 5G-4+64GB Smartphone ohne Vertragy, Starker 5000 mAh-Akku, Dimensity 810 5G-Prozessor Android Handy, 33W D...",
	"realme Narzo 50-6/128GB Smartphone ohne Vertragy, Starker 5000 mAh-Akku, Dimensity 810 5G-Prozessor Android Handy, 33W Dart Charge, Ultraflüssiges 90 Hz-Display, NFC, Dual SIM, Hyper Blue, One Size",
	"realme Narzo 50A Prime - 4+64GB Smartphone 16,7 cm (6,6'') FHD+-Vollbildschirm, 50 MP KI-Dreifach-Kamera, Starker 5000-mAh-Akku, Starker Unisoc T612-Prozessor, Flash Black, ohne Netzteil",
	"realme REA DS C31 4+64 UK LSIL RMX3501",
	"Redmi 9A Smartphone 2GB 32GB 6.53\" HD+ DotDrop Display 5000mAh (typ) 13 MP AI Rear Camera [Global Version] Green",
	"Redmi 9C Sunrise Orange 3GB RAM 64GB ROM",
	"Redmi Note 9 All Carriers Polar White 3GB RAM 64GB ROM – Deutschland",
	"Samsung A035G Galaxy A03 64 GB (Black) ohne Simlock, ohne Branding",
	"Samsung A13 5G 128 GB Black EU",
	"Samsung A127F/DS A12 6GB/128GB Black EU",
	"Samsung A137F Galaxy A13 Unlocked 64 GB (Light Blue) débloqué sans Branding",
	"Samsung A145R Galaxy A14 64GB/4GB RAM Dual-SIM grün",
	"Samsung A226 Galaxy A22 5G, Smartphone, 5G, Android 11, Capacité: 1000 GB, Brand Tim, [Italia]",
	"Samsung A236F/DS A23 5G, Awesome White, 4-GB/128-GB",
	"Samsung F936B Galaxy Z Fold 4 256GB/12GB Dual-SIM graygreen",
	"Samsung G A22 64/4 Black SM-A225FZKDEUB",
	"Samsung Galaxy A02S SMD",
	"Samsung Galaxy A13 Android-Smartphone, 6,6 Zoll Infinity-V Display, Android 12, 4 GB RAM und 128 GB interner erweiterbarer² Speicher, Akku 5.000 mAh³, Schwarz",
	"Samsung Galaxy A13 SM-A137F, Android Smartphone, 6,6 Zoll Infinity-V TFT Display, 5.000 mAh Akku, 4 GB RAM/ 64 GB Speicher, Dual-SIM, White, inkl. 30 Monate Garantie [Exclusiv bei Amazon]",
	"Samsung Galaxy A14 LTE Android Smartphone ohne Vertrag, 64 GB, 4 GB RAM, 6,6 Zoll Dynamic AMOLED Display, 5.000 mAh Akku, Black, Handy inkl. 30 Monate Herstellergarantie [Exklusiv bei Amazon]",
	"Samsung Galaxy A32 4G, Schwarz",
	"Samsung Galaxy A34 5G, Android Smartphone, 6,6 Zoll Dynamic AMOLED Display, 5.000 mAh Akku, 128 GB/6 GB RAM Handy in Awesome Silver inkl. 30 Monate Herstellergarantie [Exklusiv bei Amazon]",
	"Samsung Galaxy M52 5G Smartphone Android 128 GB Weiß",
	"Samsung Galaxy S20 FE Cloud Navy G780F Dual-SIM 128GB Android 10.0 Smartphone SM-G780FZBDEUB",
	"Samsung Galaxy S22 SM-S901B 15.5 cm (6.1) Dual SIM Android 12 5G USB Type-C 8 GB 128 GB 3700 mAh Black",
	"Samsung Galaxy S22 S908 Ultra EU 128GB, Android, phantom white",
	"Samsung Galaxy S22+, Android Smartphone, 6,6 Zoll Dynamic AMOLED Display, 4.500 mAh Akku, 128 GB/8 GB RAM, Handy in Phanto...",
	"Samsung Galaxy Xcover Pro Enterprise Edition - Smartphone 64GB, 4GB, Black",
	"Samsung Galaxy XCover6 Pro, robustes Android Smartphone ohne Vertrag, 16,72 cm/ 6,6 Zoll Display, 4.050 mAh Akku, 128 GB/6...",
	"Samsung Galaxy XCover6 Pro, robustes Android Smartphone ohne Vertrag, 16,72 cm/ 6,6 Zoll Display, 4.050 mAh Akku, 128 GB/6...",
	"Samsung Galaxy-A33 5G - 128GB Enterprise Edition (30 Monate Garantie), Awesome Black",
	"Samsung M135F Galaxy M13 128GB/4GB RAM Dual-SIM Light-Blue",
	"Samsung M135F/DS M13 4GB/64GB Green EU",
	"Samsung Sam Galaxy A13 EU-DS-64-4-4G-bu Galaxy A13 EU 64/4GB Light Blue",
	"Samsung Smartphone Galaxy S10+ (Hybrid SIM) 128GB - Prism Blue (Generalüberholt)",
	"Samsung Smartphone Marke Modell Galaxy A13 5G 64 GB Schwarz",
	"SIPP5 Motorola Moto G82 5G 128GB/6GB RAM Dual-SIM White-Lily",
	"Smartphone ohne Vertrag OUKITEL C25 Android 11 5000mAh Akku Handy ohne Vertrag Günstig 4GB/32GB 256GB Erweiterbar 6.52 Zol...",
	"Smartphone ohne Vertrag, OSCAL C20 Pro Android Handy 32GB 128GB Erweiter, Handy Günstig 6.1 Zoll HD+ Display, Octa Core Pr...",
	"Smartfon Samsung Galaxy A13 5G 4/128GB Dual SIM Niebieski (SM-A136BLBVEUB)",
	"Sony Xperia 10 IV (5G Smartphone, 6 Zoll, OLED-Display , Dreifach-Kamera, 3,5-mm-Audioanschluss, 5.000mAh Akku, Dual SIM hybrid) 24+6 Monate Garantie [Amazon Exklusiv] lavendel",
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
	"XIA DS REDMI 10C 3+64 Tim BLU",
	"Xia Redmi 10A 128-4-4G-bu Xiaomi LTE 128/4 Sky Blue, 4+128",
	"Xiaomi 10 2022 Grau 4GB RAM 128GB ROM MZB0A65EU 4GB + 128GB",
	"Xiaomi 11T Cinemagic 5G Meteorite Gray 128GB Dual SIM, grau",
	"Xiaomi all carriers , Redmi Note 12 128GB Handy, schwarz, Onyx Gray, Android 13, Dual SIM",
	"Xiaomi M5 Smartphone+Kopfhörer, 4+128GB Handy ohne Vertrag, 6.58” 90Hz FHD+ DotDrop Display, MediaTek Helio G99, 50MP AI Dreifach-Kamera, 5000mAh, NFC, Black (DE Version + 2 Jahre Garantie) – Deutschland, F...",
	"Xiaomi Mi 9T Pro Smartphone (16,23cm (6.39 Zoll) FHD+ AMOLED Display, 128GB interner Speicher + 6GB RAM, 48MP 3fach-KI-Rückkamera, 20MP Pop-up-Selfie-Frontkamera, Dual-SIM, Android 9.0) Carbon Black",
	"Xiaomi Note 10S Pebble White 6G RAM 128G ROM, Large",
	"Xiaomi Note 11 Pro 6+128G 2201116TG Grau",
	"Xiaomi Poco C40 17 cm (6.71) Dual SIM Android 11 4G USB Type-C 4 GB 64 GB 6000 mAh Green",
	"Xiaomi Redmi 9A 2+32 Aurora Green",
	"Xiaomi REDMI 9C 6,53'' HD+ 128Gb 4Gb Gris MZB0AK2EU",
	"Xiaomi Redmi 9C NFC 2+32 Lavender Purple",
	"Xiaomi Redmi 10C (Ocean Blue) Dual SIM 6.71\" IPS LCD 720x1650/2.4GHz&1.9GHz/128GB/4GB RAM/Android 11/microSDXC/WiFi.BT.4G.MZB0B2OEU",
	"Xiaomi Redmi 10C Mint Green 4GB RAM 128GB ROM 38590",
	"Xiaomi Redmi Note 8 Pro Telefon 6 Gb Ram + 64 Gb Rom,6,53 \"Vollbild,Mtk Helio G90T Octa-Core-Prozessor,20 Mp Front- Und 64 Mp Rückkamera (Globale Version) Blau",
	"Xiaomi Redmi Note 10 Pro Gradient Bronze 128GB Dual SIM, 6+128GB",
	"Xiaomi Redmi Note 10 Pro Onyx Gray 64GB Dual SIM",
	"Xiaomi Redmi Note 10 Pro(6,67\"), leistungsstarkes Android Smartphone ohne Vertrag+Kopfhörer, 120Hz AMOLED Display, 6GB RAM+128GB Speicher,Dual-SIM, 108MP Quad-Kamera, Grey-[Exklusiv bei Amazon]",
	"Xiaomi Redmi Note 10S Grau 6+64G",
	"Xiaomi Redmi Note 11 128gb Graphite Grey, Graphitgrau, 4G 128GB",
	"Xiaomi Redmi Note 11 Pro + 5G, Smartphone + Kopfhörer, 6 + 128 GB Handy ohne Vertrag, 6,67'' 120 Hz FHD + AMOLED Display, 120 W HyperCharge, 108 MP Kamera, Graphite Gray (DE Version, Amazon Exclusive)",
	"Xiaomi Redmi Note 11S, Smartphone + Kopfhörer, 6 + 64 GB Handy, 6,43'' 90 Hz FHD + AMOLED DotDisplay, MediaTek Helio G96, ...",
	"Xiaomi Redmi Note 12 Pro Glacier Blue 8GB RAM 256GB ROM",
	"Xiaomi Redmi Note 12 Pro Graphite Gray 6GB RAM 128GB ROM",
	"Xiaomi Telefonas Poco M5/64GB Žalias MZB0CA1EU Poco",
	"Xiaomi TELEFONO MOVIL REDMI 10 Blue 6.5\"-OC2.0-4GB-64GB",
	"Xiaomi TELEFONO MOVIL REDMI 10 White 6.5\"-OC2.0-4GB-64GB",
	"Xiaomi XIA DS REDMI Note 12 4+128 GLO Gry",
	"XIAOMI Xia Poco F4 128-6-5G gn 128G Nebula Green, MZB0BMSEU",
	"Xiaomi Xia Redmi A2 32-2-4G-gn Redmi A2 Dual SIM 32GB 2GB Grün",
	"ZTE Blade V40 Vita Smartphone 6,74 Zoll HD + 90 Hz, 4 GB RAM, 128 GB Speicher, 5130 mAh, Schnellladung 22,5 W, Dreifachkamera 48 MP, NFC, Rot",
	"ZTE Blade V40 VITA Zeus Black / 4+128GB / 6.75' 90HZ HD+",
	"ZTE Blade A51 GRIS/ 8-CORE /2GB/32GB/6.52\" HD+/DUAL SIM, black",
	"ZTE Smartphone Blade A32 (13, 84cm (5, 45 Zoll) HD Display, 4G LTE, 2GB RAM und 32GB interner Speicher, 5MP Hauptkamera Frontkamera, Dual-SIM, Android R GO) schwarz, 123402901021",
	"ZTE Smartphone Blade V40 S 4G (16,94cm (6,67 Zoll) FHD+ Display, 4G LTE, 4GB RAM und 128GB interner Speicher, 50MP Hauptka...",
	"ZTE Smartphone Blade V40 (16,94cm (6,67 Zoll) FHD+ Display, 4G LTE, 4GB RAM und 128GB interner Speicher, 48MP Hauptkamera und 8MP Frontkamera, Dual-SIM, Android 11) schwarz 123401201022",
	"ZTE Smartphone Blade V40 vita Buds Weiss (17,13cm (6,75 Zoll) HD+ Display, 4G LTE, 4GB RAM und 128GB interner Speicher, 48MP Hauptkamera und 8MP Frontkamera, Dual-SIM, Android R) schwarz",
	"ZTE Smartphone V40 Vita 6,74' 4 GB RAM 128 GB",
}

var amazonNamesExpected = []string{
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi Note 12",
	"Samsung Galaxy A34",
	"Apple iPhone SE (2022)",
	"Apple iPhone 12",
	"Apple iPhone 13",
	"Apple iPhone 13",
	"Apple iPhone 13 mini",
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
	"Apple iPhone 6S",
	"Bewinner S22 Ultra",
	"Bewinner S22 Ultra",
	"Blackview 5200 Pro",
	"Blackview A100",
	"Blackview A100",
	"Blackview A52",
	"Blackview A55",
	"Blackview A55 Pro",
	"Blackview A55 Pro",
	"Blackview A55",
	"Blackview A55",
	"Blackview A55",
	"Blackview A70",
	"Blackview A85",
	"Blackview A90",
	"Blackview A95",
	"Blackview A95",
	"Blackview BV4900",
	"Blackview BV4900",
	"Blackview BV4900s",
	"Blackview BV5200",
	"Blackview BV5200",
	"Blackview BV7100",
	"Blackview BV7200",
	"Blackview A55",
	"Blackview A55",
	"Blackview A55",
	"Blackview BV4900 Pro",
	"Blackview BV6600 Pro",
	"Blackview BV5200",
	"Blackview A55 Pro",
	"Wiko Y81",
	"Cubot J10",
	"Cubot Max 3",
	"Cubot Note 30",
	"Cubot P60",
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
	"emporiaSMART.4",
	"Samsung Galaxy M13",
	"Google Pixel 7 Pro",
	"Google Pixel 7 Pro",
	"Google Pixel 7 Pro",
	"Google Pixel 7 Pro",
	"Google Pixel 7",
	"OUKITEL C25",
	"HONOR Magic5 Lite",
	"HONOR Magic4 Lite",
	"HUAWEI Mate 10 lite",
	"HUAWEI Mate 20",
	"HUAWEI P10 lite",
	"HUAWEI P40 lite",
	"HUAWEI P40 lite",
	"HUAWEI P10 lite",
	"IIIF150 R2022",
	"KXD A1",
	"Lenovo P2",
	"Samsung Galaxy A04s",
	"Xiaomi POCO X5",
	"Xiaomi Redmi Go",
	"motorola moto g",
	"motorola moto one",
	"motorola moto one",
	"motorola edge 30 neo",
	"motorola edge 30 ultra",
	"motorola edge 20 lite",
	"motorola moto e20",
	"motorola moto g5 plus",
	"motorola moto g72",
	"realme X50 Pro",
	"Nokia G21",
	"Nokia 230",
	"Nokia 230",
	"Nokia X30",
	"Nothing Phone (1)",
	"Nothing Phone (1)",
	"Nothing Phone (1)",
	"OPPO Reno4 Z",
	"OSCAL C20 Pro",
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
	"OUKITEL C31 Pro",
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
	"realme 6",
	"realme 7",
	"realme 8",
	"realme 9",
	"realme 9",
	"realme 9 Pro+",
	"realme 9 Pro+",
	"realme GT",
	"realme narzo 50",
	"realme narzo 50",
	"realme narzo 50A Prime",
	"realme C31",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi Note 9",
	"Samsung Galaxy A03",
	"Samsung Galaxy A13",
	"Samsung Galaxy A12",
	"Samsung Galaxy A13",
	"Samsung Galaxy A14",
	"Samsung Galaxy A22",
	"Samsung Galaxy A23",
	"Samsung Galaxy Z Fold 4",
	"Samsung Galaxy A22",
	"Samsung Galaxy A02s",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A14",
	"Samsung Galaxy A32",
	"Samsung Galaxy A34",
	"Samsung Galaxy M52",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22 Ultra",
	"Samsung Galaxy S22+",
	"Samsung Galaxy XCover Pro",
	"Samsung Galaxy XCover 6 Pro",
	"Samsung Galaxy XCover 6 Pro",
	"Samsung Galaxy A33",
	"Samsung Galaxy M13",
	"Samsung Galaxy M13",
	"Samsung Galaxy A13",
	"Samsung Galaxy S10+",
	"Samsung Galaxy A13",
	"motorola moto g82",
	"OUKITEL C25",
	"OSCAL C20 Pro",
	"Samsung Galaxy A13",
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
	"Ulefone Armor X10 Pro",
	"Ulefone Note 14 Pro",
	"Ulefone Note 14 Pro",
	"Ulefone Note 14 Pro",
	"Ulefone Power Armor 13",
	"Ulefone Power Armor 14 Pro",
	"Ulefone Power Armor X11 Pro",
	"UMIDIGI Power 7S",
	"UMIDIGI G1 Max",
	"UMIDIGI G1 Max",
	"vivo Y76",
	"XGODY Y13",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10A",
	"Xiaomi Redmi 10 2022",
	"Xiaomi 11T",
	"Xiaomi Redmi Note 12",
	"Xiaomi POCO M5",
	"Xiaomi Mi 9T Pro",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi POCO C40",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi Note 8 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi Redmi Note 12 Pro",
	"Xiaomi POCO M5",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi Note 12",
	"Xiaomi POCO F4",
	"Xiaomi Redmi A2",
	"ZTE Blade V40 Vita",
	"ZTE Blade V40 Vita",
	"ZTE Blade A51",
	"ZTE Blade A32",
	"ZTE Blade V40s",
	"ZTE Blade V40",
	"ZTE Blade V40 Vita",
	"ZTE Blade V40 Vita",
}

func TestLint(t *testing.T) {
	for i, name := range amazonNames {
		if _name := shop.AmazonCleanFn(name); _name != amazonNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, amazonNamesExpected[i], name)
		}
	}
}
