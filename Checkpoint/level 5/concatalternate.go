package checkpoint

func ConcatAlternate(slice1, slice2 []int) []int {
	output := []int{}

	min := len(slice1)
	if len(slice2) < min {
		min = len(slice2)
	}

	for i := 0; i < min; i++ {
		if len(slice1) < len(slice2) {
			output = append(output, slice2[i])
			output = append(output, slice1[i])
		} else {
			output = append(output, slice1[i])
			output = append(output, slice2[i])
		}
	}

	if len(slice1) > len(slice2) {
		for i := min; i < len(slice1); i++ {
			output = append(output, slice1[i])
		}
	} else if len(slice2) > len(slice1) {
		for i := min; i < len(slice2); i++ {
			output = append(output, slice2[i])
		}
	}

	return output
}