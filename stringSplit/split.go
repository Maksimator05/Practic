package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "MTUCI - the best university in IT"
	words := strings.Fields(str)
	for i, word := range words {

		fmt.Printf("%d: %s\n", i+1, word)

	}
}
