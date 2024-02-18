func runningSum(nums []int) []int {
    var ans []int = make([]int, 0);
    preSum := 0;
    for _, num := range nums{
        preSum += num;
        ans = append(ans, preSum);
    }
    return ans;
}