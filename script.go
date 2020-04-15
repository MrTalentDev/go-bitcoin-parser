package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
)

const (
	checksumLength = 4
)

var secp256k1P = []byte{
	0xff,0xff,0xff,0xff,
	0xff,0xff,0xff,0xff,
	0xff,0xff,0xff,0xff,
	0xff,0xff,0xff,0xff,
	0xff,0xff,0xff,0xff,
	0xff,0xff,0xff,0xff,
	0xff,0xff,0xff,0xfe,
	0xff,0xff,0xfc,0x2f,
}


type ScriptType int

func (st *ScriptType) MarshalJSON() ([]byte, error) {
	res, err := json.Marshal(scriptTypeToName[*st])
	if err != nil {
		return nil, err
	}
	return res, nil
}

type PKScript struct {
	PkScript Hash				`json:"pkscript"`
	Pops	 []*parsedOpcode	`json:"pops"`
	Stype ScriptType			`json:"stype"`
	Addresses []string			`json:"addresses"`
}

func DeserializePKScript(input []byte, network *Params) (*PKScript, error){
	pk := &PKScript{
		PkScript: input,
		Pops:	nil,
		Stype:    0,
		Addresses:  []string{},
	}
	err := pk.Analysis(network)
	if err != nil {
		return nil, err
	}
	return pk, nil
}

func (pk *PKScript) Serialize() []byte{
	return pk.PkScript
}

func (pk *PKScript) String() (string, error){
	pks, err := json.MarshalIndent(pk,"","	")
	if err != nil {
		return "", err
	}
	return string(pks) + "\n", nil}

func (pk *PKScript) Analysis(network *Params) error {
	var stype ScriptType
	var addresses []string
	pops, err := ParseScript(pk.PkScript)
	if err != nil {
		return err
	}
	if IsPubKey(pops) {
		pubkeys, err := ParsePubKey(pops, PubKey)
		if err != nil {
			return  err
		}
		addr , err:= publicKeyToAddr(pubkeys[0], network)
		if err != nil {
			return err
		}
		addresses = append(addresses,addr)
		stype = PubKey
	} else if isPubkeyHash(pops) {
		hashs, err := ParsePubKey(pops, PubKeyHash)
		if err != nil {
			return err
		}
		addr := publicKeyHashToAddr(hashs[0], network)
		addresses = append(addresses, addr)
		stype = PubKeyHash
	} else if isMultiSig(pops) {
		pubkeys, err :=	ParsePubKey(pops, Multisig)
		if err != nil {
			return err
		}
		for _, pubkey := range pubkeys {
			addr, err := publicKeyToAddr(pubkey, network)
			if err != nil {
				return err
			}
			addresses = append(addresses, addr)
		}
		stype = Multisig
	} else if isScriptHash(pops) {
		hashs, err := ParsePubKey(pops, ScriptHash)
		if err != nil {
			return err
		}
		addr := scriptHashToAddr(hashs[0], network)
		addresses = append(addresses, addr)
		stype = ScriptHash
	} else if isWitnessPubKeyHash(pops) {
		hashs, err := ParsePubKey(pops, WitnessPubKeyHash)
		if err != nil {
			return err
		}

		addr, err := publicKeyHashToSegwitAddr(hashs[0], network)
		if err != nil{
			return  err
		}
		addresses = append(addresses, addr)
		stype = WitnessPubKeyHash
	} else if isWitnessScriptHash(pops) {
		hashs, err := ParsePubKey(pops, WitnessScriptHash)
		if err != nil {
			return err
		}
		addr, err := publicKeyHashToSegwitAddr(hashs[0], network)
		if err != nil{
			return  err
		}
		addresses = append(addresses, addr)
		stype = WitnessScriptHash
	}

	pk.Pops = pops
	pk.Stype = stype
	pk.Addresses = addresses
	return nil
}

const (
	PubKey ScriptType = iota
	PubKeyHash
	Multisig
	ScriptHash
	WitnessPubKeyHash
	WitnessScriptHash
	NullData
	NonStandard
)

var scriptTypeToName = []string{
	PubKey : 				"PubKey",
	PubKeyHash:				"PubKeyHash",
	Multisig:				"Multisig",
	ScriptHash:				"ScriptHash",
	WitnessPubKeyHash: 		"WitnessPubKeyHash",
	WitnessScriptHash:		"WitnessScriptHash",
	NullData:				"NullData",
	NonStandard:			"NonStandard",
}


func (s ScriptType) String() string {
	return scriptTypeToName[s]
}


func verifyPubKeyLength(data []byte) bool {
	return len(data) == 33 || len(data) == 65
}

