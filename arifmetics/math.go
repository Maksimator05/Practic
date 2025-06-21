package main

import (
	"fmt"
)

func main() {

	var a, b float64

	fmt.Println("Демонстрация арифметики")
	fmt.Println("Введите число 1")
	fmt.Scanln(&a)
	fmt.Println("Введите число 2")
	fmt.Scanln(&b)

	fmt.Println("Основные операции:")
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f * %.2f = %.2f = %.2f\n", a, b, a*b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	if b != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Printf("%.2f / %.2f = ошибка (деление на ноль)\n", a, b)
	}
	fmt.Printf("\nОперации сравнения:\n")
	fmt.Printf("%.2f == %.2f: %t\n", a, b, a == b)
	fmt.Printf("%.2f != %.2f: %t\n", a, b, a != b)
	fmt.Printf("%.2f > %.2f: %t\n", a, b, a > b)
	fmt.Printf("%.2f >= %.2f: %t\n", a, b, a >= b)
	fmt.Printf("%.2f < %.2f: %t\n", a, b, a < b)
	fmt.Printf("%.2f <= %.2f: %t\n", a, b, a <= b)

	// Инкремент и декремент
	fmt.Printf("\nОперации инкремента и декремента:\n")
	num := a
	fmt.Printf("Исходное значение num: %.2f\n", num)
	num++
	fmt.Printf("После num++: %.2f\n", num)
	num--
	fmt.Printf("После num--: %.2f\n", num)
}
