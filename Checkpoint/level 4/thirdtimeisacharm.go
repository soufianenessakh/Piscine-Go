package checkpoint

func ThirdTimeIsACharm(str string) string {
	if len(str)<3{
		return "\n"
	}
	output:=""
	for i:=2 ;i<len(str);i+=3{
		output+=string(str[i])
	}
	output+="\n"
	return output
}