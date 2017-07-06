package main

import (
	"fmt"
)

func main() {
	nextInt := inc()
	for i := 0; i < 10; i++ {
		fmt.Println(nextInt())
	}
}

func inc() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
