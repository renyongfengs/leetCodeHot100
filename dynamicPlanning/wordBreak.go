package dynamicPlanning

//139. 单词拆分
//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
//注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//链接：https://leetcode.cn/problems/word-break/?envType=study-plan-v2&envId=top-100-liked
func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	maxLen := 0
	n := len(s)
	for _, word := range wordDict {
		words[word] = true
		maxLen = max(maxLen, len(word))
	}
	result := make([]bool, n+1)
	result[0] = true
	for i := 0; i <= n; i++ {
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if result[j] && words[s[j:i]] {
				result[i] = true
				break
			}
		}
	}

	return result[n]
}
