package database

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	badger "github.com/dgraph-io/badger/v4"
)

const (
	dbPathEnv   = "GOTRACK_DB_PATH"
	defaultPath = "~/gotrack"
)

func Persist(key string, val []byte) error {
	db, err := open()
	if err != nil {
		return fmt.Errorf("cannot save to DB, %w", err)
	}
	defer db.Close()
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(val))
		return err
	})
	if err != nil {
		return fmt.Errorf("cannot save to DB, %w", err)
	}
	return nil
}

func Delete(key string) error {
	db, err := open()
	if err != nil {
		return fmt.Errorf("cannot delete from DB, %w", err)
	}
	defer db.Close()
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
	if err != nil {
		return fmt.Errorf("cannot delete from DB, %w", err)
	}
	return nil
}

func Get(key string) ([]byte, error) {
	db, err := open()
	if err != nil {
		return nil, fmt.Errorf("cannot get key from DB, %w", err)
	}
	defer db.Close()
	output := make([]byte, 0)
	err = db.View(func(txn *badger.Txn) error {
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
		return nil, fmt.Errorf("cannot get key from DB, %w", err)
	}
	return output, nil
}

func open() (*badger.DB, error) {
	path, err := dbPath()
	if err != nil {
		return nil, fmt.Errorf("cannot open database, %w", err)
	}
	db, err := badger.Open(badger.DefaultOptions(path).WithLogger(nil))
	if err != nil {
		return nil, fmt.Errorf("cannot open database, %w", err)
	}
	return db, nil
}

func dbPath() (string, error) {
	path, isPresent := os.LookupEnv(dbPathEnv)
	if !isPresent {
		path = defaultPath
	}
	if strings.Contains(path, "~") {
		newPath, err := replaceHomePath(path)
		if err != nil {
			return "", err
		}
		path = newPath
	}
	return path, nil
}

func replaceHomePath(path string) (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("could not get current user, %w", err)
	}
	return strings.Replace(path, "~", currentUser.HomeDir, 1), nil
}
