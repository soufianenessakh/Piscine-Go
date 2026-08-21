package level02

func CheckNumber(arg string)bool{
	for _,c:=range(arg){
		if (c>='1'&& c<='9'){
			return true
		}
	}
	return false
}