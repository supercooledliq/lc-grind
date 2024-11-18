func twoSum(nums []int, target int) []int {
    cache := make(map[int]int)

    for i, j := range nums {
        if k, ok := cache[j]; ok {
            return []int{i, k}
        }

        cache[target-j] = i
    }

    // fmt.Println("Cache has: %+v", cache)
    return []int{}
}
