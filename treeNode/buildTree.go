package treeNode

import (
	"slices"
)

//105. 从前序与中序遍历序列构造二叉树
//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
//链接：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/?envType=study-plan-v2&envId=top-100-liked
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}

	leftIndex := slices.Index(inorder, preorder[0])
	left := buildTree(preorder[1:1+leftIndex], inorder[:leftIndex])
	right := buildTree(preorder[1+leftIndex:], inorder[leftIndex+1:])
	return &TreeNode{preorder[0], left, right}
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	inorderVM := make(map[int]int, n)
	for i, v := range inorder {
		inorderVM[v] = i
	}
	var dfs func(preL, preR, inoL, inoR int) *TreeNode

	dfs = func(preL, preR, inoL, inoR int) *TreeNode {
		if preL == preR {
			return nil
		}
		leftSize := inorderVM[preorder[preL]] - inoL

		left := dfs(preL+1, preL+1+leftSize, inoL, inoL+leftSize)
		right := dfs(preL+1+leftSize, preR, inoL+leftSize+1, inoR)

		return &TreeNode{preorder[preL], left, right}
	}

	return dfs(0, n, 0, n)
}
