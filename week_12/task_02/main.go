package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Tags  []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("=== Работа с JSON ===")

	// Создаём пользователей
	users := []User{
		{ID: 1, Name: "Иван", Email: "ivan@mail.ru", Tags: []string{"admin", "active"}},
		{ID: 2, Name: "Мария", Email: "maria@mail.ru"},
		{ID: 3, Name: "Пётр", Email: "petr@mail.ru", Tags: []string{"user"}},
	}

	// Сериализация в JSON
	data, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println("JSON:")
	fmt.Println(string(data))

	// Сохраняем в файл
	os.WriteFile("users.json", data, 0644)

	// Читаем обратно
	fileData, _ := os.ReadFile("users.json")
	var loaded []User
	json.Unmarshal(fileData, &loaded)

	fmt.Println("\nЗагружено из файла:")
	for _, u := range loaded {
		fmt.Printf("  %d: %s <%s>\n", u.ID, u.Name, u.Email)
	}
}

