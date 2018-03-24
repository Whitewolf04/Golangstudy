package main

import "fmt"

func main() {
	//m := make(map[string]int)
	m := map[string]int{}
	m["k1"] = 7
	m["k2"] = 49
	fmt.Println("map: ", m)
	delete(m, "k1") //To delete an element in the map
	fmt.Println("map: ", m)
	_, prs := m["k1"]
	fmt.Println("prs: ", prs)
}
