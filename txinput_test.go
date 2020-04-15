package main

import (
	"bytes"
	"testing"
)

var TxInTestBytes = []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
	0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
	0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
	0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
	0xff,0xff,0xff,0xff,0x07,0x04,0xff,0xff,
	0x00,0x1d,0x01,0x04,0xff,0xff,0xff,0xff,
}

var TxInTest = 	&TxInput{
	Hash:            []byte{
		0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
		0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
		0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
		0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
	},
	Index:           4294967295,
	ScriptBytes:     7,
	SignatureScript: []byte{0x04,0xff,0xff,0x00,0x1d,0x01,0x04},
	Sequence:        4294967295,
}

func TestTxInput_Serialize(t *testing.T) {
	cbytes, err := TxInTest.Serialize()
	if err != nil {
		t.Errorf("txin serialize err: %s", err)
	}
	if bytes.Compare(cbytes, TxInTestBytes) != 0 {
		t.Errorf("Serialize test error, we got %x, but the answer is %x", cbytes, TxInTestBytes)
	}
}

func TestDeserializeTxIn(t *testing.T) {
	txin, size, err := DeserializeTxIn(TxInTestBytes, &MainnetParams)
	if err != nil {
		t.Errorf("DeserializeTxIn test error: %s",err)
	}
	if size != len(TxInTestBytes) {
		t.Errorf("txIn return size is not correct, we got:%d, but answer is:%d", size, len(TxInTestBytes))
	}
	CompareTxIn(txin, TxInTest, t)
}

func CompareTxIn(got *TxInput, ans *TxInput, t *testing.T){
	if bytes.Compare(got.Hash, ans.Hash) != 0 {
		t.Errorf("hash is not correct, we got:%x, but answer is:%x", got.Hash, ans.Hash)
	}
	if got.Index != ans.Index {
		t.Errorf("index is not correct, we got:%d, but answer is:%d", got.Index, ans.Index)
	}
	if bytes.Compare(got.SignatureScript, ans.SignatureScript) != 0 {
		t.Errorf("signatureScript is not correct, we got:%x, but answer is:%x", got.SignatureScript, ans.SignatureScript)
	}
	if got.ScriptBytes != ans.ScriptBytes {
		t.Errorf("scriptBytes is not correct, we got:%d, but answer is:%d", got.ScriptBytes, ans.ScriptBytes)
	}
	if got.Sequence != ans.Sequence {
		t.Errorf("sequence is not correct, we got:%d, but answer is:%d", got.Sequence, ans.Sequence)
	}
}
