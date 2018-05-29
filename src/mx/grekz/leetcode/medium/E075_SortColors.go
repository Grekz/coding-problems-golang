package medium

// @author grekz
func sortColors(nums []int)  {
    j, k := 0, len(nums) - 1
    for i := 0; i <= k; i++ {
        if nums[i] == 0 {
            swap(nums, i, j)
            j += 1
        }else if nums[i] == 2 {
            swap(nums, i, k)
            k -= 1
            i -= 1
        }
    }
}
func swap( nums []int, a, b int) {
    tmp := nums[b]
    nums[b] = nums[a]
    nums[a] = tmp
}