package main

import (
	"errors"
	"fmt"
	"strings"
)

const CHARSET = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

var gen = []int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}


func bech32Decode(bech string) (string, []byte, error) {
	for _, char := range bech {
		if char < 33 || char > 126{
			return "", nil, fmt.Errorf("Invalid char: %c %d ", char, char)
		}
	}
	lower := strings.ToLower(bech)
	if lower != bech && strings.ToUpper(bech) != bech {
		return "", nil, fmt.Errorf("strings are not all upper or lower, we got: %s ", bech)
	}
	bech = lower
	pos := strings.LastIndexByte(bech,'1')
	if pos < 1 || pos +7 > len(bech) || len(bech) > 90 {
		return "", nil, fmt.Errorf("invalid position, we got pos: %d", pos)
	}
	data := make([]byte, 0, len(bech[pos+1:]))
	hrp := bech[:pos]
	for _, char := range bech[pos+1:]{
		in := strings.IndexByte(CHARSET, byte(char))
		if in < 0 {
			return "", nil, fmt.Errorf("char is not in CHARSET, we got char: %c", char)
		}
		data =  append(data,byte(in))
	}
	if !bech32VerifyChecksum(hrp, data) {
		return "", nil, fmt.Errorf("verify checksum failed")
	}
	return hrp, data[:len(data)-6], nil
}

func bech32Encode(hrp string, data []byte) (string, error){
	combined := append(data,bech32CreateChecksum(hrp, data)...)
	chars, err := toCharSet(combined)
	if err !=  nil {
		return "", err
	}
	return hrp + "1" + chars, nil
}


func toCharSet(data []byte) (string, error){
	res := make([]byte, 0, len(data))
	for i, b := range data {
		if int(b) >= len(CHARSET){
			return "", errors.New(fmt.Sprintf("byte is to big, we got %d at data[%d], but max is %d",int(b), i,len(CHARSET)))
		}
		res = append(res, CHARSET[int(b)])
	}
	return string(res), nil
}

func ConvertBits(data []byte, fromBits, toBits uint8, pad bool) ([]byte, error){
	acc := 0
	bits := uint8(0)
	ret := make([]byte,0)
	maxv := (1 << toBits ) - 1
	maxAcc := (1 << (fromBits + toBits -1)) -1
	for _, value := range data {
		if value < 0  || (value >> fromBits) < 0 {
			return nil, fmt.Errorf("value is not correct, we got %x", value)
		}
		acc = ((acc << fromBits) | int(value)) & maxAcc
		bits += fromBits
		for bits >= toBits {
			bits -= toBits
			ret = append(ret, byte((acc >> bits) & maxv))
		}
	}
	if pad {
		if bits > 0 {
			ret = append(ret, byte((acc << (toBits - bits)) & maxv))
		}
	} else if bits >= fromBits || (acc << (toBits - bits) & maxv) > 0 {
		return nil, fmt.Errorf("zero padding error ")
	}
	return ret, nil
}


func bech32CreateChecksum(hrp string, data []byte) []byte {
	ints := make([]int, len(data))
	for i, b := range data {
		ints[i] = int(b)
	}
	values := append(bech32HrpExpand(hrp), ints...)
	values = append(values, []int{0, 0, 0, 0, 0, 0}...)
	polymod := bech32Polymod(values) ^ 1
	var res []byte
	for i := 0; i < 6; i++ {
		res = append(res, byte((polymod>>uint(5*(5-i)))&31))
	}
	return res
}

func bech32Polymod(values []int) int {
	chk := 1
	for _, v := range values {
		b := chk >> 25
		chk  = (chk & 0x1ffffff) << 5 ^ v
		for i := 0; i < 5; i++ {
			if (b >> uint(i)) & 1 == 1 {
				chk ^= gen[i]
			}
		}
	}
	return chk
}

func bech32HrpExpand(s string) []int {
	res := make([]int, 0,len(s)*2+1)
	for _, x := range s {
		res = append(res, int(x>>5))
	}
	res = append(res,0)
	for _, x := range s {
		res = append(res, int(x&31))
	}
	return res
}

func bech32VerifyChecksum(hrp string, data []byte) bool {
	ints := make([]int, len(data))
	for i, b := range data {
		ints[i] = int(b)
	}
	tmp := append(bech32HrpExpand(hrp), ints...)
	return bech32Polymod(tmp) == 1
}

func SegwitEncode(hrp string, witver , witprog[]byte) (string, error){
	wit := make([]byte, 0)
	cb, err:= ConvertBits(witprog, 8, 5, true)
	if err != nil {
		return "", err
	}
	wit = append(witver, cb...)
	res, err := bech32Encode(hrp, wit)
	if err != nil {
		return "", err
	}
	return res, nil
}

func SegwitDecode(hrp, addr string) (int, []byte, error){
	hrpgot, data, err := bech32Decode(addr)
	if err != nil {
		return 0, nil, err
	}
	if strings.Compare(hrp, hrpgot) != 0 {
		return 0, nil, fmt.Errorf("hrp is not equal, we got: %s, but ans is: %s",hrp, hrpgot)
	}
	decoded, err := ConvertBits(data[1:],5,8,false)
	if err != nil {
		return 0, nil, err
	}
	if len(decoded) < 2 || len(decoded) > 40 {
		return 0, nil, fmt.Errorf("decoded's length is not correct, length is %d", len(decoded))
	}
	if data[0] > 16 {
		return 0, nil, fmt.Errorf("version is bigger then 16, version:%d", data[0])
	}
	if data[0] == 0 && len(decoded) != 20 && len(decoded) != 32 {
		return 0, nil, fmt.Errorf("version is 0 and decoded's length is not 20 and 32, length is %d", len(decoded))
	}
	return int(data[0]), decoded, nil
}