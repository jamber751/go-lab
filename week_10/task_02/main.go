package main

import (
	"fmt"
	"week_10/task_02/stringutils"
)

func main() {
	fmt.Println("=== Пакет stringutils ===")

	// Инверсия строки
	s := "Привет"
	fmt.Printf("Инверсия '%s': '%s'\n", s, stringutils.Reverse(s))

	// Проверка палиндрома
	words := []string{"казак", "привет", "шалаш", "Go"}
	for _, w := range words {
		fmt.Printf("'%s' палиндром: %v\n", w, stringutils.IsPalindrome(w))
	}

	// Подсчёт символов
	text := "hello"
	counts := stringutils.CountChars(text)
	fmt.Printf("Символы в '%s': %v\n", text, counts)
}

