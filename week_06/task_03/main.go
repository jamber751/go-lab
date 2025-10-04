package main

import (
	"fmt"
	"math/rand"
	"time"
)

// countSevens подсчитывает количество вхождений числа 7 в двумерном массиве
func countSevens(matrix [][]int) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 7 {
				count++
			}
		}
	}
	return count
}

// createRandomMatrix создает случайную матрицу размером n x m
func createRandomMatrix(n, m int) [][]int {
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = rand.Intn(10) // числа от 0 до 9
		}
	}
	return matrix
}

// printMatrix выводит матрицу на экран
func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("=== Подсчет вхождений числа 7 в двумерном массиве ===")
	
	// Размеры матрицы
	n, m := 4, 5
	
	// Создаем случайную матрицу
	matrix := createRandomMatrix(n, m)
	
	fmt.Printf("Матрица %dx%d:\n", n, m)
	printMatrix(matrix)
	
	// Подсчитываем количество семерок
	sevensCount := countSevens(matrix)
	
	fmt.Printf("\nКоличество вхождений числа 7: %d\n", sevensCount)
	
	// Демонстрация с заранее известной матрицей
	fmt.Println("\n--- Демонстрация с тестовой матрицей ---")
	testMatrix := [][]int{
		{1, 7, 3, 4, 5},
		{7, 2, 7, 8, 9},
		{0, 1, 2, 7, 4},
		{5, 6, 7, 8, 9},
	}
	
	fmt.Println("Тестовая матрица:")
	printMatrix(testMatrix)
	
	testCount := countSevens(testMatrix)
	fmt.Printf("Количество вхождений числа 7 в тестовой матрице: %d\n", testCount)
}
