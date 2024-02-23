class Solution:
    def balancedStringSplit(self, s: str) -> int:
        Bal = 0
        count = 0;
        for word in s:
            if word == 'R':
                Bal += 1;
            else:
                Bal -= 1;
            if Bal == 0:
                count += 1;

        return count;
        