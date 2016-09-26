package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {

	/*
		o := &opt.Options{
			OpenFilesCacheCapacity: openFilesCacheCapacity,
			DisableBufferPool:      !enableBufferPool,
			DisableBlockCache:      !enableBlockCache,
			ErrorIfExist:           true,
			Compression:            opt.NoCompression,
			}

		db, err := leveldb.Open(stor, o)
	*/

	/*
		o := &opt.Options{
		    Filter: filter.NewBloomFilter(10),
		}
		db, err := leveldb.OpenFile("path/to/db", o)
		defer db.Close()
	*/

	db, err := leveldb.OpenFile("data/db", nil)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	/*
		err = db.Put([]byte("key"), []byte("value"), nil)
		// err = db.Delete([]byte("key"), nil)
		if err != nil {
			fmt.Println(err)
		}
	*/

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		fmt.Println(string(key), string(value))
	}
	iter.Release()
	err = iter.Error()

}
