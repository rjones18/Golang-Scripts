package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 90
	var f float64 = float64(i)
	var s string = strconv.Itoa(i)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%q", s)
}
