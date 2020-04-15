package main

import (
	"bytes"
	"errors"
	"fmt"
)

type Params struct {
	PubkeyPrefix byte
	ScriptPrefix byte
	SegwitPrefix string
	WitnessVersion byte
	MagicBytes []byte
}

var MainnetParams = Params{
	PubkeyPrefix: byte(0x00),
	ScriptPrefix: byte(0x05),
	SegwitPrefix: "bc",
	WitnessVersion: byte(0x00),
	MagicBytes: []byte{0xf9, 0xbe, 0xb4, 0xd9},
}

var TestnetParams = Params{
	PubkeyPrefix:   byte(0x6f),
	ScriptPrefix:   byte(0xc4),
	SegwitPrefix:   "tb",
	WitnessVersion: byte(0x00),
	MagicBytes: []byte{0x0b,0x11,0x09,0x07},
}


func MagicBytesToParams(magic []byte ) (*Params, error) {
	switch  {
	case bytes.Compare(magic, MainnetParams.MagicBytes) == 0:
		return &MainnetParams, nil
	case bytes.Compare(magic, TestnetParams.MagicBytes) == 0:
		return &TestnetParams, nil
	default:
		return nil, errors.New(fmt.Sprintf("magic bytes not match any params, bytes are: %x", magic))
	}
}