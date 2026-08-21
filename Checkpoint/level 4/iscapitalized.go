package checkpoint

func IsCapitalized(s string) bool {
	if len(s)==0{
		return false
	}
	wordstar:=true
	for  _,c:=range (s){
		if wordstar{
			if c>='a'&& c<='z'{
				return false
			}
			wordstar=false
		}	
		if c==' '{
			wordstar=true
		}
	}
	return true
}