// ch2.1 insertion sort, loop invariant

package insertsort

// 2.1-2 nonincreasing insertion sort

// InsertionSort sort an []int
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

/*
2.1-3 Linear search for element
 pseudocode for index:
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
