package main

import "fmt"

func foo() {
	for i := 0; i <= 1; i++ {
		fmt.Println(i)
	}
}

func zoo() {
	k := 0
	for {
		fmt.Println(k)
		k++
		if k >= 10 {
			break
		}
	}
}

func main() {
	t := 0
	for {
		t++
		if t%2 == 0 {
			continue
		}
		fmt.Println(t)
		if t > 50 {
			break
		}
	}
}
