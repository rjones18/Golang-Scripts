package main

import (
	"fmt"
)

func main() {
	var fruits [2]string = [2]string{"apples", "oranges"}
	fmt.Println(fruits)

	names := [...]string{"Alice", "Bob", "Charlie"}
	fmt.Println(names)

	for i, v := range names {
		fmt.Println(i, v)
	}

}
