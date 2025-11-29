package main

import "fmt"

// Book представляет книгу в библиотеке
type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

// String возвращает форматированное описание книги
func (b Book) String() string {
	return fmt.Sprintf("%s (%s, %d) - %.2f руб.", b.Title, b.Author, b.Year, b.Price)
}

// Discount применяет скидку к цене книги
func (b *Book) Discount(percent float64) {
	b.Price = b.Price * (1 - percent/100)
}

func main() {
	fmt.Println("=== Работа со структурой Book ===")

	// Создание книги
	book := Book{
		Title:  "Война и мир",
		Author: "Л.Н. Толстой",
		Year:   1869,
		Price:  1500.00,
	}

	fmt.Println("Книга:", book)

	// Применяем скидку 20%
	book.Discount(20)
	fmt.Println("После скидки 20%:", book)

	// Создаём срез книг
	library := []Book{
		{"Мастер и Маргарита", "М.А. Булгаков", 1967, 800},
		{"1984", "Дж. Оруэлл", 1949, 600},
		{"Преступление и наказание", "Ф.М. Достоевский", 1866, 700},
	}

	fmt.Println("\nБиблиотека:")
	for _, b := range library {
		fmt.Println(" -", b)
	}
}

