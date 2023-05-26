package String

func Length(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}
