package back

//7. 电话号码的字母组合
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number/?envType=study-plan-v2&envId=top-100-liked
func letterCombinations(digits string) []string {
	mapping := [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	result := make([]string, 0)
	path := make([]byte, len(digits))
	var dfs func(int)
	dfs = func(i int) {
		if i == len(digits) {
			result = append(result, string(path))
			return
		}

		for _, c := range mapping[digits[i]-'0'] {
			path[i] = byte(c)
			dfs(i + 1)
		}
	}
	dfs(0)
	return result
}
