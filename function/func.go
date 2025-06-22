package main

import (
	"fmt"
	"sort"
)

// Поиск элемента в срезе (возвращает индекс и наличие)
func FindElement(slice []int, target int) (int, bool) {
	for i, v := range slice {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

// Сортировка среза чисел (возрастание)
func SortSliceAscending(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

// Сортировка среза чисел (убывание)
func SortSliceDescending(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	return sorted
}

// Фильтрация среза (возвращает элементы, удовлетворяющие условию)
func FilterSlice(slice []int, condition func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}

	// Поиск элемента
	if idx, found := FindElement(numbers, 8); found {
		fmt.Printf("Элемент 8 найден по индексу %d\n", idx)
	} else {
		fmt.Println("Элемент не найден")
	}

	// Сортировка
	asc := SortSliceAscending(numbers)
	fmt.Println("Сортировка по возрастанию:", asc)

	desc := SortSliceDescending(numbers)
	fmt.Println("Сортировка по убыванию:", desc)

	// Фильтрация (четные числа)
	evenNumbers := FilterSlice(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Четные числа:", evenNumbers)
}
