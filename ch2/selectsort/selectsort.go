// ch2.2, bigO, selection sort

package selectsort

/*
2.2-1. big O notation for n^3/1000 - 100n^2 - 100n - 3, θ(n^3)
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
