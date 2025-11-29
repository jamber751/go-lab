package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Select с несколькими каналами ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Горутины с разными задержками
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Сообщение 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Сообщение 2"
	}()

	// Получаем сообщения с помощью select
	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Канал 1:", msg)
		case msg := <-ch2:
			fmt.Println("Канал 2:", msg)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Таймаут!")
		}
	}

	fmt.Println("Завершено!")
}

