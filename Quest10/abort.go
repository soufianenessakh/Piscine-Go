package Quest10

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	for i := 0; i < 5; i++ {
		smaller := 0
		bigger := 0

		for j := 0; j < 5; j++ {
			if nums[j] < nums[i] {
				smaller++
			} else if nums[j] > nums[i] {
				bigger++
			}
		}

		if smaller == 2 && bigger == 2 {
			return nums[i]
		}
	}

	return 0
}