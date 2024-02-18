func numIdenticalPairs(nums []int) int {
    var count int = 0;
    var arr []int = make([]int, 101);
    for _, i := range nums{
        count += arr[i];
        arr[i]++;
    }
    return count;  
}