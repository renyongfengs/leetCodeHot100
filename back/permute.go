package back

//46. 全排列
//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//链接：https://leetcode.cn/problems/permutations/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/n-queens/solutions/2079586/hui-su-tao-lu-miao-sha-nhuang-hou-shi-pi-mljv/?envType=study-plan-v2&envId=top-100-liked
func permute(nums []int) [][]int {
	length := len(nums)
	path := make([]int, length)
	onPath := make([]bool, length)
	result := make([][]int, 0)
	var dfs func(int)

	dfs = func(i int) {
		//i表示填第几个元素
		if i == length {
			result = append(result, append([]int(nil), path...)) //注意闭包引用
			return
		}
		//按顺序遍历used,找到未被使用的值
		for idx, val := range onPath {
			if !val {
				path[i] = nums[idx]
				onPath[idx] = true
				dfs(i + 1)
				onPath[idx] = false
			}
		}

	}

	dfs(0)

	return result
}
