package easy

// @author grekz
func intersect(nums1 []int, nums2 []int) []int {
    m1 := map[int]int{}
    for _, e := range nums1 {
        if x, ok:= m1[e]; ok{
            m1[e] = x + 1
        }else{
            m1[e] = 1
        }
    }
    res := []int{}
    for _, e := range nums2 {
        if x, ok:= m1[e]; ok && x > 0{
            m1[e] = x - 1
            res = append(res, e)
        }
    }
    return res
}