/*
 * @lc app=leetcode id=721 lang=golang
 *
 * [721] Accounts Merge
 */

// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
	var mails []string
	var mailBelongTo []int
	var personSets []int
	for personID, account := range accounts {
		personSets = append(personSets, personID)
		for _, mail := range account[1:] {
			mails = append(mails, mail)
			mailBelongTo = append(mailBelongTo, personID)
		}
	}

	// find same mail
	aggregatedMails := make(map[string][]int)
	for mailID, mail := range mails {
		aggregatedMails[mail] = append(aggregatedMails[mail], mailID)
	}

	// union-find
	for _, mailIDs := range aggregatedMails {
		if len(mailIDs) > 1 {
			for i := 0; i < len(mailIDs)-1; i++ {
				personA := mailBelongTo[mailIDs[i]]
				personB := mailBelongTo[mailIDs[i+1]]
				personAName := accounts[personA][0]
				personBName := accounts[personB][0]
				if personAName == personBName {
					personARoot := find(personSets, personA)
					personBRoot := find(personSets, personB)
					if personARoot != personBRoot {
						union(personSets, personARoot, personBRoot)
					}
				}
			}
		}
	}

	result := make(map[int]map[string]bool)
	for personID := range accounts {
		root := find(personSets, personID)
		if _, exists := result[root]; !exists {
			result[root] = make(map[string]bool)
		}
		for _, mail := range accounts[personID][1:] {
			result[root][mail] = true
		}
	}

	var ret [][]string
	for personID, mailSet := range result {
		row := []string{accounts[personID][0]}
		var mails []string
		for mail := range mailSet {
			mails = append(mails, mail)
		}
		sort.Strings(mails)
		row = append(row, mails...)
		ret = append(ret, row)
	}

	return ret
}

func find(sets []int, id int) int {
	for id != sets[id] {
		id = sets[id]
	}
	return id
}

func union(sets []int, a, b int) {
	sets[a] = b
}
// @lc code=end

