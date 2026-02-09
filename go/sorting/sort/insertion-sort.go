package sort

func InsertionSort(arr []int) []int {
	size := len(arr)

	for i:=1; i<size; i++ {
		j := i - 1
		temp := arr[i]
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	return arr
}