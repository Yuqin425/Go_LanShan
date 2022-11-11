package main

import "fmt"

var (
	n int
	l int
	r int
	a []int
)

func sort(a []int, l, r int) {
	x := 0
	for i := l; i <= r; i++ {
		for j := i; j <= r; j++ {
			if a[i] > a[j] {
				x = a[i]
				a[i] = a[j]
				a[j] = x
			}
		}
	}
}

func main() {
	fmt.Scanf("%d%d%d", &n, &l, &r)
	fmt.Scanln()
	a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	sort(a, l, r)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
