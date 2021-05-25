package test

func IsOne(i int) bool {
	if i == 1 {
		return true
	}
	return false
}

func Average(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}

	average := int(total / len(s))

	return average
}