func IsPubKey(codes []*parsedOpcode) bool {
	return len(codes) == 2 && verifyPubKeyLength(codes[0].Data) && codes[1].Opcode.Value == OP_CHECKSIG
}

func isPubkeyHash(pops []*parsedOpcode) bool {
	return len(pops) == 5 &&
		pops[0].Opcode.Value == OP_DUP &&
		pops[1].Opcode.Value == OP_HASH160 &&
		pops[2].Opcode.Value == OP_DATA_20 &&
		pops[3].Opcode.Value == OP_EQUALVERIFY &&
		pops[4].Opcode.Value == OP_CHECKSIG
}

func isSmallInt(op *opcode) bool {
	if op.Value == OP_0 || (op.Value >= OP_1 && op.Value <= OP_16) {
		return true
	}
	return false
}

// asSmallInt returns the passed opcode, which must be true according to
// isSmallInt(), as an integer.
func opToSmallInt(op *opcode) int {
	if op.Value == OP_0 {
		return 0
	}

	return int(op.Value - (OP_1 - 1))
}

func isMultiSig(pops []*parsedOpcode) bool {
	// The absolute minimum is 1 pubkey:
	// OP_0/OP_1-16 <pubkey> OP_1 OP_CHECKMULTISIG
	l := len(pops)
	if l < 4 {
		return false
	}
	if !isSmallInt(pops[0].Opcode) {
		return false
	}
	if !isSmallInt(pops[l-2].Opcode) {
		return false
	}
	if pops[l-1].Opcode.Value != OP_CHECKMULTISIG {
		return false
	}

	// Verify the number of pubkeys specified matches the actual number
	// of pubkeys provided.
	if l-3 != opToSmallInt(pops[l-2].Opcode) {
		return false
	}

	for _, pop := range pops[1 : l-2] {
		// Valid pubkeys are either 33 or 65 bytes.
		if verifyPubKeyLength(pop.Data) == false {
			return false
		}
	}
	return true
}

func isScriptHash(pops []*parsedOpcode) bool {
	return len(pops) == 3 &&
		pops[0].Opcode.Value == OP_HASH160 &&
		pops[1].Opcode.Value == OP_DATA_20 &&
		pops[2].Opcode.Value == OP_EQUAL
}

func isWitnessPubKeyHash(pops []*parsedOpcode) bool {
	return len(pops) == 2 &&
		pops[0].Opcode.Value == OP_0 &&
		pops[1].Opcode.Value == OP_DATA_20
}

func isWitnessScriptHash(pops []*parsedOpcode) bool {
	return len(pops) == 2 &&
		pops[0].Opcode.Value == OP_0 &&
		pops[1].Opcode.Value == OP_DATA_32
}

func ParseScript(script []byte) ([]*parsedOpcode, error) {
	var pops []*parsedOpcode
	for i:=0; i<(len(script));{
		b := script[i]
		op := &opcodeArray[b]
		pop := parsedOpcode{Opcode: op}
		switch {
		case op.length == 1:
			i++
		case op.length > 1:
			data := script[i+1: i+op.length]
			pop.Data = data
			i = i + op.length
		case op.length < 0:
			var value uint
			switch op.length {
			case -1:
				value = uint(script[i+1])
			case -2:
				value =  uint(binary.LittleEndian.Uint64(script[i+1: i+3]))
			case -4:
				value =  uint(binary.LittleEndian.Uint64(script[i+1: i+5]))
			default:
				return nil, errors.New("invalid opcode length")
			}
			pop.Data = script[i+1: i+1+int(value)]
			i = i + 1 - op.length + int(value)
		}
		pops = append(pops, &pop)
	}
	return pops, nil
}

func SegwitAddrToPublicKeyHash(addr string) (byte, []byte, error){
	hrp := addr[:2]
	version, data, err := SegwitDecode(hrp, addr)
	if err != nil {
		return 0, nil, err
	}

	return byte(version), data, nil
}

func ParsePubKey(pops []*parsedOpcode, stype ScriptType) ([][]byte, error) {
	switch stype {
	case PubKey:
		return [][]byte{pops[0].Data}, nil
	case PubKeyHash:
		return [][]byte{pops[2].Data}, nil
	case Multisig:
		var arr [][]byte
		for _, pop := range pops[1 : len(pops)-2] {
			arr = append(arr, pop.Data)
		}
		return arr, nil
	case ScriptHash:
		return [][]byte{pops[1].Data}, nil
	case WitnessPubKeyHash:
		return [][]byte{pops[1].Data}, nil
	case WitnessScriptHash:
		return [][]byte{pops[1].Data}, nil
	default:
		return [][]byte{}, errors.New("not support type")
	}
}
