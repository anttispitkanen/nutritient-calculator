package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("Täältä pesee :DD -> ")
	log.SetFlags(0)

	names := []string{"Poetkoe :DD", "Poempylae :DD", "Poelloevaari :DDDD"}

	// message, err := greetings.Hello("Poetkoe :DD")

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
