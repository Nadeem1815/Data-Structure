package Array

// func MoveZeroes(nums []int) {
// 	// Initialize a variable to keep track of the next non-zero index
// 	nextNonZero := 0

// 	// Traverse the array
// 	for _, num := range nums {

// 		// If the current element is non-zero, move it to the next non-zero index
// 		if num != 0 {
// 			nums[nextNonZero] = num
// 			nextNonZero++
// 			// fmt.Print(nums)
// 		}
// 	}
// 	fmt.Print(nextNonZero)
// 	// Fill the remaining elements with zeros
// 	for i := nextNonZero; i < len(nums); i++ {
// 		nums[i] = 0
// 	}
// }

func Zeroend(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
