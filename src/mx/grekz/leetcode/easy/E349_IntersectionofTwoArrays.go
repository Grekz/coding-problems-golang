package easy

// @author grekz
func intersection(nums1 []int, nums2 []int) []int {
    if len(nums1) < 1 || len(nums2) < 1 { return []int{} }
    m := mapFromArray(nums1)
    res := []int{}
    for _, x := range nums2 {
        if _, ok := m[x]; ok{
            m[x] = true
        }        
    }
    for k, v := range m { 
        if v {
            res = append(res, k)
        }
    }
    return res
}
func mapFromArray(a []int) map[int]bool{
    m := map[int]bool{}
    for _, x := range a {
        m[x] = false
    }
    return m
}