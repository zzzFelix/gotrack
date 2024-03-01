package database

import (
	"log"
	"os"

	badger "github.com/dgraph-io/badger/v4"
)

const (
	dbPathEnv   = "GOTRACK_DB_PATH"
	defaultPath = "gotrack"
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

func Get(key string) ([]byte, error) {
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
	defer db.Close()

	if err != nil {
		return nil, err
	}
	return output, nil
}

func open() *badger.DB {
	path := dbPath()
	db, err := badger.Open(badger.DefaultOptions(path).WithLogger(nil))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func dbPath() string {
	path, isPresent := os.LookupEnv(dbPathEnv)
	if !isPresent {
		return defaultPath
	}
	return path
}
