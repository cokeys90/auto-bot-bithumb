package test

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		t.Log("testStart")

		Sum(1, 2, 3, 4, 5)
	})
}

// Sum -
func Sum(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
		fmt.Println(sum)
	}
	return sum
}
