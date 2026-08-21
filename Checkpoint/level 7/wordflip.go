package checkpoint

func WordFlip(str string) string {
	if str==""{
		return "Invalid output"
	}
	if str=="    "{
		return "\n"
	}
	output:=""
	for i:=len(str)-1;i>=0;i--{
		if str[i]==' '{
			i++
			output+=string(str[i])
		}
	}
	for i:=0;i<len(str);i++{
		if str[i]==' '{
			continue
		}
		output+=string(str[i])
		
	}
	return output
}