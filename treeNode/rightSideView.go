package treeNode

//199. 二叉树的右视图
//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//链接：https://leetcode.cn/problems/binary-tree-right-side-view/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/binary-tree-right-side-view/solutions/2015061/ru-he-ling-huo-yun-yong-di-gui-lai-kan-s-r1nc/?envType=study-plan-v2&envId=top-100-liked
func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	var dfs func(n *TreeNode, deep int)
	dfs = func(n *TreeNode, deep int) {
		if n == nil {
			return
		}
		if deep == len(result) { // 这个深度首次遇到
			result = append(result, n.Val)
		}

		dfs(n.Right, deep+1) // 先递归右子树，保证首次遇到的一定是最右边的节点
		dfs(n.Left, deep+1)
	}
	dfs(root, 0)
	return result
}
