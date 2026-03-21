package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	name := []string{"Kim", "Winter", "RR"}
	message, err := greetings.Hellos(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	fmt.Println(quote.Go())
}
