package easy

// @author grekz
func search(nums []int, target int) int {
    return bSearch(nums, target, 0, len(nums) - 1)
}
func bSearch(nums []int, target, ini, end int) int {
    if ini > end {
        return -1
    }
    mid := ini + (end - ini) / 2
    if nums[mid] == target {
        return mid
    }
    if nums[mid] > target {
        return bSearch(nums, target, ini, mid - 1)
    }
    return bSearch(nums, target, mid + 1, end)
}