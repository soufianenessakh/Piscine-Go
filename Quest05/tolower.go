package piscine

func ToLower(s string) string {
	output:=""
	for _,char:=range(s){
		if char>='A' && char<='Z'{
			char=(char+32)
		}
		output+=string(char)
	}
	return output
}