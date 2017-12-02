package easy

func twoSum(nums []int, target int) []int {
    tmp := make(map[int]int)
    for i, v := range nums {
        if x, ok := tmp[target - v]; ok {
            return []int{x, i}
        }
        tmp[v] = i
        i++
    }		
    return []int{0,0}
}