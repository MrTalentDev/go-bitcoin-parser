package main

import "testing"

func TestDeserializeBlockIndex(t *testing.T) {
	test := []byte{
		0x89,0xfe,0x04,0x86,0xa8,0x1a,0x1d,0x07,
		0x01,0x08,0xc3,0x16,0x01,0x00,0x00,0x00,
		0x33,0xdf,0x29,0x01,0xba,0xdf,0x41,0xa4,
		0x01,0x65,0x9e,0x99,0xe3,0x7d,0x24,0xa8,
		0xab,0x0c,0x4d,0x8a,0xcf,0x06,0xdc,0x54,
		0x31,0x45,0x00,0x00,0x00,0x00,0x00,0x00,
		0x54,0xca,0x4b,0xc3,0xde,0x04,0x07,0x7e,
		0xb1,0x48,0xc2,0xc5,0x64,0x3b,0x60,0x35,
		0x83,0x6b,0xfa,0xea,0x87,0xa9,0x12,0x39,
		0x58,0xa7,0xcf,0x10,0x03,0x25,0x8d,0xb7,
		0x91,0x65,0xb4,0x4d,0xac,0xb5,0x00,0x1b,
		0x05,0xb3,0xe3,0x18,
	}

	ans := &BlockIndex{
		Height:      119962,
		Status:      29,
		Txs:         7,
		File:        1,
		DataPos:     8,
		UndoPos:     8726,
		BlockHeader: &BlockHeader{
			Version:             1,
			PrevBlockHeaderHash: []byte{
				0x00,0x00,0x00,0x00,0x00,0x00,0x45,0x31,
				0x54,0xdc,0x06,0xcf,0x8a,0x4d,0x0c,0xab,
				0xa8,0x24,0x7d,0xe3,0x99,0x9e,0x65,0x01,
				0xa4,0x41,0xdf,0xba,0x01,0x29,0xdf,0x33,
			},
			MerkleRootHash:      []byte{
				0xb7,0x8d,0x25,0x03,0x10,0xcf,0xa7,0x58,
				0x39,0x12,0xa9,0x87,0xea,0xfa,0x6b,0x83,
				0x35,0x60,0x3b,0x64,0xc5,0xc2,0x48,0xb1,
				0x7e,0x07,0x04,0xde,0xc3,0x4b,0xca,0x54,
			},
			Time:                1303668113,
			NBits:               453031340,
			Nonce:               417575685,
		},
	}

	got, err := DeserializeBlockIndex(test)
	if err != nil {
		t.Errorf("TestDeserializeBlockIndex error: %x", err)
	}
	CompareBlockIndex(got, ans, t)
}

func CompareBlockIndex(got *BlockIndex, ans *BlockIndex, t *testing.T){
	if got.Height != ans.Height {
		t.Errorf("height is not correct: we got: %d, but answer is: %d", got.Height, ans.Height)
	}
	if got.Status != ans.Status {
		t.Errorf("status is not correct: we got: %d, but answer is: %d", got.Status, ans.Status)
	}
	if got.Txs != ans.Txs {
		t.Errorf("txs is not correct: we got: %d, but answer is: %d", got.Txs, ans.Txs)
	}
	if got.File != ans.File {
		t.Errorf("file is not correct: we got: %d, but answer is: %d", got.File, ans.File)
	}
	if got.DataPos != ans.DataPos {
		t.Errorf("dataPos is not correct: we got: %d, but answer is: %d", got.DataPos, ans.DataPos)
	}
	if got.UndoPos != ans.UndoPos {
		t.Errorf("undoPos is not correct: we got: %d, but answer is: %d", got.UndoPos, ans.UndoPos)
	}
	CompareBlockHeader(got.BlockHeader, ans.BlockHeader, t)
}

func TestDeserializeFileIndex(t *testing.T) {
	test := []byte{
		0xab,0x44,0xbe,0xfc,0xea,0x3a,0x86,0xe0,
		0x89,0x45,0x86,0xff,0x07,0x87,0xad,0x0e,
		0x83,0xee,0xe2,0xe0,0x24,0x83,0xf0,0x92,
		0xa2,0x72,
	}

	ans := &FileIndex{
		NBlocks:      5700,
		NSize:        134182330,
		NUndoSize:    16270661,
		NHeightFirst: 131079,
		NHeightLast:  136974,
		NTimeFirst:   1308160164,
		NTimeLast:    1311035890,
	}

	got := DeserializeFileIndex(test)
	CompareFileIndex(got, ans, t)
}

func CompareFileIndex(got *FileIndex, ans *FileIndex, t *testing.T){
	if got.NBlocks != ans.NBlocks {
		t.Errorf("nBlocks is not correct: we got: %d, but answer is: %d", got.NBlocks, ans.NBlocks)
	}
	if got.NSize != ans.NSize {
		t.Errorf("nSize is not correct: we got: %d, but answer is: %d", got.NSize, ans.NSize)
	}
	if got.NUndoSize != ans.NUndoSize {
		t.Errorf("nUndoSize is not correct: we got: %d, but answer is: %d", got.NUndoSize, ans.NUndoSize)
	}
	if got.NHeightFirst != ans.NHeightFirst {
		t.Errorf("nHeightFirst is not correct: we got: %d, but answer is: %d", got.NHeightFirst, ans.NHeightFirst)
	}
	if got.NHeightLast != ans.NHeightLast {
		t.Errorf("nHeightLast is not correct: we got: %d, but answer is: %d", got.NHeightLast, ans.NHeightLast)
	}
	if got.NTimeFirst != ans.NTimeFirst {
		t.Errorf("nTimeFirst is not correct: we got: %d, but answer is: %d", got.NTimeFirst, ans.NTimeFirst)
	}
	if got.NTimeLast != ans.NTimeLast {
		t.Errorf("nTimeLast is not correct: we got: %d, but answer is: %d", got.NTimeLast, ans.NTimeLast)
	}
}