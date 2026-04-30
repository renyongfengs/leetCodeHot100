package greedy

//763. 划分字母区间
//给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。例如，字符串 "ababcc" 能够被分为 ["abab", "cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。
//注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
//返回一个表示每个字符串片段的长度的列表。
//链接：https://leetcode.cn/problems/partition-labels/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/partition-labels/solutions/2806706/ben-zhi-shi-he-bing-qu-jian-jian-ji-xie-ygsn8/?envType=study-plan-v2&envId=top-100-liked
func partitionLabels(s string) (ans []int) {
	last := [26]int{}
	for i, c := range s {
		last[c-'a'] = i // 每个字母最后出现的下标
	}

	start, end := 0, 0
	for i, c := range s {
		end = max(end, last[c-'a']) // 更新当前区间右端点的最大值
		if end == i {               // 当前区间合并完毕
			ans = append(ans, end-start+1) // 区间长度加入答案
			start = i + 1                  // 下一个区间的左端点
		}
	}
	return
}
