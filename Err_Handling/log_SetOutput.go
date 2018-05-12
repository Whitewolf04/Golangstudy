package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("err.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("nofile.txt")
	if err != nil {
		log.Println(err)
	}
}
