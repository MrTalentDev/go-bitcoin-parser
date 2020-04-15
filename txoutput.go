package main

import (
	"encoding/binary"
	"encoding/json"
)

type TxOutput struct {
	Value			int64  	  `json:"value"`	 		// 8bytes
	PkScriptBytes   uint   	  `json:"pkScriptBytes"`	// compactSize
	PkScript		*PKScript `json:"pkScript"`
}

func DeserializeTxOut(input []byte, network *Params) (*TxOutput, int, error){
	pkScriptBytesSize, pkScriptBytes, err := DecodeVarint(input[8:])
	if err != nil {
		return nil, 0, err
	}
	pkScript ,err := DeserializePKScript(input[8+pkScriptBytesSize:8+pkScriptBytesSize+int(pkScriptBytes)], network)
	if err != nil {
		return nil, 0, err
	}
	return &TxOutput{
		Value: int64(binary.LittleEndian.Uint64(input[0:8])),
		PkScriptBytes: pkScriptBytes,
		PkScript: pkScript,
	}, 8+pkScriptBytesSize+int(pkScriptBytes), nil
}

func (txOut TxOutput) Serialize() ([]byte, error) {
	pkScriptBytes, err := EncodeVarint(txOut.PkScriptBytes)
	if err != nil {
		return nil, err
	}
	return ConcatCopy(
		IntToLittleEndianBytes(txOut.Value),
		pkScriptBytes,
		txOut.PkScript.Serialize(),
		), nil
}

func (txout TxOutput) String() (string, error){
	txoutstr, err := json.MarshalIndent(txout,"","	")
	if err != nil {
		return "", err
	}
	return string(txoutstr) + "\n", nil
}