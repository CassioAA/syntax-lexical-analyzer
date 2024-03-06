package main

// import "fmt"

	/*
		line := GetUserSingleLine()
		userLineCharacters := StringToSliceOfCharacters(line)

		CheckSyntax(userLineCharacters)
	*/

	/*
		correctSyntax := CheckSyntax(userLineCharacters)

		if correctSyntax {
			// I got a warning writing fmt.Println("\nAccepted!\n")
			fmt.Println("\nAccepted!")
			fmt.Println()
		} else {
			fmt.Println("\nRejected!")
			fmt.Println()
		}
	*/
	
func main() {

	lines := GetUserLines()

	for _, line := range lines {

		userLineCharacters := StringToSliceOfCharacters(line)

		CheckSyntax(userLineCharacters)

	}

}