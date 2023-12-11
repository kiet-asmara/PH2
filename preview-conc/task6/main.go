package main

import (
	"fmt"
)

func main() {
	chOdd := make(chan int)
	chEven := make(chan int)
	done := make(chan bool)
	chError := make(chan error)

	go func() {
		for i := 1; i <= 22; i++ {
			if i > 20 {
				chError <- fmt.Errorf("number %d is greater than 20", i)
				continue
			}
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
		case err := <-chError:
			fmt.Println("Error:", err)
		case <-done:
			return
		}
	}

}
