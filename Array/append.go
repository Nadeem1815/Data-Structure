package Array


func Add(arr []int, value int) []int {
	newarr := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		newarr[i] = arr[i]
	}
	newarr[len(newarr)-1] = value

	return newarr

}
