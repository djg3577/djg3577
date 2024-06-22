func longestConsecutive(nums []int) int {

    set := make(map[int]struct{})
    longest := 0

    
    for _, number := range nums{
        set[number] = struct{}{}
    }

    for _, number := range nums{
        if _, exists := set[number-1]; !exists{
            length:= 1
            current := number +1

            for {
                if _, exists := set[current]; exists{
                    length++
                    current++
                } else{
                    break
                }
            }
            longest = max(longest, length)
        }
    }
    return longest

}