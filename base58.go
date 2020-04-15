package main

import (
	"bytes"
	"math/big"
)

var base58Char = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
var base58Len = big.NewInt(int64(len(base58Char)))

// https://zh.wikipedia.org/wiki/Base58
func Base58Encode(input []byte) []byte{
	bigInt := new(big.Int)
	bigInt.SetBytes(input)
	var result []byte
	mod := &big.Int{}
	for bigInt.Sign() > 0{
		bigInt.DivMod(bigInt, base58Len, mod)
		result = append(result, base58Char[mod.Int64()])
	}
	if input[0] == 0x00 {
		result = append(result,base58Char[0])
	}
	result = ReverseBytes(result)
	return result
}

func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	for _, b:= range input {
		num := bytes.IndexByte(base58Char,b)
		result.Mul(result,base58Len)
		result.Add(result,big.NewInt(int64(num)))
	}
	ans := result.Bytes()
	if input[0] == base58Char[0] {
		ans = append([]byte{0x00}, ans...)
	}
	return ans
}