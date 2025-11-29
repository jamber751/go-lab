package main

// FibRecursive - рекурсивное вычисление Фибоначчи
func FibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

// FibIterative - итеративное вычисление Фибоначчи
func FibIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

