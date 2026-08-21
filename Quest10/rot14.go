package Quest10

func Rot14(s string) string {
	output := ""
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			pos := int(c - 'a')
			pos = (pos + 14) % 26
			output += string(rune(pos + int('a')))
		} else if c >= 'A' && c <= 'Z' {
			pos := int(c - 'A')
			pos = (pos + 14) % 26
			output += string(rune(pos + int('A')))
		} else {
			output += string(c)
		}
	}
	return output
}