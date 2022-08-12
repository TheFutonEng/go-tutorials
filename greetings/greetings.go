package greetings

import "fmt"

// Hello function to output a greeting
func Hello(name string) string {
	// return greeting that embeds a name in the greeting
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
