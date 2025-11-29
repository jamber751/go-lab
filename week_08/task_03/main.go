package main

import (
	"fmt"
	"sync"
	"time"
)

func timer(name string, interval time.Duration, count int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= count; i++ {
		time.Sleep(interval)
		fmt.Printf("[%s] Тик %d (интервал: %v)\n", name, i, interval)
	}
}

func main() {
	fmt.Println("=== Конкурентные таймеры ===")

	var wg sync.WaitGroup

	// Запускаем таймеры с разными интервалами
	timers := []struct {
		name     string
		interval time.Duration
		count    int
	}{
		{"Быстрый", 200 * time.Millisecond, 5},
		{"Средний", 500 * time.Millisecond, 3},
		{"Медленный", 1 * time.Second, 2},
	}

	for _, t := range timers {
		wg.Add(1)
		go timer(t.name, t.interval, t.count, &wg)
	}

	wg.Wait()
	fmt.Println("Все таймеры завершены!")
}

