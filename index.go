package main

import (
	"encoding/json"
	"errors"
)

const (
	BLOCK_HAVE_DATA = 8
	BLOCK_HAVE_UNDO = 16

)

type BlockIndex struct {
	Height 		int				`json:"height"`
	Status 		int				`json:"status"`
	Txs	   		int				`json:"txs"`
	File   		int				`json:"file"`
	DataPos 	int				`json:"datapos"`
	UndoPos		int				`json:"undopos"`
	BlockHeader *BlockHeader	`json:"blockHeader"`
}

func DeserializeBlockIndex(data []byte) (*BlockIndex, error){
	 var pos 	 int
  	 var file 	 int
	 var dataPos int
	 var undoPos int
	 _, i := DecodeVarIntForIndex(data[pos:])
  	 pos += i
  	 height, i :=  DecodeVarIntForIndex(data[pos:])
  	 pos += i
  	 status, i :=  DecodeVarIntForIndex(data[pos:])
  	 pos += i
  	 txs, i :=  DecodeVarIntForIndex(data[pos:])
  	 pos += i
  	 if status & (BLOCK_HAVE_DATA |BLOCK_HAVE_UNDO) != 0 {
		file, i = DecodeVarIntForIndex(data[pos:])
		pos += i
	 } else {
	 	file = -1
	 }
	 if status & BLOCK_HAVE_DATA != 0 {
	 	dataPos, i = DecodeVarIntForIndex(data[pos:])
	 	pos += i
	 } else {
	 	dataPos = -1
	 }
	 if status & BLOCK_HAVE_UNDO != 0 {
	 	undoPos, i = DecodeVarIntForIndex(data[pos:])
	 	pos += i
	 }
	 if len(data[pos:]) != 80 {
	 	return nil, errors.New("blockheader's length is less than 80")
	 }
	 bh := DeserializeBlockHeader(data[pos:])
	return &BlockIndex{
		Height:      height,
		Status:      status,
		Txs:         txs,
		File:        file,
		DataPos:     dataPos,
		UndoPos:     undoPos,
		BlockHeader: bh,
	}, nil
}

func (i *BlockIndex) String() (string, error) {
	istr, err := json.MarshalIndent(i,"","	")
	if err != nil {
		return "", err
	}
	return string(istr) + "\n", nil
}

type FileIndex struct {
	NBlocks 		int		`json:"nBlocks"`
	NSize			int		`json:"nSize"`
	NUndoSize		int		`json:"nUndoSize"`
	NHeightFirst	int		`json:"nHeightFirst"`
	NHeightLast		int		`json:"nHeightLast"`
	NTimeFirst		int		`json:"nTimeFirst"`
	NTimeLast		int		`json:"nTimeLast"`
}

func DeserializeFileIndex(data []byte) *FileIndex{
	var pos 	 	 int

	nBlocks, i := DecodeVarIntForIndex(data[pos:])
	pos += i
	nSize, i :=  DecodeVarIntForIndex(data[pos:])
	pos += i
	nUndoSize, i :=  DecodeVarIntForIndex(data[pos:])
	pos += i
	nHeightFirst, i :=  DecodeVarIntForIndex(data[pos:])
	pos += i
	nHeightLast, i := DecodeVarIntForIndex(data[pos:])
	pos += i
	nTimeFirst, i :=  DecodeVarIntForIndex(data[pos:])
	pos += i
	nTimeLast, i :=  DecodeVarIntForIndex(data[pos:])
	pos += i

	return &FileIndex{
		NBlocks: nBlocks,
		NSize: nSize,
		NUndoSize: nUndoSize,
		NHeightFirst: nHeightFirst,
		NHeightLast: nHeightLast,
		NTimeFirst: nTimeFirst,
		NTimeLast: nTimeLast,
	}
}

func (i *FileIndex) String() (string, error) {
	istr, err := json.MarshalIndent(i,"","	")
	if err != nil {
		return "", err
	}
	return string(istr) + "\n", nil
}