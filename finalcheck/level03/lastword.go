package level03


func LastWord(s string) string {
	output := ""
	end := len(s) - 1

	for end >= 0 && s[end] == ' ' {
		end--
	}

	for i := end; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		output = string(s[i]) + output
	} 

	return output 
}