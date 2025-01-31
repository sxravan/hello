package main

import (
	"fmt"

	"github.com/sxravan/greetings"
)

func main() {
	message := greetings.Hello("Shravan")
	fmt.Println(message)
}
