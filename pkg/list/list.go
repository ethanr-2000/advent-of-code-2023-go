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
