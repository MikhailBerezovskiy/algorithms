package mergesort

import (
	"fmt"
	"sort"
	"testing"
)

func TestMergeNoSentinels(t *testing.T) {
	var A []int
	A = []int{1, 5, 7, 9, 2, 4, 6, 7}
	MergeNoSentinels(A, 0, 3, len(A)-1)
	fmt.Printf("A = %+v\n", A)
	if !sort.IntsAreSorted(A) {
		t.Errorf("A not sorted: %v", A)
	}

	A = []int{1, 3, 5, 2, 4, 6, 8, 9}
	MergeNoSentinels(A, 0, 2, len(A)-1)
	fmt.Printf("A = %+v\n", A)
	if !sort.IntsAreSorted(A) {
		t.Errorf("A not sorted: %v", A)
	}

	A = []int{33, 22, 2, 8, 10, 3, 9, 33, 44, 27}
	MergeNoSentinels(A, 2, 4, len(A)-2)
	testSlice := make([]int, len(A)-3)
	j := 0
	for i := 2; i <= len(A)-2; i++ {
		testSlice[j] = A[i]
		j++
	}
	fmt.Printf("testSlice = %+v\n", testSlice)
	fmt.Printf("A = %+v\n", A)
	if !sort.IntsAreSorted(testSlice) {
		t.Errorf("Test slice not sorted: %v\n Source array: %v", testSlice, A)
	}
}
