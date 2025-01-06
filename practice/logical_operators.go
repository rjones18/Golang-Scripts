package main

import (
	"fmt"
)

func main() {
	var x int = 10
	fmt.Println((x < 100) && (x > 5))
	fmt.Println((x < 5) || (x > 50))
	fmt.Println(!(x == 10))

}
