package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Работа с файлами ===")

	// Создаём тестовый файл
	content := `Это первая строка текста.
Это вторая строка с несколькими словами.
А это третья строка.`

	err := os.WriteFile("input.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}

	// Читаем и анализируем файл
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка открытия:", err)
		return
	}
	defer file.Close()

	var lines, words, chars int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
		chars += len(line)
	}

	// Выводим и сохраняем результат
	result := fmt.Sprintf("Строк: %d\nСлов: %d\nСимволов: %d\n", lines, words, chars)
	fmt.Print(result)

	os.WriteFile("stats.txt", []byte(result), 0644)
	fmt.Println("Статистика сохранена в stats.txt")
}

