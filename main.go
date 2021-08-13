package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
)


func main() {
	fmt.Println("Hello world!")
	fmt.Println(quote.Go())
	fmt.Println("The time is", time.Now())
	Hello("Vasile")
	fmt.Println("Bye world!")
}


