func countOfAtoms(formula string) string {
	count := DoCount(formula)
	var keys []string
	for k, _ := range count {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	ret := ""
	for _, k := range keys {
		if count[k] == 1 {
			ret = ret + k
		} else {
			ret = fmt.Sprintf("%s%s%d", ret, k, count[k])
		}
	}

	return ret
}

func DoCount(formula string) map[string]int {
	if len(formula) == 0 {
		return map[string]int{}
	}
	ret := make(map[string]int)
	f := []byte(formula)
	for len(f) > 0 {
		r := make(map[string]int)
		// find count
		digit := 1
		count := 0
		if '0' <= f[len(f)-1] && f[len(f)-1] <= '9' {
			i := len(f) - 1
			for i >= 0 && f[i] >= '0' && f[i] <= '9' {
				count = count + digit*int(f[i]-'0')
				digit = digit * 10
				i--
			}
			f = f[0 : i+1]
		} else {
			count = 1
		}
		// find token
		if f[len(f)-1] == ')' {
			i := len(f) - 2
			pc := 1
			for ; i >= 0 && pc > 0; i-- {
				if f[i] == ')' {
					pc++
				}
				if f[i] == '(' {
					pc--
				}
			}
			r = DoCount(string(f[i+2 : len(f)-1]))
			f = f[0 : i+1]
		} else {
			i := len(f) - 1
			for ; !('A' <= f[i] && f[i] <= 'Z'); i-- {
			}
			r = map[string]int{
				string(f[i:]): 1,
			}
			f = f[0:i]
		}
		// combine result
		for k, v := range r {
			if value, e := ret[k]; e {
				ret[k] = value + v*count
			} else {
				ret[k] = v * count
			}
		}
	}

	return ret
}
