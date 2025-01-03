//nolint:gosec
package cast

// Suite of casting functions to speed up solutions
// This is NOT idiomatic Go... but AOC isn't about that...

import (
	"fmt"
	"strconv"
)

// ToInt will case a given arg into an int type.
// Supported types are:
//   - string
func ToInt(arg interface{}) int {
	var val int
	switch arg := arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg)
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

// ToString will case a given arg into an int type.
// Supported types are:
//   - int
//   - byte
//   - rune
func ToString(arg interface{}) string {
	var str string
	switch arg := arg.(type) {
	case int:
		str = strconv.Itoa(arg)
	case byte:
		b := arg
		str = string(rune(b))
	case rune:
		str = string(arg)
	default:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	}
	return str
}

const (
	ASCIICodeCapA   = int('A') // 65
	ASCIICodeCapZ   = int('Z') // 65
	ASCIICodeLowerA = int('a') // 97
	ASCIICodeLowerZ = int('z') // 97
)

// ToASCIICode returns the ascii code of a given input
func ToASCIICode(arg interface{}) int {
	var asciiVal int
	switch arg := arg.(type) {
	case string:
		str := arg
		if len(str) != 1 {
			panic("can only convert ascii Code for string of length 1")
		}
		asciiVal = int(str[0])
	case byte:
		asciiVal = int(arg)
	case rune:
		asciiVal = int(arg)
	}

	return asciiVal
}

// ASCIIIntToChar returns a one character string of the given int
func ASCIIIntToChar(code int) string {
	return string(rune(code))
}

// Converts an array of int to an array of string
func IntArrayToStringArray(arr []int) []string {
	stringArray := make([]string, len(arr))

	for i, num := range arr {
		stringArray[i] = ToString(num)
	}

	return stringArray
}
