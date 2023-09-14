package main

func main() {
	// Create a new channel
	ch := make(chan string)

	go func() {
		// Send a string through the channel
		ch <- "Hello World!"
	}()

	// Receive a string from the channel
	msg := <-ch
	println(msg)

}
