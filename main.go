package main

import (
	"fmt"
	"main/test"
)

// テスト

func main() {
	fmt.Println(test.IsOne(1))
	fmt.Println(test.IsOne(0))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(test.Average(s))
}
