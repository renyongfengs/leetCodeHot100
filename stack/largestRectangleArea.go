package stack

//84. 柱状图中最大的矩形
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//链接：https://leetcode.cn/problems/largest-rectangle-in-histogram/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/largest-rectangle-in-histogram/solutions/2695467/dan-diao-zhan-fu-ti-dan-pythonjavacgojsr-89s7/?envType=study-plan-v2&envId=top-100-liked
func largestRectangleArea(heights []int) (ans int) {
	heights = append(heights, -1) // 最后大火收汁，用 -1 把栈清空
	st := []int{-1}               // 在栈中只有一个数的时候，栈顶的「下面那个数」是 -1，对应 left[i] = -1 的情况
	for right, h := range heights {
		for len(st) > 1 && heights[st[len(st)-1]] >= h {
			i := st[len(st)-1] // 矩形的高（的下标）
			st = st[:len(st)-1]
			left := st[len(st)-1] // 栈顶下面那个数就是 left
			ans = max(ans, heights[i]*(right-left-1))
		}
		st = append(st, right)
	}
	return
}
