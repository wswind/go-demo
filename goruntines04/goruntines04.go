package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads: %v", runtime.GOMAXPROCS(-1))
}
