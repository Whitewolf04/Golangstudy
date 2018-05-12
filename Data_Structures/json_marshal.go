package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	First       string
	Last        string `json: "-"`
	Age         int    `json: "wisdom score"`
	notexported int
}

func main() {
	p1 := people{"James", "Bond", 42, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
