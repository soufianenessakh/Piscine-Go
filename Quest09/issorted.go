package Quest09

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	direction := 0

	for i := 0; i < len(a)-1; i++ {
		result := f(a[i], a[i+1])

		if result == 0 {
			continue
		}

		if direction == 0 {
			direction = result
			continue
		}

		if (direction > 0 && result < 0) || (direction < 0 && result > 0) {
			return false
		}
	}

	return true
}