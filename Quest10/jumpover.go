package Quest10

func JumpOver(str string) string {
	output:=""
	if len(str)<3{
		output=string(rune('\n'))
		return output
	}
	for i:=2;i<len(str);i+=3{
		output+=string(str[i])
	}
	if output == "" {
		return "\n"
	}
	return output+"\n"
}