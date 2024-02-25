func smallerNumbersThanCurrent(nums []int) []int {
    var count []int = make([]int, 101);
    for i := 0; i < len(nums); i++ {
        count[nums[i]]++;
    }
    for i:= 1; i < len(count); i++ {
        count[i] += count[i-1];
    }
    
    for i, j := range nums{
        if j == 0{
            nums[i] = 0;
        }else{
            nums[i] = count[j-1];
        }   
    }
    return nums;
}