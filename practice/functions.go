package main

import (
	"fmt"
)

func operation(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	sum, difference := operation(5, 10)
	fmt.Println(sum, difference)

}
