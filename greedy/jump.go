package greedy

//45. 跳跃游戏 II
//给定一个长度为 n 的 0 索引整数数组 nums。初始位置在下标 0。
//每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在索引 i 处，你可以跳转到任意 (i + j) 处：
//0 <= j <= nums[i] 且
//i + j < n
//返回到达 n - 1 的最小跳跃次数。测试用例保证可以到达 n - 1。
//链接：https://leetcode.cn/problems/jump-game-ii/description/?envType=study-plan-v2&envId=top-100-liked
func jump(nums []int) int {
	curRight := 0
	nextRight := 0
	result := 0
	for i, num := range nums[0 : len(nums)-1] {
		nextRight = max(nextRight, i+num)
		if i == curRight {
			curRight = nextRight
			result++
		}
	}

	return result
}
