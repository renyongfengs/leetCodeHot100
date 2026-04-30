package dynamicPlanning

//70. 爬楼梯
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//链接：https://leetcode.cn/problems/climbing-stairs/description/?envType=study-plan-v2&envId=top-100-liked
func climbStairs(n int) int {
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		tmp := f0 + f1
		f0 = f1
		f1 = tmp
	}
	return f1
}
