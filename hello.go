package main

import (
	"fmt"
	"log"

	"github.com/sxravan/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Shravan", "Sudhveer", "Rachna"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range message {
		fmt.Println(msg)
	}
}
