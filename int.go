// Package sort contains sort algorithms for integer slice (asending order) without modifing the original slice
// Note: any algorithm with recursion involved will consumes more space due to data copying
package sort

// Selection sort algorithm sorts an array by repeatedly finding the minimum element (considering ascending order) from unsorted part and putting it at the beginning. The algorithm maintains two subarrays in a given array.
// 1) The subarray which is already sorted.
// 2) Remaining subarray which is unsorted.
// In every iteration of selection sort, the minimum element (considering ascending order) from the unsorted subarray is picked and moved to the sorted subarray.
func Selection(s []int) []int {
	length := len(s)
	sorted := make([]int, length)
	unsorted := make([]int, length)
	copy(unsorted, s)
	for i := 0; i < length; i++ {
		idx := 0
		min := unsorted[0]
		for j, e := range unsorted {
			if e < min {
				min = e
				idx = j
			}
		}
		sorted[i] = min
		copy(unsorted[idx:], unsorted[idx+1:]) // Shift a[i+1:] left one index.
		unsorted = unsorted[:len(unsorted)-1]
	}
	return sorted
}

// Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in wrong order.
func Bubble(s []int) []int {
	length := len(s)
	sorted := make([]int, length)
	copy(sorted, s)
	var swapped bool
	// sweep through the slice at most length-1 times and all bigger int will be in the right place
	for pass := 0; pass < length-1; pass++ {
		swapped = false
		for i := 0; i < length-1; i++ {
			if sorted[i] > sorted[i+1] {
				sorted[i], sorted[i+1] = sorted[i+1], sorted[i]
				swapped = true
			}
		}
		if !swapped {
			return sorted
		}
	}
	return sorted
}

// RecursiveBubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in wrong order.
func RecursiveBubble(s []int) []int {
	length := len(s)
	sorted := make([]int, length)
	copy(sorted, s)
	if len(sorted) <= 1 {
		return sorted
	}
	copy(sorted[1:], RecursiveBubble(sorted[1:]))
	for i := 0; i < length-1; i++ {
		if sorted[i] > sorted[i+1] {
			sorted[i], sorted[i+1] = sorted[i+1], sorted[i]
		}
	}
	return sorted
}

// Insertion sort is a simple sorting algorithm that works the way we sort playing cards in our hands.
// Algorithm
// // Sort an arr[] of size n
// insertionSort(arr, n)
// Loop from i = 1 to n-1.
// ……a) Pick element arr[i] and insert it into sorted sequence arr[0…i-1]
func Insertion(s []int) []int {
	length := len(s)
	sorted := make([]int, length)
	copy(sorted, s)
	// cheak sorted[i], insert it to the right place if there is value bigger than it
	for i := 1; i < length; i++ {
		temp := sorted[i]
		for j := i - 1; j >= 0; j-- {
			if sorted[j] > temp {
				sorted[j+1] = sorted[j]
				sorted[j] = temp
			}
		}
	}
	return sorted
}

// RecursiveInsertion sort is a simple sorting algorithm that works the way we sort playing cards in our hands.
// Recursion Idea:
// Base Case: If array size is 1 or smaller, return.
// Recursively sort first n-1 elements.
// Insert last element at its correct position in sorted array.
func RecursiveInsertion(s []int) []int {
	length := len(s)
	sorted := make([]int, length)
	copy(sorted, s)
	if len(sorted) <= 1 {
		return sorted
	}
	copy(sorted[:length-1], RecursiveInsertion(sorted[:length-1]))
	for i := length - 1; i > 0; i-- {
		if sorted[i] < sorted[i-1] {
			sorted[i], sorted[i-1] = sorted[i-1], sorted[i]
		}
	}
	return sorted
}

