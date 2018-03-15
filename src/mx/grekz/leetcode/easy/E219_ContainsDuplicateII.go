package easy

func containsNearbyDuplicate(nums []int, k int) bool {
    if(k == 0 || nums == nil || len(nums) < 2 ) { return false }
    values := map[int]int{}
    for i, x := range nums {
        if cur, ok := values[x]; ok && i - cur <= k {
            return true
        }
        values[x] = i
    }
    return false
}