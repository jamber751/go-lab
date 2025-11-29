package main

import (
	"fmt"
	"math"
)

// Shape - интерфейс геометрической фигуры
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle - круг
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle - прямоугольник
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// printShapeInfo выводит информацию о фигуре
func printShapeInfo(name string, s Shape) {
	fmt.Printf("%s: площадь=%.2f, периметр=%.2f\n", name, s.Area(), s.Perimeter())
}

func main() {
	fmt.Println("=== Интерфейс Shape ===")

	circle := Circle{Radius: 5}
	rect := Rectangle{Width: 4, Height: 6}

	printShapeInfo("Круг (r=5)", circle)
	printShapeInfo("Прямоугольник (4x6)", rect)

	// Срез фигур
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
	}

	fmt.Println("\nВсе фигуры:")
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	fmt.Printf("Общая площадь: %.2f\n", totalArea)
}

