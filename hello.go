package main

import (
	"fmt"
	"log"

	"github.com/sxravan/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Shravan")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
