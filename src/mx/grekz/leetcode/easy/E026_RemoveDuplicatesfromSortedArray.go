package easy

func removeDuplicates(nums []int) int {
    if (nums == nil || len(nums) < 1 ) { return 0; }
    if ( len(nums) == 1 ) { return 1; }
    le, base, count, cur := len(nums), nums[0], 0, 1
    for i := 1; i < le ; i++ {
        if ( base == nums[i] ) {
            count+=1;
            nums[i] = -1;
        }else{
            base = nums[i];
            nums[cur] = nums[i];
            cur += 1;
        }
    }
    return le - count;
}