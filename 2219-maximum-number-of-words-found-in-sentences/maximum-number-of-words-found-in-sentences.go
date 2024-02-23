func maxCount(a, b int) int{
    if a >= b{
        return a;
    }
    return b;
}

func mostWordsFound(sentences []string) int {
    max := 0;
    for _, words := range sentences{
        count := len(strings.Split(words, " "));
        max = maxCount(max, count); 
    }
    return max;
}