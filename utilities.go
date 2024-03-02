package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PopStringsFromSlice(slice *[]string, checkpointBackwards int) {
	// everything before the checkpointBackwards
	// [ everything : checkpoint ]
	*slice = (*slice)[:checkpointBackwards]
}

func PushCharactersStringIntoSlice(slice *[]string, str string) {
	// a string is not made of characteres, but bytes that represent characteres
	// a slice is a mutable-sized array, preferable for further use
	for index := 0; index < len(str); index++ {
		*slice = append(*slice, string(str[index]))
	}

}

func GetUserSingleCommand() string {

	fmt.Println("give me a command, mr. user:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	return line

}

func StringToSliceOfCharacters(str string) []string {

	var sliceOfCharacters []string
	for i := 0; i < len(str); i++ {
		sliceOfCharacters = append(sliceOfCharacters, string(str[i]))
	}
	return sliceOfCharacters

}

// the function converts slice to string
func SliceOfCharactersToString(slice []string) string {
	str := strings.Join(slice, "")
	return str
}
