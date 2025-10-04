package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для ввода массива с клавиатуры
func inputArray() []float64 {
	array := make([]float64, 12)
	fmt.Println("Введите 12 элементов массива (вещественные числа):")
	
	for i := 0; i < 12; i++ {
		fmt.Printf("Элемент %d: ", i+1)
		_, err := fmt.Scanf("%f", &array[i])
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте еще раз.")
			i-- // Повторяем ввод для этого элемента
		}
	}
	
	return array
}

// Функция для генерации случайного массива (для демонстрации)
func generateRandomArray() []float64 {
	array := make([]float64, 12)
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < 12; i++ {
		// Генерируем числа от -100 до 100 с двумя знаками после запятой
		array[i] = float64(rand.Intn(20000)-10000) / 100.0
	}
	
	return array
}

// Функция сортировки массива по убыванию с подсчетом перестановок
// Использует алгоритм пузырьковой сортировки
func bubbleSortDescending(array []float64) ([]float64, int) {
	n := len(array)
	swaps := 0
	arr := make([]float64, n)
	copy(arr, array) // Создаем копию, чтобы не изменять исходный массив
	
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] { // Сравниваем для сортировки по убыванию
				// Меняем местами элементы
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaps++
				swapped = true
			}
		}
		// Если за проход не было перестановок, массив отсортирован
		if !swapped {
			break
		}
	}
	
	return arr, swaps
}

// Функция для вывода массива
func printArray(array []float64, title string) {
	fmt.Printf("\n%s:\n", title)
	for i, value := range array {
		fmt.Printf("arr[%2d] = %8.2f\n", i, value)
	}
}

// Функция для выбора режима ввода
func chooseInputMode() int {
	var mode int
	fmt.Println("Выберите режим ввода:")
	fmt.Println("1 - Ввод с клавиатуры")
	fmt.Println("2 - Случайная генерация")
	fmt.Print("Ваш выбор: ")
	fmt.Scanf("%d", &mode)
	return mode
}

func main() {
	fmt.Println("=== Сортировка массива по убыванию с подсчетом перестановок ===")
	fmt.Println("Программа сортирует массив из 12 элементов по убыванию")
	fmt.Println("и подсчитывает количество перестановок при сортировке.\n")
	
	var array []float64
	mode := chooseInputMode()
	
	switch mode {
	case 1:
		array = inputArray()
	case 2:
		array = generateRandomArray()
		fmt.Println("Сгенерирован случайный массив:")
	default:
		fmt.Println("Неверный выбор. Используется случайная генерация.")
		array = generateRandomArray()
		fmt.Println("Сгенерирован случайный массив:")
	}
	
	// Выводим исходный массив
	printArray(array, "Исходный массив")
	
	// Сортируем массив и подсчитываем перестановки
	sortedArray, swaps := bubbleSortDescending(array)
	
	// Выводим отсортированный массив
	printArray(sortedArray, "Отсортированный массив (по убыванию)")
	
	// Выводим количество перестановок
	fmt.Printf("\nКоличество перестановок: %d\n", swaps)
	
	// Дополнительная информация
	fmt.Printf("\nАнализ сортировки:\n")
	fmt.Printf("- Размер массива: %d элементов\n", len(array))
	fmt.Printf("- Количество перестановок: %d\n", swaps)
	fmt.Printf("- Эффективность: %.2f%% (перестановок от максимально возможных)\n", 
		float64(swaps)/float64(len(array)*(len(array)-1)/2)*100)
}
