package piscine

func ToUpper(s string) string {
	output:=""
 for _,char:=range(s){
	if char>='a'&& char<='z'{
		char= (char-32)
	}
	output += string(char)
 }
 return output
}