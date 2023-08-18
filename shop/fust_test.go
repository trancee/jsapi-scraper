package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var fustNames = []string{
	"Fairphone 4 - 128 GB, Grey, 6.3'', 48 MP, 5G",
	"Fairphone 4 - 256 GB, Green, 6.3'', 48 MP, 5G",
	"Fairphone 4 - 256 GB, Grey, 6.3'', 48 MP, 5G",
	"Fairphone Fairphone 4 5G 256 GB Grün",
	"Fust A62 Lite - 64 GB, Black, 6.1'', 13 MP",
	"Fust A72 - 64 GB, Space Gray, 6.5'', 13 MP",
	"Google Pixel 6a 128 GB Charcoal",
	"Google Pixel 7 - 128 GB, Black, 6.3\", 50 MP, 5G",
	"Google Pixel 7 - 128 GB, Snow white, 6.3\", 50 MP, 5G",
	"Google Pixel 7 256 GB Obsidian",
	"Google Pixel 7 Pro - 128 GB, Black, 6.7\", 50 MP, 5G",
	"Huawei Mate S Champagne",
	"Huawei nova 9 - 128 GB, Black, 6.57\", 50 MP, 4G",
	"Huawei nova 9 - 128 GB, Starry Blue, 6.57\", 50 MP, 4G",
	"Huawei Nova gold",
	"Huawei P smart 2020 - 128 GB, Green, 6.21\", 13 MP, 4G",
	"Huawei P smart 2021 - 128 GB, Crush Green, 6.67\" 48 MP, 4G",
	"Huawei P smart 2021 - 128 GB, Midnight Black, 6.67\" 48 MP, 4G",
	"Huawei P30 - 128 GB, Aurora, 6.1\", 40 MP, 4G",
	"Huawei P30 - 128 GB, Breathing Crystal, 6.1\", 40 MP, 4G",
	"Huawei P30 lite Black + Freebuds 3i",
	"Huawei P30 Pro Black + Freebuds 3i",
	"Huawei P30 Pro New Blk + Freebuds 3i",
	"Huawei P30 Pro New Slv + Freebuds 3i",
	"Huawei P40 - 128 GB, Black, 6.1\", 50 MP, 5G",
	"Huawei P40 - 128 GB, Silver Frost, 6.1\", 50 MP, 5G",
	"Huawei P40 5G Black - Ohne Google Services und FreeBuds 3i Bluetooth-Headset",
	"Huawei P40 5G Blush Gold-Ohne Google Services und FreeBuds 3i Bluetooth-Headset",
	"Huawei P40 5G Frost Silver-Ohne Google Services und FreeBuds 3i Bluetooth-Headset",
	"Huawei P40 Lite Black + Freebuds 3i",
	"Huawei P40 Lite E Blk + Freebuds 3i",
	"Huawei P40 Pro - 256 GB, Black, 6.58\", 50 MP, 5G",
	"Huawei P40 Pro 5G Black - Ohne Google Services und FreeBuds 3i Bluetooth-Headset",
	"Huawei P40 Pro 5G Silver - Ohne Google Service und FreeBuds 3i Bluetooth-Headset",
	"Huawei P9 lite black",
	"Huawei Y6P Midnight Bk + Freebuds 3i",
	"Motorola Moto g32 - 128 GB, Dove Gray, 6.5\", 50 MP, 4G",
	"Nokia 3.4 Grey",
	"Nokia 4.2 Pink",
	"Nokia C21 32 GB Blau",
	"Nokia G11 - 32 GB, Charcoal, 6.52\", 13 MP, 4G",
	"Nokia G21 128 GB Dusk",
	"Nokia G50 - 128 GB, Midnight Sun, 6.82\", 48 MP, 5G",
	"Nokia G50 - 128 GB, Ocean Blue, 6.82\", 48 MP, 5G",
	"Nokia X20 - 128 GB, Midnight Sun, 6.67\", 64 MP, 5G",
	"Nokia X20 - 128 GB, Nordic Blue, 6.67\", 64 MP, 5G",
	"Nokia X20 n.blue 8/128 GB",
	"Nokia X30 5G 128 GB Cloudy blue",
	"Nokia X30 5G 128 GB Ice white",
	"Nokia X30 5G 256 GB Cloudy blue",
	"Nothing Phones Phone (1) 8 GB / 256 GB",
	"OnePlus 11 5G 128 GB Titan Black",
	"OnePlus 11 5G 256 GB Titan Black",
	"Oppo A16s - 64 GB, Crystal Black, 6.5\", 13 MP, 4G",
	"Oppo A5 2020 - 64 GB, Dazzling White, 6.5\", 12 MP, 4G",
	"Oppo A5 2020 Black",
	"Oppo A5 2020 White + Band Sport",
	"Oppo A54 - 64 GB, Fantastic Purple, 6.5\", 48 MP, 5G",
	"Oppo A54 - 64 GB, Fluid Black, 6.5\", 48 MP, 5G",
	"Oppo A54s - 128 GB, Crystal Black, 6.5\", 50 MP, 4G",
	"Oppo A57s - 128 GB, Starry Black, 6.56\", 50 MP, 4G",
	"Oppo A57s 128 GB Sky Blue",
	"Oppo A74 5G Space Silver",
	"Oppo A9 2020 Marine Green",
	"Oppo A9 2020 Space Purple",
	"Oppo A91 - 128 GB, Blazing Blue, 6.4\", 48 MP, 4G",
	"Oppo A91 - 128 GB, Crystal Black, 6.4\", 48 MP, 4G",
	"Oppo A94 5G 128 GB Cosmo Blue",
	"Oppo A94 5G Fluid Black",
	"Oppo A96 - 128 GB, starry black, 6.59\", 50 MP, 4G",
	"Oppo A96 128 GB Sunset Blue",
	"Oppo Find X2 Lite Moonlight Black",
	"Oppo Find X2 Neo - 256 GB, Starry Blue, 6.5\", 48 MP, 5G",
	"Oppo Find X2 Neo - Moonlight Black, 6.5\", 48 MP, 5G",
	"Oppo Find X2 Pro schwarz",
	"Oppo Find X3 Lite - 128 GB, Starry Black, 6.43\", 64 MP, 5G",
	"Oppo Find X3 Lite Astral Blue",
	"Oppo Find X3 Lite Galactic Silver",
	"Oppo Find X3 Neo – 256 GB, Starlight Black, 6.5'', 50MP, 5G",
	"Oppo Find X3 Neo - 256 GB, Galactic Silver, 6.5\", 50 MP, 5G",
	"Oppo Find X3 Neo Starlight Black",
	"Oppo Find X3 Pro - 256 GB, Gloss Black, 6.7\", 50 MP, 5G",
	"Oppo Find X3 Pro 256 GB Blue",
	"Oppo Find X5 - 256 GB, Black, 6.55\", 50 MP, 5G",
	"Oppo Find X5 - 256 GB, White, 6.55\", 50 MP, 5G",
	"Oppo Find X5 Lite - 256 GB, Lite Black, 6.43\", 64 MP, 5G",
	"Oppo Find X5 Lite - 256 GB, Lite Blue, 6.43\", 64 MP, 5G",
	"Oppo Find X5 Pro - 256 GB, Ceramic White, 6.7\", 50 MP, 5G",
	"Oppo Find X5 Pro - 256 GB, Glaze Black, 6.7\", 50 MP, 5G",
	"Oppo Reno Green + Band Sport",
	"Oppo Reno - 256 GB, Ocean Green, 6\", 48 MP, 4G",
	"Oppo Reno 2 Black",
	"Oppo Reno 4 Pro - 256 GB, Space Black, 6.5\", 48 MP, 5G",
	"Oppo Reno 4 Pro 5G Galactic Blue",
	"Oppo Reno 6 5G Black + Band Sport",
	"Oppo Reno 6 Pro 5G Lunar Grey",
	"Oppo Reno 8 - 256 GB, shimmer black, 6.43\", 50 MP, 5G",
	"Oppo Reno 8 - 256 GB, shimmer gold, 6.43\", 50 MP, 5G",
	"Oppo Reno 8 Lite 5G cosmic black",
	"Oppo Reno 8 Lite 5G rainbow spectrum",
	"Oppo Reno 8 Pro - 256 GB, glazed black, 6.7\", 50 MP, 5G",
	"Oppo Reno 8 Pro - 256 GB, glazed green, 6.7\", 50 MP, 5G",
	"Oppo Reno2 256GB Ocean Blue",
	"Oppo Reno4 - 128 GB, Space Black, 6.4\", 48 MP, 5G",
	"Oppo Reno4 Z - 128 GB, Dew White, 6.57\", 48 MP, 5G",
	"Oppo Reno4 Z - 128 GB, Ink Black, 6.57\", 48 MP, 5G",
	"Oppo Reno 4Z 5G Black + Band Sport",
	"Oppo Reno6 - 128 GB, Arctic Blue, 6.43\", 64 MP, 5G",
	"Oppo Reno6 - 128 GB, Stellar Black, 6.43\", 64 MP, 5G",
	"Oppo Reno6 Pro - 256 GB, Arctic Blue, 6.5\", 50 MP, 5G",
	"Realme 8 5G Black 4+64GB",
	"Realme 8 5G Blue 4+64GB",
	"Realme C11 Pepper Grey",
	"Samsung Galaxy A13 128 GB CH Black",
	"Samsung Galaxy A14 128 GB CH Black",
	"Samsung Galaxy A14 128 GB CH Lime Green",
	"Samsung Galaxy A14 128 GB CH Silver",
	"Samsung Galaxy A14 5G 128 GB CH Black",
	"Samsung Galaxy A14 5G 128 GB CH Lime Green",
	"Samsung Galaxy A14 5G 128 GB CH Silver",
	"Samsung Galaxy A14, 64GB, Black, 6.6\", 50MP",
	"Samsung Galaxy A14, 64GB, Black, 6.6\", 50MP + gratis Charger",
	"Samsung Galaxy A14, 64GB, Silver, 6.6\", 50MP",
	"Samsung Galaxy A14, 64GB, Silver, 6.6\", 50MP + gratis Charger",
	"Samsung Galaxy A33 - 128 GB, Awesome Black, 6.4\", 48 MP, 5G",
	"Samsung Galaxy A33 - 128 GB, Awesome White, 6.4\", 48 MP, 5G",
	"Samsung Galaxy A34 - 128 GB, Awesome Graphite, 6.6\", 48 MP, 5G",
	"Samsung Galaxy A34 - 128 GB, Awesome Graphite, 6.6\", 48 MP, 5G + gratis Charger",
	"Samsung Galaxy A34 - 128 GB, Awesome Lime, 6.6\", 48 MP, 5G",
	"Samsung Galaxy A34 - 128 GB, Awesome Lime, 6.6\", 48 MP, 5G + gratis Charger",
	"Samsung Galaxy A34 - 128 GB, Awesome Violet, 6.6\", 48 MP, 5G",
	"Samsung Galaxy A34 - 128 GB, Awesome Violet, 6.6\", 48 MP, 5G + gratis Charger",
	"Samsung Galaxy A34 - 128 GB, Awesome White, 6.6\", 48 MP, 5G",
	"Samsung Galaxy A34 - 128 GB, Awesome White, 6.6\", 48 MP, 5G + gratis Charger",
	"Samsung Galaxy A34 5G 128 GB CH Enterprise Edition Awesome Graphite",
	"Samsung Galaxy A54 - 128 GB, Awesome Graphite, 6.4\", 50 MP, 5G",
	"Samsung Galaxy A54 - 128 GB, Awesome Graphite, 6.4\", 50 MP, 5G + gratis Charger",
	"Samsung Galaxy A54 - 128 GB, Awesome Lime, 6.4\", 50 MP, 5G",
	"Samsung Galaxy A54 - 128 GB, Awesome Lime, 6.4\", 50 MP, 5G + gratis Charger",
	"Samsung Galaxy A54 - 128 GB, Awesome Violet, 6.4\", 50 MP, 5G",
	"Samsung Galaxy A54 - 128 GB, Awesome Violet, 6.4\", 50 MP, 5G + gratis Charger",
	"Samsung Galaxy A54 - 128 GB, Awesome White, 6.4\", 50 MP, 5G",
	"Samsung Galaxy A54 - 128 GB, Awesome White, 6.4\", 50 MP, 5G + gratis Charger",
	"Samsung Galaxy A54 5G 256 GB CH Awesome Lime",
	"Samsung Galaxy S21 FE - 128 GB, Graphite, 6.4\", 12 MP, 5G",
	"Samsung Galaxy S22 - 128 GB, Green, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S22 - 128 GB, Phantom Black, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S22 - 128 GB, Pink Gold, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S22 - 256 GB, Phantom Black, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S22 Ultra - 256 GB, Phantom Black, 6.8\", 108 MP, 5G",
	"Samsung Galaxy S23 - 128 GB, Cream, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S23 - 128 GB, Green, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S23 - 128 GB, Lavender, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S23 - 128 GB, Phantom Black, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S23 - 256 GB, Cream, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S23 - 256 GB, Green, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S23 - 256 GB, Lavender, 6.1\", 50 MP, 5G",
	"Samsung Galaxy S23 Ultra - 256 GB, Cream, 6.8\", 200 MP, 5G",
	"Samsung Galaxy S23 Ultra - 256 GB, Green, 6.8\", 200 MP, 5G",
	"Samsung Galaxy S23 Ultra - 256 GB, Lavender, 6.8\", 200 MP, 5G",
	"Samsung Galaxy S23 Ultra - 512 GB, Cream, 6.8\", 200 MP, 5G",
	"Samsung Galaxy S23 Ultra - 512 GB, Green, 6.8\", 200 MP, 5G",
	"Samsung Galaxy S23 Ultra - 512 GB, Lavander, 6.8\", 200 MP, 5G",
	"Samsung Galaxy S23 Ultra - 512 GB, Phantom Black, 6.8\", 200 MP, 5G",
	"Samsung Galaxy S23+ - 256 GB, Cream, 6.6\", 50 MP, 5G",
	"Samsung Galaxy S23+ - 256 GB, Green, 6.6\", 50 MP, 5G",
	"Samsung Galaxy S23+ - 256 GB, Lavender, 6.6\", 50 MP, 5G",
	"Samsung Galaxy S23+ - 256 GB, Phantom Black, 6.6\", 50 MP, 5G",
	"Samsung Galaxy S23+ - 512 GB, Cream, 6.6\", 50 MP, 5G",
	"Samsung Galaxy S23+ - 512 GB, Green, 6.6\", 50 MP, 5G",
	"Samsung Galaxy S23+ - 512 GB, Lavender, 6.6\", 50 MP, 5G",
	"Samsung Galaxy S23+ - 512 GB, Phantom Black, 6.6\", 50 MP, 5G",
	"Samsung Galaxy XCover 5 Enterprise Edition - 64 GB, Black, 5.3\", 16 MP, 4G",
	"Samsung Galaxy XCover 6 Pro Enterprise Edition CH",
	"Samsung Galaxy Z Flip3 - 256 GB, Cream, 6.7\", 12 MP, 5G",
	"Samsung Galaxy Z Flip4 - 128 GB, Graphite, 6.7\", 12 MP, 5G",
	"Samsung Galaxy Z Flip4 - 128 GB, Purple, 6.7\", 12 MP, 5G",
	"Samsung Galaxy Z Flip4 - 256 GB, Blue, 6.7\", 12 MP, 5G",
	"Samsung Galaxy Z Flip4 - 256 GB, Graphite, 6.7\", 12 MP, 5G",
	"Samsung Galaxy Z Fold4 - 512 GB, Black, 7.6\", 50 MP, 5G",
	"Samsung Galaxy Z Fold4 - 512 GB, Graygreen, 7.6\", 50 MP, 5G",
	"Samsung Speicherkarte + Motorola Moto e40",
	"Wiko View 3 Lite Blue",
	"Wiko View 3 Lite Arctic Bleen",
	"Xiaomi 11 Lite 5G NE - 128 GB, Truffle Black, 6.55\", 64 MP, 5G",
	"Xiaomi 11T - 128 GB, Meteorite Grey, 6.67\", 108 MP, 5G",
	"Xiaomi 13 256 GB Schwarz",
	"Xiaomi 13 256 GB Weiss",
	"Xiaomi 13 Lite 128 GB Blau",
	"Xiaomi 13 Lite 128 GB Pink",
	"Xiaomi 13 Lite 128 GB Schwarz",
	"Xiaomi Mi 10 Lite - 128 GB, Cosmic Grey, 6.57\", 48 MP, 5G",
	"Xiaomi Mi 10T Lite - 128 GB, Pearl Grey, 6.67\", 64 MP, 5G",
	"Xiaomi Mi 11 Lite - 128 GB, Truffle Black, 6.55\", 64 MP, 5G",
	"Xiaomi Mi 8 Black",
	"Xiaomi Mi 8 Lite black",
	"Xiaomi Mi Mix 2S - 64 GB, White, 5.99\", 12 MP, 4G",
	"Xiaomi Poco F3 - 256 GB, Moonlight Silver, 6.67\", 48 MP, 5G",
	"Xiaomi Poco F3 - 256 GB, Night Black, 6.67\", 48 MP, 5G",
	"Xiaomi Poco X3 Shadow Gray",
	"Xiaomi Redmi 10 2022 128 GB Carbon Gray",
	"Xiaomi Redmi 10C - 64GB, Graphite Gray, 6.71\", 50MP, 4G/LTE",
	"Xiaomi Redmi 10C - 64GB, Mint Green, 6.71\", 50MP, 4G/LTE",
	"Xiaomi Redmi 10C - 64GB, Ocean Blue, 6.71\", 50MP, 4G/LTE",
	"Xiaomi Redmi 12C 128 GB Gray",
	"Xiaomi Redmi 8A Ocean Blue",
	"Xiaomi Redmi 9A - 32 GB, Granite Grey, 6.53\", 13 MP, 4G",
	"Xiaomi Redmi 9A - 32 GB, Peacock Green, 6.53\", 13 MP, 4G",
	"Xiaomi Redmi 9C - 64 GB, Midnight Grey, 6.53\", 13 MP, 4G",
	"Xiaomi Redmi 9C 128 GB Midnight Grey",
	"Xiaomi Redmi 9T - 64 GB, Carbon Grey, 6.53\", 48 MP, 4G",
	"Xiaomi Redmi Note 10 5G 4/128GB Graphite Gray",
	"Xiaomi Redmi Note 10S - 128 GB, Onyx Grey, 6.43\", 64 MP, 4G",
	"Xiaomi Redmi Note 11 - 128 GB, Graphite Grey, 6.43\", 50 MP, 4G",
	"Xiaomi Redmi Note 11 Pro - 128GB, Graphite Gray, 6.67\", 108 MP, 5G",
	"Xiaomi Redmi Note 11S 128 GB Twilight Blue",
	"Xiaomi Redmi Note 12 - 128 GB, Onyx Gray, 6.67'', 50 MP, 4G",
	"Xiaomi Redmi Note 12 128 GB Blau",
	"Xiaomi Redmi Note 12 128 GB Grün",
	"Xiaomi Redmi Note 7",
	"Xiaomi Redmi Note 8T - 32 GB, Moonshadow Grey, 6.3\", 48 MP, 4G",
}

