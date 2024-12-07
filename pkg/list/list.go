<<<<<<< Updated upstream
//nolint:gosec
package list

import "slices"

func ListOfListsAreEqual[T rune | int](list1 [][]T, list2 [][]T) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if !slices.Equal[[]T](list1[i], list2[i]) {
			return false
		}
	}

	return true
}

func CountOfOccurencesOfStringInList(l []string, searchString string) int {
	count := 0
	for _, v := range l {
		if v == searchString {
			count++
		}
	}
	return count
}

func ReplaceAllInstancesOfStringInList(l []string, oldStr string, newStr string) []string {
	for i, s := range l {
		if s == oldStr {
			l[i] = newStr
		}
	}
	return l
}

// deletes values in slice at all indices, and returns a new list
func DeleteAtIndices[T any](s []T, is []int) []T {
	slices.Sort[[]int](is)
	slices.Reverse[[]int](is)
	newS := make([]T, len(s))
	copy(newS, s)
	for _, i := range is {
		newS = slices.Delete(newS, i, i+1)
	}
	return newS
}

// repeats a slice a given number of times. duplication = 0 means no change
func Repeat[T any](slice []T, duplication int) []T {
	if duplication < 1 || len(slice) == 0 {
		return slice
	}

	result := make([]T, 0, len(slice)*duplication)
	for i := 0; i <= duplication; i++ {
		result = append(result, slice...)
	}
	return result
}

func Sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
=======
//nolint:gosec
package list

import "slices"

func ListOfListsOfIntAreEqual(list1 [][]int, list2 [][]int) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if slices.Compare[[]int](list1[i], list2[i]) != 0 {
			return false
		}
	}

	return true
}

func CountOfOccurencesOfStringInList(l []string, searchString string) int {
	count := 0
	for _, v := range l {
		if v == searchString {
			count++
		}
	}
	return count
}

func ReplaceAllInstancesOfStringInList(l []string, oldStr string, newStr string) []string {
	for i, s := range l {
		if s == oldStr {
			l[i] = newStr
		}
	}
	return l
}

// deletes values in slice at all indices, and returns a new list
func DeleteAtIndices[T any](s []T, is []int) []T {
	slices.Sort[[]int](is)
	slices.Reverse[[]int](is)
	newS := make([]T, len(s))
	copy(newS, s)
	for _, i := range is {
		newS = slices.Delete(newS, i, i+1)
	}
	return newS
}

// repeats a slice a given number of times. duplication = 0 means no change
func Repeat[T any](slice []T, duplication int) []T {
	if duplication < 1 || len(slice) == 0 {
		return slice
	}

	result := make([]T, 0, len(slice)*duplication)
	for i := 0; i <= duplication; i++ {
		result = append(result, slice...)
	}
	return result
}

func Sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
>>>>>>> Stashed changes
