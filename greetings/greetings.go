package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello function to output a greeting
func Hello(name string) (string, error) {
	// If no name was given, return an error message
	if name == "" {
		return "", errors.New("empty name")
	}

	// return greeting that embeds a name in the greeting
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets inital values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	//slice of message formats
	formats := []string{
		"Hi, %v. welcome!",
		"Great to see you %v!",
		"Hail %v, well met",
	}

	// return a randomly selected message by selecting a
	// random index
	return formats[rand.Intn(len(formats))]
}
