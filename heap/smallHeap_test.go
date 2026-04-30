package heap

import (
	"fmt"
	"testing"
)

func TestNewSmallHeap(t *testing.T) {
	testList := []int{
		3, 5, 6, 1, 4, 7, 2, 0,
	}

	smallHeap := NewSmallHeap()
	for _, val := range testList {
		smallHeap.Push(val)
	}

	for !smallHeap.isEmpty() {
		fmt.Println(smallHeap.Pop())
	}
}
