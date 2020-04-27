var partitions [][]string
func partition(s string) [][]string {
	partitions = nil
	backtracking([]byte(s), []string{})
	return partitions
}

func backtracking(s []byte, partition []string) {
	if len(s) == 0 {
		partitions = append(partitions, makeSliceCopy(partition))
	} else {
		for i := 1; i <= len(s); i++ {
			if isPalindrome(s[0:i]) {
				backtracking(s[i:], append(partition, string(s[0:i])))
			}
		}
	}
}

func makeSliceCopy(s []string) []string {
	ret := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		ret[i] = s[i]
	}
	return ret
}

func isPalindrome(s []byte) bool {
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}