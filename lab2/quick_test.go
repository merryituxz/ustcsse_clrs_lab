package lab2

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
