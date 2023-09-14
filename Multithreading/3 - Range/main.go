package main

func main() {
	// Create a new channel
	ch := make(chan int)

	// Publish
	go publish(ch)

	// Receive
	receive(ch)

}

func publish(ch chan<- int) {
	defer close(ch)
	for x := 0; x < 100; x++ {
		ch <- x
	}
}

func receive(ch <-chan int) {
	for x := range ch {
		println(x)
	}
}
