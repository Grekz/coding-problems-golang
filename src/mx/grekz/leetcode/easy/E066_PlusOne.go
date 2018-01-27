package easy

func plusOne(digits []int) []int {
    le := len(digits)
    for i := le - 1; i >= 0; i--{
        if ( digits[i] < 9 ){
            digits[i] += 1
            return digits
        }
        digits[i] = 0
    }
    digits = append([]int{1}, digits...)
    return digits  
}