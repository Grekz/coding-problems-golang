package easy

func rob(nums []int) int {
    a, b := 0, 0
    for i, x := range nums {
        if i % 2 == 0 {
            a = Max(a+x, b)
        }else{
            b = Max(a, b+x)
        }
    }
    return Max(a,b)
}
func Max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}