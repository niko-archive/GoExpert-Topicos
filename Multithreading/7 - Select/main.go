package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	select {
	case x := <-c1:
		println(x)
	case x := <-c2:
		println(x)
	case <-time.After(time.Second * 3):
		println("timeout")
	}

}
