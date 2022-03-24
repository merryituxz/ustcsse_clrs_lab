package lab2

func QuickSort(arr []int, p, r int) {
	if p < r {
		q := Partition(arr, p, r)
		QuickSort(arr, p, q-1)
		QuickSort(arr, q+1, r)
	}
}

func Partition(arr []int, p, r int) int {
	x := arr[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
