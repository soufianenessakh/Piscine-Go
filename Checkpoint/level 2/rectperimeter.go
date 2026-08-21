package checkpoint

func RectPerimeter(w, h int) int {
	output:=0
	if w<0 || h<0{
		output=-1
		return output
	}
	output=(w+h)*2
	return output
}