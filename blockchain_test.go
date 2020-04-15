package main

import (
	"bytes"
	"path/filepath"
	"testing"
)



var testUnorderHashes = [][]byte{
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

var testOrderHashes = [][]byte{
	{
		0x00,0x00,0x00,0x00,0x00,0x00,0x0c,0xaf,
		0xdf,0x2f,0x34,0x30,0x37,0x7a,0xf5,0x89,
		0x55,0xcd,0x4c,0x0d,0xba,0x2e,0xce,0xcd,
		0xb2,0xa9,0x6e,0x2d,0x19,0xee,0x51,0x91,
	},
	{
		0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x32,
		0x2c,0xee,0x8a,0x77,0x4b,0x3b,0xc5,0x0b,
		0xd3,0x04,0xf9,0x2e,0xe6,0x3a,0x7d,0x8c,
		0x7c,0x01,0xbc,0x9e,0x71,0xf1,0x82,0xae,
	},
}

var testHeight = map[int][]byte{
	131301: {
		0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x32,
		0x2c,0xee,0x8a,0x77,0x4b,0x3b,0xc5,0x0b,
		0xd3,0x04,0xf9,0x2e,0xe6,0x3a,0x7d,0x8c,
		0x7c,0x01,0xbc,0x9e,0x71,0xf1,0x82,0xae,
	},
	131079: {
		0x00,0x00,0x00,0x00,0x00,0x00,0x0c,0xaf,
		0xdf,0x2f,0x34,0x30,0x37,0x7a,0xf5,0x89,
		0x55,0xcd,0x4c,0x0d,0xba,0x2e,0xce,0xcd,
		0xb2,0xa9,0x6e,0x2d,0x19,0xee,0x51,0x91,
	},
}

var testLastHeight = 136974

func NewTestBlockchain(t *testing.T) *Blockchain {
	bpath := "./test-data/blocks"
	ipath := "./test-data/index"
	blockPath, err := filepath.Abs(bpath)
	if err != nil {
		t.Errorf("create filepath %s error: %s", bpath, err)
	}
	indexPath, err := filepath.Abs(ipath)
	if err != nil {
		t.Errorf("create filepath %s error: %s", ipath, err)
	}
	blockchain, err := NewBlockchain(blockPath, indexPath)
	if err != nil {
		t.Fatalf("TestNewBlockchain error: %s", err)
	}
	return blockchain
}

func TestBlockchain_GetUnorderBlocks(t *testing.T) {
	blockchain := NewTestBlockchain(t)
	defer blockchain.Close()
	blocks, err := blockchain.GetUnorderBlocks(0,2,2)
	if err != nil {
		t.Errorf("TestBlockchain_GetUnorderBlocks error: %s", err)
	}
	for i, block := range blocks {
		CompareBlockHash(block.Hash, testUnorderHashes[i], t)
	}
}

func TestBlockchain_GetOrderBlocks(t *testing.T) {
	blockchain := NewTestBlockchain(t)
	defer blockchain.Close()
	blocks, err := blockchain.GetOrderBlocks(0,2,2)
	if err != nil {
		t.Errorf("TestBlockchain_GetOrderBlocks error: %s", err)
	}
	for i, block := range blocks {
		CompareBlockHash(block.Hash, testOrderHashes[i], t)
	}
}

func TestBlockchain_GetBlockByHash(t *testing.T) {
	blockchain := NewTestBlockchain(t)
	defer blockchain.Close()
	hash := testOrderHashes[0]
	block, err := blockchain.GetBlockByHash(hash)
	if err != nil {
		t.Errorf("TestBlockchain_GetBlockByHash error: %s", err)
	}
	CompareBlockHash(block.Hash, hash, t)
}

func TestBlockchain_GetBlockByHeight(t *testing.T) {
	blockchain := NewTestBlockchain(t)
	defer blockchain.Close()
	for height, hash := range testHeight {
		block, err := blockchain.GetBlockByHeight(height)
		if err != nil {
			t.Errorf("TestBlockchain_GetBlockByHeight error: %s", err)
		}
		CompareBlockHash(block.Hash, hash, t)
	}
}

func TestBlockchain_GetLastHeight(t *testing.T) {
	blockchain := NewTestBlockchain(t)
	defer blockchain.Close()
	last, err := blockchain.GetLastHeight()
	if err != nil {
		t.Errorf("TestBlockchain_GetLastHeight error: %s", err)
	}
	if last != testLastHeight {
		t.Errorf("last height is not correct, we got: %d, but answer is: %d",last, testLastHeight)
	}
}

func CompareBlockHash(got []byte, ans []byte, t *testing.T) {
	if bytes.Compare(got, ans) != 0 {
		t.Errorf("block hash is not correct, we got: %x, but answer is: %x",got, ans)
	}
}



