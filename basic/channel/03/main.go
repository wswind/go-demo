package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	//receive
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()

	//send
	for j := 0; j < 5; j++ {
		go func() {
			ch <- 42
			fmt.Println(<-ch)
			wg.Done()
		}()
	}

	wg.Wait()
}
