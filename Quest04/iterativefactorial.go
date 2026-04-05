package piscine

func IterativeFactorial(nb int) int {
if nb<0{
	return 0
}
output:=1
for i :=1; i<=nb;i++{
	output=output*i
}
return output
}