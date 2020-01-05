var digits []string
var tens []string
var thousands []string

func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    digits = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
    tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
    thousands = []string{"", "Thousand", "Million", "Billion"}
 
    // get english representation every 3 digits
    var results []string
    for num != 0 {
        results = append(results, toWords(num%1000))
        num /= 1000
    }
    reverse(results)
    
    // add unit 
    var ret []string
    for i := 0; i < len(results); i++ {
        if results[i] == "" {
            ret = append(ret, "")
        } else {
            ret = append(ret, strings.Trim(results[i] + " " + thousands[len(results)-1-i], " "))
        }
    }
   
    // concat
    r := strings.Builder{}
    for _, s := range ret {
        if s != "" {
            r.WriteString(s)
            r.WriteByte(' ')
        } 
    }
    
    return strings.Trim(r.String(), " ")
}

func toWords(n int) string {
    hundred := n / 100
    ten := (n - hundred * 100) / 10
    one := n % 10
    
    ret := strings.Builder{}
    if hundred != 0 {
        ret.WriteString(digits[hundred])
        ret.WriteByte(' ')
        ret.WriteString("Hundred")
    }
    if ten != 0 {
        ret.WriteByte(' ')
        if ten == 1 {
            ret.WriteString(tens[ten*10+one])
        } else {
            ret.WriteString(tens[ten])
            ret.WriteByte(' ')
            ret.WriteString(digits[one])
        }
    } else {
        if one != 0 {
            ret.WriteByte(' ')
            ret.WriteString(digits[one])
        }
    }
    return strings.Trim(ret.String(), " ")
}

func reverse(s []string) {
    for i := 0; i < len(s)/2; i++ {
        s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
    }
}
