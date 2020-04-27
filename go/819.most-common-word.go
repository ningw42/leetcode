func mostCommonWord(paragraph string, banned []string) string {
    var builder strings.Builder
    bannedDict := make(map[string]struct{})
    for _, bannedWord := range banned {
        bannedDict[bannedWord] = struct{}{}
    }
    counter := make(map[string]int)
    target := ""
    for i := 0; i < len(paragraph); i++ {
        if inAlphabet, char := toLowerCase(paragraph[i]); !inAlphabet {
            // not in alphabet
            if builder.Len() > 0 {
                word := builder.String()
                if _, isBanned := bannedDict[word]; !isBanned {
                    counter[word]++
                    if target == "" {
                        target = word
                    } else {
                        if counter[word] > counter[target] {
                            target = word
                        }
                    }
                }
                builder.Reset()
            }
        } else {
            // in alphabet
            builder.WriteByte(char)
        }
    }
    if builder.Len() > 0 {
                word := builder.String()
                if _, isBanned := bannedDict[word]; !isBanned {
                    counter[word]++
                    if target == "" {
                        target = word
                    } else {
                        if counter[word] > counter[target] {
                            target = word
                        }
                    }
                }
                builder.Reset()
    }
    
    return target
}

func toLowerCase(char byte) (bool, byte) {
    if 'a' <= char && char <= 'z' {
        return true, char
    }
    if 'A' <= char && char <= 'Z' {
        return true, char + 32
    }
    return false, 0
}