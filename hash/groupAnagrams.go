package hash

import (
	`sort`
)

//49. 字母异位词分组
//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//链接:https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	if strs == nil {
		return result
	}
	strMap := make(map[string][]string)
	for _, str := range strs {
		tempStr := sortStr(str)
		strMap[tempStr] = append(strMap[tempStr], str)
	}

	for _, strList := range strMap {
		result = append(result, strList)
	}
	return result
}

func sortStr(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
