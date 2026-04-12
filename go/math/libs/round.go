package libs

func Round(num float64) float64 {
	if num >= 0 {
		return float64(int(num+ + 0.5))
	}
	return float64(int(num - 0.5))
}