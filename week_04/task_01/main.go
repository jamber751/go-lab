package main

import (
	"fmt"
	"math"
)

// calculateSeries вычисляет значение ряда X - X³/3 + X⁵/5 + ... + (-1)ᴺX^(2N+1)/(2N+1)
// Использует цикл for для суммирования членов ряда
func calculateSeries(x float64, n int) float64 {
	result := 0.0
	
	// Цикл for для вычисления суммы ряда
	for k := 0; k <= n; k++ {
		// Вычисляем знак: (-1)^k
		sign := math.Pow(-1, float64(k))
		
		// Вычисляем степень: X^(2k+1)
		power := math.Pow(x, float64(2*k+1))
		
		// Вычисляем знаменатель: (2k+1)
		denominator := float64(2*k + 1)
		
		// Добавляем текущий член ряда к результату
		term := sign * power / denominator
		result += term
		
		// Выводим информацию о текущем члене для наглядности
		fmt.Printf("k=%d: %.6f * %.6f / %.0f = %.6f\n", 
			k, sign, power, denominator, term)
	}
	
	return result
}

func main() {
	var x float64
	var n int
	
	fmt.Println("=== Вычисление математического ряда ===")
	fmt.Println("Ряд: X - X³/3 + X⁵/5 + ... + (-1)ᴺX^(2N+1)/(2N+1)")
	fmt.Println()
	
	// Ввод данных от пользователя
	fmt.Print("Введите вещественное число X: ")
	fmt.Scan(&x)
	
	fmt.Print("Введите целое число N (> 0): ")
	fmt.Scan(&n)
	
	// Проверка корректности ввода
	if n <= 0 {
		fmt.Println("Ошибка: N должно быть положительным числом!")
		return
	}
	
	fmt.Printf("\nВычисление ряда для X = %.6f, N = %d:\n", x, n)
	fmt.Println("Члены ряда:")
	
	// Вычисление и вывод результата
	result := calculateSeries(x, n)
	
	fmt.Printf("\nРезультат: %.10f\n", result)
	
	// Дополнительная информация
	fmt.Printf("\nДля сравнения: arctan(%.6f) ≈ %.10f\n", x, math.Atan(x))
	fmt.Printf("Разность: %.10f\n", math.Abs(result - math.Atan(x)))
}
