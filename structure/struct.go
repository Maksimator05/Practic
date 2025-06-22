package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name   string
	Age    int
	Course int
	Mark   float64
}

// Определение структуры Student
func NewStudent(name string, age int, course int, mark float64) Student {
	return Student{name, age, course, mark}
}

// Функция вывода информации о студенте
func (s Student) PrintInfo() {
	fmt.Printf("Student Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Course: %d\n", s.Course)
	fmt.Printf("Mark: %.2f\n", s.Mark)
}

// Повышение курса
func (s *Student) Promote() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

// Функция проверки, является ли студент отличником
func (s Student) IsExcellent() bool {
	return s.Mark >= 4.5
}

// Функция изменения среднего балла
func (s *Student) UpdateGrade(newGrade float64) {
	if newGrade >= 0 && newGrade <= 5 {
		s.Mark = newGrade
		fmt.Printf("Средний балл %s обновлен: %.2f\n", s.Name, s.Mark)
	} else {
		fmt.Println("Ошибка: средний балл должен быть между 0 и 5")
	}
}

// Функция проверки корректности данных студента
func (s Student) Validate() error {
	var errors []string

	if s.Name == "" {
		errors = append(errors, "имя не может быть пустым")
	}
	if s.Age < 16 || s.Age > 100 {
		errors = append(errors, "возраст должен быть между 16 и 100")
	}
	if s.Course < 1 || s.Course > 6 {
		errors = append(errors, "курс должен быть между 1 и 6")
	}
	if s.Mark < 0 || s.Mark > 5 {
		errors = append(errors, "средний балл должен быть между 0 и 5")
	}

	if len(errors) > 0 {
		return fmt.Errorf("ошибки валидации: %s", strings.Join(errors, ", "))
	}
	return nil
}

func main() {
	// Создаем нового студента
	student := NewStudent("Иван Петров", 20, 2, 4.3)

	// Проверяем валидность данных
	if err := student.Validate(); err != nil {
		fmt.Println(err)
		return
	}

	// Выводим информацию
	student.PrintInfo()

	// Проверяем, отличник ли
	if student.IsExcellent() {
		fmt.Println("Этот студент - отличник!")
	} else {
		fmt.Println("Этот студент - не отличник")
	}

	// Повышаем курс
	student.Promote()

	// Обновляем средний балл
	student.UpdateGrade(4.7)
	student.PrintInfo()

	// Пытаемся установить некорректный балл
	student.UpdateGrade(5.5)
}
