package medium

// @author grekz
func searchRange(nums []int, target int) []int {
    if len(nums) < 1 {
        return []int{-1,-1}
    }
    idx := bSearch(nums, target, true)
    if nums[idx % len(nums)] != target {
        return []int{-1,-1}
    }
    rIdx := bSearch(nums, target, false) - 1
    return []int{idx, rIdx}
}
func bSearch( n []int, t int, l bool ) int {
    a, b, m := 0, len(n), 0
    for a < b {
        m = a + ( b - a ) / 2
        if n[m] > t || ( l && n[m] == t ) {
            b = m
        }else{
            a = m + 1
        }
    }
    return a
}