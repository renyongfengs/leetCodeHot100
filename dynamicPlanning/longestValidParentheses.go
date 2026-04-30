package dynamicPlanning

//32. 最长有效括号
//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号 子串 的长度。
//左右括号匹配，即每个左括号都有对应的右括号将其闭合的字符串是格式正确的，比如 "(()())"。
// 回溯
func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n)
	ans := 0
	for i := 1; i < n; i++ {
		if s[i] == '(' { // ...(
			dp[i] = 0
		} else if s[i-1] == '(' { // ...()
			dp[i] = 2
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			}
		} else if dp[i-1] > 0 { //   ..(..))
			if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' { // .((..))
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 {
					dp[i] += dp[i-dp[i-1]-2] // (..)((..))
				}
			}

		}
		ans = max(ans, dp[i])
	}
	return ans
}

func longestValidParentheses2(s string) int {
	length := len(s)
	dp := make([]int, length)
	ans := 0
	for i := 1; i < length; i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else if s[i-1] == '(' {
			dp[i] = 2
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			}
		} else if dp[i-1] > 0 {
			if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 {
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
		}

		ans = max(ans, dp[i])
	}

	return ans

}
