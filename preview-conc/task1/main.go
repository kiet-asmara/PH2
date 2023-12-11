package main

import (
	"fmt"
	"time"
)

func main() {
	let := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	go PrintLetters(let)
	go PrintNumbers()
	time.Sleep(time.Second)
}

func PrintNumbers() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
}

func PrintLetters(let []string) {
	for _, l := range let {
		go fmt.Println(l)
	}

}
