package treeNode

//104. 二叉树的最大深度
//给定一个二叉树 root ，返回其最大深度。
//二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
//链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/maximum-depth-of-binary-tree/submissions/675331170/?envType=study-plan-v2&envId=top-100-liked
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	out := 0
	var dfs func(node *TreeNode, deep int)
	dfs = func(node *TreeNode, deep int) {
		if node == nil {
			return
		}
		deep++
		out = max(out, deep)
		dfs(node.Left, deep)
		dfs(node.Right, deep)
	}
	dfs(root, 0)
	return out
}
