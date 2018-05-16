package main

import (
	"errors"
	"fmt"
	"log"
)

var errSyntax = errors.New("Syntax Error")

func exp(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("Syntax Error")
	}
	return n * n, error(nil)
}

func main() {
	n, err := exp(10)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(errSyntax)
	} else {
		fmt.Println(n)
	}
}
