package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	//"golang.org/x/text/message"
)

func main() {
    // Get a greeting message and print it.
    
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)
	// message, err := greetings.Hello("Adonis")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)


	names := []string {"Nini", "Manetha", "Her"}
	messages, err := greetings.IterHello(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
