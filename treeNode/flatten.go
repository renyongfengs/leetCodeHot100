package treeNode

//114. 二叉树展开为链表
//给你二叉树的根结点 root ，请你将它展开为一个单链表：
//展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
//展开后的单链表应该与二叉树 先序遍历 顺序相同。
//链接：https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
func flatten(root *TreeNode) {
	var head *TreeNode
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}

		dfs(n.Right)
		dfs(n.Left)
		n.Left = nil
		n.Right = head
		head = n
	}

	dfs(root)
}
