package selectsort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := make([][]int, 3)
	tests[0] = []int{1, 3, 2, 4, 5, 6, 4, 7, 8, 3}
	tests[1] = []int{3, 2, 1, 5, 3, 2}
	tests[2] = []int{1, 2, 3, 4}

	for i := 0; i < len(tests); i++ {
		A := SelectionSort(tests[i])
		fmt.Printf("A = %+v\n", A)
		if !sort.IntsAreSorted(A) {
			t.Errorf("A is not sorted %v", A)
		}
	}
}
