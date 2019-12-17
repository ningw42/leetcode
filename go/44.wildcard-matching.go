/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	// return matchByBackTracking(s, p)
	return matchByFiniteStateMachine(s, p)
}

// https://leetcode.com/problems/wildcard-matching/discuss/138878/
func matchByFiniteStateMachine(s, p string) bool {
	// 1. build up the finite state machine
	// state-input-next_state
	fsm := make(map[int]map[byte]int)
	var state int
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			if _, ok := fsm[state]; !ok {
				fsm[state] = map[byte]int{p[i]: state}
			}
		} else {
			if _, ok := fsm[state]; !ok {
				fsm[state] = map[byte]int{p[i]: state+1}
			} else {
				fsm[state][p[i]] = state + 1
			}
			state++
		}
	}

	// 2. set accept state
	acceptState := state

	// 3. apply finite state machine to s
	// 3.1 set initial state
	currentStates := map[int]bool{0: true}
	for i := 0; i < len(s); i++ {
		nextStates := make(map[int]bool)
		// every possible state
		for state := range currentStates {
			// every matched token
			for _, matchedToken := range []byte{s[i], '*', '?'} {
				// if state transfer path exists
				if nextState, ok := fsm[state][matchedToken]; ok {
					// collect all possible next states
					nextStates[nextState] = true
				}
			}
		}
		currentStates = nextStates
	}

	_, accept := currentStates[acceptState]
	return accept
}

// naive backtracking wont work, time limit exceeded
func matchByBackTracking(s, p string) bool {
	return backtracking(s, removeContinuousKleeneStar(p))
}

func removeContinuousKleeneStar(s string) string {
	var ret []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if i == 0 || s[i-1] != '*' {
				ret = append(ret, '*')
			}
		} else {
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}

func backtracking(s, p string) bool {
	if s == "" && p == "" {
		return true
	} else if s != "" && p == "" {
		return false
	} else if s == "" && p != "" {
		token := nextToken(p)
		if token == '*' {
			for i := 0; i < len(p); i++ {
				if p[i] != '*' {
					return false
				}
			}
			return true
		} else {
			return false
		}
	} else {
		token := nextToken(p)
		switch token {
		case '?':
			return backtracking(s[1:], p[1:])
		case '*':
			for i := 0; i <= len(s); i++ {
				if backtracking(s[i:], p[1:]) {
					return true
				}
			}
			return false
		default:
			return s[0] == p[0] && backtracking(s[1:], p[1:])
		}
	}
}

func nextToken(p string) byte {
	return p[0]
}
// @lc code=end

