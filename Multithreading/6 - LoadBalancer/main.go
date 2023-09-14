package main

import (
	"fmt"
	"sync"
	"time"
)

const TOTAL = 100
const WORKERS = 3

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(TOTAL)

	go publish(ch, &wg)
	for i := 0; i < WORKERS; i++ {
		go worker(i, ch)
	}

	wg.Wait()

}

func publish(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	for x := 1; x <= TOTAL; x++ {
		ch <- x
		wg.Done()
	}
}

func worker(id int, ch <-chan int) {
	for x := range ch {
		msg := fmt.Sprintf("Worker %d: %d", id, x)
		time.Sleep(time.Millisecond * 100)
		println(msg)

	}
}
