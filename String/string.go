package String

import "strings"

func Name(str string) string {
	newstr := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		if i%2 == 0 {
			s := string(str[i])
			s = strings.ToUpper(s)
			newstr[i] = []byte(s)[0]
		} else {

			newstr[i] = str[i]

		}

	}
	return string(newstr)
}
