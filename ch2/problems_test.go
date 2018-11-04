package ch2sort

import (
	"reflect"
	"testing"
)

func TestInsertionSortSubArray(t *testing.T) {
	type args struct {
		A []int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty", args{[]int{}, 2}, []int{}},
		{"k>n", args{[]int{4, 3, 2, 1}, 5}, []int{1, 2, 3, 4}},
		{"k=1", args{[]int{4, 3, 2, 1}, 1}, []int{4, 3, 2, 1}},
		{"k=2", args{[]int{3, 1, 4, 2, 5, 3, 6, 4}, 2}, []int{1, 3, 2, 4, 3, 5, 4, 6}},
		{"k=3, small last chunck", args{[]int{3, 1, 4, 2, 5, 3, 6, 4}, 3}, []int{1, 3, 4, 2, 3, 5, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSortSubArray(tt.args.A, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSortSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSortSubArrays(t *testing.T) {
	type args struct {
		A []int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"a(8), k(2)", args{[]int{1, 2, 5, 6, 3, 4, 7, 8}, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"a(8), k(2)", args{[]int{5, 6, 1, 2, 7, 8, 3, 4}, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"a(9), k(3)", args{[]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"a(3), k(2)", args{[]int{2, 3, 1}, 2}, []int{1, 2, 3}},
		{"a(10), k(3)", args{[]int{4, 5, 6, 7, 8, 9, 1, 2, 3, 10}, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortSubArrays(tt.args.A, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortSubArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
