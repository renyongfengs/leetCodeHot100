package subString

//76. 最小覆盖子串
//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//注意：
//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。
//链接：https://leetcode.cn/problems/minimum-window-substring/description/?envType=study-plan-v2&envId=top-100-liked

func minWindow(s string, t string) string {
	tVM := [128]int{} //记录字符出现的数量
	less := 0         //记录字符的种类
	for _, value := range t {
		if tVM[value] == 0 {
			less++
		}

		tVM[value]++
	}

	resultLeft, resultRight := -1, len(s)
	left := 0

	for right, value := range s {
		tVM[value]--
		if tVM[value] == 0 {
			less--
		}

		for less == 0 {
			if right-left < resultRight-resultLeft {
				resultLeft, resultRight = left, right
			}

			//滑动窗口左移
			temp := s[left] //左边界值
			if tVM[temp] == 0 {
				less++
			}

			tVM[temp]++
			left++

		}

	}

	if resultLeft < 0 {
		return ""
	}

	return s[resultLeft : resultRight+1]
}
