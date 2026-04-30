package treeNode

//102. 二叉树的层序遍历
//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//链接：https://leetcode.cn/problems/binary-tree-level-order-traversal/description/?envType=study-plan-v2&envId=top-100-liked
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)

	cur := []*TreeNode{root}

	for len(cur) > 0 {
		next := make([]*TreeNode, 0)
		val := make([]int, len(cur))
		for idx, node := range cur {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
			val[idx] = node.Val
		}

		result = append(result, val)
		cur = next
	}

	return result
}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)

	cur := []*TreeNode{root}
	for len(cur) > 0 {
		val := make([]int, len(cur))
		for idx, _ := range val {
			val[idx] = cur[0].Val

			if cur[0].Left != nil {
				cur = append(cur, cur[0].Left)
			}

			if cur[0].Right != nil {
				cur = append(cur, cur[0].Right)
			}

			cur = cur[1:]
		}

		result = append(result, val)

	}

	return result
}
