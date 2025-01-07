package main

import (
	"fmt"
	"time"
)

func calculateSquare(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {

	start := time.Now()
	for i := 1; i <= 10; i++ {
		calculateSquare(i)
	}
	elasped := time.Since(start)
	fmt.Println("Total time taken: ", elasped)
}
