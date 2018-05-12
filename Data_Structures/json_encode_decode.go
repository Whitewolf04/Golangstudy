package main

import (
	"encoding/json"
	"os"
	"strings"
	"fmt"
)

type humanity struct {
	Name   string
	Age    int
	Status string
}

func secondary() {
	p1 := humanity{"Tuan", 25, "Married"}
	json.NewEncoder(os.Stdout).Encode(p1)
}

func main(){
	var p1 humanity
	rdr := strings.NewReader(`{"Name": "Tuan", "Age": 25, "Status": "Married"}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.Status)
}