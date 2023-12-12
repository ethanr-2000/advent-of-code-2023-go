//nolint:gosec
package string_util

func ChangeRuneAtIndex(s string, i int, c rune) string {
	strRune := []rune(s)
	strRune[i] = c
	return string(strRune)
}

// repeats a string a given number of times. times=0 means no repeats
// use sep="" for no separator
func Repeat(str string, times int, sep string) string {
	if times == 0 {
		return str
	}

	new := str
	for i := 0; i < times; i++ {
		new += sep + str
	}
	return new
}
