// Package stringutils предоставляет утилиты для работы со строками.
//
// Пакет содержит функции для:
//   - Инверсии строк
//   - Проверки палиндромов
//   - Подсчёта символов
package stringutils

import "strings"

// Reverse инвертирует строку.
//
// Пример:
//
//	Reverse("hello") // возвращает "olleh"
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome проверяет, является ли строка палиндромом.
// Регистр игнорируется.
//
// Пример:
//
//	IsPalindrome("Анна") // возвращает true
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	return s == Reverse(s)
}

// CountChars возвращает карту с количеством каждого символа в строке.
func CountChars(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	return counts
}

