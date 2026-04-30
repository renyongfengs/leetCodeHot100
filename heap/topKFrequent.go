package heap

import (
	"container/heap"
)

//347. 前 K 个高频元素
//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//链接：https://leetcode.cn/problems/top-k-frequent-elements/description/?envType=study-plan-v2&envId=top-100-liked

func topKFrequent(nums []int, k int) []int {
	numFre := make(map[int]int)
	for _, n := range nums {
		numFre[n]++
	}

	recordList := new(RecordTimesList)
	heap.Init(recordList)
	for value, fre := range numFre {
		if recordList.Len() < k {
			heap.Push(recordList, RecordTimes{
				Value: value,
				Times: fre,
			})
		} else {
			top := heap.Pop(recordList).(RecordTimes)
			if top.Times >= fre {
				heap.Push(recordList, top)
			} else {
				heap.Push(recordList, RecordTimes{
					Value: value,
					Times: fre,
				})
			}
		}

	}

	result := make([]int, len(*recordList))

	for index, record := range *recordList {
		result[index] = record.Value
	}

	return result

}

type RecordTimes struct {
	Value int
	Times int
}

type RecordTimesList []RecordTimes

func (r RecordTimesList) Len() int {
	return len(r)
}

func (r RecordTimesList) Less(i, j int) bool {
	return r[i].Times < r[j].Times
}

func (r RecordTimesList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r *RecordTimesList) Push(x interface{}) {
	*r = append(*r, x.(RecordTimes))
}

func (r *RecordTimesList) Pop() interface{} {
	old := *r
	ans := old[len(old)-1]
	*r = old[0 : len(old)-1]
	return ans
}

func topKFrequent2(nums []int, k int) []int {
	freVM := make(map[int]int)
	maxFre := 0
	for _, num := range nums {
		freVM[num]++
		maxFre = max(maxFre, freVM[num])
	}

	bucket := make([][]int, maxFre+1)
	for num, fre := range freVM {
		bucket[fre] = append(bucket[fre], num)
	}

	result := make([]int, 0, k)
	for i := maxFre; i >= 0 && len(result) < k; i-- {
		result = append(result, bucket[i]...)
	}

	return result

}
