func balancedStringSplit(s string) int {
    bal := 0;
    count := 0;
    for _, j := range s{
        if string(j) == "R"{
            bal++;
        }else{
            bal--;
        }
        if bal == 0{
            count++;
        }
    }
    return count;
}