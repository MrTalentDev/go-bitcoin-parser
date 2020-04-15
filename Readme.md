Bitcoin Blockchain Parser
------

This repository contains the golang code of bitcoin blockchain parser, you can use cli interface
to interact with parser or use it as library.

Contents
--------
- [Feature](#feature)
- [Setup](#setup)
- [Example](#example)
- [Test](#test)

Feature
--------

- Deserialize block (provide raw data)
- Deserialize transaction (provide raw data)
- Get Block by hash (provide bitcoin blk folder and index folder)
- Get Block by height (provide bitcoin blk folder and index folder)
- Get unorder blocks (provide bitcoin blk folder)
- Get order blocks (provide bitcoin blk folder and index folder)
- Get last height (provide bitcoin blk folder and index folder)

Setup
------

## Setup config

We use config to set blockpath and indexpath, if you want to use deserialize block or transaction, 
you don't need to set config. but if you want to use function involved blk file, you need to set 
config in file or specify it in cli's variable.

### config path

Environmental variables: BITCOIN_PARSER_CONFIG_PATH

if you doesn't set env, then default path is parser.conf in current folder.

### parser.conf

#### Example
```
blockpath=./test-data/blocks
indexpath=./test-data/index
```

## Build 

```shell script
go get github.com/syndtr/goleveldb/leveldb
go get golang.org/x/crypto/ripemd160
go bulid
```

## Run

```shell script
./bitcoin-parser
```

Example
-------

## Deserialize block

### code

```go

    var BlockTestBytes = []byte{
        0x01,0x00,0x00,0x00,0x6f,0xe2,0x8c,0x0a,
        0xb6,0xf1,0xb3,0x72,0xc1,0xa6,0xa2,0x46,
        0xae,0x63,0xf7,0x4f,0x93,0x1e,0x83,0x65,
        0xe1,0x5a,0x08,0x9c,0x68,0xd6,0x19,0x00,
        0x00,0x00,0x00,0x00,0x98,0x20,0x51,0xfd,
        0x1e,0x4b,0xa7,0x44,0xbb,0xbe,0x68,0x0e,
        0x1f,0xee,0x14,0x67,0x7b,0xa1,0xa3,0xc3,
        0x54,0x0b,0xf7,0xb1,0xcd,0xb6,0x06,0xe8,
        0x57,0x23,0x3e,0x0e,0x61,0xbc,0x66,0x49,
        0xff,0xff,0x00,0x1d,0x01,0xe3,0x62,0x99,
        0x01,0x01,0x00,0x00,0x00,0x01,0x00,0x00,
        0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
        0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
        0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,
        0x00,0x00,0x00,0x00,0x00,0x00,0xff,0xff,
        0xff,0xff,0x07,0x04,0xff,0xff,0x00,0x1d,
        0x01,0x04,0xff,0xff,0xff,0xff,0x01,0x00,
        0xf2,0x05,0x2a,0x01,0x00,0x00,0x00,0x43,
        0x41,0x04,0x96,0xb5,0x38,0xe8,0x53,0x51,
        0x9c,0x72,0x6a,0x2c,0x91,0xe6,0x1e,0xc1,
        0x16,0x00,0xae,0x13,0x90,0x81,0x3a,0x62,
        0x7c,0x66,0xfb,0x8b,0xe7,0x94,0x7b,0xe6,
        0x3c,0x52,0xda,0x75,0x89,0x37,0x95,0x15,
        0xd4,0xe0,0xa6,0x04,0xf8,0x14,0x17,0x81,
        0xe6,0x22,0x94,0x72,0x11,0x66,0xbf,0x62,
        0x1e,0x73,0xa8,0x2c,0xbf,0x23,0x42,0xc8,
        0x58,0xee,0xac,0x00,0x00,0x00,0x00,
    }
    block, err := DeserializeBlock(BlockTestBytes, &MainnetParams)
    if err != nil {
	t.Errorf("DeserializeBlock error:%s", err)
    }
    str, err := block.String()
    if err != nil {
        t.Errorf("block to string error:%s", err)
    }
    fmt.Println("%s", str)
```

### cli
```shell script
./bitcoin-parser deserializeblock -net mainnet -raw 010000006fe28c0ab6f1b372c1a6a246ae63f74f931e8365e15a089c68d6190000000000982051fd1e4ba744bbbe680e1fee14677ba1a3c3540bf7b1cdb606e857233e0e61bc6649ffff001d01e362990101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0704ffff001d0104ffffffff0100f2052a0100000043410496b538e853519c726a2c91e61ec11600ae1390813a627c66fb8be7947be63c52da7589379515d4e0a604f8141781e62294721166bf621e73a82cbf2342c858eeac00000000
```

### output
```json
{
	"hash": "00000000839a8e6886ab5951d76f411475428afc90947ee320161bbf18eb6048",
	"blockHeader": {
		"version": 1,
		"prevBlockHeaderHash": "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f",
		"merkleRootHash": "0e3e2357e806b6cdb1f70b54c3a3a17b6714ee1f0e68bebb44a74b1efd512098",
		"time": 1231469665,
		"nbits": 486604799,
		"nonce": 2573394689
	},
	"txnCount": 1,
	"txns": [
		{
			"hash": "0e3e2357e806b6cdb1f70b54c3a3a17b6714ee1f0e68bebb44a74b1efd512098",
			"version": 1,
			"txinCount": 1,
			"txin": [
				{
					"hash": "0000000000000000000000000000000000000000000000000000000000000000",
					"index": 4294967295,
					"scriptBytes": 7,
					"signatureScript": "04ffff001d0104",
					"sequence": 4294967295
				}
			],
			"txOutCount": 1,
			"txOut": [
				{
					"value": 5000000000,
					"pkScriptBytes": 67,
					"pkScript": {
						"pkscript": "410496b538e853519c726a2c91e61ec11600ae1390813a627c66fb8be7947be63c52da7589379515d4e0a604f8141781e62294721166bf621e73a82cbf2342c858eeac",
						"pops": [
							{
								"opcode": {
									"value": "41",
									"name": "OP_DATA_65"
								},
								"data": "0496b538e853519c726a2c91e61ec11600ae1390813a627c66fb8be7947be63c52da7589379515d4e0a604f8141781e62294721166bf621e73a82cbf2342c858ee"
							},
							{
								"opcode": {
									"value": "ac",
									"name": "OP_CHECKSIG"
								},
								"data": ""
							}
						],
						"stype": "PubKey",
						"addresses": [
							"12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX"
						]
					}
				}
			],
			"witness": null,
			"lockTime": 0
		}
	]
}
```

## Deserialize Transaction

#### code

```go

    transactionTestBytes = []byte{
	0x01,0x00,0x00,0x00,0x00,0x01,0x01,0x9b,
	0x34,0xe2,0x9c,0x56,0x2d,0xe2,0x0e,0xf3,
	0x26,0x68,0x2f,0xa5,0x82,0x67,0x98,0x47,
	0xa4,0xdc,0x4a,0xc7,0xd5,0x26,0x4d,0x32,
	0x28,0xa1,0x7c,0x8f,0xef,0x91,0x0c,0x01,
	0x00,0x00,0x00,0x00,0xff,0xff,0xff,0xff,
	0x01,0x16,0xaa,0x0f,0x00,0x00,0x00,0x00,
	0x00,0x16,0x00,0x14,0x52,0x15,0x83,0x08,
	0xca,0x2e,0x51,0x49,0xd9,0x39,0x73,0x10,
	0xec,0x1f,0xe2,0xa4,0xf8,0x8a,0xfb,0x07,
	0x02,0x48,0x30,0x45,0x02,0x21,0x00,0xce,
	0x94,0x6d,0x85,0x63,0x0e,0x5d,0x4b,0xb2,
	0x27,0xf6,0x96,0x78,0xfa,0xff,0x2d,0x39,
	0x7a,0xdb,0x25,0x61,0x7c,0xc9,0xb4,0xa5,
	0x32,0x3f,0x4b,0x76,0x01,0x21,0xe8,0x02,
	0x20,0x4d,0x05,0x0a,0x9c,0x08,0x83,0x30,
	0xf2,0x36,0xae,0xc8,0x80,0x75,0xde,0x5c,
	0xa6,0x31,0xa5,0xb7,0xa3,0xae,0xa6,0x92,
	0x37,0x1d,0x0f,0x99,0xc9,0x62,0xd7,0x8d,
	0x85,0x01,0x21,0x03,0x2c,0xc6,0x3e,0xba,
	0x4d,0x1f,0x35,0xb6,0x38,0xe7,0x36,0x86,
	0xf7,0x68,0xc2,0xc9,0x63,0x99,0xe6,0x9d,
	0xc7,0x3f,0xcc,0x39,0xfe,0x30,0xb5,0x43,
	0xed,0x85,0x9c,0x19,0x00,0x00,0x00,0x00,
    }
    tx, _, err := DesserializeTransaction(transactionTestBytes, &MainnetParams)
    if err != nil {
	t.Errorf("TestDesserializeTransaction error: %s", err)
    }
    str, err := tx.String()
    if err != nil {
        t.Errorf("transaction to string error:%s", err)
    }
    fmt.Println("%s", str)
```

### cli

```shell script
./bitcoin-parser deserializetransaction -net mainnet -raw 010000000001019b34e29c562de20ef326682fa582679847a4dc4ac7d5264d3228a17c8fef910c0100000000ffffffff0116aa0f000000000016001452158308ca2e5149d9397310ec1fe2a4f88afb0702483045022100ce946d85630e5d4bb227f69678faff2d397adb25617cc9b4a5323f4b760121e802204d050a9c088330f236aec88075de5ca631a5b7a3aea692371d0f99c962d78d850121032cc63eba4d1f35b638e73686f768c2c96399e69dc73fcc39fe30b543ed859c1900000000
```

### output

```json
{
	"hash": "ff857d64d03182391620a933fc8477372af4bac3a3f737883a57882d7c7d8f36",
	"version": 1,
	"txinCount": 1,
	"txin": [
		{
			"hash": "9b34e29c562de20ef326682fa582679847a4dc4ac7d5264d3228a17c8fef910c",
			"index": 1,
			"scriptBytes": 0,
			"signatureScript": "",
			"sequence": 4294967295
		}
	],
	"txOutCount": 1,
	"txOut": [
		{
			"value": 1026582,
			"pkScriptBytes": 22,
			"pkScript": {
				"pkscript": "001452158308ca2e5149d9397310ec1fe2a4f88afb07",
				"pops": [
					{
						"opcode": {
							"value": "00",
							"name": "OP_0"
						},
						"data": ""
					},
					{
						"opcode": {
							"value": "14",
							"name": "OP_DATA_20"
						},
						"data": "52158308ca2e5149d9397310ec1fe2a4f88afb07"
					}
				],
				"stype": "WitnessPubKeyHash",
				"addresses": [
					"bc1q2g2cxzx29eg5nkfewvgwc8lz5nug47c8ta5ene"
				]
			}
		}
	],
	"witness": [
		{
			"WitnessStackCount": 2,
			"WitnessScript": [
				{
					"WitnessScriptBytes": 72,
					"WitnessScript": "3045022100ce946d85630e5d4bb227f69678faff2d397adb25617cc9b4a5323f4b760121e802204d050a9c088330f236aec88075de5ca631a5b7a3aea692371d0f99c962d78d8501"
				},
				{
					"WitnessScriptBytes": 33,
					"WitnessScript": "032cc63eba4d1f35b638e73686f768c2c96399e69dc73fcc39fe30b543ed859c19"
				}
			]
		}
	],
	"lockTime": 0
}
```

## Get Block by hash

### cli
```shell script
./bitcoin-parser -hash 00000000000000322cee8a774b3bc50bd304f92ee63a7d8c7c01bc9e71f182ae -indexpath ./test-data/index -blockpath ./test-data/blocks
```

## Get Block by height

### cli
```shell script
./bitcoin-parser getblockbyheight -height 131301 -indexpath ./test-data/index -blockpath ./test-data/blocks
```

## Get unorder blocks

blocks in blk files are not in order, this instruction will show original order blocks in blk files. 

### cli
```shell script
./bitcoin-parser getunorderblocks -start 0 -end 999999 -name 2 -blockpath ./test-data/blocks
```

## Get order blocks

this instruction will show blocks in blk files in order by height.

### cli
```shell script
./bitcoin-parser getorderblocks -start 0 -end 999999 -name 2 -indexpath ./test-data/index -blockpath ./test-data/blocks
```

## Get last height
### cli
```shell script
./bitcoin-parser getlastheight -indexpath ./test-data/index -blockpath ./test-data/blocks
```

Test
-----
if you want find more use case, you can find in test case.

```shell script
go test
```