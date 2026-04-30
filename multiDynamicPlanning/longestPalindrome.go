package multiDynamicPlanning

//5. 最长回文子串
//给你一个字符串 s，找到 s 中最长的 回文 子串。
//链接：https://leetcode.cn/problems/longest-palindromic-substring/?envType=study-plan-v2&envId=top-100-liked
func longestPalindrome(s string) string {
	n := len(s)
	ansLeft, ansRight := 0, 0

	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		// 循环结束后，s[l+1] 到 s[r-1] 是回文串
		if r-l-1 > ansRight-ansLeft {
			ansLeft = l + 1
			ansRight = r // 左闭右开区间
		}
	}

	return s[ansLeft:ansRight]
}
