func moveZeroes(nums []int) {
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			flag := false
			for j := 0; j < len(nums); j++ {
				if nums[j] == 0 && j < i {
					zero = j
					flag = true
					break
				}
			}
			if flag {
				nums[zero], nums[i] = nums[i], nums[zero]
			}
		}
	}
}

// Jugaad ad-hoc solution:
// - Track the last seen zero basically
// - When you see the element that is non zero
//     - Look for the zero that you saw last time (this can be cached if you think hard enough)
//     - Now we need to make sure that the elements are actually to be swapped
//         - We track via flag and also see if actually 'j' < 'i', that means, 0 does occur before a non-0 element
//         - In case flag is set
//         - Swap
//     - Swap and move on

func moveZeroes(nums []int) {
	zero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != zero {
				nums[zero], nums[i] = nums[i], nums[zero]
			}
			zero++
		}
	}
}

// Optimized solution:
// Keep track of where the next non-zero element should go. 
