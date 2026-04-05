package piscine

func TrimAtoi(s string) int {
	output := 0
	sign := false
	for _, char := range s {
		if char == '-' {
			if output == 0 {
				sign = true
			}
		}
		if char <= '0' || char > '9' {
			continue
		} else {
			output = output * 10
			output += int(char - '0')
		}
	}
	if sign == true {
		return -output
	}
	return output
}
