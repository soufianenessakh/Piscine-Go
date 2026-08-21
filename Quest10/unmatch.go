package Quest10

func Unmatch(a []int) int {
	
	for i:=0;i<=len(a)-1;i++{
		count:=0
		for j:=0;j<=len(a)-1;j++{
			if a[i]==a[j] {
			count++
			}
		}
		if count==1{
			return a[i]
		}
	}
	return -1
}