package treeNode

//236. 二叉树的最近公共祖先
//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，
//最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//链接：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/solutions/2023872/fen-lei-tao-lun-luan-ru-ma-yi-ge-shi-pin-2r95/?envType=study-plan-v2&envId=top-100-liked
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root // 找到 p 或 q 就不往下递归了
	}
	left := lowestCommonAncestor(root.Left, p, q)   //	左节点寻找
	right := lowestCommonAncestor(root.Right, p, q) //	右节点寻找
	if left != nil && right != nil {
		//都不为nil，当前节点就是最近公共祖先
		return root
	}

	if left != nil {
		return left
	}

	return right
}
