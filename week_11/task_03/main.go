package main

import "fmt"

func main() {
	fmt.Println("=== Фибоначчи ===")
	n := 20
	fmt.Printf("Fib(%d) рекурсивно: %d\n", n, FibRecursive(n))
	fmt.Printf("Fib(%d) итеративно: %d\n", n, FibIterative(n))
}

