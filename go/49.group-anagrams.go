
type Bytes []byte
func (b Bytes) Len() int {
	return len(b)
}
func (b Bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b Bytes) Less(i, j int) bool {
	return b[i] < b[j]
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		bs := Bytes([]byte(str))
		sort.Sort(bs)
		s := string(bs)
		anagrams[s] = append(anagrams[s], str)
	}

	var ret [][]string
	for _, s := range anagrams {
		ret = append(ret, s)
	}
	return ret
}