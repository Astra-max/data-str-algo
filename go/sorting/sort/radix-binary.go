package sort

func RadixSortBinary(a []int) []int {
	if len(a) == 0 {
		return a
	}

	b := []int{} // second stack

	size := len(a)

	// find max value
	max := 0
	for _, v := range a {
		if v > max {
			max = v
		}
	}

	// number of bits needed
	bits := 0
	for (max >> bits) != 0 {
		bits++
	}

	// radix passes
	for bit := 0; bit < bits; bit++ {

		count := size

		for count > 0 {

			num := a[0] // peek top

			// bit check
			if ((num >> bit) & 1) == 0 {
				b = append([]int{num}, b...)
				a = a[1:]
			} else {
				// rotate
				a = append(a[1:], num)
			}

			count--
		}

		for len(b) > 0 {
			a = append([]int{b[0]}, a...)
			b = b[1:]
		}
	}

	return a
}