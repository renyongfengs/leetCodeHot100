package twoPointers

// 11. 盛最多水的容器
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 链接：https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-100-liked

// 1.假设：当前容器的容量由左右边界决定(当前左右边界的高度较小值*宽度，宽度最大，如果宽度变小，而中间没有更高的边界，容量只会变小)
// 2.移动：移动较短的边界，寻找更高的边界以增加容量
// 3.终止：左右边界相遇时终止
func maxArea(height []int) int {
	if len(height) < 2 || height == nil {
		return 0
	}
	left, right := 0, len(height)-1
	result := 0
	h := 0
	for left < right {
		width := right - left
		if height[left] > height[right] {
			h = height[right]
			right--
		} else {
			h = height[left]
			left++
		}
		if h*width > result {
			result = h * width
		}
	}

	return result

}

// maxArea2 减少了一些变量的使用
func maxArea2(height []int) int {
	if len(height) < 2 || height == nil {
		return 0
	}
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		result = max(min(height[left], height[right])*(right-left), result)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}

	}

	return result

}
