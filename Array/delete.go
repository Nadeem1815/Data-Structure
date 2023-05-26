package Array

func Delete(arr []int, index int) []int {
	newarr := make([]int, len(arr)-1)
	for i := 0; i < index; i++ {
		newarr[i] = arr[i]

	}
	for i := index; i < len(newarr); i++ {
		newarr[i] = arr[i+1]
	}

	return newarr
}
