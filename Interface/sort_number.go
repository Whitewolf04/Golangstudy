package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{4, 9, 8, 12, 56, 45, 87, 135, 124, 14, 21}
	sort.Slice(n, func(i, j int) bool {
		return n[i] < n[j]
	})
	fmt.Println(n)
}
