package main

import "fmt"

func main() {
	fmt.Println("=== Проверка простых чисел ===")
	for i := 1; i <= 20; i++ {
		if IsPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

