package main

import (
	"fmt"
)

func main() {
	chOdd := make(chan int)
	chEven := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 20; i++ {
			if i%2 == 0 {
				chEven <- i
			} else {
				chOdd <- i
			}
		}
		close(done)

	}()

	for {
		select {
		case x := <-chEven:
			fmt.Println("Received an even number:", x)
		case x := <-chOdd:
			fmt.Println("Received an odd number:", x)
		case <-done:
			return
		}
	}

}
