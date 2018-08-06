package main

import (
	"fmt"
)

type Maybe struct {
	wordEntry string
}

// allows "safe" access to maybe value
func (m *Maybe) getOrElse(defaultValue string) string {
	if len(m.wordEntry) < 1 {
		m.setValue(defaultValue)
		return "set value"
	}
	return m.wordEntry
}

// Mutates maybeValue through passing by refernce 
func (m *Maybe) setValue(defaultValue string) {
	m.wordEntry = defaultValue 
}

var maybeValue = &Maybe{}

func main () {

	pointers()
	maps()

	inCorrectPasscode, err := passCode("derp")
	if err != nil {
		fmt.Println("Error occured")
	} else {
		fmt.Println(inCorrectPasscode)
	}

	correctPasscode, err := passCode("passcode")
	if err != nil {
		fmt.Println("Error occured")
	} else {
		fmt.Println(correctPasscode)
	}

	fmt.Println(maybeValue.getOrElse("GNAR"))
	fmt.Println(maybeValue)
}

/**	
 * Returns a string | error
 */
func passCode (word string) (string, error) {
	if word != "passcode" {
		return "", fmt.Errorf("not the right string")
	}
	return word, nil
}

func pointers () {
	var value = 12
	var valuePointer = &value

	// mutates the value in mem at pointer address
	*valuePointer += 5

	fmt.Println(value) // 17
}

func maps () {
	mymap := map[string]string{"name": "titan", "b": "banana"}
	fmt.Println(mymap)

	for k, v := range mymap {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

