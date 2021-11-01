package db

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

// DB .
type DB struct {
	bolt *bolt.DB
}

// New .
func New(bolt *bolt.DB) *DB {
	return &DB{
		bolt: bolt,
	}
}

// Save .
func (db *DB) Save(id string, text string) error {
	return db.bolt.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("ideas"))
		if err != nil {
			return err
		}

		log.Printf("Save idea with ID %s, %s", id, text)
		err = b.Put([]byte(id), []byte(text))
		return err
	})
}
