package Quest10

func DescendAppendRange(max, min int) []int {
	output:=[]int{}
	if max<min{
		return output
	}
	for i:=max;i>min;i--{
		output = append(output, i)
	}
	return output
}