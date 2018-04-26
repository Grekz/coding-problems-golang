package easy

// @author grekz
func numberOfLines(widths []int, S string) []int {
    lastLength := 0
    lines := 0
    for _, c := range S {
        curWidth := widths[c - 97];
        if (lastLength + curWidth > 100){
            lines += 1
            lastLength = curWidth
        } else {
            lastLength += curWidth
        }
    }
    if ( lastLength < 100 ) {
        lines += 1
    }
    return []int{lines, lastLength}
}