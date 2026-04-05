package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
sign:= 1
output:=0
star:=0
if s[0]=='-'{
	sign=-1
	star=1
}else if s[0]== '+'{
	star=1
}
if star == len(s) {
	return 0
}
 for i:=star ; i<len(s);i++ {
if s[i] <'0' || s[i] >'9'{
	return 0
}
output = output*10+(int(s[i])-'0')
} 
return output*sign
}