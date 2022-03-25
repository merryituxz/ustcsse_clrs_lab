package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n, c int
	fmt.Scanf("%d %d", &n, &c)
	v, w := make([]int, n), make([]int, c)
	f := make([]int, c+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &v[i], &w[i])
	}
	for i := 0; i < n; i++ {
		for j := c; j >= v[i]; j-- {
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}
	fmt.Println(f[c])
}
