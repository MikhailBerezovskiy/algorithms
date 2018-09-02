package insertsort

import (
	"sort"
	"testing"
)

func TestInsertSort(t *testing.T) {
	var A1 = []int{10, 50, 23, 35, 62, 42, 25, 66, 7}
	var A2 = []int{34, 43, 53, 35, 65, 56, 7, 1, 4, 7}
	var A3 = []int{1}
	// insertion sort test
	InsertionSort(A1)
	InsertionSort(A2)
	InsertionSort(A3)

	if !sort.IntsAreSorted(A1) {
		t.Errorf("A1:%v not sorted", A1)
	}
	if !sort.IntsAreSorted(A2) {
		t.Errorf("A2:%v not sorted", A2)
	}
	if !sort.IntsAreSorted(A3) {
		t.Errorf("A3:%v not sorted", A3)
	}
}

func TestIndex(t *testing.T) {
	var A1 = []int{10, 50, 23, 35, 62, 42, 25, 66, 7}
	//var A2 = []int{34, 43, 53, 35, 65, 56, 7, 1, 4, 7}
	if Index(A1, 23) != 2 {
		t.Error("index of 23 in A not 2")
	}
	if Index(A1, 1) != -1 {
		t.Error("find not existing index")
	}
}
