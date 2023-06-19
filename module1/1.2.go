package main

import (
	"fmt"
	"time"
)

func main() {
	Queue()
}

func Queue() {
	queue := make(chan int, 10)
	go produce(queue)
	consumer(queue)
}

func produce(queue chan int) {
	for i := 0; i < 10; i++ {
		queue <- i
		fmt.Println("生产者", i)
		time.Sleep(time.Second)
	}

}
func consumer(queue chan int) {
	for data := range queue {
		fmt.Println("消费者", data)
		time.Sleep(time.Second)
	}
}
