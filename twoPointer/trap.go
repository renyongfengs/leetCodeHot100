package twoPointers

// 42. 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 链接：https://leetcode.cn/problems/trapping-rain-water/description/?envType=study-plan-v2&envId=top-100-liked
// 思路： 使用前缀最大值和后缀最大值数组，计算每个位置能接的雨水量，需要知道每个点的左右最高点
func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	length := len(height)
	//前缀最大值
	preMax := make([]int, length)
	preMax[0] = height[0]
	for i := 1; i < length; i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}
	//后缀最大值
	sufMax := make([]int, length)
	sufMax[length-1] = height[length-1]
	for i := length - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}

	result := 0
	for i := 0; i < length; i++ {
		result += min(preMax[i], sufMax[i]) - height[i]
	}
	return result
}

func trap2(height []int) int {
	if len(height) < 2 {
		return 0
	}

	var (
		result            = 0
		left, right       = 0, len(height) - 1
		leftMax, rightMax = 0, 0
	)

	for left <= right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			result += leftMax - height[left]
			left++
		} else {
			result += rightMax - height[right]
			right--
		}
	}
	return result
}
