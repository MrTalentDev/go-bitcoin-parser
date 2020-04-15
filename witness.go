package main

import (
	"encoding/json"
)

type Witness struct {
	WitnessStackCount uint		   `json:witnessStackCOunt`		// compactSize
	WitnessScript []*WitnessScript `json:witnessScript`
}

type WitnessScript struct {
	WitnessScriptBytes uint  `json:witnessScriptBytes` // compactSize
	WitnessScript	   Hash  `json:witnessScript`
}

func DeserializeWitness(input []byte, network *Params) (*Witness, int, error) {
	var witness Witness
	var witnessScriptLen int
	witnessStackCountSize, witnessStackCount, err := DecodeVarint(input)
	if err != nil {
		return nil, 0, err
	}
	_, input = CutBytes(input, witnessStackCountSize)
	witness.WitnessStackCount = witnessStackCount
	for i := 0; i < int(witnessStackCount); i++ {
		var witnessScript WitnessScript
		witnessScriptBytesSize, witnessScriptBytes, err := DecodeVarint(input)
		if err != nil {
			return nil, 0, err
		}
		witnessScript.WitnessScriptBytes = witnessScriptBytes
		if err != nil {
			return nil, 0 ,err
		}
		_, input = CutBytes(input, witnessScriptBytesSize)
		witnessScript.WitnessScript, input = CutBytes(input, int(witnessScriptBytes))
		witness.WitnessScript = append(witness.WitnessScript, &witnessScript)
		witnessScriptLen += witnessScriptBytesSize + int(witnessScriptBytes)
	}
	witnessLen := witnessStackCountSize + witnessScriptLen
	return &witness, witnessLen, nil
}

func (w *Witness) Serialize() ([]byte, error) {
	var witnessScripts []byte
	witnessStackCount, err := EncodeVarint(w.WitnessStackCount)
	if err != nil {
		return []byte{}, err
	}
	for _, witnessScript := range w.WitnessScript {
		witnessScriptBytes, err := EncodeVarint(witnessScript.WitnessScriptBytes)
		if err != nil {
			return []byte{}, err
		}
		witnessScripts = append(witnessScripts, witnessScriptBytes...)
		witnessScripts = append(witnessScripts, witnessScript.WitnessScript...)
	}
	return ConcatCopy(
		witnessStackCount,
		witnessScripts), nil
}

func (w *Witness) String() (string, error) {
	wstr, err := json.MarshalIndent(w,"","	")
	if err != nil {
		return "", err
	}
	return string(wstr) + "\n", nil
}