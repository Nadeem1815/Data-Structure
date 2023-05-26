package Recursion

import "fmt"

func BinRec(arr []int, val int) int {
	star := 0
	end := len(arr) - 1
	mid := (star + end) / 2
	fmt.Println(mid)
	if val == arr[mid] {
		return mid + 1
	} else if arr[mid] < val {
		return BinRec(arr[mid+1:], val)
	} else {
		return BinRec(arr[:mid-1], val)
	}
}
