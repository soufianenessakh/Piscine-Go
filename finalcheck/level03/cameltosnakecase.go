package level03

func CamelToSnakeCase(s string) string{
	
	if (s[len(s)-1]>='A' && s[len(s)-1]<='Z')||(s[len(s)-1]>='0'&&s[len(s)-1]<='9'){
		return s
	}
	output:=""
	for i:=0;i<len(s);i++{
		if s[i]>='A'&& s[i]<='Z'{
			if i!=0 {
				output+="_"
			}	
		}
		output+=string(s[i])
	}
	return output
}