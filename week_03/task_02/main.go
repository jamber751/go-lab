package main

import (
	"fmt"
	"math"
)

const PI = 3.14

// Структура для хранения элементов окружности
type Circle struct {
	Radius      float64
	Diameter    float64
	Circumference float64
	Area        float64
}

// Вычисляет все элементы окружности по радиусу
func calculateFromRadius(radius float64) Circle {
	return Circle{
		Radius:        radius,
		Diameter:      2 * radius,
		Circumference: 2 * PI * radius,
		Area:          PI * radius * radius,
	}
}

// Вычисляет все элементы окружности по диаметру
func calculateFromDiameter(diameter float64) Circle {
	radius := diameter / 2
	return calculateFromRadius(radius)
}

// Вычисляет все элементы окружности по длине окружности
func calculateFromCircumference(circumference float64) Circle {
	radius := circumference / (2 * PI)
	return calculateFromRadius(radius)
}

// Вычисляет все элементы окружности по площади
func calculateFromArea(area float64) Circle {
	radius := math.Sqrt(area / PI)
	return calculateFromRadius(radius)
}

// Выводит все элементы окружности в заданном порядке
func printCircle(c Circle) {
	fmt.Printf("Радиус: %.2f\n", c.Radius)
	fmt.Printf("Диаметр: %.2f\n", c.Diameter)
	fmt.Printf("Длина окружности: %.2f\n", c.Circumference)
	fmt.Printf("Площадь круга: %.2f\n", c.Area)
}

func main() {
	var elementNumber int
	var value float64

	fmt.Println("=== Вычисление элементов окружности ===")
	fmt.Println("Выберите известный элемент:")
	fmt.Println("1 - Радиус (R)")
	fmt.Println("2 - Диаметр (D)")
	fmt.Println("3 - Длина окружности (L)")
	fmt.Println("4 - Площадь круга (S)")
	fmt.Print("Введите номер элемента (1-4): ")
	
	_, err := fmt.Scanf("%d", &elementNumber)
	if err != nil || elementNumber < 1 || elementNumber > 4 {
		fmt.Println("Ошибка: введите число от 1 до 4")
		return
	}

	fmt.Print("Введите значение: ")
	_, err = fmt.Scanf("%f", &value)
	if err != nil {
		fmt.Println("Ошибка: введите корректное число")
		return
	}

	if value <= 0 {
		fmt.Println("Ошибка: значение должно быть положительным")
		return
	}

	var circle Circle

	// Вычисляем все элементы в зависимости от введенного
	switch elementNumber {
	case 1:
		fmt.Printf("\nИзвестен радиус: %.2f\n", value)
		circle = calculateFromRadius(value)
	case 2:
		fmt.Printf("\nИзвестен диаметр: %.2f\n", value)
		circle = calculateFromDiameter(value)
	case 3:
		fmt.Printf("\nИзвестна длина окружности: %.2f\n", value)
		circle = calculateFromCircumference(value)
	case 4:
		fmt.Printf("\nИзвестна площадь круга: %.2f\n", value)
		circle = calculateFromArea(value)
	}

	fmt.Println("\nРезультат:")
	fmt.Println("==========")
	printCircle(circle)
}
