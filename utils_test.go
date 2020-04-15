package main

import (
	"bytes"
	"math"
	"testing"
)

func TestDecodeVarint(t *testing.T) {
	svarint, size, err := DecodeVarint([]byte{0xfc})
	if err != nil {
		t.Errorf("TestDecodeVarint error: %s", err)
	}
	if size != 252 {
		t.Errorf("size are not correct, we got: %d, but answer is: %d", size, 252)
	}
	if svarint != 1 {
		t.Errorf("size are not correct, we got: %d, but answer is: %d", svarint, 1)
	}

	svarint, size, err = DecodeVarint([]byte{0xfd,0x12,0x34})
	if err != nil {
		t.Errorf("TestDecodeVarint error: %s", err)
	}
	if size != 13330 {
		t.Errorf("size are not correct, we got: %d, but answer is: %d", size, 13330)
	}
	if svarint != 3 {
		t.Errorf("size are not correct, we got: %d, but answer is: %d", svarint, 3)
	}

	svarint, size, err = DecodeVarint([]byte{0xfe,0x12,0x34,0x56,0x78})
	if err != nil {
		t.Errorf("TestDecodeVarint error: %s", err)
	}
	if size != 2018915346 {
		t.Errorf("size are not correct, we got: %d, but answer is: %d", size, 2018915346)
	}
	if svarint != 5 {
		t.Errorf("size are not correct, we got: %d, but answer is: %d", svarint, 5)
	}

	svarint, size, err = DecodeVarint([]byte{0xff,0x12,0x34,0x56,0x78,0x12,0x34,0x56,0x78})
	if err != nil {
		t.Errorf("TestDecodeVarint error: %s", err)
	}
	if size != 8671175386481439762 {
		t.Errorf("size are not correct, we got: %d, but answer is: %d", size, 8671175386481439762)
	}
	if svarint != 9 {
		t.Errorf("size are not correct, we got: %d, but answer is: %d", svarint, 9)
	}
}

func TestAddHeadSlice(t *testing.T) {
	in := []byte{0x12,0x34}
	head := []byte{0xfd}
	ans := []byte{0xfd, 0x12, 0x34}
	got := AddHeadSlice(head, in)
	if bytes.Compare(got, ans) != 0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",got, ans)
	}
}

func TestIntToLittleEndianBytes(t *testing.T) {
	var uint8test  uint8  = math.MaxUint8
	var uint16test uint16 = math.MaxUint16
	var uint32test uint32 = math.MaxUint32
	var uint64test uint64 = math.MaxUint64
	var int32test  int32  = math.MaxInt32
	var int64test  int64  = math.MaxInt64
	uint8testBytes  := []byte{0xff}
	uint16testBytes := []byte{0xff,0xff}
	uint32testBytes := []byte{0xff,0xff,0xff,0xff,}
	uint64testBytes := []byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}
	int32testBytes := []byte{0xff,0xff,0xff,0x7f,}
	int64testBytes := []byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0x7f,}
	testbytes := IntToLittleEndianBytes(uint8test)
	if bytes.Compare(testbytes, uint8testBytes) != 0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",testbytes, uint8testBytes)
	}
	testbytes = IntToLittleEndianBytes(uint16test)
	if bytes.Compare(testbytes, uint16testBytes) != 0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",testbytes, uint16testBytes)
	}
	testbytes = IntToLittleEndianBytes(uint32test)
	if bytes.Compare(testbytes, uint32testBytes) != 0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",testbytes, uint32testBytes)
	}
	testbytes = IntToLittleEndianBytes(uint64test)
	if bytes.Compare(testbytes, uint64testBytes) != 0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",testbytes, uint64testBytes)
	}
	testbytes = IntToLittleEndianBytes(int32test)
	if bytes.Compare(testbytes, int32testBytes) != 0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",testbytes, int32testBytes)
	}
	testbytes = IntToLittleEndianBytes(int64test)
	if bytes.Compare(testbytes, int64testBytes) != 0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",testbytes, int64testBytes)
	}
}

