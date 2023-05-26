package Binary

func Bina(arr []int, target int) int {

	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			l = mid + 1

		} else if arr[mid] > target {
			r = mid - 1

		}
	}
	return -1
}
