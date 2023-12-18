package main

import (
	"fmt"
	"regexp"
)

func CodelandUsernameValidation(str string) string {

	// code goes here
	regexCheck := func(word string) bool {
		re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
		return re.MatchString(word)
	}

	if len(str) >= 4 && len(str) <= 25 && regexCheck(str) {
		str = "true"
	} else {
		str = "false"
	}

	return str

}

func main() {
	fmt.Println("halo guys")
	fmt.Println(CodelandUsernameValidation("_a"))

}
