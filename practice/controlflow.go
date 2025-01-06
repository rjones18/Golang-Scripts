package main

import (
	"fmt"
)

func main() {
	var city string = "London"
	var city_2 string = "Paris"
	if city == city_2 {
		fmt.Println("Equal")
	} else if city != city_2 {
		fmt.Println("Not Equal")
	} else {
		fmt.Println("Help")
	}
}
