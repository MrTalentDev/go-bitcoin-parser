package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
)

var blknamereg = ".+blk[0*]+([1-9]{1}[0-9]*).dat"


type Blk struct {
	path string
	name int
}

func NewBlk(path string) (*Blk, error) {
	name, err := BlkPathToName(path)
	if err != nil {
		return nil, err
	}
	return &Blk{
		path: path,
		name: name,
	}, nil
}



func DeserializeBlkBlocks(input []byte) ([]*BlkBlock, error){
	var blocks []*BlkBlock
	calSize := 0
	for calSize < len(input) {
		blkBlock, size, err := DeserializeBlkBlock(input[calSize:])
		if err != nil {
			return nil, err
		}
		// last bytes are all zero, we don't need continue read blocks
		if blkBlock == nil {
			return blocks, nil
		}
		blocks = append(blocks,blkBlock)
		calSize += size
	}
	return blocks, nil
}


func (b *Blk) ReadBlkFile() ([]*Block, error) {
	var blocks []*Block
	var blkBlocks []*BlkBlock
	data, err:= ioutil.ReadFile(b.path)
	if err != nil {
		return nil, err
	}
	blkBlocks, err = DeserializeBlkBlocks(data)
	if err != nil {
		return nil, err
	}
	for _, blkblock := range blkBlocks {
		blocks = append(blocks, blkblock.Block)
	}
	return blocks, nil
}

func BlkPathToName(path string) (int, error) {
	r, _ := regexp.Compile(blknamereg)
	n := r.FindStringSubmatch(path)

	if len(n)<1 {
		return 0, nil
	}
	name, err := strconv.Atoi(n[1])
	if err != nil {
		return 0, err
	}
	return name, nil
}