package checkpoint

func HashCode(dec string) string {
	output:=""
	size:=len(dec)
	for _,c:=range dec{
		newchar:=(int(c)+size)%127
		if newchar<33{
			newchar+=33
		}
		output+=string(rune(newchar))
	}
	return output
}