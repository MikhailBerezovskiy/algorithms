package dnc

import (
	"math"
)

/*
Divide and Conquer
exercises for Maximum Subarray Problem
*/

/* pseudocode for divide and conquer maximum subarray problem

Find-Max-Crossing-Subarray(A, low, mid, high)
	leftSum = -inf
	sum = 0
	for i = mid downto low
		sum = sum + A[i]
		if sum > leftSum
			leftSum = sum
			maxLeft = i
	rightSum = +inf
	sum = 0
	for i = mid to high
		sum = sum + A[i]
		if sum > rightSum
			rightSum = sum
			maxRight = i
	return (maxLeft, maxRight, leftSum + rightSum)

Find-Maximum-Subarray(A, low, high)
	if high == low
		return (low, high, A[low])
	else
		mid = floor((low+high)/2)
		(leftLow, leftHigh, leftSum) = Find-Maximum-Subarray(A, low, mid)
		(rightLow, rightHigh, rightSum) = Find-Maximum-Subarray(A, mid, high)
		(crossLow, crossRight, crossSum) = Find-Max-Crossing-Subarray(A, low, mid, hight)
		if leftSum >= rightSum && leftSum >= crossSum
			return (leftLow, leftHigh, leftSum)
		else if rightSum >= leftSum && rightSum >= crossSum
			return (rightLow, rightHigh, rightSum)
		else
			return (crossLow, crossHigh, crossSum)

*/

// FindMaxCrossingSubarray returns left and right indexes for greates from middle point of array
func FindMaxCrossingSubarray(A []int, low, high, mid int) (maxLeft, maxRight, sum int) {
	leftSum := int(math.Inf(0))
	sum = 0
	for i := mid; i >= 0; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := int(math.Inf(0))
	sum = 0
	for i := mid + 1; i < high; i++ {
		sum += A[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}

// FindMaxSubarray returns maximum subarray from given array with given low, high indexes
func FindMaxSubarray(A []int, low, high int) (slow, shigh, sum int) {
	if low == high {
		return low, high, A[low]
	}
	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := FindMaxSubarray(A, low, mid)
	rightLow, rightHigh, rightSum := FindMaxSubarray(A, mid+1, high)
	crossLow, crossHigh, crossSum := FindMaxCrossingSubarray(A, low, high, mid)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}

// 4.1-1 What does find-maximum-subarray return when all elemnts of A are negative
