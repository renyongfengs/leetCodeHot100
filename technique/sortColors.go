package technique

//75. 颜色分类
//给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//必须在不使用库内置的 sort 函数的情况下解决这个问题。
//链接：https://leetcode.cn/problems/sort-colors/description/?envType=study-plan-v2&envId=top-100-liked
//插入排序
func sortColors(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0 && nums[j] >= nums[j+1]; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}
}
