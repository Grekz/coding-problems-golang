package medium

// @author grekz
func singleNonDuplicate(nums []int) int {
    i, j := 0, len(nums) - 1
    for i < j {
        mid := i + ( j - i ) / 2
        if nums[mid] == nums[mid ^ 1]{
            i = mid + 1
        }else{
            j = mid
        }
    }
    return nums[i]
}