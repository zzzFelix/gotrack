package database

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

func Persist(key string, val string) {
	db := open()

	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(val))
		return err
	})
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
}

func Get(key string) {
	db := open()
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			fmt.Printf("%s\n", val)
			return nil
		})
		return err
	})
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
}

func open() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions("/tmp/gotrack").WithLogger(nil))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
