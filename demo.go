package main

import (
	"fmt"
	"math/rand"
)

func main() {
	println(Sum(1, 2))
	for i := 0; i < 10; i++ {
		if i == 1 {
			fmt.Println("%d / ", i)
		}
		a := rand.Int()
		fmt.Println("%d / ", a)
		fmt.Println("%d / ", &a)
	}
	println("hello world")
}
func Sum(a, b int) int {
	return a + b
}
