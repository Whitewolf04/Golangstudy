package main

import (
	"fmt"
	"log"
)

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("Square root of negative number: %v", f)
	}

	return 42, nil
}

func main() {
	_, err := sqrt(-10.34)
	if err != nil {
		log.Fatalln(err)
	}
}
