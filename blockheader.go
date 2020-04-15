package main

import (
	"encoding/binary"
	"encoding/json"
)

type BlockHeader struct {
	Version int32 				`json:"version"`				// 4 byte little endian
	PrevBlockHeaderHash  Hash	`json:"prevBlockHeaderHash"`  	// 32 byte little endian
	MerkleRootHash Hash 		`json:"merkleRootHash"`			// 32 byte little endian
	Time uint32 				`json:"time"`					// 4 byte little endian
	NBits uint32 				`json:"nbits"`					// 4 byte little endian
	Nonce uint32 				`json:"nonce"`					// 4 byte little endian
}


func DeserializeBlockHeader(input []byte) *BlockHeader {
	return &BlockHeader{
		Version:    int32(binary.LittleEndian.Uint32(input[0:4])),
		PrevBlockHeaderHash:  ReverseBytes(input[4:36]),
		MerkleRootHash: ReverseBytes(input[36:68]),
		Time:  binary.LittleEndian.Uint32(input[68:72]),
		NBits:       binary.LittleEndian.Uint32(input[72:76]),
		Nonce:      binary.LittleEndian.Uint32(input[76:80]),
	}
}

func (bh *BlockHeader) Serialize() []byte{

	return ConcatCopy(
		IntToLittleEndianBytes(bh.Version),
		ReverseBytes(bh.PrevBlockHeaderHash),
		ReverseBytes(bh.MerkleRootHash),
		IntToLittleEndianBytes(bh.Time),
		IntToLittleEndianBytes(bh.NBits),
		IntToLittleEndianBytes(bh.Nonce),
		)
}

func (bh *BlockHeader) String() (string, error) {
	bhs, err := json.MarshalIndent(bh,"","	")
	if err != nil {
		return "", err
	}
	return string(bhs) + "\n", nil

}
