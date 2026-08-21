package Quest10

func Join(strs []string, sep string) string {
	output := ""

	for i := 0; i < len(strs); i++ {
		output += strs[i]
		if i != len(strs)-1 {
			output += sep
		}
	}

	return output
}