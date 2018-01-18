package easy

func removeElement(nums []int, val int) int {
    count := 0
    for i := 0; i < len(nums); i++{
        if(nums[i] != val){
            temp := nums[i]
            nums[i] = nums[count]
            nums[count] = temp
            count += 1
        }
    } 
    return count
}