func isMatch(s string, p string) bool {
	return match(s, p)
}

func match(s, p string) bool {
	if s == "" && p == "" {
		return true
	} else if s != "" && p == "" {
		return false
	} else if s == "" && p != "" {
		// s is empty, all the rest p should be consist of wildcard token to match
		token := nextToken(p)
		if len(token) == 2 {
			return match(s, p[2:])
		} else {
			return false
		}
	} else {
		token := nextToken(p)
		if len(token) == 1 {
			// token with length 1
			if token == "." {
				// single wildcard token
				return match(s[1:], p[1:])
			} else {
				// single alphabet token
				return s[0] == token[0] && match(s[1:], p[1:])
			}
		} else {
			// token with length 2, "a*", ".*" ish
			if token[0] == '.' {
				// wildcard kleene star token
				for i := 0; i <= len(s); i++ {
					if match(s[i:], p[2:]) {
						return true
					}
				}
				return false
			} else {
				// alphabet kleene star token
				for i := 0; i <= len(s) && (i == 0 || s[i-1] == token[0]); i++ {
					if match(s[i:], p[2:]) {
						return true
					}
				}
				return false
			}
		}
	}
}

func nextToken(s string) string {
	if len(s) > 1 {
		if s[1] == '*' {
			return s[:2]
		} else {
			return s[:1]
		}
	} else {
		return s[:1]
	}
}