package Array

import "fmt"

func Miss(arr []int) {

	large := arr[0]
	for i := 1; i < len(arr); i++ {
		if large < arr[i] {
			large = arr[i]
		}

	}
	fmt.Print(large)

	for k := 1; k < large; k++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[j] != k {
				count++

			} else {
				break
			}
		}
		if count == len(arr) {
			fmt.Println(k)

		}

	}
}
