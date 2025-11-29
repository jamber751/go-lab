package main

import (
	"fmt"
	"time"
)

// producer отправляет числа в канал
func producer(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("Отправка: %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

// consumer получает и обрабатывает числа
func consumer(ch <-chan int, done chan<- bool) {
	sum := 0
	for num := range ch {
		fmt.Printf("Получено: %d\n", num)
		sum += num
	}
	fmt.Printf("Сумма: %d\n", sum)
	done <- true
}

func main() {
	fmt.Println("=== Producer-Consumer ===")

	ch := make(chan int)
	done := make(chan bool)

	go producer(ch, 5)
	go consumer(ch, done)

	<-done
	fmt.Println("Завершено!")
}

