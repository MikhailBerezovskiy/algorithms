package mergesort

import (
	"fmt"
	"reflect"
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

func TestInsertionSortRecursive(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"empty", args{[]int{}}, []int{}},
		{"2 elements", args{[]int{2, 1}}, []int{1, 2}},
		{"test1", args{[]int{2, 3, 1, 4, 5}}, []int{1, 2, 3, 4, 5}},
		{"test2", args{[]int{5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSortRecursive(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSortRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	type args struct {
		A   []int
		pos int
		x   int
	}
	tests := []struct {
		name    string
		args    args
		wantPos int
	}{
		// TODO: Add test cases.
		{"in array", args{[]int{1, 2, 3, 4}, 0, 4}, 3},
		{"in array", args{[]int{1, 2, 3, 4}, 0, 2}, 1},
		{"in array", args{[]int{1, 2, 3, 4}, 0, 1}, 0},
		{"in array", args{[]int{1, 2, 3, 4}, 0, 3}, 2},
		{"even array", args{[]int{1, 2, 3, 4, 5}, 0, 3}, 2},
		{"not in array", args{[]int{1, 2, 3, 4}, 0, 5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPos := BinarySearchRecursive(tt.args.A, tt.args.pos, tt.args.x); gotPos != tt.wantPos {
				t.Errorf("BinarySearchRecursive() = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}
