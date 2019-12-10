package main

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
)

func main() {

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_ = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))

		return err
	})

	_ = db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})

	_ = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
}
