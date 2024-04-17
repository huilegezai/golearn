package main

import (
	"fmt"
)

const s string = "constant"

func main() {
	fmt.Print(s)
	const n = 5000000000
	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))
}
