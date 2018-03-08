package easy

func containsDuplicate(nums []int) bool {
    tmp := map[int]bool{}
    for _, x := range nums {
        tmp[x] = true
    }
    return len(nums) != len(tmp)
}