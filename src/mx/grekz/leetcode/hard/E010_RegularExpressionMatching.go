package hard

// @author grekz
func isMatch(s string, p string) bool {
    tn, pn := len(s), len(p)
    dp := make([][]bool, tn + 1)
    for i := range dp {
        dp[i] = make([]bool, pn + 1)
    }
    dp[tn][pn] = true
    for i := tn ; i >= 0; i-- { 
        for j := pn - 1; j >= 0; j-- { 
            fm := i < tn && ( s[i] == p[j] || p[j] == '.' )
            if j + 1 < pn && p[j+1] == '*' {
                dp[i][j] = dp[i][j+2] || fm && dp[i+1][j]
            }else{
                dp[i][j] = fm && dp[i+1][j+1]
            }
        }
    }
    return dp[0][0]
}