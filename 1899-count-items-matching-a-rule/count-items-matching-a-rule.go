func ternary(a, b string) int{
    if a == b{
        return 1
    }
    return 0;
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    count := 0;
    for _, item := range items{
        if ruleKey == "type"{
            count += ternary(item[0], ruleValue);
        }else if ruleKey == "color"{
            count += ternary(item[1], ruleValue);
        }else{
            count += ternary(item[2], ruleValue);
        }
    }
    return count;
}