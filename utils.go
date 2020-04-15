package main

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"math"
	"os"
	"reflect"
)

func DecodeVarint(data []byte) (int, uint, error){
	switch size := data[0]; {
	case size < 253:
		return 1, uint(size), nil
	case size == 253:
		num := binary.LittleEndian.Uint16(data[1:3])
		return 3, uint(num), nil
	case size == 254:
		num := binary.LittleEndian.Uint32(data[1:5])
		return 5, uint(num), nil
	case size == 255:
		num := binary.LittleEndian.Uint64(data[1:9])
		return 9, uint(num), nil
	default:
		return 0, 0, errors.New("can't decode this")
	}
}

func DecodeVarIntForIndex(input []byte)  (n int, pos int) {
	for {
		data := input[pos]
		pos += 1
		n = (n << 7) | int(data & 0x7f)
		if data & 0x80 == 0 {
			return n, pos
		}
		n += 1
	}
}

func EncodeVarint(pkScriptBytes uint) ([]byte, error){
	switch {
	case pkScriptBytes >= 0 && pkScriptBytes < 253 :
		return IntToLittleEndianBytes(uint8(pkScriptBytes)), nil
	case pkScriptBytes >= 253 && pkScriptBytes <= math.MaxUint16:
		bytes := IntToLittleEndianBytes(uint16(pkScriptBytes))
		result := AddHeadSlice([]byte{0xfd},bytes)
		return result, nil
	case pkScriptBytes > math.MaxUint16 && pkScriptBytes <= math.MaxUint32:
		bytes := IntToLittleEndianBytes(uint32(pkScriptBytes))
		result := AddHeadSlice([]byte{0xfe},bytes)
		return result, nil
	case pkScriptBytes > math.MaxUint32 && pkScriptBytes <= math.MaxUint64:
		bytes := IntToLittleEndianBytes(uint64(pkScriptBytes))
		result := AddHeadSlice([]byte{0xff},bytes)
		return result, nil
	default:
		return nil, errors.New("not in range")
	}
}

func AddHeadSlice(head []byte, slice[]byte) []byte{
	result := make([]byte, len(head)+len(slice))
	copy(result[0:len(head)], head)
	copy(result[len(head):], slice)
	return result
}

func CutBytes(input []byte, index int) ([]byte, []byte) {
	result := input[0:index]
	input = input[index:]
	return result, input
}

func ConcatCopy(slices ...[]byte) []byte {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	result := make([]byte, totalLen)
	var i int
	for _, s := range slices {
		i += copy(result[i:], s)
	}
	return result
}

func ReverseBytes(input []byte) []byte{
	temp := make([]byte, len(input))
	copy(temp, input)
	for i, j := 0, len(temp)-1 ; i < j ; i, j = i+1, j-1{
		temp[i], temp[j] = temp[j], temp[i]
	}
	return temp
}

func DoubleSha256(input []byte) []byte {
	sha := sha256.Sum256(input)
	doubleSha:= sha256.Sum256(sha[:])
	return doubleSha[:]
}

func IntToLittleEndianBytes(input interface{}) []byte{
	switch t, v := reflect.TypeOf(input), reflect.ValueOf(input); t.Kind() {
	case reflect.Uint8:
		return  []byte{v.Interface().(uint8)}
	case reflect.Uint16:
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, v.Interface().(uint16))
		return buf
	case reflect.Uint32:
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, v.Interface().(uint32))
		return buf
	case reflect.Uint64:
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, v.Interface().(uint64))
		return buf
	case reflect.Int32:
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, uint32(v.Interface().(int32)))
		return buf
	case reflect.Int64:
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, uint64(v.Interface().(int64)))
		return buf
	default:
		panic(errors.New("input type is not supported"))
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}