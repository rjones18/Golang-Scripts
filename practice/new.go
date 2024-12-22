package main

import (
	"fmt"
)

var name string // This is a global variable

// Initialize or any other descriptive name would be better than "new"
func Initialize() {
	// Add your implementation here
	fmt.Println("Initializing...")
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)
}

// Add the main function as the entry point
func main() {
	Initialize() // Call your Initialize function
}
