package lab2

func CountingSort(A, B []int, k int) {
	C := make([]int, k+1)
	for j := 0; j < len(A); j++ {
		C[A[j]]++
	}
	for i := 1; i <= k; i++ {
		C[i] += C[i-1]
	}

	for j := len(A) - 1; j >= 0; j-- {
		B[C[A[j]]-1] = A[j]
		C[A[j]]--
	}
}
