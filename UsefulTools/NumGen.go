package UsefulTools

func NumGen()[]int{
	numList := make([]int, 101)
	for i := 0; i <= 1000; i++{
		numList = append(numList, i)
	}
	return numList
}

func NumGen2()[]int{
	numList := make([]int, 1000001)
	for k:=0; k<=1000000; k++{
		numList = append(numList, k)
	}
	return numList
}
