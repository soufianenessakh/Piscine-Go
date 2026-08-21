package level02


func RetainFirstHalf(str string)string{
	if str=="" || len(str)==1{
		return str
	}
	half:=len(str)/2
	return str[:half]
}