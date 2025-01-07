package main

import (
	"fmt"
)

func main() {
	x := func(l int, w int) int {
		return l * w
	}

	fmt.Println(x(20, 30))
}
