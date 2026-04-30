package technique

//169. 多数元素
//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//链接：https://leetcode.cn/problems/majority-element/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/majority-element/solutions/3744717/on-mo-er-tou-piao-fa-yan-jin-zheng-ming-ww1zv/?envType=study-plan-v2&envId=top-100-liked
func majorityElement(nums []int) (ans int) {
	hp := 0
	for _, x := range nums {
		if hp == 0 { // x 是初始擂主，生命值为 1
			ans, hp = x, 1
		} else if x == ans { // 比武，同门加血，否则扣血
			hp++
		} else {
			hp--
		}
	}
	return
}
