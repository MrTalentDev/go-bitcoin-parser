package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var TestBlk = &Blk{
	name: 2,
}

var TestPath = map[string]int{
	"/home/ubuntu/test-data/blocks/blk00002.dat":2,
	"/home/ubuntu/test-data/blocks/blk00001":0,
	"/home/ubuntu/test-data/blocks/blk00015.dat":15,
	"/home/ubuntu/test-data/blocks/blk01015.dat":1015,
}

var TestBlkHash = [][]byte{
	{
		0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x32,
		0x2c,0xee,0x8a,0x77,0x4b,0x3b,0xc5,0x0b,
		0xd3,0x04,0xf9,0x2e,0xe6,0x3a,0x7d,0x8c,
		0x7c,0x01,0xbc,0x9e,0x71,0xf1,0x82,0xae,
	},
	{
		0x00,0x00,0x00,0x00,0x00,0x00,0x0c,0xaf,
		0xdf,0x2f,0x34,0x30,0x37,0x7a,0xf5,0x89,
		0x55,0xcd,0x4c,0x0d,0xba,0x2e,0xce,0xcd,
		0xb2,0xa9,0x6e,0x2d,0x19,0xee,0x51,0x91,
	},
}

func NewTestBlk(t *testing.T) *Blk{
	path, err := filepath.Abs("./test-data/blocks/blk00002.dat")
	if err != nil {
		t.Errorf("create path error: %s", err)
	}
	blk, err := NewBlk(path)
	if err != nil {
		t.Fatalf("TestNewBlk error:%s", err)
	}
	if blk.name != TestBlk.name {
		t.Fatalf("name is not correct, we got: %d, but answer is: %d",blk.name, TestBlk.name)
	}
	return blk
}

func TestBlkPathToName(t *testing.T) {
	for path, ans := range TestPath {
		name, err := BlkPathToName(path)
		if err != nil {
			t.Errorf("TestBlkPathToName error:%s", err)
		}
		if name != ans {
			t.Errorf("name is not correct, we got: %d, but answer is: %d",name, ans)
		}
	}
}

func TestDeserializeBlkBlocks(t *testing.T) {
	blk := NewTestBlk(t)
	file, err := ioutil.ReadFile(blk.path)
	if err != nil {
		t.Errorf("ioutil.ReadFile error: %s", err)
	}
	blocks, err := DeserializeBlkBlocks(file)
	if err != nil {
		t.Errorf("TestDeserializeBlkBlocks error: %s", err)
	}
	for i, block := range blocks {
		if bytes.Compare(block.Block.Hash ,TestBlkHash[i]) != 0 {
			t.Errorf("name is not correct, we got: %x, but answer is: %x", block.Block.Hash, TestBlkHash[i])
		}
	}
}

func TestBlk_ReadBlkFile(t *testing.T) {
	blk := NewTestBlk(t)
	blocks, err := blk.ReadBlkFile()
	if err != nil {
		t.Errorf("TestBlk_ReadBlkFile error: %s", err)
	}
	for i, block := range blocks {
		if bytes.Compare(block.Hash,TestBlkHash[i]) != 0 {
			t.Errorf("name is not correct, we got: %x, but answer is: %x", block.Hash, TestBlkHash[i])
		}
	}
}