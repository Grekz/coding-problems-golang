package easy

import "regexp"

func numJewelsInStones(J string, S string) int {
    return len(regexp.MustCompile("[^"+J+"]").ReplaceAllString(S, ""))
}