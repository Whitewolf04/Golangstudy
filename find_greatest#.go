package main

import "fmt"

func greatest(abcd ...float64) float64 {
	var b float64
	for _, v := range abcd {
		if v > b {
			b = v
		}
	}
	return b
}
func main() {
	data := []float64{36, 75, 65, 89, 53, 64, 90, 76}
	n := greatest(data...)
	fmt.Println(n)
}
