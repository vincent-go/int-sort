package sort

import (
	"testing"
)

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, j := range a {
		if b[i] != j {
			return false
		}
	}
	return true
}

type sortTest struct {
	input  []int
	expect []int
}

var testCases = []sortTest{
	{[]int{2, 4, 4, 6, 5, 1}, []int{1, 2, 4, 4, 5, 6}},
	{[]int{2, 1, 0, -1, -4}, []int{-4, -1, 0, 1, 2}},
	{[]int{3, 2}, []int{2, 3}},
	{[]int{2}, []int{2}},
	{[]int{}, []int{}},
}

func runSortFuncTest(t *testing.T, f func(s []int) []int, funcName string, testCases []sortTest) {
	for _, test := range testCases {
		result := f(test.input)
		if !isEqual(test.input, test.input) {
			t.Errorf("%v changed the original slice!\nFrom: %v\nTo: %v", funcName, test.input, test.input)
		}
		if !isEqual(result, test.expect) {
			t.Errorf("%v failed!\nExpect: %v\nGot: %v", funcName, test.expect, result)
		}
	}
}

func TestSelection(t *testing.T) { runSortFuncTest(t, Selection, "Selection sort", testCases) }
func TestBubble(t *testing.T)    { runSortFuncTest(t, Bubble, "Bubble sort", testCases) }
func TestRecursiveBubble(t *testing.T) {
	runSortFuncTest(t, RecursiveBubble, "RecursiveBubble sort", testCases)
}
func TestInsertion(t *testing.T) { runSortFuncTest(t, Insertion, "Insertion sort", testCases) }
func TestRecursiveInsertion(t *testing.T) {
	runSortFuncTest(t, RecursiveInsertion, "RecursiveInsertion sort", testCases)
}
func TestQuick(t *testing.T) {
	runSortFuncTest(t, Quick, "Quick sort", testCases)
}
func TestHeap(t *testing.T) {
	runSortFuncTest(t, Heap, "Heap sort", testCases)
}
func TestMerge(t *testing.T) {
	runSortFuncTest(t, Merge, "Merge sort", testCases)
}
