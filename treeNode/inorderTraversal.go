package treeNode

//94. 二叉树的中序遍历
//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//链接：https://leetcode.cn/problems/binary-tree-inorder-traversal/description/?envType=study-plan-v2&envId=top-100-liked
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := inorderTraversal(root.Left)
	left = append(left, root.Val)
	right := inorderTraversal(root.Right)

	return append(left, right...)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
