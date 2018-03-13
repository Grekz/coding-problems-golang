package easy

import "strconv"

func fizzBuzz(n int) []string {
    res := make([]string, 0)
    for i := 1; i <= n; i++ {
        x := strconv.Itoa(i)
        if (i % 3 == 0 && i % 5 == 0) {
            x = "FizzBuzz"
        } else if (i % 3 == 0) {
            x = "Fizz"
        } else if (i % 5 == 0) {
            x = "Buzz"
        }
        res = append(res, x)
    }
    return res
}