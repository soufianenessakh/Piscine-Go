package piscine

func Compare(a, b string) int {
	output:=0
if a>b{
	return 1
}
if a==b{
	return 0
}
if a<b{
	return -1
}
return output
}