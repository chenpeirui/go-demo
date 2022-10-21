package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}

	msg := fmt.Sprintf(randomFormat(), name)
	// msg := fmt.Sprintf(randomFormat())
	return msg, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with meesages.
	msgs := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each time.
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved massage with
		// the name.
		msgs[name] = msg
	}

	return msgs, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
