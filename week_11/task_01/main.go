package main

import "fmt"

func main() {
	fmt.Println("=== Калькулятор ===")
	fmt.Printf("2 + 3 = %d\n", Add(2, 3))
	fmt.Printf("4 * 5 = %d\n", Multiply(4, 5))

	if result, err := Divide(10, 2); err == nil {
		fmt.Printf("10 / 2 = %d\n", result)
	}
}

