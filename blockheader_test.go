package main

import (
	"bytes"
	"testing"
)

var BlockHeaderTestBytes = []byte{
	0x01,0x00,0x00,0x00,0x6f,0xe2,0x8c,0x0a,
	0xb6,0xf1,0xb3,0x72,0xc1,0xa6,0xa2,0x46,
	0xae,0x63,0xf7,0x4f,0x93,0x1e,0x83,0x65,
	0xe1,0x5a,0x08,0x9c,0x68,0xd6,0x19,0x00,
	0x00,0x00,0x00,0x00,0x98,0x20,0x51,0xfd,
	0x1e,0x4b,0xa7,0x44,0xbb,0xbe,0x68,0x0e,
	0x1f,0xee,0x14,0x67,0x7b,0xa1,0xa3,0xc3,
	0x54,0x0b,0xf7,0xb1,0xcd,0xb6,0x06,0xe8,
	0x57,0x23,0x3e,0x0e,0x61,0xbc,0x66,0x49,
	0xff,0xff,0x00,0x1d,0x01,0xe3,0x62,0x99,
}

var BlockHeaderTest =  &BlockHeader{
	Version:	1,
	PrevBlockHeaderHash: []byte{
	0x00,0x00,0x00,0x00,0x00,0x19,0xd6,0x68,
	0x9c,0x08,0x5a,0xe1,0x65,0x83,0x1e,0x93,
	0x4f,0xf7,0x63,0xae,0x46,0xa2,0xa6,0xc1,
	0x72,0xb3,0xf1,0xb6,0x0a,0x8c,0xe2,0x6f,
	},
	MerkleRootHash: []byte{
	0x0e,0x3e,0x23,0x57,0xe8,0x06,0xb6,0xcd,
	0xb1,0xf7,0x0b,0x54,0xc3,0xa3,0xa1,0x7b,
	0x67,0x14,0xee,0x1f,0x0e,0x68,0xbe,0xbb,
	0x44,0xa7,0x4b,0x1e,0xfd,0x51,0x20,0x98,
	},
	Time:				 1231469665,
	NBits:               486604799,
	Nonce:               2573394689,
}


func TestBlockHeader_Serialize(t *testing.T) {
	b := BlockHeaderTest.Serialize()
	if bytes.Compare(b,  BlockHeaderTestBytes) != 0 {
		t.Errorf("Serialize test error, we got %x, but the answer is %x", b, BlockHeaderTestBytes)
	}
}

func TestDeserializeBlockHeader(t *testing.T) {
	bh := DeserializeBlockHeader(BlockHeaderTestBytes)
	CompareBlockHeader(bh, BlockHeaderTest, t)
}


func CompareBlockHeader(got *BlockHeader, ans *BlockHeader, t *testing.T) {
		if got.Version != ans.Version {
			t.Errorf("version is not correct, we got:%d, but answer is:%d", got.Version, ans.Version)
		}
		if got.Nonce != ans.Nonce {
			t.Errorf("nonce is not correct, we got:%d, but answer is:%d", got.Nonce, ans.Nonce)
		}
		if got.NBits != ans.NBits {
			t.Errorf("nbits is not correct, we got:%d, but answer is:%d", got.NBits, ans.NBits)
		}
		if got.Time != ans.Time {
			t.Errorf("time is not correct, we got:%d, but answer is:%d", got.Time, ans.Time)
		}
		if bytes.Compare(got.MerkleRootHash, ans.MerkleRootHash) != 0 {
			t.Errorf("merkleRootHash is not correct, we got:%x, but answer is:%x",got.MerkleRootHash, ans.MerkleRootHash)
		}
		if bytes.Compare(got.PrevBlockHeaderHash, ans.PrevBlockHeaderHash) != 0 {
			t.Errorf("prevBlockHeaderHsah is not correct, we got:%x, but answer is:%x", got.PrevBlockHeaderHash, ans.PrevBlockHeaderHash)
		}
}