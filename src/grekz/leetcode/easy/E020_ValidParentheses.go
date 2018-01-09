package easy

import "strings"

func isValid(s string) bool {
    if ( len(s) < 1 ) {return true}
    if ( len(s) == 1 ) {return false}
    ope := "({["
    clo := ")}]"
    stack := make([]rune, 0)
    
    for _,cur := range s {
        if (strings.Index(ope, string(cur)) != -1 ) {
            stack = append(stack, cur)
        }else {
            leSt := len(stack)
            if (leSt < 1) { return false }
            popped := string(stack[leSt-1])
            curSt := string(cur)
            stack = stack[:leSt-1]
            if ( strings.Index(ope, popped) != strings.Index(clo, curSt) ){
                return false
            }
        }
    }
    return len(stack) == 0
}