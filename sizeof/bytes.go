package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 1. Логический тип
	var b bool
	fmt.Printf("bool: %d байт\n", unsafe.Sizeof(b))

	// 2. Целочисленные типы
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	fmt.Printf("int8: %d байт\n", unsafe.Sizeof(i8))
	fmt.Printf("int16: %d байт\n", unsafe.Sizeof(i16))
	fmt.Printf("int32: %d байт\n", unsafe.Sizeof(i32))
	fmt.Printf("int64: %d байт\n", unsafe.Sizeof(i64))

	// 3. Числа с плавающей точкой
	var f32 float32
	var f64 float64
	fmt.Printf("float32: %d байт\n", unsafe.Sizeof(f32))
	fmt.Printf("float64: %d байт\n", unsafe.Sizeof(f64))

	// 4. Комплексные числа
	var c64 complex64
	var c128 complex128
	fmt.Printf("complex64: %d байт\n", unsafe.Sizeof(c64))
	fmt.Printf("complex128: %d байт\n", unsafe.Sizeof(c128))

	// 5. Строки
	var s string
	fmt.Printf("string: %d байт (базовая структура)\n", unsafe.Sizeof(s))

}
