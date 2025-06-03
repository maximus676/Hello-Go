package main

import (
	"fmt"
	"time"
)

// Пример работы горутины   не совсем понятно для чего это обсудить с Жекой  ????????

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(2 * time.Second)
	fmt.Println("Working...")
	done <- true // Отправка значения в канал done
}

func main() {
	done := make(chan bool, 1) // Создание канала с буфером размером 1
	go worker(done)            // Запуск горутины worker

	<-done // Ожидание получения значения из канала done

	fmt.Println("Work completed")

	// так же пример задержки к примеру дял того что бы успела выполниться гроутина
	time.Sleep(2 * time.Second)
	fmt.Println("Прошло еще 2 секунды")

}
