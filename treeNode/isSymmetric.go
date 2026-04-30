package treeNode

//101. 对称二叉树
//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//链接：https://leetcode.cn/problems/symmetric-tree/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/symmetric-tree/solutions/2015063/ru-he-ling-huo-yun-yong-di-gui-lai-kan-s-6dq5/?envType=study-plan-v2&envId=top-100-liked
func isSymmetric(root *TreeNode) bool {
	return isSameTreeForIsSymmetric(root.Left, root.Right)

}

func isSameTreeForIsSymmetric(tree1, tree2 *TreeNode) bool {
	//两棵树有一个为nil，另一个不是nil
	if tree1 == nil || tree2 == nil {
		return tree1 == tree2
	}
	//相同的情况，值相同，左右子节点相同
	return tree1.Val == tree2.Val && isSameTreeForIsSymmetric(tree1.Left, tree2.Right) && isSameTreeForIsSymmetric(tree1.Right, tree2.Left)
}
