package hard

func longestValidParentheses(s string) int {
    if ( len(s) < 2) {return 0}
    res := 0
    stack := make([]int,0)
    stack = append(stack,-1)
    for i, x := range s {
        if (x == rune('(')){
            stack = append(stack,i)
        }else{
            stack = stack[:len(stack)-1]
            if (len(stack) > 0){
                res = Max(res, i - stack[len(stack)-1] )
            }else{
                stack = append(stack,i)
            }
        }
    }
    return res
}
func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}