package main

import (
	"gobank/api"
	"log"
)

func main() {
	store, err := api.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewApiServer(":3000", store)
	server.Run()
}
