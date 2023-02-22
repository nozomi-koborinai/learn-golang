package main

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Ptr[T any](s T) *T {
	fmt.Println(&s)
	return &s
}

func main() {
	Print[string]([]string{"Hello, ", "playground\n"})
	Print[int]([]int{10, 20})

	ptr := Ptr[bool](false)
	fmt.Println(ptr)
}
