package Quest10

func ShoppingSummaryCounter(str string) map[string]int {
	word:=""
	summary := make(map[string]int)
	for  _,char:=range str{
		if char!=' '{
			word+=string(char)
		}else if char==' '{
			summary[word]++
			word=""
		}
	}
	if word!=""{
		summary[word]++
	}
	return summary
}