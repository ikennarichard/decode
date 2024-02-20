package main

import (
	"fmt"
	"strings"
)

var morseCodes = map[string]string{
		".-":   "A",
		"-...": "B",
		"-.-.": "C",
		"-..":  "D",
		".":    "E",
		"..-.": "F",
		"--.":  "G",
		"....": "H",
		"..":   "I",
		".---": "J",
		"-.-":  "K",
		".-..": "L",
		"--": "M",
		"-.": "N",
		"---": "O",
		".--.": "P",
		"--.-": "Q",
		".-.": "R",
		"...": "S",
		"-": "T",
		"..-": "U",
		"...-": "V",
		".--": "W",
		"-..-": "X",
		"-.--": "Y",
		"--..": "Z",
		"-----":	"0",
		".----":	"1",
		"..---":	"2",
		"...--":	"3",
		"....-":	"4",
		".....":	"5",
		"-....":	"6",
		"--...":	"7",
		"---..":	"8",
		"----.":	"9",
}

func decodeChar(char string)string {
	return morseCodes[char]
}

func decodeWord(morseCode string)string{
	splits := strings.Fields(morseCode)
	var word string

	for _, code := range splits {
		word += morseCodes[code]
	}
	return word
}

func decodeMessage(morseCode string)string{
	splits := strings.Split(morseCode, "  ")
	var message string

	for _, code := range splits {
		message += decodeWord(code)+" "
	}
	return strings.TrimSpace(message)
}

func main() {
	fmt.Println(decodeChar(".-"))
	fmt.Println(decodeWord("--  -.--"))
	fmt.Println(decodeMessage(".-   -... --- -..-   ..-. ..- .-.. .-..   --- ..-.   .-. ..- -... .. . ..."))
}