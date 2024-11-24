func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }

    i, j := 0, 0 

    for ; j < len(nums); {
        if nums[i] != nums[j] {
            nums[i+1] = nums[j]
            i++
        }
        j++
    }

    // for _, val := range nums {
    //     fmt.Println(val)
    // }
    return i+1
}
