package main

import (
	"fmt"
)

func main() {
	a := []int{}
	printSlice(a)
	a = append(a, 1)
	printSlice(a)
	a = append(a, 2, 3, 4)
	printSlice(a)
	a2 := make([]int, len(a), cap(a)*2)
	copy(a, a2)
	printSlice(a2)

}

func printSlice(a []int) {
	fmt.Printf("Slice:%v\nLength:%d\nCap:%d\n", a, len(a), cap(a))
}
