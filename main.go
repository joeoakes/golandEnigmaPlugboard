package main

import (
	"fmt"
	"strings"
)

type Rotor struct {
	wiring   string
	position int
}

func (r *Rotor) encrypt(input rune) rune {
	offset := 'A' - rune(r.position)
	index := (int(input) - 'A' + int(offset)) % 26
	encrypted := r.wiring[index]
	return rune((int(encrypted)-'A'-int(offset)+26)%26 + 'A')
}

type EnigmaMachine struct {
	plugboard map[rune]rune
	rotors    []*Rotor
}

func NewRotor(wiring string) *Rotor {
	return &Rotor{
		wiring:   wiring,
		position: 0,
	}
}

func (e *EnigmaMachine) SetRotorPositions(positions []int) {
	for i, rotor := range e.rotors {
		rotor.position = positions[i]
	}
}

func (e *EnigmaMachine) encrypt(input string) string {
	encrypted := ""
	for _, char := range input {
		// Pass through plugboard
		if plug, ok := e.plugboard[char]; ok {
			char = plug
		}

		// Pass through rotors
		for _, rotor := range e.rotors {
			char = rotor.encrypt(char)
		}

		// Pass through reflector (not implemented in this simple example)
		// In a real Enigma machine, there is a reflector that sends the signal back through the rotors.

		// Reverse pass through rotors
		for i := len(e.rotors) - 1; i >= 0; i-- {
			char = e.rotors[i].encrypt(char)
		}

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

	rotor1 := NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ")
	rotor2 := NewRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE")
	rotor3 := NewRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO")

	machine := &EnigmaMachine{
		plugboard: plugboard,
		rotors:    []*Rotor{rotor1, rotor2, rotor3},
	}

	plaintext := "HELLO" //H=72, E=69, L=76, O=79
	encrypted := machine.encrypt(strings.ToUpper(plaintext))

	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Encrypted: %s\n", encrypted)

	// Reset rotor positions
	machine.SetRotorPositions([]int{0, 0, 0})

	decrypted := machine.encrypt(encrypted)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
