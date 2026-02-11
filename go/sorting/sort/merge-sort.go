package sort

func MergeSort(arr []int, first, last int) []int {
	if first >= last {
		return []int{arr[first]}
	}
	
    mid := (first + last) / 2
	left := MergeSort(arr, first, mid)
	right := MergeSort(arr, mid+1, last)
	return Merge(left, right)
}

func Merge(arr1, arr2 []int) []int {
	n := len(arr1)
	m := len(arr2)
	i, j := 0,0
	temp := []int{}

	for i< n && j < m {
		if arr1[i] < arr2[j] {
			temp = append(temp, arr1[i])
			i++
		} else {
			temp = append(temp, arr2[j])
			j++
		}
	}
	temp = append(temp, arr1[i:]...)
	temp = append(temp, arr2[j:]...)
	return temp
}