package main

import "fmt"

// Person - базовая структура человека
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Привет, меня зовут %s!", p.Name)
}

// Employee - сотрудник с встроенной структурой Person
type Employee struct {
	Person   // Встраивание
	Position string
	Salary   float64
}

func (e Employee) Info() string {
	return fmt.Sprintf("%s, %d лет, должность: %s, зарплата: %.2f",
		e.Name, e.Age, e.Position, e.Salary)
}

func main() {
	fmt.Println("=== Встраивание структур ===")

	person := Person{Name: "Иван", Age: 25}
	fmt.Println(person.Greet())

	employee := Employee{
		Person:   Person{Name: "Мария", Age: 30},
		Position: "Разработчик",
		Salary:   150000,
	}

	// Доступ к встроенным полям напрямую
	fmt.Println("Имя сотрудника:", employee.Name)
	fmt.Println(employee.Greet()) // Метод от Person
	fmt.Println(employee.Info())
}

