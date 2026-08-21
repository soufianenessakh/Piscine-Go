package Quest10

func StringToIntSlice(str string) []int {
	output:=[]int{}
	for _,char:=range str{
		output = append(output, int(char))
	}
	return output
}