package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; ожидалось 5", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)
	if result != 20 {
		t.Errorf("Multiply(4, 5) = %d; ожидалось 20", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) вернул ошибку: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; ожидалось 5", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) должен вернуть ошибку")
	}
}

