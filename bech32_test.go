package main

import (
	"bytes"
	"strings"
	"testing"
)

var testChecksum = []struct {
	str   string
	valid bool
}{
	{"A12UEL5L", true},
	{"a12uel5l", true},
	{"an83characterlonghumanreadablepartthatcontainsthenumber1andtheexcludedcharactersbio1tt5tgs", true},
	{"abcdef1qpzry9x8gf2tvdw0s3jn54khce6mua7lmqqqxw", true},
	{"11qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqc8247j", true},
	{"split1checkupstagehandshakeupstreamerranterredcaperred2y9e3w", true},                                // invalid checksum
	{"1nwldj5", false},
	{"an84characterslonghumanreadablepartthatcontainsthenumber1andtheexcludedcharactersbio1569pvx", false},
	{"pzry9x0s0muk", false},
	{"1pzry9x0s0muk", false},
	{"x1b4n0q5v", false},
	{"li1dgmt3", false},
	{"de1lg7wt", false},
	{"A1G7SGD8", false},
	{"10a06t8", false},
	{"1qzzfhee", false},
}

var testSegwit = []struct{
	addr string
	hrp  string
	witver []byte
	witprog []byte
}{
	{
		addr:"bc1pw508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7k7grplx",
		hrp: "bc",
		witver: []byte{0x01},
		witprog:[]byte{
			0x75,0x1e,0x76,0xe8,0x19,0x91,
			0x96,0xd4,0x54,0x94,0x1c,0x45,0xd1,0xb3,
			0xa3,0x23,0xf1,0x43,0x3b,0xd6,0x75,0x1e,
			0x76,0xe8,0x19,0x91,0x96,0xd4,0x54,0x94,
			0x1c,0x45,0xd1,0xb3,0xa3,0x23,0xf1,0x43,
			0x3b,0xd6,
		},
	},
	{
		addr:"bc1zw508d6qejxtdg4y5r3zarvaryvg6kdaj",
		hrp: "bc",
		witver: []byte{0x02},
		witprog: []byte{
			0x75,0x1e,0x76,0xe8,0x19,0x91,
			0x96,0xd4,0x54,0x94,0x1c,0x45,0xd1,0xb3,
			0xa3,0x23,
		},
	},
	{
		addr:"bc1q2g2cxzx29eg5nkfewvgwc8lz5nug47c8ta5ene",
		hrp: "bc",
		witver: []byte{0x00},
		witprog: []byte{
			0x52,0x15,0x83,0x08,0xca,0x2e,0x51,0x49,
			0xd9,0x39,0x73,0x10,0xec,0x1f,0xe2,0xa4,
			0xf8,0x8a,0xfb,0x07,
		},
	},
}

func TestChecksum(t *testing.T) {
	for _, tc := range testChecksum {
		_, _, err := bech32Decode(tc.str)
		if err != nil {
			if tc.valid == true {
				t.Errorf("TestValidChecksum error: %s, str:%s", err,tc.str)
			}else {
				t.Logf("TestValidChecksum error: %s, str:%s", err, tc.str)
			}
		}
	}
}

func TestSegwitEncode(t *testing.T) {
	for _, test := range testSegwit {
		addr, err := SegwitEncode(test.hrp, test.witver, test.witprog)
		if err != nil {
			t.Errorf("TestSegwitEncode error: %s", err)
		}
		if strings.Compare(addr, test.addr) != 0 {
			t.Errorf("address is not correct, we got: %s, but answer is : %s", addr, test.addr)
		}
	}
}

func TestSegwitDecode(t *testing.T) {
	for _, test := range testSegwit {
		ver, decode, err := SegwitDecode(test.hrp, test.addr)
		if err != nil {
			t.Errorf("TestSegwitDncode error: %s", err)
		}
		if ver != int(test.witver[0]) {
			t.Errorf("witver is not correct, we got: %d, but ans is: %d", ver, test.witver[0])
		}
		if bytes.Compare(decode, test.witprog) != 0 {
			t.Errorf("decoded bytes is not correct, we got: %x, but ans is: %x", decode, test.witprog)
		}
	}
}