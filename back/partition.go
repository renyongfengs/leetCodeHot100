package back

//131. 分割回文串
//给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//链接：https://leetcode.cn/problems/palindrome-partitioning/description/
//题解：https://leetcode.cn/problems/palindrome-partitioning/solutions/2059414/hui-su-bu-hui-xie-tao-lu-zai-ci-pythonja-fues/
func partition(s string) [][]string {
	result := make([][]string, 0)
	path := make([]string, 0)
	n := len(s)

	var dfs func(i, start int)
	dfs = func(i, start int) {
		// s 分割完毕
		if i == len(s) {
			result = append(result, append([]string(nil), path...))
			return
		}

		// 不分割，不选 i 和 i+1 之间的逗号
		if i < n-1 { // i=n-1 时只能分割
			dfs(i+1, start)
		}

		//选当前值
		subStr := s[start : i+1]
		if isPalindromeStr(subStr) {
			path = append(path, subStr)
			// 考虑 i+1 后面的逗号怎么选
			// start=i+1 表示下一个子串从 i+1 开始
			dfs(i+1, i+1)
			//恢复
			path = path[:len(path)-1]
		}

	}

	dfs(0, 0)

	return result
}

func isPalindromeStr(subStr string) bool {
	for i, j := 0, len(subStr)-1; i < j; i, j = i+1, j-1 {
		if subStr[i] != subStr[j] {
			return false
		}
	}
	return true
}
