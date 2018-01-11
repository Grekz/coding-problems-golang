package medium

func generateParenthesis(n int) []string {
    res := make([]string,0)
    btHelper("",n,n,&res)
    return res
}

func btHelper(st string, ope int, clo int, res *[]string) *[]string {
    if ( ope > 0) {
        btHelper(st + "(", ope - 1, clo, res)
    }
    if ( clo > ope ) {
        btHelper(st + ")", ope, clo - 1, res)
    }
    if ( clo < 1 ) {
        *res = append(*res, st)
    }
    return res
}