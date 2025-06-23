package main

import (
	"fmt"
	"strings"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Course    int
	Mark      float64
}

// Конструктор для создания нового студента
func NewStudent(name string, birthYear int, course int, mark float64) Student {
	return Student{
		Name:      name,
		BirthYear: birthYear,
		Course:    course,
		Mark:      mark,
	}
}

// Метод для вычисления возраста студента
func (s Student) CalculateAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

// Метод для получения статуса студента
func (s Student) GetStatus() string {
	switch {
	case s.Mark >= 4.5:
		return "отличник"
	case s.Mark >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

// Метод для вывода информации о студенте
func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.CalculateAge())
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.Mark)
	fmt.Printf("Статус: %s\n", s.GetStatus())
}

// Метод для повышения курса
func (s *Student) Promote() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

// Метод для изменения оценки
func (s *Student) UpdateGrade(newGrade float64) {
	if newGrade >= 0 && newGrade <= 5 {
		s.Mark = newGrade
		fmt.Printf("Средний балл %s обновлен: %.2f\n", s.Name, s.Mark)
	} else {
		fmt.Println("Ошибка: средний балл должен быть между 0 и 5")
	}
}

// Метод для валидации данных студента
func (s Student) Validate() error {
	var errors []string

	if s.Name == "" {
		errors = append(errors, "имя не может быть пустым")
	}
	age := s.CalculateAge()
	if age < 16 || age > 100 {
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
	student := NewStudent("Иван Петров", 2005, 2, 4.3)

	// Проверяем валидность данных
	if err := student.Validate(); err != nil {
		fmt.Println(err)
		return
	}

	// Выводим информацию
	student.PrintInfo()

	// Повышаем курс
	student.Promote()

	// Обновляем средний балл и проверяем статус
	student.UpdateGrade(4.7)
	fmt.Printf("Новый статус: %s\n", student.GetStatus())

	// Пытаемся установить некорректный балл
	student.UpdateGrade(5.5)
}
