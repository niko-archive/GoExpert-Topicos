package main

func main() {
	// Create a new channel
	forever := make(chan bool)

	// Generate a Deadlock:
	// The channel is waiting for a value to be sent, but never receives it.
	// The program will be blocked forever. Or Go compiler will detect it and report an error.
	<-forever

}
