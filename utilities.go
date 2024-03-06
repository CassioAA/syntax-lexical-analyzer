package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// a very simple way to check what's happening
func Test(message string, optionalValue *string) {

	hasOptionalValue := optionalValue != nil

	if hasOptionalValue {
		fmt.Println(message, " at ", *optionalValue, "° time")
	} else {
		fmt.Println(message)
	}

}

func getUserLine() string {

	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')

	gotErrorReading := err != nil

	if gotErrorReading {
		fmt.Println("Error reading input:", err)
		return ""
	}

	// it comes with these escape sequences
	line = strings.TrimSuffix(line, "\r\n")

	return line

}

func GetUserLines() []string {
	
	var lines []string

	fmt.Println("give me commands, mr. user:")

	for line := getUserLine(); line != "exit"; line = getUserLine() {
		lines = append(lines, line)
	}

	return lines

}

func GetUserSingleLine() string {

	line := getUserLine()

	return line

}

func StringToSliceOfCharacters(str string) []string {

	var sliceOfCharacters []string

	for i := 0; i < len(str); i++ {

		characterString := string(str[i])
		sliceOfCharacters = append(sliceOfCharacters, characterString)

	}

	return sliceOfCharacters

}

// the function converts slice to string
func SliceOfCharactersToString(slice []string) string {
	str := strings.Join(slice, "")
	return str
}
