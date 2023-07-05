package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var mobiledeviceNames = []string{
	"Apple iPhone SE (2020) 128GB - (PRODUCT)RED",
	"Apple iPhone SE (2020) 256GB - Black EU",
	"Apple iPhone SE (2020) 256GB - Red EU",
	"Apple iPhone SE (2020) 64GB - schwarz",
	"Apple iPhone XR 64GB (Product)Red",
	"Apple iPhone XR 64GB Schwarz",
	"Google Pixel 4 64GB - Clearly White DE",
	"Google Pixel 6 128GB - stormy black (schwarz)",
	"Google Pixel 6a 5G 128GB - Charcoal DE",
	"Google Pixel 6a 5G 128GB - grÃ¼n (Sage Green)",
	"Google Pixel 6a 5G 128GB - Weiss (Chalk White)",
	"Honor 70 5G Dual Sim 8GB RAM 128GB - Emerald Green EU",
	"Honor 70 5G Dual Sim 8GB RAM 128GB - Midnight Black",
	"Honor 70 5G Dual Sim 8GB RAM 256GB - Emerald Green EU",
	"Honor 70 5G Dual Sim 8GB RAM 256GB - Midnight Black",
	"Honor Magic4 Lite 5G Dual Sim 6GB RAM 128GB - Blau (Ocean Blue)",
	"Honor Magic5 Lite 5G Dual Sim 6GB RAM 128GB - Midnight Black",
	"HTC Desire 22 Pro 5G Dual Sim 8GB RAM 128GB - Gold (Flowing Gold)",
	"HTC Desire 22 Pro 8GB RAM 128GB - schwarz (flowing Black)",
	"Huawei P30 Dual Sim 128GB - Aurora Blue EU",
	"Motorola XT2083-9 Defy (2021) Dual Sim 4GB RAM 64GB - Schwarz",
	"Motorola XT2169-1 Moto G71 5G Dual Sim 6GB RAM 128GB - Iron Black",
	"Motorola XT2239-18 Moto E22i Dual Sim 2GB RAM 32GB - Graphite Grey",
	"Nokia X30 Dual Sim 5G 6GB RAM 128GB - weiss (Ice White)",
	"Nokia XR20 Dual Sim 5G 4GB RAM 64GB - Granite Grey",
	"Nokia XR20 Dual Sim 5G 4GB RAM 64GB - Ultra Blue",
	"Nothing Phone 1 5G Dual Sim 8GB RAM 128GB - schwarz",
	"Nothing Phone 1 5G Dual Sim 8GB RAM 256GB - Black EU",
	"Oppo A57s Dual Sim 4GB RAM 128GB - schwarz (Starry Black)",
	"Oppo Find X3 Lite 5G 8GB RAM 128GB - schwarz (Starry Black)",
	"Oppo Reno6 5G Dual Sim 8GB RAM 128GB - blau (Blue)",
	"Oppo Reno6 5G Dual Sim 8GB RAM 128GB - schwarz (Stellar Black)",
	"Oppo Reno8 Lite 5G Dual Sim 8GB RAM 128GB - Cosmic Black EU",
	"Realme 9 Dual Sim 8GB RAM 128GB - Stargaze White",
	"Realme 9 Pro 5G Dual Sim 6GB RAM 128GB - Aurora Green",
	"Realme 9 Pro 5G Dual Sim 6GB RAM 128GB - Midnight Black EU",
	"Realme 9 Pro 5G Dual Sim 6GB RAM 128GB - Sunrise Blue",
	"Realme 9 Pro+ 5G Dual Sim 6GB RAM 128GB - Aurora Green EU",
	"Realme 9 Pro+ 5G Dual Sim 6GB RAM 128GB - Midnight Black",
	"Realme 9 Pro+ 5G Dual Sim 6GB RAM 128GB - Sunrise Blue",
	"Realme GT 2 5G Dual Sim 12GB RAM 256GB - Paper White",
	"Realme GT 2 5G Dual Sim 8GB RAM 128GB - Paper Green EU",
	"realme GT Master Edition Dual Sim 6GB RAM 128GB - Luna White",
	"Samsung Galaxy A04S A047 (2022) Dual Sim 3GB RAM 32GB - schwarz",
	"Samsung Galaxy A13 A136 5G Dual Sim 4GB RAM 128GB - Blau",
	"Samsung Galaxy A13 A136 5G Dual Sim 4GB RAM 128GB - weiss",
	"Samsung Galaxy A13 A136 5G Dual Sim 4GB RAM 64GB - Blue EU",
	"Samsung Galaxy A13 A137 Dual Sim 3GB RAM 32GB - Black EU",
	"Samsung Galaxy A13 A137 Dual Sim 4GB RAM 64GB - schwarz EU",
	"Samsung Galaxy A23 5G A236 Dual Sim 4GB RAM 128GB - Black EU",
	"Samsung Galaxy A23 5G A236 Dual Sim 4GB RAM 64GB - Blue",
	"Samsung Galaxy A32 A325 LTE Dual Sim 4GB RAM 128GB Enterprise Edition - Black EU",
	"Samsung Galaxy A33 5G A336 Dual Sim 6GB RAM 128GB - Awesome White",
	"Samsung Galaxy A52s 5G A528 Dual Sim 6GB RAM 128GB Enterprise Edition - schwarz EU",
	"Samsung Galaxy A53 5G A536 Dual Sim 6GB RAM 128GB - schwarz EU",
	"Samsung Galaxy A53 5G A536 Dual Sim 6GB RAM 128GB Enterprice Edition - schwarz (Awesome Black) EU",
	"Samsung Galaxy A53 5G A536 Dual Sim 8GB RAM 256GB - Awesome Peach",
	"Samsung Galaxy A53 5G A536 Dual Sim 8GB RAM 256GB - Awesome White EU",
	"Samsung Galaxy A53 5G A536 Dual Sim 8GB RAM 256GB - blau (Awesome Blue) EU",
	"Samsung Galaxy A53 5G A536 Dual Sim 8GB RAM 256GB - schwarz (DE)",
	"Samsung Galaxy M23 M236 5G Dual Sim 4GB RAM 128GB - Light Blue EU",
	"Samsung Galaxy M33 M336 5G Dual Sim 6GB RAM 128GB - Brown EU",
	"Samsung Galaxy S20 FE G781 5G Dual Sim 128GB - Cloud Lavender EU",
	"Samsung Galaxy S20 FE G781 5G Dual Sim 128GB - Cloud Navy (EU)",
	"Samsung Galaxy S20 FE G781 5G Dual Sim 128GB - Cloud Navy DE",
	"Samsung Galaxy X Cover 6 Pro G736 128GB Dual Sim Enterprise Edition - Black EU",
	"Samsung Galaxy XCover 5 64GB Dual Sim Enterprise Edition - Schwarz EU",
	"Xiaomi 11T Pro 5G Dual Sim 8GB RAM 256GB - Celestial Blue",
	"Xiaomi 11T Pro 5G Dual Sim 8GB RAM 256GB - Meteorite Gray",
	"Xiaomi 12 5G Dual Sim 8GB RAM 256GB - Blau",
	"Xiaomi 12 Lite 5G Dual Sim 8GB RAM 128GB - schwarz",
	"Xiaomi 12T 5G Dual Sim 8GB RAM 128GB - Blue EU",
	"Xiaomi 12T 5G Dual Sim 8GB RAM 128GB - schwarz",
	"Xiaomi 12T 5G Dual Sim 8GB RAM 128GB - Silber Eu- Modell (gibt kein CH)",
	"Xiaomi 13 Lite 5G Dual Sim 8GB RAM 256GB - schwarz",
	"Xiaomi Poco C40 Dual Sim 4GB RAM 64GB - POCO Yellow EU",
	"Xiaomi Poco F4 5G Dual Sim 8GB RAM 256GB - Moonlight Silver",
	"Xiaomi Poco F4 5G Dual Sim 8GB RAM 256GB - Night Black EU",
	"Xiaomi Poco M4 Pro Dual Sim 8GB RAM 256GB - Cool Blue EU",
	"Xiaomi Poco X5 5G Dual Sim 8GB RAM 256GB - Blue EU",
	"Xiaomi Redmi 10 (2022) Dual Sim 4GB RAM 128GB - Carbon Grey EU",
	"Xiaomi Redmi 10 (2022) Dual Sim 4GB RAM 64GB - Grey EU",
	"Xiaomi Redmi A1 Dual Sim 2GB RAM 32GB - Black EU",
	"Xiaomi Redmi Note 10 Pro Dual Sim 6GB RAM 128GB - Glacier Blue",
	"Xiaomi Redmi Note 10 Pro Dual Sim 6GB RAM 128GB - Onyx Grey",
	"Xiaomi Redmi Note 10 Pro Dual Sim 6GB RAM 128GB Gradient Bronze",
	"Xiaomi Redmi Note 10 Pro Dual Sim 6GB RAM 64GB - Glacier Blue",
	"Xiaomi Redmi Note 10 Pro Dual Sim 6GB RAM 64GB - Onyx Grey",
	"Xiaomi Redmi Note 11 Dual Sim 4GB RAM 128GB - Graphite Gray",
	"Xiaomi Redmi Note 11 Dual Sim 4GB RAM 128GB - Twilight Blue EU",
	"Xiaomi Redmi Note 11 Dual Sim 4GB RAM 64GB - Star Blue EU",
	"Xiaomi Redmi Note 11 Dual Sim 6GB RAM 128GB - Twilight Blue EU",
	"Xiaomi Redmi Note 11 Pro 5G Dual Sim 6GB RAM 128GB - Graphite Grey EU",
	"Xiaomi Redmi Note 11 Pro 5G Dual Sim 8GB RAM 128GB - Graphite Grey EU",
	"Xiaomi Redmi Note 11 Pro+ 5G Dual Sim 6GB RAM 128GB - Forest Green EU",
	"Xiaomi Redmi Note 11 Pro+ 5G Dual Sim 6GB RAM 128GB - Graphite Grey EU",
	"Xiaomi Redmi Note 11 Pro+ 5G Dual Sim 6GB RAM 128GB - Star Blue EU",
	"Xiaomi Redmi Note 11s 5G Dual Sim 4GB RAM 128GB - Midnight Black EU",
	"Xiaomi Redmi Note 11s Dual Sim 6GB RAM 128GB - Graphite Gray",
	"Xiaomi Redmi Note 11s Dual Sim 6GB RAM 128GB - Twilight Blue",
	"Xiaomi Redmi Note 11s Dual Sim 6GB RAM 64GB - Graphite Gray",
}

