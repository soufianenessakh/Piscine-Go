package checkpoint

import q "piscine/Quest04"

func FindPrevPrime(nb int) int {
	if nb <= 1 {
		return 0
	}

	for i := nb; i >= 2; i-- {
		if q.IsPrime(i) {
			return i
		}
	}

	return 0
}