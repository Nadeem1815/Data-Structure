package Array

func Eg(arr []int) []int {
	leng := make([]int, len(arr)*2)
	for i := 0; i < len(arr); i++ {
		leng[i] = arr[i]
	}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// return arr
	for i := len(arr); i < len(leng); i++ {
		leng[i] = arr[i-len(arr)]
	}

	return leng
}
