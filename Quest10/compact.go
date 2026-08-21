package Quest10

func Compact(ptr *[]string) int {
	s:=*ptr
	output:=[]string{}
	for _,v:=range s{
		if v!=""{
			output = append(output, v)
		}
	}
	*ptr=output
	return len(output)
}