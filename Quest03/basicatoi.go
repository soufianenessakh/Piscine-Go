package piscine

func BasicAtoi(s string) int {
	output:=0
	for _,nb:=range(s){
		output=output*10+int(nb-'0')
	}
	return output
}