var fustNamesExpected = []string{
	"Fairphone 4",
	"Fairphone 4",
	"Fairphone 4",
	"Fairphone 4",
	"Inoi A62 Lite",
	"Inoi A72",
	"Google Pixel 6a",
	"Google Pixel 7",
	"Google Pixel 7",
	"Google Pixel 7",
	"Google Pixel 7 Pro",
	"HUAWEI Mate S",
	"HUAWEI nova 9",
	"HUAWEI nova 9",
	"HUAWEI nova",
	"HUAWEI P smart 2020",
	"HUAWEI P smart 2021",
	"HUAWEI P smart 2021",
	"HUAWEI P30",
	"HUAWEI P30",
	"HUAWEI P30 lite",
	"HUAWEI P30 Pro",
	"HUAWEI P30 Pro",
	"HUAWEI P30 Pro",
	"HUAWEI P40",
	"HUAWEI P40",
	"HUAWEI P40",
	"HUAWEI P40",
	"HUAWEI P40",
	"HUAWEI P40 lite",
	"HUAWEI P40 lite E",
	"HUAWEI P40 Pro",
	"HUAWEI P40 Pro",
	"HUAWEI P40 Pro",
	"HUAWEI P9 lite",
	"HUAWEI Y6p",
	"motorola moto g32",
	"Nokia 3.4",
	"Nokia 4.2",
	"Nokia C21",
	"Nokia G11",
	"Nokia G21",
	"Nokia G50",
	"Nokia G50",
	"Nokia X20",
	"Nokia X20",
	"Nokia X20",
	"Nokia X30",
	"Nokia X30",
	"Nokia X30",
	"Nothing Phone (1)",
	"OnePlus 11",
	"OnePlus 11",
	"OPPO A16s",
	"OPPO A5 2020",
	"OPPO A5 2020",
	"OPPO A5 2020",
	"OPPO A54",
	"OPPO A54",
	"OPPO A54s",
	"OPPO A57s",
	"OPPO A57s",
	"OPPO A74",
	"OPPO A9 2020",
	"OPPO A9 2020",
	"OPPO A91",
	"OPPO A91",
	"OPPO A94",
	"OPPO A94",
	"OPPO A96",
	"OPPO A96",
	"OPPO Find X2 Lite",
	"OPPO Find X2 Neo",
	"OPPO Find X2 Neo",
	"OPPO Find X2 Pro",
	"OPPO Find X3 Lite",
	"OPPO Find X3 Lite",
	"OPPO Find X3 Lite",
	"OPPO Find X3 Neo",
	"OPPO Find X3 Neo",
	"OPPO Find X3 Neo",
	"OPPO Find X3 Pro",
	"OPPO Find X3 Pro",
	"OPPO Find X5",
	"OPPO Find X5",
	"OPPO Find X5 Lite",
	"OPPO Find X5 Lite",
	"OPPO Find X5 Pro",
	"OPPO Find X5 Pro",
	"OPPO Reno",
	"OPPO Reno",
	"OPPO Reno2",
	"OPPO Reno4 Pro",
	"OPPO Reno4 Pro",
	"OPPO Reno6",
	"OPPO Reno6 Pro",
	"OPPO Reno8",
	"OPPO Reno8",
	"OPPO Reno8 Lite",
	"OPPO Reno8 Lite",
	"OPPO Reno8 Pro",
	"OPPO Reno8 Pro",
	"OPPO Reno2",
	"OPPO Reno4",
	"OPPO Reno4 Z",
	"OPPO Reno4 Z",
	"OPPO Reno4 Z",
	"OPPO Reno6",
	"OPPO Reno6",
	"OPPO Reno6 Pro",
	"realme 8",
	"realme 8",
	"realme C11",
	"Samsung Galaxy A13",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
	"Samsung Galaxy A14",
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
	"Samsung Galaxy A34",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy A54",
	"Samsung Galaxy S21 FE",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22",
	"Samsung Galaxy S22 Ultra",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23 Ultra",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy S23+",
	"Samsung Galaxy XCover 5",
	"Samsung Galaxy XCover 6 Pro",
	"Samsung Galaxy Z Flip3",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Flip4",
	"Samsung Galaxy Z Fold4",
	"Samsung Galaxy Z Fold4",
	"motorola moto e40",
	"Wiko View 3 Lite",
	"Wiko View 3 Lite",
	"Xiaomi 11 Lite",
	"Xiaomi 11T",
	"Xiaomi 13",
	"Xiaomi 13",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Lite",
	"Xiaomi 13 Lite",
	"Xiaomi Mi 10 Lite",
	"Xiaomi Mi 10T Lite",
	"Xiaomi Mi 11 Lite",
	"Xiaomi Mi 8",
	"Xiaomi Mi 8 Lite",
	"Xiaomi Mi Mix 2S",
	"Xiaomi POCO F3",
	"Xiaomi POCO F3",
	"Xiaomi POCO X3",
	"Xiaomi Redmi 10 2022",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 10C",
	"Xiaomi Redmi 12C",
	"Xiaomi Redmi 8A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi 9C",
	"Xiaomi Redmi 9T",
	"Xiaomi Redmi Note 10",
	"Xiaomi Redmi Note 10S",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11S",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi Note 7",
	"Xiaomi Redmi Note 8T",
}

func TestFustClean(t *testing.T) {
	for i, name := range fustNames {
		if _name := shop.FustCleanFn(name); _name != fustNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, fustNamesExpected[i], name)
		}
	}
}
