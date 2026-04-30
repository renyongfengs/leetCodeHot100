package slidingWindow

//3. 无重复字符的最长子串
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度
//链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-100-liked
//思路：使用滑动窗口，维护一个字符出现的哈希表，右指针不断移动扩大窗口，遇到重复字符时左指针移动缩小窗口，直到没有重复字符为止

func lengthOfLongestSubstring(s string) int {
	charMaper := make(map[byte]int)
	start := 0
	result := 0
	for index, value := range []byte(s) {
		pos, ok := charMaper[value]
		if !ok {
			charMaper[value] = index
		} else {
			if pos >= start {
				start = pos + 1
			}

			charMaper[value] = index
		}

		result = max(index-start+1, result)
	}
	return result
}
