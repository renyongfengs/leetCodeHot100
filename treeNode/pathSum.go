package treeNode

//437. 路径总和 III
//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//链接：https://leetcode.cn/problems/path-sum-iii/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/path-sum-iii/solutions/2784856/zuo-fa-he-560-ti-shi-yi-yang-de-pythonja-fmzo/?envType=study-plan-v2&envId=top-100-liked
func pathSum(root *TreeNode, targetSum int) int {
	out := 0
	vm := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, ans int) {
		if n == nil {
			return
		}

		ans += n.Val
		// 把 node 当作路径的终点，统计有多少个起点
		out += vm[ans-targetSum]

		vm[ans]++
		dfs(n.Left, ans) //左节点先执行
		dfs(n.Right, ans)

		vm[ans]-- // 恢复现场（撤销 cnt[s]++，去掉左节点的值）
	}

	dfs(root, 0)

	return out
}
