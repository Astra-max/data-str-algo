package libs

func Ceil(num float64) float64 {
	if num == float64(int(num)) {
		return num
	} else if num > 0 {
		return float64(int(num) + 1)
	} else {
		return float64(int(num))
	}
}