package main

import "testing"

func TestMagicBytesToParams(t *testing.T) {
	mainmb := []byte{0xf9, 0xbe, 0xb4, 0xd9}
	testmb := []byte{0x0b,0x11,0x09,0x07}
	param, err := MagicBytesToParams(mainmb)
	if err != nil {
		t.Errorf("TestMagicBytesToParams error: %s", err)
	}
	if param != &MainnetParams{
		t.Errorf("param is not correct, we got: %p, but answer is: %p", param, &MainnetParams)
	}

	param, err = MagicBytesToParams(testmb)
	if err != nil {
		t.Errorf("TestMagicBytesToParams error: %s", err)
	}
	if param != &TestnetParams{
		t.Errorf("param is not correct, we got: %p, but answer is: %p", param, &TestnetParams)
	}
}
