package Quest09

func CountIf(f func(string) bool, tab []string) int {
	output:=0
	for _,char:=range tab{
		if f(char)==true{
		 output++
		}
	}
	return output
}
func IsNumeric(s string) bool {
	for _,nambr:=range(s){
		if (nambr<'0' || nambr>'9'){
			return  false
		}	
	}
	return true
}