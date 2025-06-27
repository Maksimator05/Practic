package main

import (
	"fmt"
)

// Функция удаления элемента из среза по индексу

func removeElement(slice []string, index int) []string {

	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// Создаем пустой срез строк для хранения фруктов
	var fruits []string

	// Добавляем элементы в срез по одному
	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fruits = append(fruits, "cherry")

	// Создаем дополнительный срез с фруктами
	moreFruits := []string{"peach", "avokado"}
	// Добавляем все элементы из moreFruits в fruits
	fruits = append(fruits, moreFruits...)

	fmt.Println("Все фрукты")
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i+1, fruit)
	}

	fmt.Printf("\nВсего фруктов: %d\n", len(fruits))

	// Удаляем элемент с индексом 2 (третий элемент, так как индексация с 0)
	fruits = removeElement(fruits, 2)
	fmt.Println("После удаления первого элемента:", fruits)
}
