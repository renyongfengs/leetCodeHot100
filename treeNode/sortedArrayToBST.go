package treeNode

//108. 将有序数组转换为二叉搜索树
//给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。
//链接：https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/?envType=study-plan-v2&envId=top-100-liked
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	length := len(nums)
	return &TreeNode{
		Val:   nums[length/2],
		Left:  sortedArrayToBST(nums[:length/2]),
		Right: sortedArrayToBST(nums[length/2+1:]),
	}
}
