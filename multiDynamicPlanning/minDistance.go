package multiDynamicPlanning

//72. 编辑距离
//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
//你可以对一个单词进行如下三种操作：
//插入一个字符
//删除一个字符
//替换一个字符
//链接：https://leetcode.cn/problems/edit-distance/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/edit-distance/solutions/2133222/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-uo5q/?envType=study-plan-v2&envId=top-100-liked
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		val := &memo[i][j]
		if *val != -1 {
			return *val
		}

		if word1[i] == word2[j] {
			*val = dfs(i-1, j-1)
		} else {
			*val = min(dfs(i, j-1), dfs(i-1, j), dfs(i-1, j-1)) + 1
		}
		return *val
	}

	return dfs(m-1, n-1)
}

func minDistance2(s, t string) int {
	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := range m {
		f[0][j+1] = j + 1
	}
	for i, x := range s {
		f[i+1][0] = i + 1
		for j, y := range t {
			if x == y {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1
			}
		}
	}
	return f[n][m]
}
