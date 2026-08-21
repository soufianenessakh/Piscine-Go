package Quest10

func ReverseMenuIndex(menu []string) []string {
	output:=[]string{}
	for i:=len(menu)-1;i>=0;i--{
		output = append(output, menu[i])
	}
	return output
}