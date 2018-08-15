package easy

// @author grekz
import "strings"

func uncommonFromSentences(A string, B string) []string {
    tmp := strings.Split( A + " " + B, " ")
    cnt := map[string]int{}
    for _, x := range tmp {
        if cur, ok := cnt[x]; ok {
            cnt[x] = cur + 1
        }else{
            cnt[x] = 1             
        }
    }
    res := make([]string, 0)
    for k, v := range cnt {
        if v == 1 {
            res = append(res, k)        
        }
    }
    return res
}