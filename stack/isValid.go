package stack

//20. 有效的括号
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//每个右括号都有一个对应的相同类型的左括号。
//链接：https://leetcode.cn/problems/valid-parentheses/submissions/?envType=study-plan-v2&envId=top-100-liked
func isValid(s string) bool {
	//长度不是偶数
	if len(s)%2 != 0 {
		return false
	}
	vm := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0)
	for _, v := range s {
		//左括号插入
		if _, ok := vm[v]; !ok {
			stack = append(stack, v)
		} else {
			//	右括号，判断是否存在相邻的左括号
			if len(stack) == 0 || stack[len(stack)-1] != vm[v] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
