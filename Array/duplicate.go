package Array

func Duplicate(dup []int) []int {
	// p:=make(map[int]bool)
	// n := []int{}
	// count := 0
	s := len(dup)
	for i := 0; i < s-1; i++ {

		if dup[i] == dup[i+1] {
			for j := i; j < s-1; j++ {
				dup[j] = dup[j+1]

				// n = append(n, )
				// j--
				// count++
				// return dup

			}
			s--
			i--
			// dup-count

		}

		// }
		// return n
		// for _;value:=range dup{

	}
	return dup[s:]
}
