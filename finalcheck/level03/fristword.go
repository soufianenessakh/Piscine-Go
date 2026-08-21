package level03

func FirstWord(s string) string {
    if s==""{
		return s
	}
	output:=""
	for i:=0 ;i<len(s);i++{
		if s[i]==' ' && output==""{
			continue
		}
		if s[i]!=' ' {
			output+=string(s[i])
		}
		if s[i]==' ' && output!=""{
			break
		}
	}
	return output
}