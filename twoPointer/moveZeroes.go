package twoPointers

//283. 移动零
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//链接：https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=top-100-liked

func moveZeroes(nums []int) {
	leftBorder := 0
	for index, val := range nums {
		if val != 0 {
			nums[leftBorder], nums[index] = nums[index], nums[leftBorder]
			leftBorder++
		}
	}
}
