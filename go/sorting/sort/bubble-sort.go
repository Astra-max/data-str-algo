package sort

func BubbleSort(arr []int) []int {
	size := len(arr)

	for i := 0; i < size; i++ {
		for j := 1; j < size; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]

			}
		}
	}
	return arr
}
