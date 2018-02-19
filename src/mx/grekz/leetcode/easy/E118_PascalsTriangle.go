package easy

func generate(numRows int) [][]int {
    res := make([][]int, 0,0)
    cur := make([]int, 0,0)
    for i := 0; i < numRows; i++ {
        for j := len(cur)-1; j > 0; j-- {
            cur[j] = cur[j] + cur[j-1]
        }
        cur = append(cur, 1)
        tmp := make([]int, len(cur))
        copy(tmp, cur)
        res = append(res, tmp)
    }
    return res
}