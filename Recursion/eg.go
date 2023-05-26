package Recursion

import "fmt"

func F(n int) int {
	if n <= 1 {
		return 1
		fmt.Println(n)

	}
	fmt.Println(n)
	return n * F(n-1)
	

}
