package main

func main() {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	println(<-ch)
	println(<-ch)
}
