package easy

func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    if ( n == 0 ) {return}
    i := m -1 
    j := n - 1
    ind := n + m  - 1
    for(ind >= 0 ){
        if(i < 0) {
            nums1[ind] = nums2[j]
            j -= 1
        }else if(j < 0){ 
            nums1[ind] = nums1[i]
            i-=1
        }else if(nums1[i] > nums2[j]){ 
            nums1[ind] = nums1[i]
            i-=1
        }else{ 
            nums1[ind] = nums2[j]
            j-=1
        }
        ind -=1
    }
}