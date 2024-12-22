package main

import (
	"fmt"
)

var she string = "She" //This is a global variable
// this is a comment
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	name := "Lisa"                // Simple and concise
	var greeting string = "Hello" // Explicitly states this is a string
	var i int = 78
	var user string
	user = "Reg"

	fmt.Println(user)
	fmt.Println(greeting, name)
	fmt.Println(name)
	fmt.Printf("Hey, %v! You have been served! %d/100", name, i)
}
