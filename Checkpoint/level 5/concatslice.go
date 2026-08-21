package checkpoint

func ConcatSlice(slice1, slice2 []int) []int {
	output:=[]int{}
	for i:=0;i<len(slice1);i++{
		output = append(output,slice1[i])
	}
	for y:=0;y<len(slice2);y++{
		output = append(output, slice2[y])
	}
	return output
}