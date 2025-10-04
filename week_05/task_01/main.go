package main

import (
	"fmt"
	"math"
)

// isPrime проверяет, является ли число простым
// Использует вложенные циклы для проверки делителей
func isPrime(n int) bool {
	// 1 не является простым числом
	if n <= 1 {
		return false
	}
	
	// 2 - единственное четное простое число
	if n == 2 {
		return true
	}
	
	// Проверяем четность
	if n%2 == 0 {
		return false
	}
	
	// Проверяем делители от 3 до sqrt(n) с шагом 2
	// Внешний цикл для перебора возможных делителей
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		// Внутренняя проверка делимости
		if n%i == 0 {
			return false
		}
	}
	
	return true
}

// findFirstNPrimes находит первые n простых чисел
// Использует вложенные циклы для поиска простых чисел
func findFirstNPrimes(n int) []int {
	primes := make([]int, 0, n)
	num := 2 // Начинаем с первого простого числа
	
	// Внешний цикл продолжается, пока не найдем n простых чисел
	for len(primes) < n {
		// Внутренняя проверка на простоту
		if isPrime(num) {
			primes = append(primes, num)
		}
		num++
	}
	
	return primes
}

// printPrimesInColumns выводит простые числа в колонках
// Демонстрирует использование вложенных циклов для форматирования
func printPrimesInColumns(primes []int, columns int) {
	fmt.Printf("Первые %d простых чисел (в %d колонках):\n\n", len(primes), columns)
	
	// Внешний цикл для строк
	for i := 0; i < len(primes); i += columns {
		// Внутренний цикл для колонок в текущей строке
		for j := 0; j < columns && i+j < len(primes); j++ {
			fmt.Printf("%4d ", primes[i+j])
		}
		fmt.Println() // Переход на новую строку
	}
}

func main() {
	fmt.Println("=== Программирование алгоритма с вложенными циклами ===")
	fmt.Println("Задача: Найти первые 100 простых чисел")
	fmt.Println()
	
	// Находим первые 100 простых чисел
	primes := findFirstNPrimes(100)
	
	// Выводим результат в 10 колонок для удобства чтения
	printPrimesInColumns(primes, 10)
	
	fmt.Println()
	fmt.Printf("Всего найдено: %d простых чисел\n", len(primes))
	fmt.Printf("Последнее простое число: %d\n", primes[len(primes)-1])
	
	// Дополнительная статистика
	fmt.Println("\n=== Статистика ===")
	fmt.Printf("Первые 10 простых чисел: %v\n", primes[:10])
	fmt.Printf("Последние 10 простых чисел: %v\n", primes[90:])
	
	// Проверяем корректность алгоритма
	fmt.Println("\n=== Проверка корректности ===")
	testNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	fmt.Println("Проверка известных простых чисел:")
	for _, num := range testNumbers {
		if isPrime(num) {
			fmt.Printf("✓ %d - простое\n", num)
		} else {
			fmt.Printf("✗ %d - НЕ простое (ошибка в алгоритме!)\n", num)
		}
	}
	
	// Проверяем составные числа
	compositeNumbers := []int{4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}
	fmt.Println("\nПроверка составных чисел:")
	for _, num := range compositeNumbers {
		if !isPrime(num) {
			fmt.Printf("✓ %d - составное\n", num)
		} else {
			fmt.Printf("✗ %d - простое (ошибка в алгоритме!)\n", num)
		}
	}
}
