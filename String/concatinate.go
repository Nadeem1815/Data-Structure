package String

func Concatinate(str string, str1 string) string {
	// newstr := make([]byte, len(str)+len(str1))
	// for i := 0; i < len(str); i++ {
	// 	newstr[i] = str[i]
	// }
	// for i := len(str); i < len(newstr); i++ {
	// 	newstr[i] = str1[i-len(str)]

	// }

	newstr:=str+""+str1
	return string(newstr)
}
