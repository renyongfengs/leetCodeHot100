package stack

//739. 每日温度
//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
//如果气温在这之后都不会升高，请在该位置用 0 来代替。
//链接：https://leetcode.cn/problems/daily-temperatures/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/daily-temperatures/?envType=study-plan-v2&envId=top-100-liked
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}

	return ans
}
