package piscine

func StrRev(s string) string {
output := ""
for i := len (s)-1; i >= 0 ; i --{
	output += string(s[i])
}
return output
}