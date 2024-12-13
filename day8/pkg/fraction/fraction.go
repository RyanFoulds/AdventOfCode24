package fraction

func ReduceFraction(numerator, denominator int) (num, den int) {
	commonFactor := getCommonFactor(numerator, denominator)

	num = numerator / commonFactor
	den = denominator / commonFactor
	return
}

func getCommonFactor(num, den int) int {
	if den == 0 {
		return num
	}
	return getCommonFactor(den, num%den)
}
