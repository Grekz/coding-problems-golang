package easy

// @author grekz
func moveZeroes(nums []int)  {
    pointer := 0
    for i, e := range nums {
        if e != 0 {
            tmp := e
            nums[i] = nums[pointer]
            nums[pointer] = tmp
            pointer += 1
        }
    }
}