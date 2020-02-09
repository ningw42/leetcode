/*
 * @lc app=leetcode id=336 lang=golang
 *
 * [336] Palindrome Pairs
 */

// @lc code=start
func palindromePairs(words []string) [][]int {
	// return naive(words)
	return trie(words)
}

type Node struct {
	Children map[byte]*Node
	Value string
	End bool
}

func (node *Node) AddWord(word []byte, index int) {
	if index == len(word) {
		node.End = true
	} else {
		char := word[index]
		if _, exists := node.Children[char]; !exists {
			node.Children[char] = &Node{
				Children: make(map[byte]*Node),
				Value: string(word[:index+1]),
				End: false,
			}
		}
		node.Children[char].AddWord(word, index+1)
	}
}

func (node *Node) Traverse(prefix string) []string {
	var result []string
	if node.End {
		result = append(result, prefix)
	}
	for char, next := range node.Children {
		result = append(result, next.Traverse(prefix+string([]byte{char}))...)
	}
	return result
}

// search reversed word among all words using trie
func search(root *Node, word string) []string {
	node := root
	var result []string
	var i int
	for i = 0; i < len(word); i++ {
		// 1. the given word is longer than the paired word
		if node.End && isPalindrome(word[i:]) {
			result = append(result, node.Value)
		}
		char := word[i]
		if next, exists := node.Children[char]; exists {
			node = next
		} else {
			// 2. not matches
			break
		}
	}
	// 3. the given word is shorter than the paired word
	if i == len(word) {
		for _, restOfWord := range node.Traverse("") {
			if isPalindrome(restOfWord) {
				result = append(result, word + restOfWord)
			}
		}
	}

	return result
}

// use trie to search answer
func trie(words []string) [][]int {
	// build trie
	mapping := make(map[string]int)
	root := &Node{
		Children: make(map[byte]*Node),
		Value: "",
		End: false,
	}
	for i, word := range words {
		root.AddWord([]byte(word), 0)
		mapping[word] = i
	}

	// search for pairs
	var pairs [][]int
	for _, word := range words {
		reversedWord := reverse(word)
		pairWords := search(root, reversedWord)
		for _, pairWord := range pairWords {
			if pairWord != word {
				pairs = append(pairs, []int{mapping[pairWord], mapping[word]})
			}
		}
	}

	return pairs
}

func reverse(s string) string {
	chars := []byte(s)
	for i := 0; i < len(chars) / 2; i++ {
		chars[i], chars[len(chars)-1-i] = chars[len(chars)-1-i], chars[i]
	}
	return string(chars)
}

// AC with poor time complexity
func naive(words []string) [][]int {
    var answers [][]int
    for i := 0; i < len(words); i++ {
        for j := i + 1; j < len(words); j++ {
            if w := words[i] + words[j]; isPalindrome(w) {
                answers = append(answers, []int{i, j})
            }
            if w := words[j] + words[i]; isPalindrome(w) {
                answers = append(answers, []int{j, i})
            }
        }
    }
    return answers
}

func isPalindrome(s string) bool {
    for i, j := 0, len(s) - 1; i < j; {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}
// @lc code=end

