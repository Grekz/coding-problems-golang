package medium

// @author grekz
func canVisitAllRooms(rooms [][]int) bool {
    seen := make([]bool, len(rooms))
    stack := []int{0}
    seen[0] = true
    for len(stack) != 0 {
        l := len(stack)
        cur := stack[l-1]
        stack = stack[:l-1]
        for _, x := range(rooms[cur]) {
            if !seen[x] {
                seen[x] = true
                stack = append(stack, x)
            }
        }
    }
    for _, x := range(seen) {
        if ( !x ) {
            return false
        }
    }
    return true
}