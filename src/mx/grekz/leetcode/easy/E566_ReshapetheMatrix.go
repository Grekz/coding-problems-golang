package easy

// @author grekz
func matrixReshape(nums [][]int, r int, c int) [][]int {
    m := len(nums)
    n := len(nums[0])
    if n * m != r * c {
        return nums
    }
    //Create multidimensional array golang
    res := make([][]int, r)
    i := 0
    for j := range res {
        res[j] = make([]int, c)
    }
    for _, curArr := range nums {
        for _, cur := range curArr {
            res[i/c][i%c] = cur
            i += 1
        }
    }
    return res
}