package test

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップさせる")
	}

	v := IsOne(i)

	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}

func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("スキップさせる")
	}

	v := Average([]int{1, 2, 3, 4, 5})

	if v != 3 {
		t.Errorf("%v != %v", v, 3)
	}
}
