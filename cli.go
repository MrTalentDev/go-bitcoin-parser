package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type CLI struct {
	blockchain *Blockchain
}

func(cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("deserializeblock -raw RAW -net NET - deserialize block by raw block")
	fmt.Println("deserializetransaction -raw RAW -net NET - deserialize block by raw transaction")
	fmt.Println("getblockbyhash -hash HASH -indexpath INDEXPATH -blockpath BLOCKPATH - get block by block hash")
	fmt.Println("getblockbyheight -height HEIGHT -indexpath INDEXPATH -blockpath BLOCKPATH - get block by height")
	fmt.Println("getunorderblocks -start START -end END -name NAME -blockpath BLOCKPATH - get unorder blocks ")
	fmt.Println("getorderblocks -start START -end END -name NAME -indexpath INDEXPATH -blockpath BLOCKPATH - get order blocks ")
	fmt.Println("getlastheight -indexpath INDEXPATH -blockpath BLOCKPATH - get last height in indexdb")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) loadConfig(){
	path := os.Getenv("BITCOIN_PARSER_CONFIG_PATH")
	if path == "" || fileExists(path) == false {
		currentPath, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = currentPath + "/parser.conf"
		if fileExists(path) == false {
			fmt.Printf("doesn't set config\n")
		} else {
			blockchain, err := loadConfigFile(path)
			if err != nil {
				fmt.Printf("load config file occur error, path is: %s, error is : %s\n",path , err)
				return
			}
			cli.blockchain = blockchain
		}
	}else {
		blockchain, err := loadConfigFile(path)
		if err != nil {
			fmt.Printf("load config file occur error, path is: %s, error is : %s\n",path , err)
			return
		}
		cli.blockchain = blockchain
	}
}

func loadConfigFile(path string) (*Blockchain, error){
	var blockpath string
	var indexpath string
	configsb, err := ioutil.ReadFile(path)
	if err != nil{
		return nil, err
	}
	configs := strings.Split(string(configsb), "\n")
	for _, config := range configs {
		kv := strings.Split(config,"=")
		switch kv[0] {
		case "blockpath":
			blockpath = kv[1]
		case "indexpath":
			indexpath = kv[1]
		default:
			break
		}
	}
	blockchain, err := NewBlockchain(blockpath,indexpath)
	if err != nil {
		return nil, err
	}
	return blockchain, nil
}

