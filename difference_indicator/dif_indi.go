package main

import "fmt"

func modifyByValue(val int) {

	val = 100
	fmt.Printf("Inside modifyByValue: val = %d \n", val)

}
func modifyByPointer(ptr *int) {

	*ptr = 100
	fmt.Printf("Inside modifyByPointer: ptr = %d \n", *ptr)

}

func main() {
	// Демонстрация с простым типом (int)
	value := 10
	fmt.Println("\n--- Работа с базовым типом ---")
	fmt.Printf("Before modifyByValue: value = %d\n", value)
	modifyByValue(value)
	fmt.Printf("After modifyByValue: value = %d\n", value)

	fmt.Printf("\nBefore modifyByPointer: value = %d\n", value)
	modifyByPointer(&value)
	fmt.Printf("After modifyByPointer: value = %d\n", value)
}
