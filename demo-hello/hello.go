package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	msg, err := greetings.Hello("Perry")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)

	names := []string{"Perry", "Cassie", "Tim"}

	msgs, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// jiodsjaodsaj jiosm djsio
	fmt.Print(msgs)
}
