/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	lenA, sliceA := reverse(a)
	lenB, sliceB := reverse(b)

	var reversed, result []byte
	carry := 0
	for i := 0; i < lenA || i < lenB; i++ {
		var aa, bb int
		if i < lenA {
			aa = int(sliceA[i] - '0')
		} else {
			aa = 0
		}
		if i < lenB {
			bb = int(sliceB[i] - '0')
		} else {
			bb = 0
		}

		sum := aa + bb + carry
		digit := sum % 2
		carry = sum / 2

		reversed = append(reversed, byte(digit+'0'))
	}
	if carry != 0 {
		reversed = append(reversed, '1')
	}

	for i := len(reversed) - 1; i >= 0; i-- {
		result = append(result, reversed[i])
	}

	return string(result)
}

func reverse(str string) (int, []byte) {
	length := len(str)
	slice := make([]byte, length)
	for i, c := range str {
		slice[length-1-i] = byte(c)
	}

	return length, slice
}

// @lc code=end

