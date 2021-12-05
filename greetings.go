package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if name is empty, return an error
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formatOptions := []string{
		"Hello, %s. Nice to meet you.",
		"Welcome, %s!",
		"%s, it's an honor to finally make your acquaintance.",
	}

	r := rand.Intn(len(formatOptions))

	return formatOptions[r]
}
