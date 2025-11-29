package main

import (
	"fmt"
	"week_10/task_01/mathutils"
)

func main() {
	fmt.Println("=== Использование пакета mathutils ===")

	// Факториал
	fmt.Printf("5! = %d\n", mathutils.Factorial(5))

	// Проверка на простое число
	for _, n := range []int{2, 7, 10, 17} {
		fmt.Printf("%d простое: %v\n", n, mathutils.IsPrime(n))
	}

	// НОД
	fmt.Printf("НОД(48, 18) = %d\n", mathutils.GCD(48, 18))

	// Фибоначчи
	fmt.Print("Фибоначчи: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", mathutils.Fibonacci(i))
	}
	fmt.Println()
}

