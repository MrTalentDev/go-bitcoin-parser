package main

import (
	"encoding/json"
)

type Block struct {
	Hash		Hash			`json:"hash"`
	height		int
	BlockHeader *BlockHeader    `json:"blockHeader"`	// 80 byte
	TxnCount	uint            `json:"txnCount"`		// compactSize uint
	Txns  		[]*Transaction	`json:"txns"`
}



func DeserializeBlock(input []byte, network *Params) (*Block, error) {
	blockHeaderBytes, input := CutBytes(input,80)
	blockHeader := DeserializeBlockHeader(blockHeaderBytes)
	txnCountSize, txnCount, err := DecodeVarint(input)
	if err != nil {
		return nil, err
	}
	_ , input= CutBytes(input, txnCountSize)
	var txns []*Transaction
	for i, size := 0, 0; i <int(txnCount) ; i++ {
		txn, txnSize, err := DeserializeTransaction(input[size:], network)
		if err != nil {
			return nil, err
		}
		txns = append(txns, txn)
		size = size + txnSize
	}
	block := Block{
		BlockHeader: blockHeader,
		TxnCount: txnCount,
		Txns: txns,
	}
	block.Hash = block.newHash()
	return &block, nil
}

func (b Block) Serialize() ([]byte, error){
	blockHeader := b.BlockHeader.Serialize()
	txnCount, err := EncodeVarint(b.TxnCount)
	if err != nil {
		return nil, err
	}
	var txns []byte
	for _ ,txn := range b.Txns {
		txnBytes, err := txn.Serialize()
		if err != nil {
			return nil, err
		}
		txns = append(txns, txnBytes...)
	}
	return ConcatCopy(
		blockHeader,
		txnCount,
		txns,
	), nil

}

// https://en.bitcoin.it/wiki/Block_hashing_algorithm
func (b Block) newHash() []byte {
	bhBytes := b.BlockHeader.Serialize()
	return ReverseBytes(DoubleSha256(bhBytes))
}

func (b Block) String() (string, error){
	bs, err := json.MarshalIndent(b,"","	")
	if err != nil {
		return "", err
	}
	return string(bs) + "\n", nil
}