package main

import (
	"fmt"

	"github.com/greetings"
)

func main() {
	message := greetings.Hello("Shravan")
	fmt.Println(message)
}
