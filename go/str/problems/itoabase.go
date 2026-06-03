package problems

func Itoabase(num, base int) string {
	results := ""
	baseTo := "0123456789ABCDEF"

	if base < 2 || base > 36 {
		return "Invalid base"
	}

	for num > 0 {
		remainder := num % base
		results = string(baseTo[remainder]) + results
		num /= base
	}
	return results
}