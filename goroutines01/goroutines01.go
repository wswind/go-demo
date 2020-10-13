package main

import (
	"fmt"
	"time"
)

// func main() {
// 	go sayHello()
// 	time.Sleep(100 * time.Millisecond)

// }

// func sayHello() {
// 	fmt.Println("Hello")
// }

//----
// func main() {
// 	var msg = "Hello"
// 	go func(msg string) {
// 		fmt.Println(msg)
// 	}(msg)

// 	msg = "goodbye"

// 	time.Sleep(100 * time.Millisecond)

// }

//----
//go run -race goroutines01.go
func main() {
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()

	msg = "goodbye"

	time.Sleep(100 * time.Millisecond)

}
