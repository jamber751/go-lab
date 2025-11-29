package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d: обработка задачи %d\n", id, job)
		time.Sleep(200 * time.Millisecond) // Имитация работы
	}
}

func main() {
	fmt.Println("=== Буферизированный канал ===")

	jobs := make(chan int, 10) // Буфер на 10 задач
	var wg sync.WaitGroup

	// Запускаем 3 worker'а
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Отправляем задачи
	for j := 1; j <= 9; j++ {
		jobs <- j
		fmt.Printf("Задача %d добавлена в очередь\n", j)
	}
	close(jobs)

	wg.Wait()
	fmt.Println("Все задачи выполнены!")
}

