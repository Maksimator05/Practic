package main

import (
	"fmt"
)

func main() {
	fmt.Println("Таблица умножения от 1 до 10:")
	fmt.Println("----------------------------")

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x  %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}
