package checkpoint

func RevConcatAlternate(slice1,slice2 []int) []int {
	output:=[]int{}
	i:=len(slice1)-1
	j:=len(slice2)-1
	if len(slice1)>len(slice2){
		for i>=len(slice2){
			output = append(output,slice1[i])
			i--
		}
	}
	 if len(slice2)>len(slice1){
		for j>=len(slice1){
			output = append(output, slice2[j])
			j--
		}
	}
	for i>=0&&j>=0{
		output = append(output, slice1[i])
		output = append(output, slice2[j])
		i--
		j--
	}	
return output
}
