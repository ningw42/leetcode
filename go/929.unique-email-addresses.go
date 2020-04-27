func numUniqueEmails(emails []string) int {
    mails := make(map[string]bool)
    for _, email := range emails {
        parts := strings.Split(email, "@")
        local := parts[0]
        domain := parts[1]
        
        ppos := strings.Index(local, "+")
        if ppos != -1 {
            local = local[:ppos]
        }
        local = strings.Replace(local, ".", "", -1)
        
        mails[local + "@" + domain] = true
    }
    
    return len(mails)
}