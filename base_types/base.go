package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// 1. Логический тип (bool)
	var isActive bool = true
	isAdmin := false
	fmt.Println("Логические типы:")
	fmt.Println("isActive:", isActive, "isAdmin:", isAdmin)
	fmt.Println()

	// 2. Строки (string)
	var name string = "Иван"
	email := "ivan@example.com"
	multiline := `Многострочная строка`
	fmt.Println("Строки:")
	fmt.Println("Имя:", name)
	fmt.Println("Email:", email)
	fmt.Println("Текст:", multiline)
	fmt.Println()

	// 3. Целочисленные типы
	var age int = 30           // int (зависит от платформы)
	var count int32 = 1000     // 32-битное целое
	var bigNum int64 = 1 << 40 // 64-битное целое
	year := 2023
	fmt.Println("Целочисленные типы:")
	fmt.Println("age:", age, "count:", count, "bigNum:", bigNum, "year:", year)
	fmt.Println()

	// 4. Числа с плавающей точкой
	var price float32 = 19.99 // 32-битное
	pi := 3.141592653589793   // float64 (по умолчанию)
	fmt.Println("Числа с плавающей точкой:")
	fmt.Println("price:", price, "pi:", pi)
	fmt.Println()

	// 5. Комплексные числа
	var z complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Println("Комплексные числа:")
	fmt.Println("z:", z, "Реальная часть:", real(z), "Мнимая часть:", imag(z))
	fmt.Println()

	// 6. Байты (byte)
	var b byte = 65 // alias для uint8
	fmt.Println("Байты:")
	fmt.Println("byte:", b, "как символ:", string(b))

}
