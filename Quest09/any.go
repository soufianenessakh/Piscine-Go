package Quest09

func Any(f func(string) bool, a []string) bool {
	for _, char := range a {
		if f(char) == true {
			return true
		}
	}
	return false
}
