package checkpoint

func format(n int) string {
	if n < 10 {
		return "0" + string(rune(n+'0'))
	}
	return string(rune((n/10)+'0')) + string(rune((n%10)+'0'))
}

func FromTo(from int, to int) string {
	output := ""
	if (from > 99 || to > 99) || (from < 0 || to < 0) {
		return "Invalid\n"
	}
	if from < to {
		for i := from; i <= to; i++ {
			if i != from {
				output += ", "
			}
			output += format(i)
		}
	} else {
		for i := from; i >= to; i-- {
			if i != from {
				output += ", "
			}
			output += format(i)
		}
	}
	return output + "\n"
}
