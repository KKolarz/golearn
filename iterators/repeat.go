package iterators

import "strings"

func Repeat(char string, repeats int) string {

	var repeated strings.Builder

	for i := 0; i < repeats; i++ {
		repeated.WriteString(char)
	}

	return repeated.String()
}
