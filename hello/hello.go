package main

import (
	"example.com/greetings"
	"fmt"
	"rsc.io/quote"
)

func main() {
	greetings.Hello("Gladys")
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
