class Solution:
    def balancedStringSplit(self, s: str) -> int:
        Lword = 0;
        Rword = 0;
        count = 0;
        for word in s:
            if word == 'R':
                Rword += 1;
            elif word == 'L':
                Lword += 1;
            if Lword == Rword:
                count += 1;
        return count;
        