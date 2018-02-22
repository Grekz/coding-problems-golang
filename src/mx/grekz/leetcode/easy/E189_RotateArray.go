package easy

func rotate(nums []int, k int)  {
    n := len(nums)
    k %= n
    if n == 0 || k == 0 { return }
    reverse(nums, 0, n - 1)
    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
}
func reverse(nums []int, i int, j int) {
    for i < j {
        tmp := nums[i]
        nums[i] = nums[j]
        nums[j] = tmp
        i += 1
        j -= 1
    }
}