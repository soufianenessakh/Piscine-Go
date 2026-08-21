package checkpoint

func PrintIf(s string) string {
	if len(s) == 0 || len(s) >= 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}