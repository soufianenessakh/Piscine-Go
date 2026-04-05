package piscine

func BasicAtoi2(s string) int {
output := 0
for _, char := range s{
if char<'0' || char>'9'{
	return 0
}
output = output*10+(int(char)-'0')
} 
return output
}