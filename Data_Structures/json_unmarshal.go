package main

import (
	"encoding/json"
	"fmt"
)

type human struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 human
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First": "James", "Last": "Bond", "Age": 20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
