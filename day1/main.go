package main

import (
	"fmt"
	"time"
)

func sumation(a int) int {
	total := 0
	for i := 1; i <= a; i++ {
		total += i
	}
	return total
}

func main() {
	n := 1000000000
	st := time.Now()
	fmt.Println(sumation(n))
	et := time.Now()

	fmt.Println("Time taken: ", et.Sub(st))
}
