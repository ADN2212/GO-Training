package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	names := []string {"Nini", "Manetha", "Her"}
	messages, err := greetings.IterHello(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
