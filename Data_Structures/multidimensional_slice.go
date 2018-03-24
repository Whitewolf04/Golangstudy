package main

import "fmt"

func main() {
	records := make([][]string, 0)
	student1 := make([]string, 3)
	student1[0] = "Anya"
	student1[1] = "77.0"
	student1[2] = "90.0"
	records = append(records, student1)
	student2 := make([]string, 3)
	student2[0] = "Todd"
	student2[1] = "56.0"
	student2[2] = "76.0"
	records = append(records, student2)
	fmt.Println(records)
}
