package checkpoint

func NotDecimal(dec string) string {

	str := ""
		if dec == "" {
		return "\n"
	}
	for i, c := range dec {
		if c == '-' && i == 0 {
			str += string(c)
		}
		if c=='0' && i==0{
			continue
		}
		if ('a' <= c && c >= 'z') || ('A' <= c && c >= 'Z') {
			return dec
		}
		if c != '.' && c != '-'{
			str += string(c)
		}
	}
	return str + string('\n')
}
