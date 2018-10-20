package sort

// ch2.1 insertion sort, loop invariant

// 2.1-2 non-increasing insertion sort

// InsertionSort sort an []int
// for non-increasing change: a[i] < key
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

/*
2.1-3 Linear search for element
pseudocode for index function:
 Index(A, v)
 ind = NIL
 for i = 1 to A.length
	 if A[i] == v
		 ind = i
		   break
 return ind

index loop invariant:
	initialization: index value is initialized before loop as not found, if array is empty
then index is nil
	maintenance: each loop iteration we check element i for equivalence with given value,
that means in A[1..i-1] there is no matching and current index value is nil
	termination: loop terminates when match happens or loop reached end of array, both of
ends return index if match happens i is returned, otherwise nil
*/

// Index return index of int in []int
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

/*
2.1-4 Sum of 2 binary arrays
Problem: two n-bit binary integers stored in A and B n size arrays
the sum of two integers should be stored in binary form in (n+1) array C
state problem formally and write pseudocode for adding the two integers

 pseudocode:
	BinarySum(A, B)
	  C = [A.length+1]int
	  for i = 1 to A.length
		 sum = A[i] + B[i] + C[i]
		 C[i] = sum % 2
		 C[i+1] = sum / 2
	  return C

loop invariant:
	initialization: if n=0 then C initialized as [0] array
	maintenance: each iteration we add C[i], A[i], and B[i] which can be from 0 to 3
		we store reminder of division 2 in C[i] which can be 0 or 1, for 1 and 3 sum values
		and store quotient of division in C[i+1] which can be 1 if sum equals 3,
		and will be added in next iteration
		by end of each iteration we have C[n+1] array with sum of A and B
  termination: loop terminates after end of A array is reached, which means all elements must be summed
*/

// BinarySum return n+1 slice bits sum of two n-sized slices A and B
func BinarySum(A, B []byte) []byte {
	C := make([]byte, len(A)+1)
	for i := 0; i < len(A); i++ {
		sum := A[i] + B[i] + C[i]
		C[i] = sum % 2
		C[i+1] = sum / 2
	}
	return C
}

// ch2.2, bigO, selection sort

/*
2.2-1.
big O notation for n^3/1000 - 100n^2 - 100n - 3,
-> θ(n^3)

Selection sort
pseudocode:
	SelectionSort(A)
	for i = 1 to A.Length-1
	  min = i
	  for j = i+1 to A.Length
		  if A[j] < A[min]
			  min = j
	  exchange A[i] with A[min]
	return A

	Loop invariant maintenance: in each iteration sub-loop iterate through
remain part of array and search minimum value, and then swap minimum value with current value,
as a result at the beginning of each iteration 1..i elements are sorted
	The best and worst case for running loop in terms of bigO are same
we need to go through array n * (n / 2) times, while n is outer loop and n/2 inner loop
therefore bigO is θ(n^2)
*/

// SelectionSort returns sorted list of integers
func SelectionSort(A []int) []int {
	for i := 0; i < len(A)-1; i++ {
		min := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		A[i], A[min] = A[min], A[i]
	}
	return A
}

/*
2.2-3 linear search average and wort cases
in average case probability of finding value at index i in n array is equal 1/n
with median = (n+1)/2
in worst case element not in array or at very last index which makes algorithm run n+1 or n times
so average is (n+1)/2 and worst is n+1, with bigO θ(n) for both cases
*/

/*
2.2-4 how we can modify algorithms to have a good best cases running time?
for sort algorithms we can and check if array is sorted then return array
that makes best-case runs with θ(n) time
in general add precondition checks
*/

// ch2.3 merge sort

// 2.3-1 Illustrate merge sort on array A={3, 41, 52, 26, 38, 57, 9, 49}
//
//		 {3 9 26 38 41 49 52 57}			sorted
//         /                \
//    {3 26 41 52}        {9 38 49 57}
//      /       \          /       \
//  {3 41}    {26 52}  {38 57}    {9 49}
//  /   \     /    \    /   \      |   \
// {3} {41} {52} {26} {38} {57}   {9} {49}	input

