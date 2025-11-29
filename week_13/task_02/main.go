package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	tasks  = make(map[int]Task)
	nextID = 1
	mu     sync.Mutex
)

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		mu.Lock()
		list := make([]Task, 0, len(tasks))
		for _, t := range tasks {
			list = append(list, t)
		}
		mu.Unlock()
		json.NewEncoder(w).Encode(list)

	case "POST":
		var t Task
		json.NewDecoder(r.Body).Decode(&t)
		mu.Lock()
		t.ID = nextID
		nextID++
		tasks[t.ID] = t
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)

	case "DELETE":
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) >= 3 {
			id, _ := strconv.Atoi(parts[2])
			mu.Lock()
			delete(tasks, id)
			mu.Unlock()
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	// Добавляем тестовые данные
	tasks[1] = Task{ID: 1, Title: "Изучить Go", Done: false}
	tasks[2] = Task{ID: 2, Title: "Написать API", Done: true}
	nextID = 3

	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", tasksHandler)

	fmt.Println("REST API запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

