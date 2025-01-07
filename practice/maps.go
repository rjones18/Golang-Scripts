package main

import (
	"fmt"
)

func main() {
	codes := map[string]string{
		"en": "English",
		"fr": "French",
		"es": "Spanish",
		"it": "Italian",
		"de": "German",
	}
	fmt.Println(codes["en"])
}