func (cli *CLI) setBlockchain(blockpath string, indexpath string){
	if blockpath != "" {
		if cli.blockchain != nil {
			cli.blockchain.Close()
		}
		blockchain, err := NewBlockchain(blockpath, indexpath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cli.blockchain = blockchain
	}
	if cli.blockchain == nil {
		fmt.Printf("blockchain is not defined, you must specify indexpath and blockpath\n")
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()
	cli.loadConfig()
	deserializeblockCmd := flag.NewFlagSet("deserializeblock", flag.ExitOnError)
	deserializetransactionCmd := flag.NewFlagSet("deserializetransaction", flag.ExitOnError)
	getblockbyhashCmd := flag.NewFlagSet("getblockbyhash", flag.ExitOnError)
	getblockbyheightCmd := flag.NewFlagSet("getblockbyheight", flag.ExitOnError)
	getunorderblocksCmd := flag.NewFlagSet("getunorderblocks", flag.ExitOnError)
	getorderblocksCmd := flag.NewFlagSet("getorderblocks", flag.ExitOnError)
	getlastheightCmd := flag.NewFlagSet("getlastheightCmd", flag.ExitOnError)

	deserializeblockCmdRaw := deserializeblockCmd.String("raw","", "raw block")
	deserializeblockCmdNet := deserializeblockCmd.String("net", "mainnet", "input mainnet or testnet")
	deserializetransactionCmdRaw := deserializetransactionCmd.String("raw","","raw transaction")
	deserializetransactionCmdNet := deserializetransactionCmd.String("net","mainnet","input mainnet or testnet")
	getblockbyhashCmdHash := getblockbyhashCmd.String("hash","","block hash")
	getblockbyhashCmdBlockPath := getblockbyhashCmd.String("blockpath","","block path")
	getblockbyhashCmdIndexPath := getblockbyhashCmd.String("indexpath","", "index path")
	getblockbyheightCmdHeight := getblockbyheightCmd.Int("height",0,"block height")
	getblockbyheightCmdBlockPath := getblockbyheightCmd.String("blockpath","","block path")
	getblockbyheightCmdIndexPath := getblockbyheightCmd.String("indexpath","", "index path")
	getunorderblocksCmdStart := getunorderblocksCmd.Int("start",0,"start number")
	getunorderblocksCmdEnd := getunorderblocksCmd.Int("end",999999,"end number")
	getunorderblocksCmdName := getunorderblocksCmd.Int("name",0,"blk file name")
	getunorderblocksCmdBlockPath := getunorderblocksCmd.String("blockpath","","block path")
	getorderblocksCmdStart := getorderblocksCmd.Int("start",0,"start number")
	getorderblocksCmdEnd := getorderblocksCmd.Int("end",9999999,"end number")
	getorderblocksCmdName := getorderblocksCmd.Int("name",0,"blk file name")
	getorderblocksCmdBlockPath := getorderblocksCmd.String("blockpath","", "block path")
	getorderblocksCmdIndexPath := getorderblocksCmd.String("indexpath","", "index path")
	getlastheightCmdBlockPath := getlastheightCmd.String("blockpath","", "block path")
	getlastheightCmdIndexPath := getlastheightCmd.String("indexpath","", "index path")

	switch os.Args[1] {
	case "deserializeblock":
		err := deserializeblockCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	case "deserializetransaction":
		err := deserializetransactionCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	case "getblockbyhash":
		err := getblockbyhashCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	case "getblockbyheight":
		err := getblockbyheightCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	case "getunorderblocks":
		err := getunorderblocksCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	case "getorderblocks":
		err := getorderblocksCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	case "getlastheight":
		err := getlastheightCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if deserializeblockCmd.Parsed() {
		if *deserializeblockCmdRaw == "" {
			deserializeblockCmd.Usage()
			//os.Exit(1)
		}
		switch *deserializeblockCmdNet {
		case "mainnet":
			rawToPrintBlock(*deserializeblockCmdRaw, &MainnetParams)
		case "testnet":
			rawToPrintBlock(*deserializeblockCmdRaw, &TestnetParams)
		}
	}

	if deserializetransactionCmd.Parsed() {
		if *deserializetransactionCmdRaw == "" {
			deserializetransactionCmd.Usage()
			//os.Exit(1)
		}
		switch *deserializetransactionCmdNet {
		case "mainnet":
			rawToPrintTransaction(*deserializetransactionCmdRaw, &MainnetParams)
		case "testnet":
			rawToPrintTransaction(*deserializetransactionCmdRaw, &TestnetParams)
		}
	}

	if getblockbyhashCmd.Parsed() {
		if *getblockbyhashCmdHash == "" {
			getblockbyhashCmd.Usage()
		}
		cli.setBlockchain(*getblockbyhashCmdBlockPath, *getblockbyhashCmdIndexPath)
		hashb, err := hex.DecodeString(*getblockbyhashCmdHash)
		if err != nil {
			fmt.Printf("raw string to bytes error: %s\n", err)
			os.Exit(1)
		}
		block, err := cli.blockchain.GetBlockByHash(hashb)
		defer cli.blockchain.Close()
		if err != nil {
			fmt.Printf("getBlockByHash error: %s\n", err)
			os.Exit(1)
		}
		if block == nil {
			fmt.Println("block isn't exist in db")
			os.Exit(1)
		}
		blockstr, err := block.String()
		if err != nil {
			fmt.Printf("block to string happened error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(blockstr)
	}

	if getblockbyheightCmd.Parsed() {
		if *getblockbyheightCmdHeight == 0 {
			getblockbyheightCmd.Usage()
		}
		cli.setBlockchain(*getblockbyheightCmdBlockPath, *getblockbyheightCmdIndexPath)
		block, err:= cli.blockchain.GetBlockByHeight(*getblockbyheightCmdHeight)
		defer cli.blockchain.Close()
		if err != nil {
			fmt.Printf("getBlockByHeight error: %s\n", err)
			os.Exit(1)
		}
		if block == nil {
			fmt.Println("block isn't exist in db")
			os.Exit(1)
		}
		blockstr, err := block.String()
		if err != nil {
			fmt.Printf("block to string happened error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(blockstr)
	}

	if getunorderblocksCmd.Parsed() {
		if *getunorderblocksCmdName == 0 {
			getunorderblocksCmd.Usage()
			os.Exit(1)
		}
		cli.setBlockchain(*getunorderblocksCmdBlockPath, "")
		blocks, err := cli.blockchain.GetUnorderBlocks(*getunorderblocksCmdStart, *getunorderblocksCmdEnd, *getunorderblocksCmdName)
		defer cli.blockchain.Close()
		if err != nil {
			fmt.Printf("GetUnorderBlocks error: %s\n", err)
			os.Exit(1)
		}
		for _, block := range blocks {
			blockstr, err := block.String()
			if err != nil {
				fmt.Printf("block to string happened error: %s\n", err)
				os.Exit(1)
			}
			fmt.Println(blockstr)
		}
	}


	if getorderblocksCmd.Parsed() {
		if *getorderblocksCmdName == 0 {
			getorderblocksCmd.Usage()
			os.Exit(1)
		}
		cli.setBlockchain(*getorderblocksCmdBlockPath, *getorderblocksCmdIndexPath)
		blocks, err := cli.blockchain.GetOrderBlocks(*getorderblocksCmdStart, *getorderblocksCmdEnd, *getorderblocksCmdName)
		defer cli.blockchain.Close()
		if err != nil {
			fmt.Printf("GetOrderBlocks error: %s\n", err)
			os.Exit(1)
		}
		for _, block := range blocks {
			blockstr, err := block.String()
			if err != nil {
				fmt.Printf("block to string happened error: %s\n", err)
				os.Exit(1)
			}
			fmt.Println(blockstr)
		}
	}

	if getlastheightCmd.Parsed(){
		cli.setBlockchain(*getlastheightCmdBlockPath, *getlastheightCmdIndexPath)
		height, err := cli.blockchain.GetLastHeight()
		defer cli.blockchain.Close()
		if err != nil {
			fmt.Printf("GetLastHeight error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(height)
	}
}


func rawToPrintBlock(raw string, params *Params){
	rawb, err := hex.DecodeString(raw)
	if err != nil {
		fmt.Printf("raw string to bytes error: %s\n", err)
	}
	block, err := DeserializeBlock(rawb, params)
	if err != nil {
		fmt.Printf("raw transfer to block happened error: %s\n", err)
	}
	blockstr, err := block.String()
	if err != nil {
		fmt.Printf("block to string happened error: %s\n", err)
	}
	fmt.Println(blockstr)
}

func rawToPrintTransaction(raw string, params *Params){
	txb, err := hex.DecodeString(raw)
	if err != nil {
		fmt.Printf("raw string to bytes error: %s\n", err)
	}
	tx, _, err := DeserializeTransaction(txb, params)
	if err != nil {
		fmt.Printf("raw transfer to transaction happened error: %s\n", err)
	}
	txstr, err := tx.String()
	if err != nil {
		fmt.Printf("transaction to string happened error: %s\n", err)
	}
	fmt.Println(txstr)
}