package main

import (
	"fmt"
)

var name string // This is a global variable
var gym bool    // This is a global variable

// Initialize or any other descriptive name would be better than "new"
func Initialize2() {
	// Add your implementation here
	fmt.Println("Initializing...")
	fmt.Print("Enter your first and are you a gym rat : ")
	fmt.Scanf("%s %t", &name, &gym)
	fmt.Println(name, gym)
}

// Add the main function as the entry point
func main() {
	Initialize2() // Call your Initialize function
}
