package main

import (
	"fmt"
)

func main() {
	var dayNumber int

	fmt.Print("Введите номер дня недели (1-7): ")
	_, err := fmt.Scanln(&dayNumber)
	if err != nil {
		fmt.Println("Ошибка ввода!")
		return
	}

	dayName := ""
	switch dayNumber {
	case 1:
		dayName = "Понедельник"
	case 2:
		dayName = "Вторник"
	case 3:
		dayName = "Среда"
	case 4:
		dayName = "Четверг"
	case 5:
		dayName = "Пятница"
	case 6:
		dayName = "Суббота"
	case 7:
		dayName = "Воскресенье"
	default:
		fmt.Println("Ошибка! Введите число от 1 до 7")
		return
	}

	fmt.Printf("День недели №%d - это %s\n", dayNumber, dayName)
}
