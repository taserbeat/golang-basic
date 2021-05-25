package xxx

import "main/yyy"

const (
	Max = 100
	min = 1
)

func GetMin() int {
	return min
}

func NewUser() yyy.User {
	user := yyy.User{Name: "Tom", Age: 11}

	return user
}
