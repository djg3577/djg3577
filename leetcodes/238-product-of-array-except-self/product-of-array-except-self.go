func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))

    prefix := 1
    for index, number := range nums {
        res[index] = prefix
        prefix *= number
    }

    postfix := 1
    for i := len(nums) - 1 ; i >= 0; i-- {
        res[i] *= postfix
        postfix *= nums[i]
    }

    return res
}