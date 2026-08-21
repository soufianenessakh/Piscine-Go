package checkpoint

func FirstWord(s string) string {
	output := ""
	started := false

	for _, c := range s {
		if c != ' ' {
			started = true
			output += string(c)
		} else if started {
			break
		}
	}

	if output == "" {
		return "\n"
	}

	return output + "\n"
}