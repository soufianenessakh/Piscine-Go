package piscine

func AlphaCount(s string) int {
n:=0
	for _,char:= range(s){
		if (char>='a' && char<='z') ||(char>='A' && char<='Z'){
			n++
		}
	}
	return n
}