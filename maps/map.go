package main

import (
	"fmt"
	"os"
)

func main() {

	markStudent := make(map[string]int)

	markStudent["Петр Корпоратив"] = 5
	markStudent["Саша Тюнинг"] = 4
	markStudent["Альберт Сутенер"] = 3
	markStudent["Горячая Чикса"] = 2
	markStudent["Костя Травматолог"] = 1

	for {
		fmt.Println("\nСистема учета оценок студентов")
		fmt.Println("1. Показать все оценки")
		fmt.Println("2. Добавить/изменить оценку")
		fmt.Println("3. Найти оценку студента")
		fmt.Println("4. Удалить студента")
		fmt.Println("5. Выйти")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scanln(&choice) //Чтение с консоли
		switch choice {
		case 1:
			fmt.Println("\nСписок студентов и их оценки:")
			for name, mark := range markStudent {
				fmt.Printf("%s: %d\n", name, mark)

			}
		case 2:
			var name string
			var mark int
			fmt.Print("Введите имя студента: ")
			fmt.Scanln(&name)
			fmt.Print("Введите оценку: ")
			fmt.Scanln(&mark)

			markStudent[name] = mark
			fmt.Printf("Оценка студента %s установлена: %d\n", name, mark)
		case 3:
			var name string
			fmt.Print("Введите имя студента для поиска:")
			fmt.Scanln(&name)
			if mark, exists := markStudent[name]; exists {
				fmt.Printf("Оценка студента %s: %d\n", name, mark)
			} else {
				fmt.Printf("Студент %s не найден\n", name)
			}
		case 4:
			var name string
			fmt.Print("Введите имя студента для удаления:")
			fmt.Scanln(&name)

			if _, exists := markStudent[name]; exists {
				delete(markStudent, name)
				fmt.Printf("Студент %s удален\n", name)
			} else {
				fmt.Printf("<Студент %s не найден\n", name)
			}
		case 5:
			fmt.Println("Выход из программы")
			os.Exit(0)
		default:

			fmt.Println("Неправильный выбор, попробуйте снова")
		}
	}

}
