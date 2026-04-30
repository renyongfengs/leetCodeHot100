package heap

import (
	"fmt"
	"testing"
)

func TestNewBigHeap(t *testing.T) {
	testList := []int{
		3, 5, 6, 1, 4, 7, 2, 0,
	}

	bigHeap := NewBigHeap()
	for _, val := range testList {
		bigHeap.Push(val)
	}

	for !bigHeap.isEmpty() {
		fmt.Println(bigHeap.Pop())
	}
}
