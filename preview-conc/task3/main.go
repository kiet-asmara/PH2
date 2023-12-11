package main

import (
	"fmt"
	"time"
)

func main() {
	chOut := make(chan int)
	go produce(chOut)
	go consume(chOut)
	time.Sleep(time.Millisecond * 2)
}

func produce(out chan int) {
	for i := 1; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func consume(chOut chan int) {
	for i := range chOut {
		fmt.Println(i)
	}
}
