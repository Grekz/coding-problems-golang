package easy

func twoSum(numbers []int, target int) []int {
    
    for i, j := 0, len(numbers) - 1; i < j;{
        sum := numbers[j] + numbers[i]
        if (sum < target){ 
            i+=1 
        }else if(sum > target){ 
            j-=1 
        }else {
            return []int{i+1,j+1}
        }
    }
    return []int{-1,-1}
}