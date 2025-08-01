package greetings

import (
	"errors"
	"fmt"
	"math/rand"

	//"golang.org/x/text/message"
)

//Si el nombre de esta funcion no empezara con una mayuscula, la funcion no fuera visible desde otros modulos.
func Hello(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("a name must be given")
	}
	
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprintln("Fafefifofu")
	return message, nil
}

func IterHello(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

var formats []string = []string {
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",} 

func randomFormat() string {
	return formats[rand.Intn(len(formats))]
}
