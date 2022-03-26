package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("John")
	fmt.Println(message)
}
