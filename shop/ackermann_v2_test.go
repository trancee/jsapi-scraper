package shop_test

import (
	shop "jsapi-scraper/shop"
	"testing"
)

var ackermannV2Names = []string{
	" INOI A72 64GB Space Gray",
	" INOI Note 13S 256GB Space Gray",
	"Emporia 5 mini 64 GB",
	"Emporia SIMPLICITY LTE 4G",
	"Gigaset GS5 Senior 64 GB",
	"Gigaset GX4 64 GB Schwarz",
	"Motorola Moto 454",
	"Motorola moto e¹³",
	"Motorola Moto g14",
	"Motorola moto g32",
	"Motorola moto g¹³",
	"Nokia 128 GB Grey",
	"Nokia 64 GB Schwarz",
	"Nokia G22 64GB Lagoon Blue",
	"Nokia G22 64GB Meteor Grey",
	"Oppo 128 GB Glowing Black",
	"Oppo 128 GB Sunset Blue",
	"Oppo A78 Aqua Green",
	"Oppo A78 Mist Black",
	"Oppo Pro 5G Arctic Blue",
	"Xiaomi 10 5G 64 GB Graphite Gray",
	"Xiaomi Note 12 256 GB Schwarz",
	"Xiaomi Redmi 12 128 GB Sky blue",
	"Xiaomi Redmi 12 256 GB Midnight black",
	"Xiaomi Redmi 12 256 GB Sky blue",
	"Xiaomi Redmi 9A 32GB Sky Blue",
	"Xiaomi Redmi A2 32 GB Blau",
	"Xiaomi Redmi A2 32 GB Grün",
	"Xiaomi Redmi Note 11",
}

var ackermannV2NamesExpected = []string{
	"Inoi A72",
	"Inoi Note 13S",
	"emporia5 mini",
	"emporiaSIMPLICITY",
	"Gigaset GS5 senior",
	"Gigaset GX4",
	"motorola moto g54",
	"motorola moto e13",
	"motorola moto g14",
	"motorola moto g32",
	"motorola moto g13",
	"Nokia",
	"Nokia",
	"Nokia G22",
	"Nokia G22",
	"OPPO",
	"OPPO",
	"OPPO A78",
	"OPPO A78",
	"OPPO Pro",
	"Xiaomi 10",
	"Xiaomi Redmi Note 12",
	"Xiaomi Redmi 12",
	"Xiaomi Redmi 12",
	"Xiaomi Redmi 12",
	"Xiaomi Redmi 9A",
	"Xiaomi Redmi A2",
	"Xiaomi Redmi A2",
	"Xiaomi Redmi Note 11",
}

func TestAckermannV2Clean(t *testing.T) {
	for i, name := range ackermannV2Names {
		if _name := shop.AckermannV2CleanFn(name); _name != ackermannV2NamesExpected[i] {
			t.Errorf("given name \"%s\" not match expected name \"%s\"\n%s\n", _name, ackermannV2NamesExpected[i], name)
		}
	}
}
