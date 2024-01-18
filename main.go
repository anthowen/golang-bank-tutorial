package main

import (
	"fmt"
	"gobank/api"
	"log"
)

func main() {
	// Some testing (not related to API)
	cIn, cOut := PrimeChannelTest()
	cIn <- primemsg{100, false}

	store, err := api.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	msg := <-cOut
	fmt.Println(msg.num, msg.isPrime)

	server := api.NewApiServer(":3000", store)
	server.Run()
}
