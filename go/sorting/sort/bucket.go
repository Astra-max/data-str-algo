package sort

import "sort"

func Bucket(arr []int, k int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}
	// find maximun
	results := []int{}
	max := Max(arr)
	// create buckets

	buckets := make([][]int, max+1)

	// add num to respective bucket

	for i:=0; i<len(arr); i++ {
		bucketIdx := (k*arr[i])/max
		buckets[bucketIdx] = append(buckets[bucketIdx], arr[i])
	}

	for i:=0; i<k; i++ {
		sort.Ints(buckets[i])
	}

	for _, bucket := range buckets {
	    results = append(results, bucket...)
	}
	return results
}

func Max(arr []int) int {
	max := arr[0]

	for _, num := range arr[1:] {
		if num > max {
			max = num
		}
	}
	return max
}