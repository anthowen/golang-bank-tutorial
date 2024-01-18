package main

import (
	"fmt"
	"math"
	"math/rand"
)

type primemsg struct {
	num     int
	isPrime bool
}

const NUMBERS = 100
const WORKERS = 3

func isPrime(cIn chan primemsg, cOut chan primemsg) {
	id := rand.Intn(1000000)

	i := 0

	for {
		msg := <-cIn
		num := msg.num
		fmt.Println(id, "is testing", num)

		sq_root := int(math.Sqrt(float64(num)))

		for i = 2; i <= sq_root; i++ {
			if num%i == 0 {
				msg.isPrime = false
				cOut <- msg
				break
			}
		}

		if i > sq_root {
			msg.isPrime = true
			cOut <- msg
		}

	}
}

func PrimeChannelTest() (chan primemsg, chan primemsg) {
	msg := primemsg{42, false}

	cIn := make(chan primemsg, NUMBERS)
	cOut := make(chan primemsg, NUMBERS)

	// Create workers
	for i := 0; i < WORKERS; i++ {
		go isPrime(cIn, cOut)
	}

	// Push numbfers to input queue
	for i := 0; i < NUMBERS; i++ {
		msg.num = rand.Intn(1000000) + 1000000
		cIn <- msg
	}

	// Read the answers
	for i := 0; i < NUMBERS; i++ {
		msg = <-cOut
		fmt.Println(msg.num, msg.isPrime)
	}

	return cIn, cOut
}
