package main

import (
	"fmt"
	"sync"
)

func main() {
	let := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	var wg sync.WaitGroup

	wg.Add(2)
	go PrintLetters(let, &wg)
	go PrintNumbers(&wg)
	wg.Wait()
}

func PrintNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func PrintLetters(let []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, l := range let {
		fmt.Println(l)
	}
}
