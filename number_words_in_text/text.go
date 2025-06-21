package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Привет мир привет страна привет город"

	// Разбиваем текст на слова
	words := strings.Split(text, " ")

	// Создаём карту для подсчёта
	count := make(map[string]int)

	// Считаем слова
	for _, word := range words {
		count[word]++
	}

	// Выводим результат
	fmt.Println("Результат подсчёта:")
	for word, num := range count {
		fmt.Printf("%s: %d\n", word, num)
	}
}
