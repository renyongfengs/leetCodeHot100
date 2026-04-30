package heap

import (
	"math"
)

type BigHeap struct {
	List []int
	size int
}

func NewBigHeap() *BigHeap {
	return &BigHeap{
		List: make([]int, 0),
		size: 0,
	}
}

func (b *BigHeap) Push(vale int) {
	b.List = append(b.List, vale)
	b.Heapinsert(b.size)
	b.size++
}

func (b *BigHeap) Pop() int {
	if b.isEmpty() {
		return math.MinInt
	}
	out := b.List[0]
	b.size--

	if !b.isEmpty() {
		b.swap(0, b.size)
		b.HeapIfy(0, b.size)
	}

	return out

}

func (b BigHeap) isEmpty() bool {
	return b.size == 0
}

//HeapIfy 从 i 的位置往下调整堆结构
func (b *BigHeap) HeapIfy(i int, heapSize int) {
	left := 2*i + 1
	for left < heapSize {
		childLargest := left
		if left+1 < heapSize && b.List[left] < b.List[left+1] {
			childLargest = left + 1
		}

		if b.List[childLargest] < b.List[i] {
			break
		}

		b.swap(childLargest, i)

		i = childLargest
		left = 2*i + 1

	}
}

//insert 把i位置的值插入，需要与父节点比较大小
func (b *BigHeap) Heapinsert(i int) {
	for b.List[(i-1)/2] < b.List[i] {
		b.swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

//swap 交换i，j的值
func (b *BigHeap) swap(i, j int) {
	b.List[i], b.List[j] = b.List[j], b.List[i]
}
