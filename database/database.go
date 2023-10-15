package database

import (
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

func Persist(key string, val []byte) {
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

func Delete(key string) {
	db := open()
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
	if err != nil {
		log.Println(err)
	}

}

func Get(key string) []byte {
	db := open()
	output := make([]byte, 0)
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			output = val
			return nil
		})
		return err
	})
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return output
}

func open() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions("/tmp/gotrack").WithLogger(nil))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
