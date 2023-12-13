package main

import (
	"fmt"
	"strings"
)

type Rotor struct {
	wiring        string
	position      int
	notch         int
	turnover      bool
	turnoverCount int
}

func NewRotor(wiring string, notch int) *Rotor {
	return &Rotor{
		wiring:   wiring,
		position: 0,
		notch:    notch,
		turnover: false,
	}
}

func (r *Rotor) rotate() {
	r.position = (r.position + 1) % 26
	r.turnoverCount++
	if r.turnoverCount == r.notch {
		r.turnover = true
		r.turnoverCount = 0
	} else {
		r.turnover = false
	}
}

func (r *Rotor) encrypt(input rune) rune {
	offset := int('A' - rune(r.position))
	index := (int(input) - 'A' + offset) % 26
	encrypted := r.wiring[index]
	return rune((int(encrypted)-'A'-offset+26)%26 + 'A')
}

type EnigmaMachine struct {
	plugboard map[rune]rune
	rotor     *Rotor
}

func NewEnigmaMachine(plugboard map[rune]rune, rotor *Rotor) *EnigmaMachine {
	return &EnigmaMachine{
		plugboard: plugboard,
		rotor:     rotor,
	}
}

func (e *EnigmaMachine) encrypt(input string) string {
	encrypted := ""
	for _, char := range input {
		// Pass through plugboard
		if plug, ok := e.plugboard[char]; ok {
			char = plug
		}

		// Rotate the rotor before encryption
		e.rotor.rotate()
		// Pass through rotor
		char = e.rotor.encrypt(char)
		// Pass through rotor again (backwards)
		char = e.rotor.encrypt(char)
		// Pass through plugboard again
		if plug, ok := e.plugboard[char]; ok {
			char = plug
		}

		encrypted += string(char)
	}
	return encrypted
}

func main() {
	plugboard := map[rune]rune{
		'A': 'E',
		'B': 'J',
		'C': 'M',
		// Add more plugboard connections as needed
	}

	// Example rotor wiring (I rotor)
	rotor := NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", 16) // The notch is 'Q'

	machine := NewEnigmaMachine(plugboard, rotor)

	plaintext := "HELLO"
	encrypted := machine.encrypt(strings.ToUpper(plaintext))

	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Encrypted: %s\n", encrypted)

	machine.rotor.position = 0
	machine.encrypt(strings.ToUpper(encrypted))
	fmt.Printf("Plaintext: %s\n", plaintext)
}
