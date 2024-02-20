class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        umap = {ch:0 for ch in jewels};
        for i in stones:
            if i in umap:
                umap[i] += 1;
        return sum(umap.values());


        