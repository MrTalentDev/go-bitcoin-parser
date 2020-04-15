package main

import (
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type IndexDB struct {
	db *leveldb.DB
}

func NewIndexDB(path string) (*IndexDB, error){
	db, err := leveldb.OpenFile(path, &opt.Options{
		Compression: opt.NoCompression,
		ReadOnly: true,
	})
	if err != nil {
		return nil, err
	}
	return &IndexDB{
		db: db,
	}, nil
}


func (in *IndexDB) getBlockIndexByBlock(hash []byte) (*BlockIndex, error) {
	rhash := ReverseBytes(hash)
	bkey := append([]byte{0x62}, rhash...)
	val, err := in.db.Get(bkey,nil)
	if err != nil {
		if err.Error() == "leveldb: not found"  {
			return nil, nil
		}
		return nil, err
	}
	index, err := DeserializeBlockIndex(val)
	if err != nil {
		return nil, err
	}
	return index, nil
}

func (in *IndexDB) getFileIndex(file uint32) (*FileIndex, error) {
	bfile := IntToLittleEndianBytes(file)
	bfile = append([]byte{0x66}, bfile...)
	val, err := in.db.Get(bfile, nil)
	if err != nil {
		return nil, err
	}
 	index := DeserializeFileIndex(val)
	return index, nil
}

func (in *IndexDB) getLastFileIndex() (*FileIndex, error){
	val, err := in.db.Get([]byte{0x6c}, nil)
	if err != nil {
		return nil, err
	}
	index, err  := in.getFileIndex(binary.LittleEndian.Uint32(val))
	if err != nil {
		return nil, err
	}
	return index, nil
}

func (in *IndexDB) getTransactionIndexByBlock(txhash []byte) ([]byte, error) {
	rhash := ReverseBytes(txhash)
	tkey := append([]byte{0x74}, rhash...)
	val, err:= in.db.Get(tkey, nil)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (in *IndexDB) Close() error {
	err := in.db.Close()
	if err != nil {
		return err
	}
	return nil
}