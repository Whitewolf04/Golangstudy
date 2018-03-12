package main

import "fmt"

func changeMe(z []string) {
	z[0] = "Tuan"
	fmt.Println(z)
}
func main() {
	m := make([]string, 1, 25)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}
