package checkpoint

func SaveAndMiss(arg string, num int) string {
	if num<=0{
		return arg
	}
	output:=""
	group:=0

	for i:=0;i<len(arg);i+=num{
		end:=i+num
		if end>len(arg){
			end=len(arg)
		}
		set:=arg[i:end]
		if group%2==0{
			output+=set
		}
		group++

	}
	return output
}