package main

import (
	"fmt"
	"math"
	"sort"
)

// Функция для ввода массива из 12 элементов
func inputArray() []float64 {
	array := make([]float64, 12)
	
	fmt.Println("Введите 12 элементов массива (действительные числа):")
	for i := 0; i < 12; i++ {
		fmt.Printf("Элемент %d: ", i+1)
		_, err := fmt.Scanf("%f", &array[i])
		if err != nil {
			fmt.Printf("Ошибка ввода для элемента %d. Попробуйте еще раз: ", i+1)
			i-- // Повторяем ввод для этого элемента
			continue
		}
	}
	
	return array
}

// Функция для вывода массива
func printArray(array []float64, title string) {
	fmt.Printf("\n%s: ", title)
	for i, val := range array {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%.2f", val)
	}
	fmt.Println()
}

// Функция для сортировки массива по убыванию
func sortDescending(array []float64) []float64 {
	// Создаем копию массива для сортировки
	sortedArray := make([]float64, len(array))
	copy(sortedArray, array)
	
	// Сортируем по убыванию
	sort.Slice(sortedArray, func(i, j int) bool {
		return sortedArray[i] > sortedArray[j]
	})
	
	return sortedArray
}

// Функция для нахождения максимального и минимального элементов
func findMaxMin(array []float64) (float64, float64) {
	if len(array) == 0 {
		return 0, 0
	}
	
	max := array[0]
	min := array[0]
	
	for _, val := range array {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	
	return max, min
}

func main() {
	fmt.Println("=== Программа работы с одномерным массивом ===")
	fmt.Println("Задание: Ввести массив из 12 элементов, отсортировать по убыванию,")
	fmt.Println("найти сумму максимального и минимального элементов")
	fmt.Println()
	
	// Ввод массива
	array := inputArray()
	
	// Вывод исходного массива
	printArray(array, "Исходный массив")
	
	// Сортировка по убыванию
	sortedArray := sortDescending(array)
	printArray(sortedArray, "Массив после сортировки по убыванию")
	
	// Поиск максимального и минимального элементов
	max, min := findMaxMin(array)
	
	// Вычисление суммы максимального и минимального элементов
	sum := max + min
	
	// Вывод результатов
	fmt.Printf("\nРезультаты:\n")
	fmt.Printf("Максимальный элемент: %.2f\n", max)
	fmt.Printf("Минимальный элемент: %.2f\n", min)
	fmt.Printf("Сумма максимального и минимального элементов: %.2f\n", sum)
	
	// Дополнительная информация
	fmt.Printf("\nДополнительная информация:\n")
	fmt.Printf("Количество элементов в массиве: %d\n", len(array))
	fmt.Printf("Среднее арифметическое: %.2f\n", calculateAverage(array))
}

// Функция для вычисления среднего арифметического (дополнительно)
func calculateAverage(array []float64) float64 {
	if len(array) == 0 {
		return 0
	}
	
	sum := 0.0
	for _, val := range array {
		sum += val
	}
	
	return sum / float64(len(array))
}
