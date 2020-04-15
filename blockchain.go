package main

import (
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

const blkreg = ".+blk[0-9]+.dat"

type Blockchain struct {
	path		string
	blks        map[int]*Blk
	index  		*IndexDB
}

// NewBlockchain return Blockchain struct can search blocks by blkfile, or block by hash, height,
// you must set blk files path (usually at "~/.bitcoin/blocks") and index files path (usually at
// "~/.bitcoin/blocks/index"), if you doesn't set indexpath, then you can only use GetUnorderBlocks
func NewBlockchain(blockpath string, indexpath string) (*Blockchain, error) {
	blks := make(map[int]*Blk)
	var indexdb *IndexDB
	var err error
	r, _ := regexp.Compile(blkreg)
	if indexpath == "" {
		indexdb = nil
	} else {
		indexpath, err = filepath.Abs(indexpath)
		if err != nil {
			return nil, err
		}
		indexdb, err = NewIndexDB(indexpath)
		if err != nil {
			return nil, err
		}
	}
	blockpath, err = filepath.Abs(blockpath)
	if err != nil {
		return nil, err
	}
	fileInfo, err := os.Stat(blockpath)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() == false {
		m := r.MatchString(blockpath)
		if m == false {
			return nil, errors.New("there is no file match")
		}
		blk, err := NewBlk(blockpath)
		if err != nil {
			return nil, err
		}
		blks[blk.name] = blk
	} else {
		err = filepath.Walk(blockpath, func(path string, info os.FileInfo, err error) error {
			m := r.MatchString(path)
			if m == true {
				blk, err := NewBlk(path)
				if err != nil {
					return err
				}
				blks[blk.name] = blk
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return &Blockchain{
		path:	   blockpath,
		blks: 	   blks,
		index:   indexdb,
	}, nil
}

func (b *Blockchain) GetBlockByHeight(height int) (*Block, error){
	if b.index == nil {
		return nil, errors.New("doesn't set index db, must specify index path")
	}
	for name, blk := range b.blks {
		findex, err := b.index.getFileIndex(uint32(name))
		if err != nil {
			return nil, err
		}
		if height >= findex.NHeightFirst && height <= findex.NHeightLast {
			blocks, err := blk.ReadBlkFile()
			if err != nil {
				return nil, err
			}
			for _ , block := range blocks {
				bindex, err := b.index.getBlockIndexByBlock(block.Hash)
				if err != nil {
					return nil, err
				}
				// not found index in db
				if bindex == nil {
					continue
				}
				if bindex.Height == height {
					return block, nil
				}
			}
		}
	}
	return nil, nil
}

func (b *Blockchain) GetBlockByHash(hash []byte) (*Block, error){
	if b.index == nil {
		return nil, errors.New("doesn't set index db, must specify index path")
	}
	index, err := b.index.getBlockIndexByBlock(hash)
	if err != nil {
		return nil, err
	}
	// hash isn't in db
	if index == nil{
		return nil, nil
	}
	blk := b.blks[index.File]
	f, err:= os.Open(blk.path)
	if err != nil {
		return nil, err
	}
	_, err = f.Seek(int64(index.DataPos-8), 0)
	if err != nil {
		return nil, err
	}
	prefix := make([]byte, 8)
	_ , err = f.Read(prefix)
	if err != nil {
		return nil, err
	}
	params, err := MagicBytesToParams(prefix[:4])
	if err != nil {
		return nil, err
	}
	blocksize := binary.LittleEndian.Uint32(prefix[4:])
	blockbytes := make([]byte, blocksize)
	_, err = f.Read(blockbytes)
	if err != nil {
		return nil, err
	}
	block, err := DeserializeBlock(blockbytes,params)
	if err != nil {
		return nil, err
	}
	return block ,nil
}

func (b *Blockchain) GetLastHeight() (int, error)  {
	if b.index == nil {
		return 0, errors.New("doesn't set index db, must specify index path")
	}
	index, err := b.index.getLastFileIndex()
	if err != nil {
		return 0, err
	}
	return index.NHeightLast, nil
}

// GetUnorderBlocks get unorderblocks in blk file, start and end are set how many blocks you wanna
// get, name is on behalf of blk file's name (if you want to search about blk00001.dat, then name
// is 1)
func (b *Blockchain) GetUnorderBlocks(start int , end int, name int) ([]*Block, error) {
	var blocks []*Block
	var err error
	blk := b.blks[name]
	blocks, err = blk.ReadBlkFile()
	if err != nil {
		return nil, err
	}
	start, end, err = checkStartAndEnd(start, end, len(blocks))
	if err != nil {
		return nil, err
	}
	return blocks[start:end], nil
}

// GetUnorderBlocks getorderblocks in blk file, start and end are set how many blocks you wanna
// get, name is on behalf of blk file's name (if you want to search about blk00001.dat, then name
// is 1), indexpath needs to be set
func (b *Blockchain) GetOrderBlocks(start int, end int, name int) ([]*Block, error) {
	if b.index == nil {
		return nil, errors.New("doesn't set index db, must specify index path")
	}
	var keys []int
	var blocks []*Block
	unorder, err := b.GetUnorderBlocks(0,99999999, name)
	order := make(map[int]*Block, len(unorder))
	if err != nil {
		return nil, err
	}
	for _ , block := range unorder {
		index, err := b.index.getBlockIndexByBlock(block.Hash)
		if err != nil {
			return nil, err
		}
		// if block not found, then index == nil
		if index != nil {
			block.height = index.Height
			order[index.Height] = block
		}
	}
	for k := range order {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		blocks = append(blocks, order[k])
	}
	start, end, err = checkStartAndEnd(start, end, len(blocks))
	if err != nil {
		return nil, err
	}
	return blocks[start:end], nil
}

func (b *Blockchain) Close() error{
	err := b.index.Close()
	if err != nil {
		return err
	}
	return nil
}

func checkStartAndEnd(start int, end int, limit int) (int, int, error) {
	if start < 0 || start > end || start > limit || end < 0 || end < start  {
		return 0, 0, errors.New("start or end is wrong")
	}
	if end > limit {
		end =  limit
	}
	return start, end, nil
}