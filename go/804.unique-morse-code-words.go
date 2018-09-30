func uniqueMorseRepresentations(words []string) int {
	dict := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	transformations := make(map[string]bool)
	for _, word := range words {
		code := ""
		for _, r := range word {
			code = code + dict[byte(r)-97]
		}
		transformations[code] = true
	}

	return len(transformations)
}