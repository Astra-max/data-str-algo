package sort


func CountSort(arr []int) []int {

	if len(arr) == 0 {
		return arr
	}

	k := 0

	for i:=0; i<len(arr); i++ {
		if arr[i] > k {
			k = arr[i]
		}
	}

	count := make([]int, k+1)
	results := make([]int, len(arr))

	for i:=0; i<len(arr); i++ {
		count[arr[i]]++
	}

	for i:=1; i <= k; i++ {
		count[i] = count[i] + count[i-1]
	}

	for i:=len(arr)-1; i>=0; i-- {
		results[count[arr[i]]-1] = arr[i]
	}

	copy(arr, results)

	return arr
}
