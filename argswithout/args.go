package main

import (
	"fmt"
)

var say bool
var age, name = 10, "Petr"

func main() {
	say = true
	fmt.Printf("Age: %d\tName: %s\tSay: %t\n", age, name, say)
}
