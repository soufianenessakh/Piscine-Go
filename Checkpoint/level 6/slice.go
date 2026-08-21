package checkpoint

func Slice(a []string, nbrs ...int) []string {
	start := 0
	end := len(a)

	if len(nbrs) == 1 {
		start = nbrs[0]
	} else if len(nbrs) >= 2 {
		start = nbrs[0]
		end = nbrs[1]
	}

	if start < 0 {
		start = len(a) + start
	}
	if end < 0 {
		end = len(a) + end
	}

	if start > len(a) || end > len(a) || start > end {
		return nil
	}
	return a[start:end]
}