package piscine

func IsNumeric(s string) bool {
	for _,nambr:=range(s){
		if (nambr<'0' || nambr>'9'){
			return  false
		}	
	}
	return true
}