package Array

func Bitween(arr []int, value int, index int) []int {
	newarr := make([]int, len(arr)+1)
	for i := 0; i < index; i++ {
		newarr[i] = arr[i]

	}
	newarr[index] = value
	for i := index + 1; i < len(newarr); i++ {
		newarr[i] = arr[i-1]
	}
	return newarr
}

