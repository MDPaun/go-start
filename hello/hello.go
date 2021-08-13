package main

import (
	"fmt"

	"github.com/MDPaun/go-start/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Vasile")
    fmt.Println(message)
}