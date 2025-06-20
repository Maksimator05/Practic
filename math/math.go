package main

import (
	"fmt"
	"math"
)

func main() {

	const (
		pi = math.Pi
		e  = math.E
	)

	fmt.Println("Pi = ", pi)
	fmt.Printf("E = %v\n", e)
	fmt.Printf("Длина окружности с R=5:", 2*pi*5)
	fmt.Print("\n")
	fmt.Printf("Натуральный логарифм от 10: %.5f\n", math.Log(10))
}
