package checkpoint

func CanJump(c []uint) bool {
	if len(c) == 0 {
		return false
	}

	if len(c) == 1 {
		return true
	}

	pos := 0
	last := len(c) - 1

	for pos < len(c) {
		next := pos + int(c[pos])

		if next == last {
			return true
		}

		if next > last {
			return false
		}

		pos = next
	}

	return false
}