package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Добро пожаловать на главную страницу!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Гость"
	}
	fmt.Fprintf(w, "Привет, %s!", name)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Текущее время: %s", time.Now().Format("15:04:05 02.01.2006"))
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/time", timeHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

