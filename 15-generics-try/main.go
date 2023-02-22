package main

import "fmt"

func Apply[T any](s []T, f func(int, T)) {
	for i, v := range s {
		f(i, v)
	}
}

func main() {
	var sum int
	Apply[int]([]int{10, 20}, func(i, v int) {
		sum += v
	})
	fmt.Println(sum) // 30
}
