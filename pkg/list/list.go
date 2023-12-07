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
