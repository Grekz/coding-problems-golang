package easy

// @author grekz
func findShortestSubArray(nums []int) int {
    counts, ini := make(map[int]int), make(map[int]int)
    res, max := 0, 0
    for i, x := range nums {
        if _, ok := ini[x]; !ok {
            ini[x] = i
        }
        cur, ok := counts[x]
        if !ok {
            cur = 0
        }
        cur += 1
        counts[x] = cur
        if max < cur {
            max = cur
            res = i - ini[x] + 1
        }else if max == cur {
            tmp := i - ini[x] + 1
            if res > tmp {
                res = tmp
            }
        }
    }
    return res
}