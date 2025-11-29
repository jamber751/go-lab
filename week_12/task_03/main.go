package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("=== Работа с CSV ===")

	// Создаём CSV файл
	file, _ := os.Create("data.csv")
	writer := csv.NewWriter(file)
	writer.Write([]string{"Имя", "Возраст", "Город"})
	writer.Write([]string{"Иван", "25", "Москва"})
	writer.Write([]string{"Мария", "30", "Санкт-Петербург"})
	writer.Write([]string{"Пётр", "22", "Казань"})
	writer.Flush()
	file.Close()

	// Читаем и обрабатываем CSV
	file, _ = os.Open("data.csv")
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	file.Close()

	fmt.Println("Данные:")
	totalAge := 0
	count := 0
	for i, record := range records {
		if i == 0 { // Пропускаем заголовок
			continue
		}
		age, _ := strconv.Atoi(record[1])
		totalAge += age
		count++
		fmt.Printf("  %s, %d лет, %s\n", record[0], age, record[2])
	}

	fmt.Printf("\nСредний возраст: %.1f лет\n", float64(totalAge)/float64(count))
}

