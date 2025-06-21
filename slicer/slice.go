package main

import (
	"fmt"
)

func removeElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {

	var fruits []string
	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fruits = append(fruits, "cherry")

	moreFruits := []string{"peach", "avokado"}
	fruits = append(fruits, moreFruits...)

	fmt.Println("Все фрукты")
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i+1, fruit)
	}

	fmt.Printf("\nВсего фруктов: %d\n", len(fruits))

	fruits = removeElement(fruits, 2)
	fmt.Println("После удаления первого элемента:", fruits)

}
