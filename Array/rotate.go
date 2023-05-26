package Array

func Rotate(arr []int, times int) []int {
	for i := 0; i < times; i++ {
		for j := len(arr) - 1; j > 0; j-- {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
		}

	}
	return arr
}
