func reverseBits(num uint32) uint32 {
	var ret, digitWeight uint32
	digitWeight = 0x80000000
	for i := 0; i < 32; i++ {
		if num%2 == 1 {
			ret += digitWeight
		}
		digitWeight = digitWeight / 2
		num = num / 2
	}

	return ret
}
