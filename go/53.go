func getCrossing(nums []int, left, mid, right int) int {
    sum := 0 
    leftSum := math.MinInt32
    for i := mid; i >= left; i-- {
        sum += nums[i]
        if sum > leftSum {
            leftSum = sum 
        }
    }

    sum = 0 
    rightSum := math.MinInt32
    for i := mid+1; i <= right; i++ {
        sum += nums[i]
        if sum > rightSum {
            rightSum = sum
        }
    }

    return leftSum+rightSum
}

func strike(nums []int, left, right int) int {
    if left == right {
        return nums[left]
    }

    mid := (left+right)/2
    leftSum := strike(nums, left, mid)
    rightSum := strike(nums, mid+1, right)
    crossing := getCrossing(nums, left, mid, right)

    // fmt.Println("---")
    // fmt.Println("Left sum: ", leftSum)
    // fmt.Println("Right sum: ", rightSum)
    // fmt.Println("Crossing: ", crossing)
    return max(max(leftSum, rightSum), crossing)
}

func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    return strike(nums, 0, len(nums)-1)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// ---