// 2.3-2 rewrite merge procedure without using sentinels
/* pseudocode

Merge(A,p,q,r)
	n1 = q - p + 1
	n2 = r - q
	let L[1..n1] and R [1..n2]
	for i=1 to n1
		L[i] = A[p+i-1]
	for j=1 to n2
		R[j] = A[q+j-1]
	i = 1
	j = 1
	for k=p to r
		if i > n1
			A[k] = R[j]
			j = j + 1
			continue
		if j > n2
			A[k] = L[i]
			i = i + 1
			continue
		if L[i] < R[j]
			A[k] = L[i]
			i = i + 1
		else
			A[k] = R[j]
			j = j + 1

*/

// MergeNoSentinels merge 2 sorted parts of array in place
func MergeNoSentinels(A []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	L, R := make([]int, n1), make([]int, n2)
	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = A[q+1+j]
	}
	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if i >= n1 {
			A[k] = R[j]
			j++
			continue
		}
		if j >= n2 {
			A[k] = L[i]
			i++
			continue
		}
		if L[i] < R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}

// 2.3-3 mathematical induction

// 2.3-4 recursive insertion sort

// InsertionSortRecursive - O(n^2)
func InsertionSortRecursive(A []int) []int {
	// termination condition
	if len(A) <= 1 {
		return A
	}

	n := len(A) - 1
	key := A[n]

	// get sorted sub array A[1...n-1]
	sorted := InsertionSortRecursive(A[:n])
	sorted = append(sorted, 0)

	// insert A[n] in sorted subarray A[1...n-1]
	for k, v := range sorted {
		if key <= v {
			n = k
			break
		}
	}
	copy(sorted[n+1:], sorted[n:])
	sorted[n] = key

	return sorted
}

// 2.3-5 Binary search

// BinarySearchRecursive lg(n) search in sorted array
// A - sorted list for search, x item for search,
// if ins true search return position where to insert item
func BinarySearchRecursive(A []int, x int, ins bool) (pos int) {
	// closure function only resize array, x and pos are in enclosing scope
	var f func([]int)
	f = func(a []int) {
		if len(a) == 0 {
			if ins {
				return
			}
			pos = -1
			return
		}
		n := len(a) / 2
		if a[n] == x {
			pos += n
			return
		}
		low := 0
		high := len(a)
		if x > a[n] {
			low = n + 1
			pos += n + 1
		} else {
			high = n
		}
		f(a[low:high])
	}
	f(A)
	return pos
}

// 2.3-6 Binary search in InsertionSort

// InsertionSortBinarySearch sort an []int, with binary search for inseting item
// for non-increasing change: a[i] < key, O(n*lg(n))
func InsertionSortBinarySearch(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		pos := BinarySearchRecursive(a[:j], key, true)
		// shift values
		copy(a[pos+1:j+1], a[pos:j])
		a[pos] = key
	}
	return a
}

/*
2.3-7 Descibe  O(nlgn) time algorithm, that determines whether in set S exist
sum of two items that equals exactly x
returns true if array has two elements with sum of them equals x

pseudocode
---------------------------------------------------------------
HasSumOfX(S, x)
	sorted = MergeSort(S) or InsertBinarySearch(S) # c1 n*lg(n)
	for i = 1 to S.length # c2 n
		second_summand = x - sorted[i] # c3 n
		pos = BinarySearch(sorted, second_summand) # c4 lg(n)
		if pos > -1 then # c5 n
		  return true # c6 1
	return false # c7 1
---------------------------------------------------------------
T(n) = n*lg(n) + n + n + lg(n) + n + 1 + 1 = n(4+lg(n)) + lg(n) + 2
consequently,
O(n) = n*lg(n)
*/

// HasSumOfX return true if Set of int's has two items that sum of them equals x
func HasSumOfX(S []int, x int) bool {
	// sort S with merge sort
	sorted := InsertionSortBinarySearch(S)
	for i := 0; i < len(sorted); i++ {
		secondSummand := x - sorted[i]
		pos := BinarySearchRecursive(sorted, secondSummand, false)
		if pos > -1 {
			return true
		}
	}
	return false
}
