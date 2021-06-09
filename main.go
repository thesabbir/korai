package main

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/thesabbir/korai/packages/server"
	"log"
)

func main() {
	fmt.Println("Hello")
	db, err := badger.Open(badger.DefaultOptions("db/badger"))
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *badger.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("error while closing badger")
		}
	}(db)

	server.Start()

}
