package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Maksim Shiyan"
	Time := time.Now()

	fmt.Printf("Hello %s\n", name)
	fmt.Println("Time is", Time)
}
