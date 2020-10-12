package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	//ms = &myStruct{foo: 42}
	fmt.Println(ms)
	ms = new(myStruct)
	fmt.Println(ms)
	fmt.Printf("%p", ms)
}

type myStruct struct {
	foo int
}
