package main

func gotReduced(stack *[]string, checkpointBackwardsIndex int) bool {

	subStackString := SliceOfCharactersToString((*stack)[checkpointBackwardsIndex:])

	grammar := map[string]string{
		"ExprOpExpr": "Expr",
		"(Expr)":     "Expr",
		"-Expr":      "Expr",
		"num":        "Expr",
		"+":          "Op",
		"-":          "Op",
		"*":          "Op",
	}

	respectiveValue, belongsGrammar := grammar[subStackString]

	if belongsGrammar {
		PopStringsFromSlice(stack, checkpointBackwardsIndex)
		PushCharactersStringIntoSlice(stack, respectiveValue)
	}

	return belongsGrammar

}

func goBackwardsToReduce(stack *[]string) {

	// goes backward...
	for checkpointBackwards := len(*stack) - 1; checkpointBackwards >= 0; checkpointBackwards-- {
		// ...from the right side edge to an checkpointBackwards
		// it tries to reduce and if it got reduced it tries to reduce again
		if gotReduced(stack, checkpointBackwards) {
			checkpointBackwards = len(*stack) - 1
		}
	}

}

func Check(characters []string) []string {

	var stack []string

	for _, char := range characters {
		// shift
		stack = append(stack, char)
		goBackwardsToReduce(&stack)
	}

	return stack

}
