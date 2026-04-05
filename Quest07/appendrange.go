package Quest07

func AppendRange(min, max int) []int {
	output:= []int {}
	if min>=max {
		return nil
	}
	for i:=min ;i<max;i++{
		output=append(output, i)
	}
	return output
}