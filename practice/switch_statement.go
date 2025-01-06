package main

import (
	"fmt"
)

func main() {

	var i int = 100
	switch i {
	case 10:
		fmt.Println("i is 10")
	case 100:
		fmt.Println("string 100")
	default:
		fmt.Println("default")
	}
}
