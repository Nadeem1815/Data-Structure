package String

func Sreverse(li string) string {
	sr := []byte(li)
	for i := 0; i < len(sr)/2; i++ {
		sr[i], sr[len(sr)-1-i] = sr[len(sr)-1-i], sr[i]
	}
	return string(sr)
}
