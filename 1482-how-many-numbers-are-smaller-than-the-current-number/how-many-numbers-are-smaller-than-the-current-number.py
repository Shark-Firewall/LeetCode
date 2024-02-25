class Solution:
    def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:
        umap = [0]*101
        for i in nums:
            umap[i] += 1;
        
        for i in range(1, len(umap)):
            umap[i] += umap[i-1];

        for i in range(0, len(nums)):
            if nums[i] == 0:
                nums[i] = 0;
            else:
                nums[i] = umap[nums[i] - 1];
        return nums;
        