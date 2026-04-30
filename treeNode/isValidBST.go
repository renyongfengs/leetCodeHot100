package treeNode

import (
	"math"
)

//98. 验证二叉搜索树
//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//有效 二叉搜索树定义如下：
//节点的左子树只包含 严格小于 当前节点的数。
//节点的右子树只包含 严格大于 当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//链接：https://leetcode.cn/problems/validate-binary-search-tree/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/validate-binary-search-tree/solutions/2020306/qian-xu-zhong-xu-hou-xu-san-chong-fang-f-yxvh/?envType=study-plan-v2&envId=top-100-liked
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(n *TreeNode, left, right int) bool

	dfs = func(n *TreeNode, left, right int) bool {
		if n == nil {
			return true
		}
		x := n.Val
		return left < x && x < right && dfs(n.Left, left, x) && dfs(n.Right, x, right)
	}

	return dfs(root, math.MinInt, math.MaxInt)
}

func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	pre := math.MinInt
	var dfs func(n *TreeNode) bool
	dfs = func(n *TreeNode) bool {
		if n == nil {
			return true
		}
		if !dfs(n.Left) { //
			return false
		}

		if n.Val <= pre { //中
			return false
		}
		pre = n.Val
		return dfs(n.Right)
	}

	return dfs(root)
}

func isValidBST3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(n *TreeNode) (int, int)

	dfs = func(n *TreeNode) (int, int) {
		if n == nil {
			return math.MaxInt, math.MinInt
		}
		lMin, lMax := dfs(n.Left)
		rMin, rMax := dfs(n.Right)
		if n.Val <= lMax || n.Val >= rMin {
			return math.MinInt, math.MaxInt
		}
		return min(lMin, n.Val), max(rMax, n.Val)
	}

	_, val := dfs(root)

	return val != math.MaxInt
}
