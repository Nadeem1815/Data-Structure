package String

func Linear(li string, target byte) int {
	for i := 0; i < len(li); i++ {
		if li[i] == target {
			return i
		}
	}
	return -1
}
