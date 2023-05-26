package Array

func Sum(arr []int) []int {
	sum := 0
	temp := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		sum = arr[i] + sum
		temp[i] = sum

	}
	return temp
}

