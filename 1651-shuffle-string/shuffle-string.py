class Solution:
    def restoreString(self, s: str, indices: List[int]) -> str:
        ans = [""]*len(s);
        count = 0;
        for i in indices:
            ans[i] = s[count]
            count += 1;
        return "".join(ans);