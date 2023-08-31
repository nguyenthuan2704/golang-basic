package main

import (
	"fmt"
	"strings"
)

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func mapInt(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func mapAny[K, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func main() {
	si := []int{10, 25, 49, 56, 78}
	fmt.Println(Index(si, 56))

	st := []string{"Hello", "World", "Golang"}
	fmt.Println(Index(st, "Golang"))

	arr := []int{1, 2, 3, 4, 5}
	str := []string{"Senior", "Software", "Engineer"}
	rs := mapInt(arr, func(v int) int {
		return v * 2
	})
	ra := mapAny(arr, func(v int) int {
		return v * 2
	})
	sr := mapAny(str, func(v string) string {
		return strings.ToUpper(v)
	})
	fmt.Println(rs)
	fmt.Println(ra)
	fmt.Println(sr)
}
