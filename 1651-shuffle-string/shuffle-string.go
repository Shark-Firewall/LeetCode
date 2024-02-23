func restoreString(s string, indices []int) string {
    ans := make([]byte, len(s))
    for j, i := range indices{
        ans[i] = s[j]
    }
    return string(ans);
}