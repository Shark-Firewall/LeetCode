class Solution:
    def countMatches(self, items: List[List[str]], ruleKey: str, ruleValue: str) -> int:
        count = 0
        for i in items:
            if ruleKey == "type":
                count += 1 if i[0] == ruleValue else 0;
            elif ruleKey == "color":
                count += 1 if i[1] == ruleValue else 0;
            elif ruleKey == "name":
                count += 1 if i[2] == ruleValue else 0;
        return count;
        