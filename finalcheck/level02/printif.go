package level02


func PrintIf(str string)string{
	if str=="" || len(str)>=3{
		return "G"
	}
	return "Invalid Input"
}