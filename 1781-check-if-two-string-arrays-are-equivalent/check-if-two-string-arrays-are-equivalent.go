func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    string1 := "";
    string2 := "";
    for _, str := range word1{
        string1 += str;
    }
    for _, str := range word2{
        string2 += str;
    }
    if string1 == string2{
        return true;
    }
    return false;
}