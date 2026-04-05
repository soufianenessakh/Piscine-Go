package Quest07

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	nbrsilce := max - min
	output := make([]int, nbrsilce)
	for j := 0; j < nbrsilce; j++ {
		output[j] = min + j
	}
	return output
}
