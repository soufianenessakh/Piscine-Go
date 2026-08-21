package level03

func RepeatAlpha(s string) string {
	output := ""
	for _, char := range s {

		if char >= 'a' && char <= 'z' {
			newchar := int(char-'a') + 1
			for i := 1; i <= newchar; i++ {
				output += string(rune(char))
			}

		} else if char >= 'A' && char <= 'Z' {
			newchar := int(char-'A') + 1
			for i := 1; i <= newchar; i++ {
				output += string(rune(char))
			}

		} else {
			output += string(rune(char))
		}
	}
	return output
}