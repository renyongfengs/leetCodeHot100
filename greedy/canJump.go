package greedy

//55. 跳跃游戏
//给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
//链接：https://leetcode.cn/problems/jump-game/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/jump-game/solutions/2798996/liang-chong-li-jie-fang-shi-wei-hu-zui-y-q67s/?envType=study-plan-v2&envId=top-100-liked
func canJump(nums []int) bool {
	mx := 0
	for i, jump := range nums {
		if i > mx { // 无法到达 i
			return false
		}
		mx = max(mx, i+jump) // 从 i 最右可以跳到 i + jump
	}
	return true
}
