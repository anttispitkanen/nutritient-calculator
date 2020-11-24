package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello is a helloing function :DD
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns a randomly selected greeting message
func randomFormat() string {
	formats := []string{
		"Hi, welcome %v!",
		"Great to see you, %v!",
		"Jou jou mitä %v! :DD",
	}

	return formats[rand.Intn(len(formats))]
}
