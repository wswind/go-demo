package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var msg = "Hello"
// 	go func(msg string) {
// 		fmt.Println(msg)
// 	}(msg)

// 	msg = "goodbye"

// 	time.Sleep(100 * time.Millisecond)

// }

var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)

	msg = "goodbye"
	wg.Wait()
}
