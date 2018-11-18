package dnc

import (
	"testing"
)

func TestFindMaxCrossingSubarray(t *testing.T) {
	type args struct {
		A    []int
		low  int
		high int
		mid  int
	}
	tests := []struct {
		name         string
		args         args
		wantMaxLeft  int
		wantMaxRight int
		wantSum      int
	}{
		{"pos and neg", args{[]int{10, -2, -5, 5, 1, 2, -3}, 0, 6, 3}, 0, 5, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMaxLeft, gotMaxRight, gotSum := FindMaxCrossingSubarray(tt.args.A, tt.args.low, tt.args.high, tt.args.mid)
			if gotMaxLeft != tt.wantMaxLeft {
				t.Errorf("FindMaxCrossingSubarray() gotMaxLeft = %v, want %v", gotMaxLeft, tt.wantMaxLeft)
			}
			if gotMaxRight != tt.wantMaxRight {
				t.Errorf("FindMaxCrossingSubarray() gotMaxRight = %v, want %v", gotMaxRight, tt.wantMaxRight)
			}
			if gotSum != tt.wantSum {
				t.Errorf("FindMaxCrossingSubarray() gotSum = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestFindMaxSubarray(t *testing.T) {
	type args struct {
		A    []int
		low  int
		high int
	}
	tests := []struct {
		name      string
		args      args
		wantSlow  int
		wantShigh int
		wantSum   int
	}{
		{"all pos", args{[]int{0, -1, -2, 3, 5}, 0, 4}, 3, 4, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSlow, gotShigh, gotSum := FindMaxSubarray(tt.args.A, tt.args.low, tt.args.high)
			if gotSlow != tt.wantSlow {
				t.Errorf("FindMaxSubarray() gotSlow = %v, want %v", gotSlow, tt.wantSlow)
			}
			if gotShigh != tt.wantShigh {
				t.Errorf("FindMaxSubarray() gotShigh = %v, want %v", gotShigh, tt.wantShigh)
			}
			if gotSum != tt.wantSum {
				t.Errorf("FindMaxSubarray() gotSum = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
