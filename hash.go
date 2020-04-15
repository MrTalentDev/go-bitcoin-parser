package main

import (
	"encoding/hex"
	"encoding/json"
)

type Hash []byte

func (h Hash) MarshalJSON() ([]byte, error) {
	res, err:= json.Marshal(hex.EncodeToString(h))
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Hex byte

func (h Hex) MarshalJSON() ([]byte, error) {
	res, err:= json.Marshal(hex.EncodeToString([]byte{byte(h)}))
	if err != nil {
		return nil, err
	}
	return res, nil
}