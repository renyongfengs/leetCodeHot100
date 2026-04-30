package slidingWindow

//438. 找到字符串中所有字母异位词
//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//链接：https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/?envType=study-plan-v2&envId=top-100-liked

func findAnagrams(s string, p string) []int {
	pLen, sLen := len(p), len(s)
	if pLen > sLen {
		return nil
	}
	result := make([]int, 0)
	pCount, sCount := [26]int{}, [26]int{}
	for idx, val := range p {
		pCount[val-'a']++    //统计p的字符
		sCount[s[idx]-'a']++ //统计s前面的字串
	}
	if pCount == sCount {
		result = append(result, 0)
	}
	for idx, val := range s[:sLen-pLen] {
		sCount[val-'a']--
		sCount[s[pLen+idx]-'a']++

		if sCount == pCount {
			result = append(result, idx+1)
		}

	}
	return result
}
