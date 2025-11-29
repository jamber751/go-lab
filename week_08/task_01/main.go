package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("Горутина %d: %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== Базовые горутины ===")

	var wg sync.WaitGroup

	// Запускаем 5 горутин
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Ждём завершения всех горутин
	wg.Wait()
	fmt.Println("Все горутины завершены!")
}

