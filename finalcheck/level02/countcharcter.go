package level02


func Countchar(str string, c rune)int{
	count:=0
	for  _,char:=range(str) {
		if char == c{
			count++
		}
	}
	return count
}