package treeNode

//230. 二叉搜索树中第 K 小的元素
//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。
//链接：https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/kth-smallest-element-in-a-bst/solutions/2952810/zhong-xu-bian-li-pythonjavaccgojsrust-by-wc02/?envType=study-plan-v2&envId=top-100-liked
func kthSmallest(root *TreeNode, k int) int {
	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return -1
		}

		leftResult := dfs(n.Left)
		if leftResult != -1 {
			return leftResult
		}
		k--
		if k == 0 {
			return n.Val
		}
		return dfs(n.Right)
	}

	return dfs(root)
}
