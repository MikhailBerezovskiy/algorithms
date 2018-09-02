package insertsort

//// index test

// insertion sort algorithm
// for nonincreasing chnage: a[i] < key
func InsertionSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
	}
	return a
}

// Index
// pseudocode for index:
// Index(A, v)
// ind = NIL
// for i = 1 to A.length
//     if A[i] == v
//         ind = i
//		   break
// return ind
//
// index loop invariant:
// initialization: index value is initialized before loop as not found, if array is empty
//		then index is nil
// maintenance: each loop iteration we check element i for equivalence with given value, that means
//		in A[1..i-1] there is no matching and current index value is nil
// termination: loop terminates when match happens or loop reached end of array, both of ends return index
//		if match happens i is returned, otherwise nil
func Index(a []int, v int) int {
	ind := -1
	for i := 0; i < len(a); i++ {
		if a[i] == v {
			ind = i
			break
		}
	}
	return ind
}

// Problem: two n-bit binary integers stored in A and B n size arrays
// the sum of two integers should be stored in binary form in (n+1) array C
// state problem formally and write pseudocode for adding the two integers
