package Array

func Finbinary(arr []int) int {
	Length := len(arr)
	Sum := 0
	for i := 0; i < Length-1; i++ {
		Sum = Sum + arr[i]
	}
	return Sum
}
