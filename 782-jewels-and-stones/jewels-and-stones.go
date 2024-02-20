func numJewelsInStones(jewels string, stones string) int {
    var umap = make(map[rune]int);
    for _, i := range jewels {
        umap[i] = 0
    }
    count := 0;
    for _, i := range stones {
        _ , tmp := umap[i]
        if tmp{
            umap[i]++;
            count++;
        }
    }
    return count;
}