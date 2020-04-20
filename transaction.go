package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

var witnessEncodeBytes = []byte{0x00, 0x01}


type Transaction struct {
	Hash		 Hash		 `json:"hash"`
	Version 	 uint32  	 `json:"version"`		// 4 byte
	TxInCount    uint        `json:"txinCount"`		// compactSize
	TxIn         []*TxInput  `json:"txin"`
	TxOutCount   uint     	 `json:"txOutCount"`	// compactSize
	TxOut        []*TxOutput `json:"txOut"`
	Witness		 []*Witness  `json:"witness"`
	LockTime	 uint32		 `json:"lockTime"`
}



func DeserializeTransaction(input []byte, network *Params) (*Transaction, int, error) {
	var witness []*Witness
	sumSize := 0
	witnessEncodeBytesSize := 0
	isWitness := false
	versionBytes, input := CutBytes(input,4)
	version := binary.LittleEndian.Uint32(versionBytes)
	if bytes.Compare(input[0:2], witnessEncodeBytes) == 0 {
		_, input = CutBytes(input,2)
		isWitness = true
		witnessEncodeBytesSize += len(witnessEncodeBytes)
	}
	txInCountSize, txInCount, err := DecodeVarint(input)
	if err != nil {
		return nil, 0, err
	}
	_, input = CutBytes(input,txInCountSize)
	var txIn []*TxInput
	txInSize := 0
	for i := 0; i <int(txInCount) ; i++ {
		txinput, txinputSize, err := DeserializeTxIn(input[txInSize:], network)
		if err != nil {
			return nil, 0, err
		}
		txIn = append(txIn, txinput)
		txInSize = txInSize + txinputSize
	}
	_, input = CutBytes(input, txInSize)
	txOutCountSize, txOutCount, err := DecodeVarint(input)
	if err != nil {
		return nil, 0, err
	}
	_, input = CutBytes(input, txOutCountSize)
   	var txOut []*TxOutput
	txOutSize := 0
	for i := 0; i <int(txOutCount); i++ {
		txout, txoutputSize, err := DeserializeTxOut(input[txOutSize:], network)
		if err != nil {
			return nil, 0, err
		}
		txOut = append(txOut, txout)
		txOutSize = txOutSize + txoutputSize
	}
	_, input = CutBytes(input, txOutSize)
	witnessSize := 0
	if isWitness == true {
		for i:= 0; i < int(txInCount); i++ {
			w, witnessLen, err:= DeserializeWitness(input[witnessSize:], network)
			if err != nil {
				return nil, 0, err
			}
			witness = append(witness, w)
			witnessSize += witnessLen
		}
		_, input = CutBytes(input, witnessSize)
	}
	lockTimeBytes, input := CutBytes(input, 4)
	lockTime := binary.LittleEndian.Uint32(lockTimeBytes)
	sumSize = len(versionBytes) + witnessEncodeBytesSize + txInCountSize + txInSize + txOutCountSize + txOutSize + witnessSize + len (lockTimeBytes)
	transaction := Transaction{
		Version:    version,
		TxInCount:  txInCount,
		TxIn:       txIn,
		TxOutCount: txOutCount,
		TxOut:      txOut,
		Witness:    witness,
		LockTime:   lockTime,
	}
	transaction.Hash, err  = transaction.newHash()
	if err != nil {
		return nil, 0, err
	}
	return &transaction, sumSize, nil
}

func (tx Transaction) Serialize() ([]byte, error) {
	 txInCount, err := EncodeVarint(tx.TxInCount)
	 if err != nil {
	 	return nil, err
	 }

	 txOutCount, err := EncodeVarint(tx.TxOutCount)
	 if err != nil {
	 	return nil, err
	 }
	 var txIn []byte
	 for _ ,tx := range tx.TxIn {
	 	txBytes, err := tx.Serialize()
	 	if err != nil {
	 		return nil, err
		}
	 	txIn = append(txIn, txBytes...)
	 }
	 var txOut []byte
	 for _ ,tx := range tx.TxOut {
		txBytes, err := tx.Serialize()
		if err != nil {
			return nil, err
		}
		txOut = append(txOut, txBytes...)
	 }

	 return ConcatCopy(
		 IntToLittleEndianBytes(tx.Version),
		 txInCount,
		 txIn,
		 txOutCount,
		 txOut,
		 IntToLittleEndianBytes(tx.LockTime),
	 ), nil

}

func (tx Transaction) newHash() ([]byte, error) {
	txBytes, err := tx.Serialize()
	if err != nil {
		return nil, nil
	}
	return ReverseBytes(DoubleSha256(txBytes)), nil
}

func (tx Transaction) String() (string, error){
	txstr, err := json.MarshalIndent(tx,"","	")
	if err != nil {
		return "", err
	}
	return string(txstr) + "\n", nil
}