func searchRange(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	lb := -1
	rb := -1

	for l <= r {
		m := int((l + r) / 2)

		if nums[m] == target {
			lb = m
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	l = 0
	r = len(nums) - 1
	for l <= r {
		m := int((l + r) / 2)

		if nums[m] == target {
			rb = m
			l = m + 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return []int{lb, rb}
}