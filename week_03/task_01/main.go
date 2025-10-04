package main

import (
	"fmt"
	"math"
)

// processNumbers обрабатывает три попарно различных действительных числа
// согласно заданному алгоритму
func processNumbers(x, y, z float64) (float64, float64, float64) {
	// Проверяем, что числа попарно различны
	if x == y || x == z || y == z {
		fmt.Println("Ошибка: числа должны быть попарно различными")
		return x, y, z
	}

	// Вычисляем сумму трех чисел
	sum := x + y + z

	// Если сумма меньше единицы
	if sum < 1.0 {
		// Находим наименьшее число и заменяем его полусуммой двух других
		min := math.Min(math.Min(x, y), z)
		
		if min == x {
			// x наименьшее, заменяем на полусумму y и z
			x = (y + z) / 2
		} else if min == y {
			// y наименьшее, заменяем на полусумму x и z
			y = (x + z) / 2
		} else {
			// z наименьшее, заменяем на полусумму x и y
			z = (x + y) / 2
		}
	} else {
		// Сумма >= 1, заменяем меньшее из x и y полусуммой двух оставшихся
		if x < y {
			// x меньше y, заменяем x на полусумму y и z
			x = (y + z) / 2
		} else {
			// y меньше или равно x, заменяем y на полусумму x и z
			y = (x + z) / 2
		}
	}

	return x, y, z
}

func main() {
	fmt.Println("=== Обработка трех попарно различных чисел ===")
	fmt.Println()

	// Тестовые случаи
	testCases := []struct {
		name string
		x, y, z float64
	}{
		{"Случай 1: сумма < 1", 0.1, 0.2, 0.3},
		{"Случай 2: сумма >= 1", 0.5, 0.6, 0.7},
		{"Случай 3: сумма = 1", 0.3, 0.3, 0.4},
		{"Случай 4: отрицательные числа", -0.1, 0.2, 0.3},
		{"Случай 5: большие числа", 10.0, 20.0, 30.0},
	}

	for _, tc := range testCases {
		fmt.Printf("%s\n", tc.name)
		fmt.Printf("Исходные числа: x=%.2f, y=%.2f, z=%.2f\n", tc.x, tc.y, tc.z)
		fmt.Printf("Сумма: %.2f\n", tc.x+tc.y+tc.z)
		
		newX, newY, newZ := processNumbers(tc.x, tc.y, tc.z)
		
		fmt.Printf("Результат: x=%.2f, y=%.2f, z=%.2f\n", newX, newY, newZ)
		fmt.Println("---")
	}

	// Интерактивный ввод
	fmt.Println("Введите три попарно различных числа:")
	var x, y, z float64
	fmt.Print("x = ")
	fmt.Scanf("%f", &x)
	fmt.Print("y = ")
	fmt.Scanf("%f", &y)
	fmt.Print("z = ")
	fmt.Scanf("%f", &z)

	fmt.Printf("\nИсходные числа: x=%.2f, y=%.2f, z=%.2f\n", x, y, z)
	fmt.Printf("Сумма: %.2f\n", x+y+z)

	newX, newY, newZ := processNumbers(x, y, z)
	fmt.Printf("Результат: x=%.2f, y=%.2f, z=%.2f\n", newX, newY, newZ)
}
