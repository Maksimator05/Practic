package main

import (
	"fmt"
	"os"
)

func main() {

	var a, b float64
	var operator string
	fmt.Println("Простой калькулятор")
	fmt.Println("Доступные операции: +, -, *, /")
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	var result float64
	var err error = nil

	switch operator {
	case "+":

		result = a + b

	case "-":
		result = a - b

	case "*":
		result = a * b

	case "/":
		if b == 0 {
			err = fmt.Errorf("Ошибка: деление на ноль невозможно")
		} else {
			result = a / b
		}
	default:
		err = fmt.Errorf("Ошибка: неизвестный оператор")
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n\n", a, operator, b, result)
}
