package main

import (
	"fmt"
	"math"
)

// calculateSum вычисляет сумму от i=1 до k: (-1)^i / (i-3)^2
// Пропускает слагаемые, равные нулю или бесконечности
func calculateSum(k int) float64 {
	var sum float64
	
	for i := 1; i <= k; i++ {
		// Проверяем, не равен ли знаменатель нулю (i-3)^2 = 0
		if i == 3 {
			fmt.Printf("Пропускаем слагаемое при i=%d (знаменатель равен нулю)\n", i)
			continue
		}
		
		// Вычисляем (-1)^i / (i-3)^2
		numerator := math.Pow(-1, float64(i))
		denominator := math.Pow(float64(i-3), 2)
		term := numerator / denominator
		
		// Проверяем на бесконечность и NaN
		if math.IsInf(term, 0) || math.IsNaN(term) {
			fmt.Printf("Пропускаем слагаемое при i=%d (результат: %f)\n", i, term)
			continue
		}
		
		sum += term
		fmt.Printf("i=%d: (-1)^%d / (%d-3)^2 = %f\n", i, i, i, term)
	}
	
	return sum
}

// calculateProduct вычисляет произведение от n=i до 2k: (n^3 - 8) / (n + 2)
// Пропускает множители, равные нулю или бесконечности
func calculateProduct(i, k int) float64 {
	product := 1.0
	
	for n := i; n <= 2*k; n++ {
		// Проверяем, не равен ли знаменатель нулю (n + 2 = 0)
		if n == -2 {
			fmt.Printf("Пропускаем множитель при n=%d (знаменатель равен нулю)\n", n)
			continue
		}
		
		// Вычисляем (n^3 - 8) / (n + 2)
		numerator := math.Pow(float64(n), 3) - 8
		denominator := float64(n) + 2
		term := numerator / denominator
		
		// Проверяем на бесконечность и NaN
		if math.IsInf(term, 0) || math.IsNaN(term) {
			fmt.Printf("Пропускаем множитель при n=%d (результат: %f)\n", n, term)
			continue
		}
		
		// Проверяем, не равен ли множитель нулю
		if term == 0 {
			fmt.Printf("Пропускаем множитель при n=%d (результат равен нулю)\n", n)
			continue
		}
		
		product *= term
		fmt.Printf("n=%d: (%d^3 - 8) / (%d + 2) = %f\n", n, n, n, term)
	}
	
	return product
}

// calculateW вычисляет W = Σ_{i=1}^{k} [(-1)^i / (i-3)^2] * Π_{n=i}^{2k} [(n^3 - 8) / (n + 2)]
func calculateW(k int) float64 {
	var result float64
	
	fmt.Printf("Вычисляем W для k = %d\n", k)
	fmt.Println("W = Σ_{i=1}^{k} [(-1)^i / (i-3)^2] * Π_{n=i}^{2k} [(n^3 - 8) / (n + 2)]")
	fmt.Println()
	
	for i := 1; i <= k; i++ {
		fmt.Printf("--- Вычисление для i = %d ---\n", i)
		
		// Вычисляем сумму для текущего i
		sumTerm := calculateSumTerm(i)
		if math.IsInf(sumTerm, 0) || math.IsNaN(sumTerm) || sumTerm == 0 {
			fmt.Printf("Пропускаем i=%d (сумма равна %f)\n", i, sumTerm)
			continue
		}
		
		// Вычисляем произведение для текущего i
		productTerm := calculateProduct(i, k)
		if math.IsInf(productTerm, 0) || math.IsNaN(productTerm) || productTerm == 0 {
			fmt.Printf("Пропускаем i=%d (произведение равно %f)\n", i, productTerm)
			continue
		}
		
		// Умножаем сумму на произведение
		term := sumTerm * productTerm
		result += term
		
		fmt.Printf("Слагаемое для i=%d: %f * %f = %f\n", i, sumTerm, productTerm, term)
		fmt.Printf("Промежуточная сумма: %f\n", result)
		fmt.Println()
	}
	
	return result
}

// calculateSumTerm вычисляет одно слагаемое суммы: (-1)^i / (i-3)^2
func calculateSumTerm(i int) float64 {
	// Проверяем, не равен ли знаменатель нулю
	if i == 3 {
		return 0 // Будет пропущено в основной функции
	}
	
	numerator := math.Pow(-1, float64(i))
	denominator := math.Pow(float64(i-3), 2)
	return numerator / denominator
}

func main() {
	var k int
	
	fmt.Println("Программа вычисления W = Σ_{i=1}^{k} [(-1)^i / (i-3)^2] * Π_{n=i}^{2k} [(n^3 - 8) / (n + 2)]")
	fmt.Print("Введите значение k: ")
	
	_, err := fmt.Scanf("%d", &k)
	if err != nil {
		fmt.Printf("Ошибка ввода: %v\n", err)
		return
	}
	
	if k <= 0 {
		fmt.Println("k должно быть положительным числом")
		return
	}
	
	fmt.Printf("\nНачинаем вычисления для k = %d\n", k)
	fmt.Println("=" * 50)
	
	result := calculateW(k)
	
	fmt.Println("=" * 50)
	fmt.Printf("Результат: W = %f\n", result)
}