var mobiledeviceNamesExpected = []string{
	"Apple iPhone SE (2020)",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE (2020)",
	"Apple iPhone SE (2020)",
	"Apple iPhone XR",
	"Apple iPhone XR",
	"Google Pixel 4",
	"Google Pixel 6",
	"Google Pixel 6a",
	"Google Pixel 6a",
	"Google Pixel 6a",
	"Honor 70",
	"Honor 70",
	"Honor 70",
	"Honor 70",
	"Honor Magic4 Lite",
	"Honor Magic5 Lite",
	"HTC Desire 22 Pro",
	"HTC Desire 22 Pro",
	"Huawei P30",
	"Motorola Defy",
	"Motorola Moto G71",
	"Motorola Moto E22i",
	"Nokia X30",
	"Nokia XR20",
	"Nokia XR20",
	"Nothing Phone (1)",
	"Nothing Phone (1)",
	"Oppo A57s",
	"Oppo Find X3 Lite",
	"Oppo Reno6",
	"Oppo Reno6",
	"Oppo Reno8 Lite",
	"Realme 9",
	"Realme 9 Pro",
	"Realme 9 Pro",
	"Realme 9 Pro",
	"Realme 9 Pro+",
	"Realme 9 Pro+",
	"Realme 9 Pro+",
	"Realme GT 2",
	"Realme GT 2",
	"realme GT",
	"Samsung Galaxy A04S",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A13",
	"Samsung Galaxy A23",
	"Samsung Galaxy A23",
	"Samsung Galaxy A32",
	"Samsung Galaxy A33",
	"Samsung Galaxy A52s",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy A53",
	"Samsung Galaxy M23",
	"Samsung Galaxy M33",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy S20 FE",
	"Samsung Galaxy XCover 6 Pro",
	"Samsung Galaxy XCover 5",
	"Xiaomi 11T Pro",
	"Xiaomi 11T Pro",
	"Xiaomi 12",
	"Xiaomi 12 Lite",
	"Xiaomi 12T",
	"Xiaomi 12T",
	"Xiaomi 12T",
	"Xiaomi 13 Lite",
	"Xiaomi Poco C40",
	"Xiaomi Poco F4",
	"Xiaomi Poco F4",
	"Xiaomi Poco M4 Pro",
	"Xiaomi Poco X5",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi 10",
	"Xiaomi Redmi A1",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 10 Pro",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11 Pro+",
	"Xiaomi Redmi Note 11s",
	"Xiaomi Redmi Note 11s",
	"Xiaomi Redmi Note 11s",
	"Xiaomi Redmi Note 11s",
}

func TestMobileDeviceClean(t *testing.T) {
	for i, name := range mobiledeviceNames {
		if _name := shop.MobileDeviceCleanFn(name); _name != mobiledeviceNamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, mobiledeviceNamesExpected[i], name)
		}
	}
}
