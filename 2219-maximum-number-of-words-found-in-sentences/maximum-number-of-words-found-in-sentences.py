class Solution:
    def mostWordsFound(self, sentences: List[str]) -> int:
        max = 0;
        for sentence in sentences:
            count = 0;
            for word in sentence:
                if word == " ":
                    count += 1;
            if count > max:
                max = count;
        return max+1;