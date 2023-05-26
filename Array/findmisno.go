package Array

func MissNo(arr []int) int {
	n := len(arr)
	sum := 0
	actualsum := (n + 1) * (n + 2) / 2
	for _, value := range arr {
		sum = sum + value
	}
	return actualsum - sum
}
