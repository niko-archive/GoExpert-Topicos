package main

import "sync"

const TOTAL = 100

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(TOTAL)

	go publish(ch, &wg)
	go receive(ch)

	wg.Wait()

}

func publish(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	for x := 0; x < TOTAL; x++ {
		ch <- x
		wg.Done()
	}
}

func receive(ch <-chan int) {
	for x := range ch {
		println(x)
	}
}
