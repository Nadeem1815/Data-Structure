package Array

func Conca(nums []int) []int {
	arr := make([]int, len(nums)*2)
	// for i := 0; i < len(nums); i++ {
	// 	arr[i] = nums[i]
	// }

	// for i := len(nums); i < len(arr); i++ {
	// 	arr[i] = nums[i-len(nums)]

	// }
	
	arr = append(nums, nums...)

	return arr
}
