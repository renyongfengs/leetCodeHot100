package array

import (
	"sort"
)

//56. 合并区间
//相关企业
//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//链接：https://leetcode.cn/problems/merge-intervals/description/?envType=study-plan-v2&envId=top-100-liked
//讲解：https://leetcode.cn/problems/merge-intervals/solutions/2798138/jian-dan-zuo-fa-yi-ji-wei-shi-yao-yao-zh-f2b3/?envType=study-plan-v2&envId=top-100-liked

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, val := range intervals {
		length := len(result)
		if len(result) != 0 && result[length-1][1] >= val[0] {
			result[length-1][1] = max(val[1], result[length-1][1]) //可以合并
		} else {
			result = append(result, val)
		}
	}

	return result
}
