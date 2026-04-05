package Quest07

func ConcatParams(args []string) string {
	output:=""
	for i:=0;i<=len(args)-1;i++{
		output += args[i]
		if i!=len(args)-1{
			output+="\n"
		}
	}
	return output
}