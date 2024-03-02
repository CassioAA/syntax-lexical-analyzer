package main

import "fmt"

func main() {

	command := GetUserSingleCommand()
	commandCharacters := StringToSliceOfCharacters(command)
	result := Check(commandCharacters)
	fmt.Println(result)

}
