package main

import (
	"fmt"
	"math"
)

// solveQuadratic решает квадратное уравнение ax² + bx + c = 0
// Возвращает количество корней и сами корни
func solveQuadratic(a, b, c float64) (int, float64, float64) {
	// Вычисляем дискриминант
	discriminant := b*b - 4*a*c
	
	// Если дискриминант отрицательный, корней нет
	if discriminant < 0 {
		return 0, 0, 0
	}
	
	// Если дискриминант равен нулю, один корень
	if discriminant == 0 {
		root := -b / (2 * a)
		return 1, root, 0
	}
	
	// Если дискриминант положительный, два корня
	sqrtDiscriminant := math.Sqrt(discriminant)
	root1 := (-b + sqrtDiscriminant) / (2 * a)
	root2 := (-b - sqrtDiscriminant) / (2 * a)
	
	return 2, root1, root2
}

func main() {
	var a, b, c float64
	
	fmt.Println("Решение квадратного уравнения ax² + bx + c = 0")
	fmt.Print("Введите коэффициент a (a ≠ 0): ")
	fmt.Scan(&a)
	
	// Проверяем, что a не равно нулю
	if a == 0 {
		fmt.Println("Ошибка: коэффициент a не может быть равен нулю!")
		return
	}
	
	fmt.Print("Введите коэффициент b: ")
	fmt.Scan(&b)
	
	fmt.Print("Введите коэффициент c: ")
	fmt.Scan(&c)
	
	fmt.Printf("\nУравнение: %.2fx² + %.2fx + %.2f = 0\n", a, b, c)
	
	// Решаем уравнение
	numRoots, root1, root2 := solveQuadratic(a, b, c)
	
	switch numRoots {
	case 0:
		fmt.Println("Действительных корней нет.")
	case 1:
		fmt.Printf("Уравнение имеет один корень: x = %.4f\n", root1)
	case 2:
		fmt.Printf("Уравнение имеет два корня:\n")
		fmt.Printf("x₁ = %.4f\n", root1)
		fmt.Printf("x₂ = %.4f\n", root2)
	}
}
