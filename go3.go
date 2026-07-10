func reverse(x int) int {
	var result int
	num := 0

	max := math.MaxInt32
	min := math.MinInt32
	for x != 0 {
		num = x % 10
		if result > max/10 || (result == max/10 && num > 7) {
			return 0
		}
		if result < min/10 || (result == min/10 && num < -8) {
			return 0
		}

		result = result*10 + num
		x = x / 10
	}
	return result
}