package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuidObj1, _ := uuid.NewUUID()
	fmt.Println("   ", uuidObj1.String())

	uuidObj2, _ := uuid.NewRandom()
	fmt.Println("   ", uuidObj2.String())
}
