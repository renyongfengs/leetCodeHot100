package treeNode

//226. 翻转二叉树
//给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
//链接：https://leetcode.cn/problems/invert-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)   // 翻转左子树
	right := invertTree(root.Right) // 翻转右子树

	root.Left, root.Right = right, left //交换左右儿子
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left //交换左右儿子

	invertTree(root.Left)  // 翻转左子树
	invertTree(root.Right) // 翻转右子树
	return root
}
