package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

type BlkBlock struct {
	MagicBytes	Hash  `json:"magicBytes"`	// 4byte
	Size  		uint32  `json:"size"`		// 4byte
	Block 		*Block	`json:block`
}

var ZeroMagicBytes = []byte{0x00,0x00,0x00,0x00}


func DeserializeBlkBlock(input []byte) (*BlkBlock, int, error){
	magicBytes := input[:4]
	// last bytes are all zero, we don't need continue read blocks
	if bytes.Compare(magicBytes, ZeroMagicBytes) == 0 {
		return nil, 0 ,nil
	}
	params, err := MagicBytesToParams(magicBytes)
	if err != nil {
		return nil, 0, err
	}

	size := binary.LittleEndian.Uint32(input[4:8])
	block, err := DeserializeBlock(input[8:8+size], params)
	if err != nil {
		return nil, 0, err
	}
	return &BlkBlock{
		MagicBytes: magicBytes,
		Size:       size,
		Block:      block,
	}, 8+int(size), nil
}

func (b *BlkBlock) Serialize() ([]byte, error) {
	sblock, err := b.Block.Serialize()
	if err != nil {
		return nil, err
	}
	return ConcatCopy(
		b.MagicBytes,
		IntToLittleEndianBytes(b.Size),
		sblock,
	), nil
}

func (b *BlkBlock) String() (string, error){
	bs, err := json.MarshalIndent(b,"","	")
	if err != nil {
		return "", err
	}
	return string(bs) + "\n", nil
}

func blockToBlkBlock(block []byte, params *Params) []byte {
	size := IntToLittleEndianBytes(uint32(len(block)))
	return ConcatCopy(params.MagicBytes, size, block)
}
