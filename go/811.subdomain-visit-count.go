func subdomainVisits(cpdomains []string) []string {
    counter := make(map[string]int)
    for _, domainCount := range cpdomains {
        tmp := strings.Split(domainCount, " ")
        count, domain := tmp[0], tmp[1]
        numericCount, _ := strconv.Atoi(count)
        for _, relatedDomain := range allDomains(domain) {
            counter[relatedDomain] += numericCount
        }
    }
    
    var ret []string
    for domain, count := range counter {
        ret = append(ret, strconv.Itoa(count) + " " + domain)
    }
    
    return ret
}

func allDomains(domain string) []string {
    ret := []string{domain}
    for cursor := len(domain) - 1; cursor >= 0; cursor-- {
        if domain[cursor] == '.' {
            ret = append(ret, domain[cursor+1:])
        }
    }
    return ret
}