package UsefulTools

func CheckPrimeNumber(num int)bool{
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}
