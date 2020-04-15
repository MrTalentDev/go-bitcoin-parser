package main

import (
	"encoding/binary"
	"encoding/json"
)

type TxInput struct {
	Hash  			Hash 		`json:"hash"`				// 32 byte
	Index  			uint32 		`json:"index"`				// 4 byte
	ScriptBytes 	uint   		`json:"scriptBytes"`		// compactSize
	SignatureScript Hash		`json:"signatureScript"`	// signatureScript
	Sequence        uint32 		`json:"sequence"`			// 4 byte
}

func DeserializeTxIn(input []byte, network *Params) (*TxInput, int, error){
	scriptBytesSize, scriptBytes,err := DecodeVarint(input[36:])

	if err != nil {
		return nil, 0, err
	}
	return &TxInput{
		Hash: 			 	input[0:32],
		Index: 			 	binary.LittleEndian.Uint32(input[32:36]),
		ScriptBytes:     	scriptBytes,
		SignatureScript: 	input[36+scriptBytesSize:36+scriptBytesSize+int(scriptBytes)],
		Sequence:			binary.LittleEndian.Uint32(input[36+scriptBytesSize+int(scriptBytes):36+scriptBytesSize+int(scriptBytes)+4]),
	}, 36+scriptBytesSize+int(scriptBytes)+4, nil
}

func (txinput TxInput) Serialize() ([]byte, error) {
	scriptBytes, err := EncodeVarint(txinput.ScriptBytes)
	if err != nil {
		return nil , err
	}
	return ConcatCopy(
		txinput.Hash,
		IntToLittleEndianBytes(txinput.Index),
		scriptBytes,
		txinput.SignatureScript,
		IntToLittleEndianBytes(txinput.Sequence),
		), nil
}

func (txIn TxInput) String() (string, error){
	txinstr, err := json.MarshalIndent(txIn,"","	")
	if err != nil {
		return "", err
	}
	return string(txinstr) + "\n", nil
}