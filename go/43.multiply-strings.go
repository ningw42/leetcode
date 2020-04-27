func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var rows []string
	var row []byte
	var carry int

	for i := len(num2) - 1; i >= 0; i-- {
		carry = 0
		digit2 := int(num2[i] - '0')
		row = []byte{}
		for padding := 0; padding < len(num2) - 1 - i; padding++ {
			row = append(row, '0')
		}
		for j := len(num1) - 1; j >= 0; j-- {
			digit1 := int(num1[j] - '0')
			result := digit1 * digit2 + carry
			carry = result / 10
			row = append(row, byte(result % 10) + '0')
		}
		if carry != 0 {
			row = append(row, byte(carry) + '0')
		}
		rows = append(rows, string(row))
	}

	return add(rows)
}

func add(rows []string) string {
	var exit bool
	var carry int
	var result []byte
	var index int
	var digitSum int

	for {
		exit = true
		digitSum = carry
		for _, row := range rows {
			if index < len(row) {
				exit = false
				digitSum += int(row[index] - '0')
			}
		}
		if exit {
			break
		}
		result = append(result, byte(digitSum % 10) + '0')
		carry = digitSum / 10
		index++
	}
	if carry != 0 {
		result = append(result, byte(carry) + '0')
	}

	return string(reverse(result))
}

func reverse(row []byte) []byte {
	for i := 0; i < len(row) / 2; i++ {
		row[i], row[len(row) - 1 - i] = row[len(row) - 1 - i], row[i]
	}
	return row
}