package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	// Правило високосного года:
	// 1. Год кратен 400 -> високосный
	// 2. Иначе, если кратен 100 -> не високосный
	// 3. Иначе, если кратен 4 -> високосный
	// 4. Иначе -> не високосный
	return (year%400 == 0) || (year%100 != 0 && year%4 == 0)
}

func main() {
	var year int

	fmt.Print("Введите год для проверки: ")
	fmt.Scanln(&year)

	if isLeapYear(year) {
		fmt.Printf("%d год - високосный\n", year)
	} else {
		fmt.Printf("%d год - не високосный\n", year)
	}

}
