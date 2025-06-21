package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "MTUCI - the best university in IT"

	charCount := utf8.RuneCountInString(str)
	fmt.Printf("1. Namber of chars: %d\n", charCount)

	substr := "best"

	index := strings.Index(str, substr)

	if index != -1 {
		fmt.Printf("2. Substring '%s' find on position %d\n", substr, index)
	} else {
		fmt.Printf("3. Substring '%s' not found\n", substr)
	}
	if strings.Contains(str, substr) {
		fmt.Printf("4. string include substring '%s'\n", substr)
	}

	lower := strings.ToLower(str)
	upper := strings.ToUpper(str)
	fmt.Printf("5. Lower of '%s' and upper of '%s'\n", lower, upper)

}
