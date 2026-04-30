package greedy

//121. 买卖股票的最佳时机
//给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-100-liked
func maxProfit(prices []int) int {
	minPrice := prices[0]
	result := 0
	for i := 1; i < len(prices); i++ {
		result = max(result, prices[i]-minPrice)

		minPrice = min(minPrice, prices[i])
	}
	return result
}
