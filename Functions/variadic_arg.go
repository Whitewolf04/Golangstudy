package main

import "fmt"

func main() {
	data := []float64{65, 54, 76, 89}
	n := avg(data...)
	fmt.Println(n)
}
func avg(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
