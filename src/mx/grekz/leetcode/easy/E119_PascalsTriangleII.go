package easy

func getRow(rowIndex int) []int {
    cur := []int{1}
    for i := 0; i < rowIndex; i++ {
        for j := len(cur)-1; j > 0; j-- {
            cur[j] = cur[j] + cur[j-1]
        }
        cur = append(cur, 1)
    }
    return cur    
}