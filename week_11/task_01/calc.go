package main

import "errors"

// Add складывает два числа
func Add(a, b int) int {
	return a + b
}

// Multiply умножает два числа
func Multiply(a, b int) int {
	return a * b
}

// Divide делит a на b, возвращает ошибку при делении на 0
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

