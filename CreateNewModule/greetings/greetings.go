package greetings

import "fmt"
import "errors"

// Hello returns a greeting for the named person.
func Hello(name string) string{
	// If no name was given, return an error with a message.
	if name==""{
		return "", errors.New("empty Name!")
	}
	var message string
	// Return a greeting that embeds the name in a message.
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}