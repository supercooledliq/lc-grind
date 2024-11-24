func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	j := 1
	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			k++
		} else {
			k = 1
		}

		if k <= 2 {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

// Here we are reading using `j` writing using `i` and counting using `k`. 
