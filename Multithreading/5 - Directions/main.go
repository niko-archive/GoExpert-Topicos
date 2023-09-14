package main

func main() {
	ch := make(chan string)

	go receiveOnly(ch)
	sendOnly("Hello", ch)

}

// chan<- string: send only
func sendOnly(nome string, ch chan<- string) {
	ch <- nome
}

// <-chan string: receive only
func receiveOnly(ch <-chan string) {
	println(<-ch)
}
