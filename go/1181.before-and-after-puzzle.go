func beforeAndAfterPuzzles(phrases []string) []string {
    // find first and last word of each phrase
    var firsts, lasts []string
    for i := 0; i < len(phrases); i++ {
        phrase := phrases[i]
        // first words
        if index := strings.Index(phrase, " "); index != -1 {
            firsts = append(firsts, phrase[:index])
        } else {
            firsts = append(firsts, phrase)
        }
        // last words
        if index := strings.LastIndex(phrase, " "); index != -1 {
            lasts = append(lasts, phrase[index+1:])
        } else {
            lasts = append(lasts, phrase)
        }
    }
   
    distinctRet := make(map[string]struct{})
    for i := 0; i < len(phrases); i++ {
        for j := 0; j < len(phrases); j++ {
            if i != j && lasts[i] == firsts[j] {
                merged := phrases[i] + phrases[j][len(firsts[j]):]
                distinctRet[merged] = struct{}{}
            }
        }
    }
  
    // sort
    ret := make([]string, len(distinctRet))
    i := 0
    for s := range distinctRet {
        ret[i] = s
        i++
    }
    sort.Strings(ret)
    return ret
}