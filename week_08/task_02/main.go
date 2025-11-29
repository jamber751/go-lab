package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=== Параллельное вычисление суммы ===")

	// Создаём большой массив
	size := 10_000_000
	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = i + 1
	}

	// Последовательная сумма
	start := time.Now()
	seqSum := 0
	for _, n := range numbers {
		seqSum += n
	}
	seqTime := time.Since(start)

	// Параллельная сумма
	start = time.Now()
	var parallelSum int64
	var wg sync.WaitGroup
	numWorkers := 4
	chunkSize := size / numWorkers

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localSum := int64(0)
			for i := start; i < end; i++ {
				localSum += int64(numbers[i])
			}
			atomic.AddInt64(&parallelSum, localSum)
		}(w*chunkSize, (w+1)*chunkSize)
	}
	wg.Wait()
	parTime := time.Since(start)

	fmt.Printf("Последовательная сумма: %d (время: %v)\n", seqSum, seqTime)
	fmt.Printf("Параллельная сумма: %d (время: %v)\n", parallelSum, parTime)
	fmt.Printf("Ускорение: %.2fx\n", float64(seqTime)/float64(parTime))
}

