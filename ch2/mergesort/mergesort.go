// ch2.3 merge sort

package mergesort

// 2.3-1 Illustrate merge sort on array A={3, 41, 52, 26, 38, 57, 9, 49}
//
//		 {3 9 26 38 41 49 52 57}			sorted
//			/				\
//    {3 26 41 52}		 {9 38 49 57}
//		/		\		   /	  \
//  {3 41}   {26 52}   {38 57}  {9 49}
//	/	\	  /	  \     /   \    |	 \
// {3} {41} {52} {26} {38} {57} {9} {49}	input

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
