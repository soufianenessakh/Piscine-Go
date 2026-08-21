package level02


func RectPerimeter(w, h int)int{
	output:=0
	if w<0 || h<0{
		return -1
	}
		output=(w+h)*2
	
	return output
}