func TestReverseBytes(t *testing.T) {
	var test = []byte{0x12, 0x34, 0x56, 0x78}
	var ans  = []byte{0x78, 0x56, 0x34, 0x12}
	got := ReverseBytes(test)
	if bytes.Compare(got, ans) != 0 {
		t.Errorf("ReverseBytes bytes are not correct, we got: %x, but answer is: %x",test, ans)
	}
}

func TestConcatCopy(t *testing.T) {
	var test = []byte{
		0x12,0x34,0x56,0x78,
	}
	var ans = []byte{
		0x12,0x34,0x56,0x78,0x12,0x34,0x56,0x78,
	}
	got := ConcatCopy(test,test)
	if bytes.Compare(got, ans) != 0 {
		t.Errorf("ConcatCopy bytes are not correct, we got: %x, but answer is: %x",test, ans)
	}
}

func TestCutBytes(t *testing.T) {
	var test = []byte{
		0x12,0x34,0x56,0x78,
	}
	var anscut = []byte{
		0x12,0x34,
	}
	var ansinput = []byte{
		0x56,0x78,
	}
	cut, input := CutBytes(test,2)
	if bytes.Compare(cut, anscut) != 0 {
		t.Errorf("cut bytes are not correct, we got: %x, but answer is: %x",cut, anscut)
	}
	if bytes.Compare(input, ansinput) != 0 {
		t.Errorf("input bytes are not correct, we got: %x, but answer is: %x",input, ansinput)
	}
}

func TestDoubleSha256(t *testing.T) {
	var test = []byte{
		0x12,0x34,0x56,0x78,0x12,0x34,0x56,0x78,
	}
	var ans = []byte{
		0xb4,0xdd,0xc3,0xad,0x04,0x23,0x75,0xcc,
		0x7c,0x3e,0xc4,0xf8,0x7b,0xb7,0x07,0xc0,
		0xf0,0x1d,0x85,0xde,0xd9,0x33,0xe0,0x19,
		0x48,0x2b,0x17,0xdd,0x24,0x92,0x00,0xe3,
	}
	got := DoubleSha256(test)
	if bytes.Compare(got, ans) != 0 {
		t.Errorf("DoubleSha256 bytes are not correct, we got: %x, but answer is: %x",got, ans)
	}
}

func TestEncodeVarint(t *testing.T) {
	var test1 uint = 252
	var test2 uint = math.MaxUint16
	var test3 uint = math.MaxUint32
	var test4 uint = math.MaxUint64
	var ans1  = []byte{0xfc}
	var ans2  = []byte{0xfd,0xff,0xff}
	var ans3  = []byte{0xfe,0xff,0xff,0xff,0xff}
	var ans4  = []byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff}

	got, err := EncodeVarint(test1)
	if err != nil {
		t.Errorf("TestEncodeVarint error: %s", err)
	}
	if bytes.Compare(got, ans1) !=0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",got, ans1)
	}

	got, err = EncodeVarint(test2)
	if err != nil {
		t.Errorf("TestEncodeVarint error: %s", err)
	}
	if bytes.Compare(got, ans2) != 0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",got, ans2)
	}

	got, err = EncodeVarint(test3)
	if err != nil {
		t.Errorf("TestEncodeVarint error: %s", err)
	}
	if bytes.Compare(got, ans3) != 0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",got, ans3)
	}

	got, err = EncodeVarint(test4)
	if err != nil {
		t.Errorf("TestEncodeVarint error: %s", err)
	}
	if bytes.Compare(got, ans4) != 0 {
		t.Errorf("bytes are not correct, we got: %x, but answer is: %x",got, ans4)
	}
}

func TestDecodeVarIntForIndex(t *testing.T) {
	test := []byte{
		0x81,0x54,0xbe,0xf9,0x99,0x17,0x88,0x9f,
		0xb6,0x3d,0x94,0xbc,0x00,0x94,0xc1,0x13,
		0x84,0xa8,0xab,0xd3,0x2e,0x84,0xa8,0xc4,
		0xa7,0x41,
	}
	n, pos:= DecodeVarIntForIndex(test)
	if n != 340  {
		t.Errorf("value are not correct, we got: %d, but answer is: %d",n, 340)
	}
	if pos != 2 {
		t.Errorf("pos are not correct, we got: %d, but answer is: %d",pos, 2)
	}
}