package ch2sort

import (
	"bytes"
	"fmt"
	"reflect"
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

func TestBinarySum(t *testing.T) {
	var A, B, C, wait []byte
	A = []byte{0, 1, 1, 1}
	B = []byte{1, 0, 0, 0}
	C = BinarySum(A, B)
	wait = []byte{1, 1, 1, 1, 0}
	if !bytes.Equal(C, wait) {
		t.Errorf("Expect %v, got %v", wait, C)
	}

	A = []byte{1, 1, 1, 1, 0, 0, 1}
	B = []byte{1, 1, 0, 0, 1, 0, 1}
	C = BinarySum(A, B)
	wait = []byte{0, 1, 0, 0, 0, 1, 0, 1}
	if !bytes.Equal(C, wait) {
		t.Errorf("Expect %v, got %v", wait, C)
	}
}

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
		x   int
		ins bool
	}
	tests := []struct {
		name    string
		args    args
		wantPos int
	}{
		{"in array", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 8, false}, 7},
		{"in array", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 1, false}, 0},
		{"in array", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8, false}, 7},
		{"in array", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, false}, 2},
		{"even array", args{[]int{1, 2, 3, 4, 5}, 3, false}, 2},
		{"not in array", args{[]int{1, 2, 3, 4}, 5, false}, -1},
		{"not in array", args{[]int{1, 2, 4, 6}, 3, true}, 2},
		{"not in array", args{[]int{1, 2, 4, 6}, 8, true}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPos := BinarySearchRecursive(tt.args.A, tt.args.x, tt.args.ins); gotPos != tt.wantPos {
				t.Errorf("BinarySearchRecursive() = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}

func TestInsertionSortBinarySearch(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"unsorted with dups", args{[]int{2, 5, 2, 5, 1, 6, 3, 6}}, []int{1, 2, 2, 3, 5, 5, 6, 6}},
		{"sorted", args{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"reversed", args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSortBinarySearch(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSortBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasSumOfX(t *testing.T) {
	type args struct {
		S []int
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty set", args{[]int{}, 10}, false},
		{"true", args{[]int{1, 2, 3, 4, 5, 6}, 7}, true},
		{"false", args{[]int{1, 1, 1, 1}, 3}, false},
		{"only 1 match", args{[]int{1, 2, 3, 4, 5}, 9}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasSumOfX(tt.args.S, tt.args.x); got != tt.want {
				t.Errorf("HasSumOfX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		A []int
		p int
		r int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty array", args{[]int{}, 0, 0}, []int{}},
		{"sort 2 elems", args{[]int{2, 1}, 0, 1}, []int{1, 2}},
		{"sort 4 elems", args{[]int{2, 1, 4, 3}, 0, 3}, []int{1, 2, 3, 4}},
		{"sorted array", args{[]int{1, 2, 3, 4, 5}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{"sort 7 elems", args{[]int{4, 3, 2, 1, 7, 6, 5}, 0, 6}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"sort part", args{[]int{4, 3, 2, 1, 7, 6, 5}, 0, 5}, []int{1, 2, 3, 4, 6, 7, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.A, tt.args.p, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
