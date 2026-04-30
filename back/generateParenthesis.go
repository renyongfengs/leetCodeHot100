package back

//22. 括号生成
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//链接：https://leetcode.cn/problems/generate-parentheses/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/generate-parentheses/solutions/2071015/hui-su-bu-hui-xie-tao-lu-zai-ci-pythonja-wcdw/?envType=study-plan-v2&envId=top-100-liked
func generateParenthesis(n int) (ans []string) {
	path := make([]byte, n*2) // 所有括号长度都是一样的 2n

	// 目前填了 left 个左括号，right 个右括号
	var dfs func(int, int)
	dfs = func(left, right int) {
		if right == n { // 填完 2n 个括号
			ans = append(ans, string(path)) // 加入答案
			return
		}
		if left < n { // 可以填左括号
			path[left+right] = '(' // 直接覆盖
			dfs(left+1, right)
		}
		if right < left { // 可以填右括号
			path[left+right] = ')' // 直接覆盖
			dfs(left, right+1)
		}
	}

	dfs(0, 0) // 一开始没有填括号
	return
}
