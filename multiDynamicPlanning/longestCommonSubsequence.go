package multiDynamicPlanning

//1143. 最长公共子序列
//给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
//两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
//链接：https://leetcode.cn/problems/longest-common-subsequence/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/longest-common-subsequence/solutions/2133188/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-lbz5/?envType=study-plan-v2&envId=top-100-liked
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	memo := make([][]int, m)
	for i := range memo {
		tmp := make([]int, n)
		for j := range tmp {
			tmp[j] = -1
		}
		memo[i] = tmp
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}

		val := &memo[i][j]
		if *val != -1 {
			return *val
		}

		if text1[i] == text2[j] {
			*val = dfs(i-1, j-1) + 1
		} else {
			*val = max(dfs(i-1, j), dfs(i, j-1))
		}
		return *val
	}
	return dfs(m-1, n-1)
}

func longestCommonSubsequence2(s, t string) int {
	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i, x := range s {
		for j, y := range t {
			if x == y {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[n][m]
}

func longestCommonSubsequence23(s, t string) int {
	m := len(t)
	f := make([]int, m+1)
	for _, x := range s {
		pre := 0
		for j, y := range t {
			tmp := f[j+1]
			if x == y {
				f[j+1] = pre + 1
			} else {
				f[j+1] = max(f[j+1], f[j])
			}
			pre = tmp
		}
	}
	return f[m]
}
