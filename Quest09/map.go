package Quest09


func Map(f func(int) bool, a []int) []bool {
	output:=make([]bool,len(a))
for i, nbr := range a {
		output[i]=f(nbr)
	}
	return output
}
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}