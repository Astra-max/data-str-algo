package sort

func RadixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	maxVal := 0

	for _, val := range arr {
		maxVal = max(maxVal, val)
	}

	pstn := 1

	for maxVal/pstn > 0 {
		customCountSort(arr, pstn)
		pstn *= 10
	}
	return arr
}

func customCountSort(arr []int, pstn int) {
	results := make([]int, len(arr))
	count := make([]int, 10)

	for _, val := range arr {
		count[(val/pstn)%10]++
	}

	for i:=1; i<10; i++ {
		count[i] += count[i-1]
	}

	for i:=len(arr)-1; i>=0; i-- {
		results[count[(arr[i]/pstn)%10]-1] = arr[i]
		count[(arr[i]/pstn)%10]--
	}
	copy(arr, results)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}