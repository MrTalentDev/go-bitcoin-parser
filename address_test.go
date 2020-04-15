package main

import "testing"



func TestPublicKeyToAddr(t *testing.T){
	pubkey := []byte{
		0x04,0x96,0xb5,0x38,0xe8,0x53,0x51,0x9c,
		0x72,0x6a,0x2c,0x91,0xe6,0x1e,0xc1,0x16,
		0x00,0xae,0x13,0x90,0x81,0x3a,0x62,0x7c,
		0x66,0xfb,0x8b,0xe7,0x94,0x7b,0xe6,0x3c,
		0x52,0xda,0x75,0x89,0x37,0x95,0x15,0xd4,
		0xe0,0xa6,0x04,0xf8,0x14,0x17,0x81,0xe6,
		0x22,0x94,0x72,0x11,0x66,0xbf,0x62,0x1e,
		0x73,0xa8,0x2c,0xbf,0x23,0x42,0xc8,0x58,
		0xee,
	}
	ans := "12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX"
	addr, err := publicKeyToAddr(pubkey, &MainnetParams)
	if err != nil {
		t.Errorf("Test publicKeyToAddr error: %s", err)
	}
	if addr != ans {
		t.Errorf("address are not correct, we got: %s, but answer is: %s",addr, ans)
	}
}

func TestPublicKeyHashToAddr(t *testing.T){
	pubkey := []byte{
		0x12,0xab,0x8d,0xc5,0x88,0xca,0x9d,0x57,
		0x87,0xdd,0xe7,0xeb,0x29,0x56,0x9d,0xa6,
		0x3c,0x3a,0x23,0x8c,
	}
	ans := "12higDjoCCNXSA95xZMWUdPvXNmkAduhWv"
	addr := publicKeyHashToAddr(pubkey, &MainnetParams)
	if addr != ans {
		t.Errorf("address are not correct, we got: %s, but answer is: %s", addr, ans)
	}
}

func TestScriptHashToAddr(t *testing.T){
	hash := []byte{
		0x74,0x82,0x84,0x39,0x0f,0x9e,0x26,0x3a,
		0x4b,0x76,0x6a,0x75,0xd0,0x63,0x3c,0x50,
		0x42,0x6e,0xb8,0x75,
	}
	ans := "3CK4fEwbMP7heJarmU4eqA3sMbVJyEnU3V"
	addr := scriptHashToAddr(hash, &MainnetParams)
	if addr != ans {
		t.Errorf("address are not correct, we got: %s, but answer is: %s", addr, ans)
	}

}

func TestPublicKeyHashToSegwitAddr(t *testing.T){
	hash := []byte{
		0x52,0x15,0x83,0x08,0xca,0x2e,0x51,0x49,
		0xd9,0x39,0x73,0x10,0xec,0x1f,0xe2,0xa4,
		0xf8,0x8a,0xfb,0x07,
	}
	ans := "bc1q2g2cxzx29eg5nkfewvgwc8lz5nug47c8ta5ene"
	addr, err := publicKeyHashToSegwitAddr(hash, &MainnetParams)
	if err != nil {
		t.Errorf("TestPublicKeyHashToSegwitAddr error: %s", err)
	}
	if addr != ans {
		t.Errorf("address are not correct, we got: %s, but answer is: %s", addr, ans)
	}
}
