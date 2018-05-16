package easy

// @author grekz
func nextGreaterElement(findNums []int, nums []int) []int {
    m := make(map[int]int)
    s := make([]int, 0)
    for _, x := range nums {
        for len(s) > 0 && s[len(s)-1] < x {
            m[s[len(s)-1]] = x
            s = s[0: len(s) -1]
        }
        s = append(s, x)
    }
    for i, x := range findNums {
        findNums[i] = get(m, x)
    }
    return findNums
}
func get(m map[int]int, k int) int {
    if value, ok := m[k]; ok {
        return value
    }
    return -1
}