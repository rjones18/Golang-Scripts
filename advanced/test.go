package main

import (
	"fmt"

	"github.com/pborman/uuid"
)

func main() {
	u := uuid.NewRandom()
	fmt.Println(u)
}
