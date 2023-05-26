package Array

func Reverse(arr1 []int) []int {
	length := len(arr1) - 1
	for i := 0; i < len(arr1)/2; i++ {
		temp := arr1[i]
		arr1[i] = arr1[length]
		arr1[length] = temp
		length--
	}

	return arr1
}
