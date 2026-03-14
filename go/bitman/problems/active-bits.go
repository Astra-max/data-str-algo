package problems


/**
 * ActiveBits - returns number of active bits
 * Return - 0 if num is 0 / num of bits
 */

func ActiveBits(num int) int {
	count :=0

	for num > 0 {
		if num & 1 == 1 {
			count++
		}
		num >>=1
	}
	return count
}
