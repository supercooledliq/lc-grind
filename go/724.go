func pivotIndex(nums []int) int {
    sum := 0 
    for _, val := range nums {
        sum += val 
    }

    leftSum := 0 
    for i, val := range nums {
        if leftSum == sum-leftSum-val {
            return i
        }
        leftSum += val
    }
    return -1 
}

// The core idea is:
// - We need to eventually check if the leftSum becomes equal to rightSum
// - And we can get the rightSum on the fly
// - We get the sum of all the elements in the array 
// - Then we keep getting the leftSum as we iterate the array
// - At each point, rightSum = sum-leftSum-currentElement 
// - If this becomes equal to current leftSum, then we know it is the answer
// 
// What I learned? 
// - "Online" calculation of prefix sums
