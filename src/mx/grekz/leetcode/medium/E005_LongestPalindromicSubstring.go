package medium

func longestPalindrome(s string) string {
    t := make([]rune, len(s)*2 + 3)
    lt := len(t)
    p := make([]int, lt)
    t[0] = '$'
    t[1] = '#'
    for i, x := range s {
        t[i*2 + 2] = x
        t[i*2 + 3] = '#'
    }
    t[lt-1] = '%'
    // fmt.Println(t)
    c, r, mc, mp, mir := 0, 0, 0, 0, 0
    for i := 1; i < lt - 1; i++ {
        if i < r {
            mir = 2*c - i
            if r-i < p[mir] {
                p[i] = r - i   
            }else{
                p[i] = p[mir]
            }
        }
        for t[i-(p[i]+1)] == t[i+(p[i]+1)]{
            p[i] += 1
        }
        if i + p[i] > r {
            c = i
            r = i + p[i]
            if p[i] > mp {
                mc = c
                mp = p[i]
            }
        }
    }
    st := (mc - mp) / 2
    en := (mc + mp) / 2
    return s[st:en]
}