// Quick sort is a divide and conquer algorithm. It first divides the input array into two smaller sub-arrays: the low elements and the high elements. It then recursively sorts the sub-arrays. The steps for in-place Quicksort are:
// Pick an element, called a pivot, from the array.
// Partitioning: reorder the array so that all elements with values less than the pivot come before the pivot, while all elements with values greater than the pivot come after it (equal values can go either way). After this partitioning, the pivot is in its final position. This is called the partition operation.
// Recursively apply the above steps to the sub-array of elements with smaller values and separately to the sub-array of elements with greater values.
// The base case of the recursion is arrays of size zero or one, which are in order by definition, so they never need to be sorted.
// The pivot selection and partitioning steps can be done in several different ways; the choice of specific implementation schemes greatly affects the algorithm's performance.
func Quick(s []int) []int {
	sorted := make([]int, len(s))
	copy(sorted, s)
	divIdx := len(sorted) - 1
OUTER:
	for i := len(sorted) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if j == i-1 && sorted[j] > sorted[i] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
				divIdx--
				break OUTER
			} else if sorted[j] > sorted[i] {
				sorted[j], sorted[i-1] = sorted[i-1], sorted[j]
				sorted[i], sorted[i-1] = sorted[i-1], sorted[i]
				divIdx--
				break
			}
		}
	}
	if len(sorted) <= 1 {
		return sorted
	}
	copy(sorted[divIdx+1:], Quick(sorted[divIdx+1:]))
	copy(sorted[:divIdx], Quick(sorted[:divIdx]))
	return sorted
}

// Heap sort algorithm involves preparing the list by first turning it into a max heap.
// The algorithm then repeatedly swaps the first value of the list with the last value,
// decreasing the range of values considered in the heap operation by one,
// and sifting the new first value into its position in the heap.
// This repeats until the range of considered values is one value in length.
func Heap(s []int) []int {
	length := len(s)
	sorted := buildMaxHeap(s)
	for i := length - 1; i > 0; i-- {
		sorted[0], sorted[i] = sorted[i], sorted[0]
		copy(sorted[:i], buildMaxHeap(sorted[:i]))
	}
	return sorted
}

// create max heap from the given slice
func buildMaxHeap(s []int) []int {
	r := []int{}
	for i := 0; i < len(s); i++ {
		r = append(r, s[i])
		r = bubbleUp(r, i)
	}
	return r
}

func bubbleUp(s []int, i int) []int {
	if i == 0 {
		return s
	}
	if s[i] > s[(i-1)/2] {
		s[(i-1)/2], s[i] = s[i], s[(i-1)/2]
		return bubbleUp(s, (i-1)/2)
	}
	return s
}

// Merge Sort is a Divide and Conquer algorithm.
// It divides input array in two halves, calls itself for the two halves and then merges the two sorted halves.
// The merge() function is used for merging two halves.
// The merge(arr, l, m, r) is key process that assumes that arr[l..m]
// and arr[m+1..r] are sorted and merges the two sorted sub-arrays into one
func Merge(s []int) []int {
	length := len(s)
	unsorted := make([]int, length)
	copy(unsorted, s)
	middle := length / 2
	if middle == 0 {
		return unsorted
	}
	copy(unsorted[:middle], Merge(unsorted[:middle]))
	copy(unsorted[middle:], Merge(unsorted[middle:]))
	return merge(unsorted[:middle], unsorted[middle:])
}

func merge(first, last []int) []int {
	// important: have to get the iteration value first
	itr := len(first) + len(last)
	r := []int{}
	// if use len(first) + len(last) instead of itr, because the first/last gets modified in the iteration
	// len(first) + len(last) will because smaller with these iteration
	// and the r will only be the a part of what it should be!
	for i := 0; i < itr; i++ {
		if first[0] < last[0] {
			r = append(r, first[0])
			first = first[1:]
			if len(first) == 0 {
				r = append(r, last...)
				return r
			}
		} else {
			r = append(r, last[0])
			last = last[1:]
			if len(last) == 0 {
				r = append(r, first...)
				return r
			}
		}
	}
	return r
}
