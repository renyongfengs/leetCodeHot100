package heap

import (
	"math"
)

type SmallHeap struct {
	List []int
	size int
}

func NewSmallHeap() *SmallHeap {
	return &SmallHeap{
		List: make([]int, 0),
		size: 0,
	}
}

func (s *SmallHeap) Push(val int) {
	s.List = append(s.List, val)
	s.HeapInsert(s.size)
	s.size++
}

func (s *SmallHeap) Pop() int {
	if s.isEmpty() {
		return math.MinInt
	}
	out := s.List[0]
	s.size--

	if !s.isEmpty() {
		s.swap(0, s.size)
		s.HeapIfy(0, s.size)
	}

	return out
}

func (s *SmallHeap) HeapIfy(i, size int) {
	left := 2*i + 1
	for left < size {
		//找出左右节点最小的
		smallIdx := left
		if left+1 < size && s.List[left+1] < s.List[left] {
			smallIdx = left + 1
		}

		if s.List[i] <= s.List[smallIdx] {
			break
		}

		s.swap(i, smallIdx)
		i = smallIdx
		left = 2*i + 1

	}

}

func (s *SmallHeap) HeapInsert(i int) {
	for s.List[i] < s.List[(i-1)/2] {
		s.swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func (s SmallHeap) isEmpty() bool {
	return s.size == 0
}

func (s *SmallHeap) swap(i, j int) {
	s.List[i], s.List[j] = s.List[j], s.List[i]
}
