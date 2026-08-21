package checkpoint

func Gcd(a, b uint) uint {
	if a==0 || b==0{
		return 0
	}
	for b!=0{
		temp:=a%b
			a=b
			b=temp
	}
	return a
}