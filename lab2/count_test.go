package lab2

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	A := []int{95, 94, 91, 98, 99, 90, 99, 93, 91, 92}
	B := make([]int, len(A))
	CountingSort(A, B, 100)
	fmt.Println(B)
}
