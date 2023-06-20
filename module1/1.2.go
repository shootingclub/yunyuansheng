package main

import (
	"fmt"
	"math/rand"
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
	defer func() {
		if r := recover(); r == nil {
			fmt.Errorf("%s", "produce to panic")
			produce(queue)
		}
	}()
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			rand.Seed(time.Now().UnixNano())
			data := rand.Intn(100)
			queue <- data
			fmt.Println("生产者", data)
		}
	}
}
func consumer(queue chan int) {
	for data := range queue {
		fmt.Println("消费者", data)
		time.Sleep(time.Second)
	}
}
