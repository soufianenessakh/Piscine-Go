package level02



func PrintIfNot(str string)string{
	if str=="" || len(str)<3{
		return "G"
	}
	return "Invalid Input"
}