package ch2sort

/*
2-1. Insertion sort on small arrays in merge sort
because insertion sort is more efficient on small arrays than merge sort
consider an algorithm that will sort k length subarrays sorted by insertion sort
and after merge them together
*/

// a. Show that insetion sort can sort the n/k sublists each of length k in 0(n*k) worst-case time

// InsertionSortSubArray split array in sub arrays with k-size
// and sort each one by InsertionSort algorithm
// T(n) = c1*n/k + c2*n/k + c3*n/k*k^2 + c4 + c5*1/k*k^2 + c6
// => n(2/k)+ k(n+1) => O(nk)
func InsertionSortSubArray(A []int, k int) []int {
	chunks := len(A) / k
	for i := 0; i < chunks; i++ { // c1 n/k
		sub := A[i*k : (i+1)*k]  // c2 n/k
		sub = InsertionSort(sub) // c3 n/k * k^2
	}
	// check if there last chunk that smaller than k size
	if chunks*k < len(A) { // c4 1
		InsertionSort(A[chunks*k : len(A)]) // c5 1/k k^2
	}
	return A // c6 1
}

// b. Show how to merge sublists in O(n lg(n/k)) worst-case time

// MergeSortSubArrays merge-sort k-sized subarrays within A array
func MergeSortSubArrays(A []int, k int) []int {
	p := 0
	q := (1 + len(A)/k) / 2 * k // split subarrays for merging by k-size border
	r := len(A)
	if len(A) > k { // lowest value to return is at most subarray size
		MergeSortSubArrays(A[p:q], k)
		MergeSortSubArrays(A[q:r], k)
		MergeNoSentinels(A, p, q-1, r-1)
	}
	return A
}

// c. What is a largest value for k as k=f(n) when modified MergeSort algorithm
//is less effective than standart merge sort
