package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	QueueTest()
}

type Queue struct {
	q    chan int
	cond *sync.Cond
}

func QueueTest() {
	queue := Queue{
		q:    make(chan int, 10),
		cond: sync.NewCond(&sync.Mutex{}),
	}

	go queue.consumer()
	for {
		time.Sleep(time.Second)
		rand.Seed(time.Now().UnixNano())
		data := rand.Intn(100)
		queue.produce(data)
	}

}

func (queue *Queue) produce(item int) {
	queue.cond.L.Lock()
	defer queue.cond.L.Unlock()
	// 开始入队
	queue.q <- item
	fmt.Println("生产者", item)
	// 入队进行判断是否队列已满
	if len(queue.q) == cap(queue.q) {
		fmt.Println("生产完成 开始消费")
		queue.cond.Wait()
	}
}

func (queue *Queue) consumer() {
	for {
		// 是否满了
		if cap(queue.q) == len(queue.q) {
			for data := range queue.q { //消费
				if len(queue.q) != 0 {
					time.Sleep(time.Second)
					fmt.Println("消费者", data)
				} else { // 消费完了
					fmt.Println("消费完成 开始生产")
					queue.cond.Signal()
					break
				}
			}
		}
	}
}
