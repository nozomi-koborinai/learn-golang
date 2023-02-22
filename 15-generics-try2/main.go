package main

import "fmt"

func Filter[T any](s []T, f func(int, T) bool) []T {
	var r []T
	for i, v := range s {
		if f(i, v) {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	ns := []int{1, 2, 3, 4}
	ms := Filter[int](ns, func(i, v int) bool {
		return v%2 == 0
	})
	fmt.Println(ms) // [2 4]
}
