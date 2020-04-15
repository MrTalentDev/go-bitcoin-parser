package main

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

func publicKeyToAddr(publicKey []byte, params *Params) (string, error){
	shaPub := sha256.Sum256(publicKey)
	ripEncoder := ripemd160.New()
	_, err := ripEncoder.Write(shaPub[:])
	if err != nil {
		return "", err
	}
	hash := ripEncoder.Sum(nil)
	// append Network ID Byte
	return publicKeyHashToAddr(hash, params), nil
}

func publicKeyHashToAddr(hash []byte, params *Params) string {
	hash = append([]byte{params.PubkeyPrefix}, hash...)
	//generate checksum
	shaHash := sha256.Sum256(hash)
	checkSum := sha256.Sum256(shaHash[:])
	address := Base58Encode(append(hash, checkSum[:checksumLength]...))
	return string(address)
}

func scriptHashToAddr(hash []byte, params *Params) string {
	hash = append([]byte{params.ScriptPrefix}, hash...)
	shaHash := sha256.Sum256(hash)
	checkSum := sha256.Sum256(shaHash[:])
	address := Base58Encode(append(hash, checkSum[:checksumLength]...))
	return string(address)
}

func publicKeyHashToSegwitAddr(witnessProgram []byte, params *Params) (string, error) {
	addr, err:= SegwitEncode(params.SegwitPrefix, []byte{params.WitnessVersion}, witnessProgram)
	if err != nil {
		return "", err
	}
	return addr, nil
}