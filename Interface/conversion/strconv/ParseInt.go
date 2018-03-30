package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("12.453", 64)
	i, _ := strconv.ParseInt("34", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, i, u)

	w := strconv.FormatBool(true)
	x := strconv.FormatFloat(12.456456, 'E', -1, 64)
	y := strconv.FormatInt(-42, 10)
	z := strconv.FormatUint(42, 16)

	fmt.Println(w, x, y, z)
}
