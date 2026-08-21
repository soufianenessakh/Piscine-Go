package checkpoint

func CountChar(str string, c rune) int {
    count:=0
	for _,char:=range str{
		if c==char{
			count++
		}
	} 
	return count
}