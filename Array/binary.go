package Array

func Binary(arr []int, target int) int {
	// arr1 := []int{1,2,3,4,5,6,7,8}
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid

		} else if target < arr[mid] {
			end = mid - 1

		} else if target > arr[mid] {
			start = mid + 1

		}

	}
	return -1
}
