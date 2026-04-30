package stack

import (
	"strings"
	"unicode"
)

//394. 字符串解码
//给定一个经过编码的字符串，返回它解码后的字符串。
//编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//测试用例保证输出的长度不会超过 105。
//链接：https://leetcode.cn/problems/decode-string/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/decode-string/solutions/3744428/di-gui-yong-zhan-mo-ni-di-gui-pythonjava-kcsv/?envType=study-plan-v2&envId=top-100-liked
func decodeString(s string) string {
	type pair struct {
		s string
		k int
	}
	stack := []pair{} // 用于模拟计算机的递归
	res := ""
	k := 0
	for _, c := range s {
		if unicode.IsLetter(c) {
			res += string(c)
		} else if unicode.IsDigit(c) {
			k = k*10 + int(c-'0')
		} else if c == '[' {
			// 模拟递归
			// 在递归之前，把当前递归函数中的局部变量 res 和 k 保存到栈中
			stack = append(stack, pair{res, k})
			// 递归，初始化 res 和 k
			res = ""
			k = 0
		} else { // ']'
			// 递归结束，从栈中恢复递归之前保存的局部变量
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 此时 res 是下层递归的返回值，将其重复 p.k 次，拼接到递归前的 p.s 之后
			res = p.s + strings.Repeat(res, p.k)
		}
	}
	return res
}
