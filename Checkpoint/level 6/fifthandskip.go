package checkpoint

func FifthAndSkip(str string) string {
	if str==""{
		return "\n"
	}
	if len(str)<5 {
		return "Invalid Input\n"
	}
	output:=""
	count:=0
	for i,c:=range(str){
		if c!=' '{
			count++
		}else if c==' '{
			continue
		}
		if count==6{
			count=0
			continue
		}
		output+=string(c)
		if count==5{
			for j := i + 1; j < len(str); j++ {
				if str[j] != ' ' {
					output += " "
					break
				}
				
			}
			
		}
	}
	return output 
}