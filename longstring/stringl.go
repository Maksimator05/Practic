package main

import "fmt"

// Функция для поиска самой длинной строки
func FindLongestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]

	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}

	return longest
}

func main() {
	// Примеры использования
	fmt.Println(FindLongestString("cat", "dog", "elephant"))
	fmt.Println(FindLongestString("apple", "banana", "pear"))
	fmt.Println(FindLongestString("one"))
	fmt.Println(FindLongestString("gg", "short", "very long"))
}
